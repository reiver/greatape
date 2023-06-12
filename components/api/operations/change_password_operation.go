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

type changePasswordOperation struct {
	SecureOperation

	run func(IContext, *ChangePasswordRequest) (*ChangePasswordResult, error)
}

func ChangePasswordOperation() IOperation {
	return &changePasswordOperation{
		run: ChangePasswordService,
	}
}

func (operation *changePasswordOperation) Id() (ID, ID) {
	return CHANGE_PASSWORD_REQUEST, CHANGE_PASSWORD_RESULT
}

func (operation *changePasswordOperation) InputContainer() Pointer {
	return new(ChangePasswordRequest)
}

func (operation *changePasswordOperation) OutputContainer() Pointer {
	return new(ChangePasswordResult)
}

func (operation *changePasswordOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*ChangePasswordRequest))
}
