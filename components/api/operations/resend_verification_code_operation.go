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
	ResendVerificationCodeRunner  func(IContext, *ResendVerificationCodeRequest) (*ResendVerificationCodeResult, error)
	ResendVerificationCodeRunners []ResendVerificationCodeRunner

	resendVerificationCodeOperation struct {
		Operation

		runners ResendVerificationCodeRunners
	}
)

func ResendVerificationCodeOperation() IOperation {
	return &resendVerificationCodeOperation{
		runners: ResendVerificationCodeRunners{
			ResendVerificationCodeService,
		},
	}
}

func (operation *resendVerificationCodeOperation) Tag() string {
	return "RESEND_VERIFICATION_CODE"
}

func (operation *resendVerificationCodeOperation) Id() (ID, ID) {
	return RESEND_VERIFICATION_CODE_REQUEST, RESEND_VERIFICATION_CODE_RESULT
}

func (operation *resendVerificationCodeOperation) InputContainer() Pointer {
	return new(ResendVerificationCodeRequest)
}

func (operation *resendVerificationCodeOperation) OutputContainer() Pointer {
	return new(ResendVerificationCodeResult)
}

func (operation *resendVerificationCodeOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*ResendVerificationCodeRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
