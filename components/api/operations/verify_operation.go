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
	VerifyRunner  func(IContext, *VerifyRequest) (*VerifyResult, error)
	VerifyRunners []VerifyRunner

	verifyOperation struct {
		Operation

		runners VerifyRunners
	}
)

func VerifyOperation() IOperation {
	return &verifyOperation{
		runners: VerifyRunners{
			VerifyService,
		},
	}
}

func (operation *verifyOperation) Tag() string {
	return "VERIFY"
}

func (operation *verifyOperation) Id() (ID, ID) {
	return VERIFY_REQUEST, VERIFY_RESULT
}

func (operation *verifyOperation) InputContainer() Pointer {
	return new(VerifyRequest)
}

func (operation *verifyOperation) OutputContainer() Pointer {
	return new(VerifyResult)
}

func (operation *verifyOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*VerifyRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
