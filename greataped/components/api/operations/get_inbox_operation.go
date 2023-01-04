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

/*
func (operation *getInboxOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
