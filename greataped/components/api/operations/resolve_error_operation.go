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

type resolveErrorOperation struct {
	SecureOperation

	run func(IContext, *ResolveErrorRequest) (*ResolveErrorResult, error)
}

func ResolveErrorOperation() IOperation {
	return &resolveErrorOperation{
		run: ResolveErrorService,
	}
}

func (operation *resolveErrorOperation) Id() (ID, ID) {
	return RESOLVE_ERROR_REQUEST, RESOLVE_ERROR_RESULT
}

func (operation *resolveErrorOperation) InputContainer() Pointer {
	return new(ResolveErrorRequest)
}

func (operation *resolveErrorOperation) OutputContainer() Pointer {
	return new(ResolveErrorResult)
}

func (operation *resolveErrorOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*ResolveErrorRequest))
}

/*
func (operation *resolveErrorOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
