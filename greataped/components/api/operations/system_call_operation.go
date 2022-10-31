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

type systemCallOperation struct {
	AdminOperation

	run func(IContext, *SystemCallRequest) (*SystemCallResult, error)
}

func SystemCallOperation() IOperation {
	return &systemCallOperation{
		run: SystemCallService,
	}
}

func (operation *systemCallOperation) Id() (ID, ID) {
	return SYSTEM_CALL_REQUEST, SYSTEM_CALL_RESULT
}

func (operation *systemCallOperation) InputContainer() Pointer {
	return new(SystemCallRequest)
}

func (operation *systemCallOperation) OutputContainer() Pointer {
	return new(SystemCallResult)
}

func (operation *systemCallOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*SystemCallRequest))
}
