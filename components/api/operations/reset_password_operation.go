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

type resetPasswordOperation struct {
	Operation

	run func(IContext, *ResetPasswordRequest) (*ResetPasswordResult, error)
}

func ResetPasswordOperation() IOperation {
	return &resetPasswordOperation{
		run: ResetPasswordService,
	}
}

func (operation *resetPasswordOperation) Id() (ID, ID) {
	return RESET_PASSWORD_REQUEST, RESET_PASSWORD_RESULT
}

func (operation *resetPasswordOperation) InputContainer() Pointer {
	return new(ResetPasswordRequest)
}

func (operation *resetPasswordOperation) OutputContainer() Pointer {
	return new(ResetPasswordResult)
}

func (operation *resetPasswordOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*ResetPasswordRequest))
}
