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

type postToInboxOperation struct {
	Operation

	run func(IContext, *PostToInboxRequest) (*PostToInboxResult, error)
}

func PostToInboxOperation() IOperation {
	return &postToInboxOperation{
		run: PostToInboxService,
	}
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
	return operation.run(context, payload.(*PostToInboxRequest))
}

/*
func (operation *postToInboxOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
