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
	GetInboxRunner  func(IContext, *GetInboxRequest) (*GetInboxResult, error)
	GetInboxRunners []GetInboxRunner

	getInboxOperation struct {
		Operation

		runners GetInboxRunners
	}
)

func GetInboxOperation() IOperation {
	return &getInboxOperation{
		runners: GetInboxRunners{
			GetInboxService,
		},
	}
}

func (operation *getInboxOperation) Tag() string {
	return "GET_INBOX"
}

func (operation *getInboxOperation) Id() (ID, ID) {
	return GET_INBOX_REQUEST, GET_INBOX_RESULT
}

func (operation *getInboxOperation) InputContainer() Pointer {
	return new(GetInboxRequest)
}

func (operation *getInboxOperation) OutputContainer() Pointer {
	return new(GetInboxResult)
}

func (operation *getInboxOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*GetInboxRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
