package test

import (
	"os"
	"testing"
	"time"

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

var (
	apiLocal, apiRemote IApi
	remoteEndpoint      = "https://greatape.social"
)

func TestMain(main *testing.M) {
	logger := logging.NewLogger(false)
	logger.SetLevel(logging.LEVEL_SUPPRESS_SYS_COMP)
	configuration := settings.NewTestConfiguration()
	operationsFactory := operations.NewFactory()
	handlersFactory := handlers.NewFactory()
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
