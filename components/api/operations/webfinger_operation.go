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

type webfingerOperation struct {
	Operation

	run func(IContext, *WebfingerRequest) (*WebfingerResult, error)
}

func WebfingerOperation() IOperation {
	return &webfingerOperation{
		run: WebfingerService,
	}
}

func (operation *webfingerOperation) Id() (ID, ID) {
	return WEBFINGER_REQUEST, WEBFINGER_RESULT
}

func (operation *webfingerOperation) InputContainer() Pointer {
	return new(WebfingerRequest)
}

func (operation *webfingerOperation) OutputContainer() Pointer {
	return new(WebfingerResult)
}

func (operation *webfingerOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*WebfingerRequest))
}

/*
func (operation *webfingerOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
