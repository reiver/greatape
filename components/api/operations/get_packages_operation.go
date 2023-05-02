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

type getPackagesOperation struct {
	Operation

	run func(IContext, *GetPackagesRequest) (*GetPackagesResult, error)
}

func GetPackagesOperation() IOperation {
	return &getPackagesOperation{
		run: GetPackagesService,
	}
}

func (operation *getPackagesOperation) Id() (ID, ID) {
	return GET_PACKAGES_REQUEST, GET_PACKAGES_RESULT
}

func (operation *getPackagesOperation) InputContainer() Pointer {
	return new(GetPackagesRequest)
}

func (operation *getPackagesOperation) OutputContainer() Pointer {
	return new(GetPackagesResult)
}

func (operation *getPackagesOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*GetPackagesRequest))
}

/*
func (operation *getPackagesOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
