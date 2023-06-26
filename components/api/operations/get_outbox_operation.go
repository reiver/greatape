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
	GetOutboxRunner  func(IContext, *GetOutboxRequest) (*GetOutboxResult, error)
	GetOutboxRunners []GetOutboxRunner

	getOutboxOperation struct {
		Operation

		runners GetOutboxRunners
	}
)

func GetOutboxOperation() IOperation {
	return &getOutboxOperation{
		runners: GetOutboxRunners{
			GetOutboxService,
		},
	}
}

func (operation *getOutboxOperation) Tag() string {
	return "GET_OUTBOX"
}

func (operation *getOutboxOperation) Id() (ID, ID) {
	return GET_OUTBOX_REQUEST, GET_OUTBOX_RESULT
}

func (operation *getOutboxOperation) InputContainer() Pointer {
	return new(GetOutboxRequest)
}

func (operation *getOutboxOperation) OutputContainer() Pointer {
	return new(GetOutboxResult)
}

func (operation *getOutboxOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*GetOutboxRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
