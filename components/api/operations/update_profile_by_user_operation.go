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
	UpdateProfileByUserRunner  func(IContext, *UpdateProfileByUserRequest) (*UpdateProfileByUserResult, error)
	UpdateProfileByUserRunners []UpdateProfileByUserRunner

	updateProfileByUserOperation struct {
		SecureOperation

		runners UpdateProfileByUserRunners
	}
)

func UpdateProfileByUserOperation() IOperation {
	return &updateProfileByUserOperation{
		runners: UpdateProfileByUserRunners{
			UpdateProfileByUserService,
		},
	}
}

func (operation *updateProfileByUserOperation) Tag() string {
	return "UPDATE_PROFILE_BY_USER"
}

func (operation *updateProfileByUserOperation) Id() (ID, ID) {
	return UPDATE_PROFILE_BY_USER_REQUEST, UPDATE_PROFILE_BY_USER_RESULT
}

func (operation *updateProfileByUserOperation) InputContainer() Pointer {
	return new(UpdateProfileByUserRequest)
}

func (operation *updateProfileByUserOperation) OutputContainer() Pointer {
	return new(UpdateProfileByUserResult)
}

func (operation *updateProfileByUserOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*UpdateProfileByUserRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
