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
	LoginRunner  func(IContext, *LoginRequest) (*LoginResult, error)
	LoginRunners []LoginRunner

	loginOperation struct {
		Operation

		runners LoginRunners
	}
)

func LoginOperation() IOperation {
	return &loginOperation{
		runners: LoginRunners{
			LoginService,
		},
	}
}

func (operation *loginOperation) Tag() string {
	return "LOGIN"
}

func (operation *loginOperation) Id() (ID, ID) {
	return LOGIN_REQUEST, LOGIN_RESULT
}

func (operation *loginOperation) InputContainer() Pointer {
	return new(LoginRequest)
}

func (operation *loginOperation) OutputContainer() Pointer {
	return new(LoginResult)
}

func (operation *loginOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*LoginRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
