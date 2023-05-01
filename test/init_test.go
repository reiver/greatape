package test

import (
	"os"
	"testing"
	"time"

	"github.com/reiver/greatape/components/api/handlers"
	"github.com/reiver/greatape/components/api/operations"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	"github.com/reiver/greatape/components/model/repository"
	"github.com/reiver/greatape/providers/outbound/email"
	"github.com/reiver/greatape/providers/outbound/sms"
	"github.com/xeronith/diamante/analytics"
	"github.com/xeronith/diamante/logging"
	"github.com/xeronith/diamante/server"
	"github.com/xeronith/diamante/settings"
)

var (
	apiLocal, apiRemote IApi
	remoteEndpoint      = "https://greatape.social"
)

func TestMain(main *testing.M) {
	// logging
	logger := logging.NewLogger(false)
	logger.SetLevel(logging.LEVEL_SUPPRESS_SYS_COMP)
	// configuration
	configuration := settings.NewTestConfiguration()
	configuration.GetPostgreSQLConfiguration().SetDatabase("greatape")
	// factories
	operationsFactory := operations.NewFactory()
	handlersFactory := handlers.NewFactory()
	// providers
	measurementsProvider := analytics.NewInfluxDbProvider(configuration, logger)
	emailProvider := email.NewProvider(logger)
	smsProvider := sms.NewProvider(logger)

	if testServer, err := server.New(configuration, operationsFactory, handlersFactory, OPCODES); err != nil {
		logger.Fatal(err)
	} else {
		if err := repository.Initialize(configuration, logger); err != nil {
			logger.Fatal(err)
		}

		if err := core.Initialize(configuration, logger); err != nil {
			logger.Fatal(err)
		}

		testServer.Localizer().Register(Errors)
		testServer.SetSecurityHandler(core.Conductor.IdentityManager())
		testServer.SetMeasurementsProvider(measurementsProvider)
		testServer.SetEmailProvider(emailProvider)
		testServer.SetSMSProvider(smsProvider)

		go testServer.Start()

		apiLocal = core.NewApi(testServer.PassiveEndpoint(), logger)
		apiRemote = core.NewApi(remoteEndpoint, logger)

		os.Exit(main.Run())
	}
}

func Run(t *testing.T, api IApi, f func(IApi) error) {
	defer func() {
		delay := time.Millisecond * 50
		time.Sleep(repository.AUTO_FLUSH_DURATION + delay)
	}()

	if err := f(api); err != nil {
		t.Fatal(err)
	}
}
