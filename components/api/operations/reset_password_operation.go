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
	ResetPasswordRunner  func(IContext, *ResetPasswordRequest) (*ResetPasswordResult, error)
	ResetPasswordRunners []ResetPasswordRunner

	resetPasswordOperation struct {
		Operation

		runners ResetPasswordRunners
	}
)

func ResetPasswordOperation() IOperation {
	return &resetPasswordOperation{
		runners: ResetPasswordRunners{
			ResetPasswordService,
		},
	}
}

func (operation *resetPasswordOperation) Tag() string {
	return "RESET_PASSWORD"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*ResetPasswordRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
