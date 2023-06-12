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

type echoOperation struct {
	Operation

	run func(IContext, *EchoRequest) (*EchoResult, error)
}

func EchoOperation() IOperation {
	return &echoOperation{
		run: EchoService,
	}
}

func (operation *echoOperation) Id() (ID, ID) {
	return ECHO_REQUEST, ECHO_RESULT
}

func (operation *echoOperation) InputContainer() Pointer {
	return new(EchoRequest)
}

func (operation *echoOperation) OutputContainer() Pointer {
	return new(EchoResult)
}

func (operation *echoOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*EchoRequest))
}
