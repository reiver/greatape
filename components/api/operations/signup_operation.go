package operations

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/api/services"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/operation"
	. "github.com/xeronith/diamante/contracts/service"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/operation"
)

type (
	SignupRunner  func(IContext, *SignupRequest) (*SignupResult, error)
	SignupRunners []SignupRunner

	signupOperation struct {
		Operation

		runners SignupRunners
	}
)

func SignupOperation() IOperation {
	return &signupOperation{
		runners: SignupRunners{
			SignupService,
		},
	}
}

func (operation *signupOperation) Tag() string {
	return "SIGNUP"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*SignupRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
