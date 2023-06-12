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

type getInboxOperation struct {
	Operation

	run func(IContext, *GetInboxRequest) (*GetInboxResult, error)
}

func GetInboxOperation() IOperation {
	return &getInboxOperation{
		run: GetInboxService,
	}
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
	return operation.run(context, payload.(*GetInboxRequest))
}
