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

type updateProfileByUserOperation struct {
	SecureOperation

	run func(IContext, *UpdateProfileByUserRequest) (*UpdateProfileByUserResult, error)
}

func UpdateProfileByUserOperation() IOperation {
	return &updateProfileByUserOperation{
		run: UpdateProfileByUserService,
	}
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
	return operation.run(context, payload.(*UpdateProfileByUserRequest))
}

/*
func (operation *updateProfileByUserOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
