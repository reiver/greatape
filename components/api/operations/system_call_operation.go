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

type (
	SystemCallRunner  func(IContext, *SystemCallRequest) (*SystemCallResult, error)
	SystemCallRunners []SystemCallRunner

	systemCallOperation struct {
		AdminOperation

		runners SystemCallRunners
	}
)

func SystemCallOperation() IOperation {
	return &systemCallOperation{
		runners: SystemCallRunners{
			SystemCallService,
		},
	}
}

func (operation *systemCallOperation) Tag() string {
	return "SYSTEM_CALL"
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
	return operation.runners[0](context, payload.(*SystemCallRequest))
}
