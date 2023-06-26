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
	CheckUsernameAvailabilityRunner  func(IContext, *CheckUsernameAvailabilityRequest) (*CheckUsernameAvailabilityResult, error)
	CheckUsernameAvailabilityRunners []CheckUsernameAvailabilityRunner

	checkUsernameAvailabilityOperation struct {
		Operation

		runners CheckUsernameAvailabilityRunners
	}
)

func CheckUsernameAvailabilityOperation() IOperation {
	return &checkUsernameAvailabilityOperation{
		runners: CheckUsernameAvailabilityRunners{
			CheckUsernameAvailabilityService,
		},
	}
}

func (operation *checkUsernameAvailabilityOperation) Tag() string {
	return "CHECK_USERNAME_AVAILABILITY"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*CheckUsernameAvailabilityRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
