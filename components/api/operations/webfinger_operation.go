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
	WebfingerRunner  func(IContext, *WebfingerRequest) (*WebfingerResult, error)
	WebfingerRunners []WebfingerRunner

	webfingerOperation struct {
		Operation

		runners WebfingerRunners
	}
)

func WebfingerOperation() IOperation {
	return &webfingerOperation{
		runners: WebfingerRunners{
			WebfingerService,
		},
	}
}

func (operation *webfingerOperation) Tag() string {
	return "WEBFINGER"
}

func (operation *webfingerOperation) Id() (ID, ID) {
	return WEBFINGER_REQUEST, WEBFINGER_RESULT
}

func (operation *webfingerOperation) InputContainer() Pointer {
	return new(WebfingerRequest)
}

func (operation *webfingerOperation) OutputContainer() Pointer {
	return new(WebfingerResult)
}

func (operation *webfingerOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*WebfingerRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
