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

type getProfileByUserOperation struct {
	SecureOperation

	run func(IContext, *GetProfileByUserRequest) (*GetProfileByUserResult, error)
}

func GetProfileByUserOperation() IOperation {
	return &getProfileByUserOperation{
		run: GetProfileByUserService,
	}
}

func (operation *getProfileByUserOperation) Id() (ID, ID) {
	return GET_PROFILE_BY_USER_REQUEST, GET_PROFILE_BY_USER_RESULT
}

func (operation *getProfileByUserOperation) InputContainer() Pointer {
	return new(GetProfileByUserRequest)
}

func (operation *getProfileByUserOperation) OutputContainer() Pointer {
	return new(GetProfileByUserResult)
}

func (operation *getProfileByUserOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*GetProfileByUserRequest))
}

/*
func (operation *getProfileByUserOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
