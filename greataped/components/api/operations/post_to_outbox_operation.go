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

/*
func (operation *postToOutboxOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
