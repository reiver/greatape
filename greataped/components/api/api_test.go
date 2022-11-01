package api_test

import (
	"os"
	"testing"

	"github.com/xeronith/diamante/analytics"
	"github.com/xeronith/diamante/logging"
	"github.com/xeronith/diamante/server"
	"github.com/xeronith/diamante/settings"
	"rail.town/infrastructure/components/api/handlers"
	"rail.town/infrastructure/components/api/operations"
	. "rail.town/infrastructure/components/api/protobuf"
	"rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
	"rail.town/infrastructure/components/model/repository"
	"rail.town/infrastructure/providers/outbound/email"
	"rail.town/infrastructure/providers/outbound/sms"
)

var api IApi

func TestReloadSystemComponentApi(test *testing.T) {
	input := &SystemCallRequest{
		Command: "",
	}

	if output, err := api.SystemCall(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestEchoApi(test *testing.T) {
	input := &EchoRequest{
		Document: nil,
	}

	if output, err := api.Echo(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

//region Initialization

func TestMain(main *testing.M) {
	logger := logging.NewLogger(false)
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

		testServer.Localizer().Register(constants.Errors)
		testServer.SetSecurityHandler(core.Conductor.IdentityManager())
		testServer.SetMeasurementsProvider(measurementsProvider)
		testServer.SetEmailProvider(emailProvider)
		testServer.SetSMSProvider(smsProvider)

		go testServer.Start()

		api = core.NewApi(testServer.PassiveEndpoint(), logger)
		os.Exit(main.Run())
	}
}

//endregion
