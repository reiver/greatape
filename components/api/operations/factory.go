package operations

import . "github.com/xeronith/diamante/contracts/operation"

type operationFactory struct{}

func (factory *operationFactory) Operations() []IOperation {
	return []IOperation{
		SystemCallOperation(),
		EchoOperation(),
		CheckUsernameAvailabilityOperation(),
		SignupOperation(),
		ResendVerificationCodeOperation(),
		VerifyOperation(),
		LoginOperation(),
		GetProfileByUserOperation(),
		UpdateProfileByUserOperation(),
		ChangePasswordOperation(),
		ResetPasswordOperation(),
		LogoutOperation(),
		WebfingerOperation(),
		GetPackagesOperation(),
		GetActorOperation(),
		FollowActorOperation(),
		AuthorizeInteractionOperation(),
		GetFollowersOperation(),
		GetFollowingOperation(),
		PostToOutboxOperation(),
		GetOutboxOperation(),
		PostToInboxOperation(),
		GetInboxOperation(),
	}
}

func NewFactory() IOperationFactory {
	return &operationFactory{}
}
