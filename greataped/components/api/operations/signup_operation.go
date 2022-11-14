package operations

import (
	. "github.com/xeronith/diamante/contracts/operation"
	. "github.com/xeronith/diamante/contracts/service"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/operation"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/api/services"
	. "rail.town/infrastructure/components/contracts"
)

type signupOperation struct {
	Operation

	run func(IContext, *SignupRequest) (*SignupResult, error)
}

func SignupOperation() IOperation {
	return &signupOperation{
		run: SignupService,
	}
}

func (operation *signupOperation) Id() (ID, ID) {
	return SIGNUP_REQUEST, SIGNUP_RESULT
}

func (operation *signupOperation) InputContainer() Pointer {
	return new(SignupRequest)
}

func (operation *signupOperation) OutputContainer() Pointer {
	return new(SignupResult)
}

func (operation *signupOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*SignupRequest))
}

/*
func (operation *signupOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
