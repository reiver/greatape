package operations

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/api/services"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/operation"
	. "github.com/xeronith/diamante/contracts/service"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/operation"
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
