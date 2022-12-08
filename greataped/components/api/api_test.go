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

func TestSignupApi(test *testing.T) {
	input := &SignupRequest{
		Username: "username",
		Email:    "email",
		Password: "password",
	}

	if output, err := api.Signup(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestVerifyApi(test *testing.T) {
	input := &VerifyRequest{
		Email: "email",
		Token: "token",
		Code:  "code",
	}

	if output, err := api.Verify(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestLoginApi(test *testing.T) {
	input := &LoginRequest{
		Email:    "email",
		Password: "password",
	}

	if output, err := api.Login(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestGetProfileByUserApi(test *testing.T) {
	input := &GetProfileByUserRequest{}

	if output, err := api.GetProfileByUser(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestUpdateProfileByUserApi(test *testing.T) {
	input := &UpdateProfileByUserRequest{
		DisplayName: "display_name",
		Avatar:      "avatar",
		Banner:      "banner",
		Summary:     "summary",
		Github:      "github",
	}

	if output, err := api.UpdateProfileByUser(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestLogoutApi(test *testing.T) {
	input := &LogoutRequest{}

	if output, err := api.Logout(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestWebfingerApi(test *testing.T) {
	input := &WebfingerRequest{
		Resource: "resource",
	}

	if output, err := api.Webfinger(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestGetActorApi(test *testing.T) {
	input := &GetActorRequest{
		Username: "username",
	}

	if output, err := api.GetActor(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestFollowActorApi(test *testing.T) {
	input := &FollowActorRequest{
		Username: "username",
		Acct:     "acct",
	}

	if output, err := api.FollowActor(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestAuthorizeInteractionApi(test *testing.T) {
	input := &AuthorizeInteractionRequest{
		Uri: "uri",
	}

	if output, err := api.AuthorizeInteraction(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestGetFollowersApi(test *testing.T) {
	input := &GetFollowersRequest{
		Username: "username",
	}

	if output, err := api.GetFollowers(input); err != nil {
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
