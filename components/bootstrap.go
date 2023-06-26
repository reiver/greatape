package components

import (
	"flag"
	"runtime"

	"github.com/reiver/greatape/components/api/handlers"
	"github.com/reiver/greatape/components/api/operations"
	. "github.com/reiver/greatape/components/constants"
	"github.com/reiver/greatape/components/core"
	"github.com/reiver/greatape/components/model/repository"
	"github.com/reiver/greatape/providers/outbound/email"
	"github.com/reiver/greatape/providers/outbound/sms"
	"github.com/xeronith/diamante/analytics"
	"github.com/xeronith/diamante/logging"
	"github.com/xeronith/diamante/server"
	"github.com/xeronith/diamante/settings"
)

var configFilePath = flag.String("config", "config.yaml", "Configuration File Path")

func Run() {
	flag.Parse()
	if !core.Dockerized {
		runtime.GOMAXPROCS(10)
	}

	logger := logging.NewLogger(core.Dockerized)

	configuration, err := settings.NewConfiguration(*configFilePath, core.Dockerized)
	if err != nil {
		logger.Fatal(err)
	}

	operationsFactory := operations.NewFactory()
	handlersFactory := handlers.NewFactory()
	measurementsProvider := analytics.NewInfluxDbProvider(configuration, logger)
	emailProvider := email.NewProvider(logger)
	smsProvider := sms.NewProvider(logger)

	if mainServer, err := server.New(configuration, operationsFactory, handlersFactory); err != nil {
		logger.Fatal(err)
	} else {
		if err := repository.Initialize(configuration, logger); err != nil {
			logger.Fatal(err)
		}

		if err := core.Initialize(configuration, logger); err != nil {
			logger.Fatal(err)
		}

		mainServer.Localizer().Register(Errors)
		mainServer.SetSecurityHandler(core.Conductor.IdentityManager())
		mainServer.SetMeasurementsProvider(measurementsProvider)
		mainServer.SetEmailProvider(emailProvider)
		mainServer.SetSMSProvider(smsProvider)

		mainServer.Start()
	}
}
