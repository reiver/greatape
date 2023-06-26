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
	LogoutRunner  func(IContext, *LogoutRequest) (*LogoutResult, error)
	LogoutRunners []LogoutRunner

	logoutOperation struct {
		SecureOperation

		runners LogoutRunners
	}
)

func LogoutOperation() IOperation {
	return &logoutOperation{
		runners: LogoutRunners{
			LogoutService,
		},
	}
}

func (operation *logoutOperation) Tag() string {
	return "LOGOUT"
}

func (operation *logoutOperation) Id() (ID, ID) {
	return LOGOUT_REQUEST, LOGOUT_RESULT
}

func (operation *logoutOperation) InputContainer() Pointer {
	return new(LogoutRequest)
}

func (operation *logoutOperation) OutputContainer() Pointer {
	return new(LogoutResult)
}

func (operation *logoutOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*LogoutRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
