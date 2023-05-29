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

type checkUsernameAvailabilityOperation struct {
	Operation

	run func(IContext, *CheckUsernameAvailabilityRequest) (*CheckUsernameAvailabilityResult, error)
}

func CheckUsernameAvailabilityOperation() IOperation {
	return &checkUsernameAvailabilityOperation{
		run: CheckUsernameAvailabilityService,
	}
}

func (operation *checkUsernameAvailabilityOperation) Id() (ID, ID) {
	return CHECK_USERNAME_AVAILABILITY_REQUEST, CHECK_USERNAME_AVAILABILITY_RESULT
}

func (operation *checkUsernameAvailabilityOperation) InputContainer() Pointer {
	return new(CheckUsernameAvailabilityRequest)
}

func (operation *checkUsernameAvailabilityOperation) OutputContainer() Pointer {
	return new(CheckUsernameAvailabilityResult)
}

func (operation *checkUsernameAvailabilityOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*CheckUsernameAvailabilityRequest))
}

/*
func (operation *checkUsernameAvailabilityOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
