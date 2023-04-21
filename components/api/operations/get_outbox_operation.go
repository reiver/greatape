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

type getOutboxOperation struct {
	Operation

	run func(IContext, *GetOutboxRequest) (*GetOutboxResult, error)
}

func GetOutboxOperation() IOperation {
	return &getOutboxOperation{
		run: GetOutboxService,
	}
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
	return operation.run(context, payload.(*GetOutboxRequest))
}

/*
func (operation *getOutboxOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
