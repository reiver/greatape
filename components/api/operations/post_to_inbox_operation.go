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
	PostToInboxRunner  func(IContext, *PostToInboxRequest) (*PostToInboxResult, error)
	PostToInboxRunners []PostToInboxRunner

	postToInboxOperation struct {
		Operation

		runners PostToInboxRunners
	}
)

func PostToInboxOperation() IOperation {
	return &postToInboxOperation{
		runners: PostToInboxRunners{
			PostToInboxService,
		},
	}
}

func (operation *postToInboxOperation) Tag() string {
	return "POST_TO_INBOX"
}

func (operation *postToInboxOperation) Id() (ID, ID) {
	return POST_TO_INBOX_REQUEST, POST_TO_INBOX_RESULT
}

func (operation *postToInboxOperation) InputContainer() Pointer {
	return new(PostToInboxRequest)
}

func (operation *postToInboxOperation) OutputContainer() Pointer {
	return new(PostToInboxResult)
}

func (operation *postToInboxOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*PostToInboxRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
