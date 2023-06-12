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

type verifyOperation struct {
	Operation

	run func(IContext, *VerifyRequest) (*VerifyResult, error)
}

func VerifyOperation() IOperation {
	return &verifyOperation{
		run: VerifyService,
	}
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
	return operation.run(context, payload.(*VerifyRequest))
}
