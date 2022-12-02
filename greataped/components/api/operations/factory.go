package operations

import . "github.com/xeronith/diamante/contracts/operation"

type operationFactory struct{}

func (factory *operationFactory) Operations() []IOperation {
	return []IOperation{
		SystemCallOperation(),
		EchoOperation(),
		SignupOperation(),
		VerifyOperation(),
		LoginOperation(),
		GetProfileByUserOperation(),
		UpdateProfileByUserOperation(),
		LogoutOperation(),
		WebfingerOperation(),
		GetActorOperation(),
		FollowActorOperation(),
	}
}

func NewFactory() IOperationFactory {
	return &operationFactory{}
}
