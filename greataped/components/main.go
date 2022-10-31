package main

import (
	"flag"
	"runtime"

	"github.com/xeronith/diamante/analytics"
	"github.com/xeronith/diamante/logging"
	"github.com/xeronith/diamante/server"
	"github.com/xeronith/diamante/settings"
	"rail.town/infrastructure/components/api/handlers"
	"rail.town/infrastructure/components/api/operations"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
	"rail.town/infrastructure/components/model/repository"
	"rail.town/infrastructure/providers/outbound/email"
	"rail.town/infrastructure/providers/outbound/sms"
)

var configFilePath = flag.String("config", "config.yaml", "Configuration File Path")

func main() {
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

	if mainServer, err := server.New(configuration, operationsFactory, handlersFactory, OPCODES); err != nil {
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
