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
	GetProfileByUserRunner  func(IContext, *GetProfileByUserRequest) (*GetProfileByUserResult, error)
	GetProfileByUserRunners []GetProfileByUserRunner

	getProfileByUserOperation struct {
		SecureOperation

		runners GetProfileByUserRunners
	}
)

func GetProfileByUserOperation() IOperation {
	return &getProfileByUserOperation{
		runners: GetProfileByUserRunners{
			GetProfileByUserService,
		},
	}
}

func (operation *getProfileByUserOperation) Tag() string {
	return "GET_PROFILE_BY_USER"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*GetProfileByUserRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
