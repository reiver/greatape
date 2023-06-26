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
	ChangePasswordRunner  func(IContext, *ChangePasswordRequest) (*ChangePasswordResult, error)
	ChangePasswordRunners []ChangePasswordRunner

	changePasswordOperation struct {
		SecureOperation

		runners ChangePasswordRunners
	}
)

func ChangePasswordOperation() IOperation {
	return &changePasswordOperation{
		runners: ChangePasswordRunners{
			ChangePasswordService,
		},
	}
}

func (operation *changePasswordOperation) Tag() string {
	return "CHANGE_PASSWORD"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*ChangePasswordRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
