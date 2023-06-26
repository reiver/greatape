package api_test

import (
	"os"
	"testing"

	"github.com/reiver/greatape/components/api/handlers"
	"github.com/reiver/greatape/components/api/operations"
	. "github.com/reiver/greatape/components/api/protobuf"
	"github.com/reiver/greatape/components/constants"
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

func TestGetServerConfigurationApi(test *testing.T) {
	input := &GetServerConfigurationRequest{}

	if output, err := api.GetServerConfiguration(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestCheckUsernameAvailabilityApi(test *testing.T) {
	input := &CheckUsernameAvailabilityRequest{
		Username: "username",
	}

	if output, err := api.CheckUsernameAvailability(input); err != nil {
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

func TestResendVerificationCodeApi(test *testing.T) {
	input := &ResendVerificationCodeRequest{
		Email: "email",
	}

	if output, err := api.ResendVerificationCode(input); err != nil {
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

func TestChangePasswordApi(test *testing.T) {
	input := &ChangePasswordRequest{
		CurrentPassword: "current_password",
		NewPassword:     "new_password",
	}

	if output, err := api.ChangePassword(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestResetPasswordApi(test *testing.T) {
	input := &ResetPasswordRequest{
		UsernameOrEmail: "username_or_email",
	}

	if output, err := api.ResetPassword(input); err != nil {
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

func TestGetPackagesApi(test *testing.T) {
	input := &GetPackagesRequest{}

	if output, err := api.GetPackages(input); err != nil {
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

func TestGetFollowingApi(test *testing.T) {
	input := &GetFollowingRequest{
		Username: "username",
	}

	if output, err := api.GetFollowing(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestPostToOutboxApi(test *testing.T) {
	input := &PostToOutboxRequest{
		Username: "username",
		Body:     nil,
	}

	if output, err := api.PostToOutbox(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestGetOutboxApi(test *testing.T) {
	input := &GetOutboxRequest{
		Username: "username",
	}

	if output, err := api.GetOutbox(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestPostToInboxApi(test *testing.T) {
	input := &PostToInboxRequest{
		Username: "username",
		Body:     nil,
	}

	if output, err := api.PostToInbox(input); err != nil {
		test.Fatal(err)
	} else if output == nil {
		test.Fail()
	}
}

func TestGetInboxApi(test *testing.T) {
	input := &GetInboxRequest{
		Username: "username",
	}

	if output, err := api.GetInbox(input); err != nil {
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

	if testServer, err := server.New(configuration, operationsFactory, handlersFactory); err != nil {
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
