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
	EchoRunner  func(IContext, *EchoRequest) (*EchoResult, error)
	EchoRunners []EchoRunner

	echoOperation struct {
		Operation

		runners EchoRunners
	}
)

func EchoOperation() IOperation {
	return &echoOperation{
		runners: EchoRunners{
			EchoService,
		},
	}
}

func (operation *echoOperation) Tag() string {
	return "ECHO"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*EchoRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
