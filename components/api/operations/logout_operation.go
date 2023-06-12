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

type logoutOperation struct {
	SecureOperation

	run func(IContext, *LogoutRequest) (*LogoutResult, error)
}

func LogoutOperation() IOperation {
	return &logoutOperation{
		run: LogoutService,
	}
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
	return operation.run(context, payload.(*LogoutRequest))
}
