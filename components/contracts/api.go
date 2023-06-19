package contracts

import . "github.com/reiver/greatape/components/api/protobuf"

type IApi interface {
	SetToken(string)
	SetDebugMode(bool)
	//API Methods
	SystemCall(*SystemCallRequest) (*SystemCallResult, error)
	Echo(*EchoRequest) (*EchoResult, error)
	CheckUsernameAvailability(*CheckUsernameAvailabilityRequest) (*CheckUsernameAvailabilityResult, error)
	Signup(*SignupRequest) (*SignupResult, error)
	ResendVerificationCode(*ResendVerificationCodeRequest) (*ResendVerificationCodeResult, error)
	Verify(*VerifyRequest) (*VerifyResult, error)
	Login(*LoginRequest) (*LoginResult, error)
	GetProfileByUser(*GetProfileByUserRequest) (*GetProfileByUserResult, error)
	UpdateProfileByUser(*UpdateProfileByUserRequest) (*UpdateProfileByUserResult, error)
	ChangePassword(*ChangePasswordRequest) (*ChangePasswordResult, error)
	ResetPassword(*ResetPasswordRequest) (*ResetPasswordResult, error)
	Logout(*LogoutRequest) (*LogoutResult, error)
	Webfinger(*WebfingerRequest) (*WebfingerResult, error)
	GetPackages(*GetPackagesRequest) (*GetPackagesResult, error)
	GetActor(*GetActorRequest) (*GetActorResult, error)
	FollowActor(*FollowActorRequest) (*FollowActorResult, error)
	AuthorizeInteraction(*AuthorizeInteractionRequest) (*AuthorizeInteractionResult, error)
	GetFollowers(*GetFollowersRequest) (*GetFollowersResult, error)
	GetFollowing(*GetFollowingRequest) (*GetFollowingResult, error)
	PostToOutbox(*PostToOutboxRequest) (*PostToOutboxResult, error)
	GetOutbox(*GetOutboxRequest) (*GetOutboxResult, error)
	PostToInbox(*PostToInboxRequest) (*PostToInboxResult, error)
	GetInbox(*GetInboxRequest) (*GetInboxResult, error)
}
