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
	PostToOutboxRunner  func(IContext, *PostToOutboxRequest) (*PostToOutboxResult, error)
	PostToOutboxRunners []PostToOutboxRunner

	postToOutboxOperation struct {
		Operation

		runners PostToOutboxRunners
	}
)

func PostToOutboxOperation() IOperation {
	return &postToOutboxOperation{
		runners: PostToOutboxRunners{
			PostToOutboxService,
		},
	}
}

func (operation *postToOutboxOperation) Tag() string {
	return "POST_TO_OUTBOX"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*PostToOutboxRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
