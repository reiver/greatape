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
