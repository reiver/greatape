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

type resendVerificationCodeOperation struct {
	Operation

	run func(IContext, *ResendVerificationCodeRequest) (*ResendVerificationCodeResult, error)
}

func ResendVerificationCodeOperation() IOperation {
	return &resendVerificationCodeOperation{
		run: ResendVerificationCodeService,
	}
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
	return operation.run(context, payload.(*ResendVerificationCodeRequest))
}
