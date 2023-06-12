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

type postToOutboxOperation struct {
	Operation

	run func(IContext, *PostToOutboxRequest) (*PostToOutboxResult, error)
}

func PostToOutboxOperation() IOperation {
	return &postToOutboxOperation{
		run: PostToOutboxService,
	}
}

func (operation *postToOutboxOperation) Id() (ID, ID) {
	return POST_TO_OUTBOX_REQUEST, POST_TO_OUTBOX_RESULT
}

func (operation *postToOutboxOperation) InputContainer() Pointer {
	return new(PostToOutboxRequest)
}

func (operation *postToOutboxOperation) OutputContainer() Pointer {
	return new(PostToOutboxResult)
}

func (operation *postToOutboxOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*PostToOutboxRequest))
}
