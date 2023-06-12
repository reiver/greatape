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
