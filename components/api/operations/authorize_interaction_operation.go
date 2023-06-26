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
	AuthorizeInteractionRunner  func(IContext, *AuthorizeInteractionRequest) (*AuthorizeInteractionResult, error)
	AuthorizeInteractionRunners []AuthorizeInteractionRunner

	authorizeInteractionOperation struct {
		Operation

		runners AuthorizeInteractionRunners
	}
)

func AuthorizeInteractionOperation() IOperation {
	return &authorizeInteractionOperation{
		runners: AuthorizeInteractionRunners{
			AuthorizeInteractionService,
		},
	}
}

func (operation *authorizeInteractionOperation) Tag() string {
	return "AUTHORIZE_INTERACTION"
}

func (operation *authorizeInteractionOperation) Id() (ID, ID) {
	return AUTHORIZE_INTERACTION_REQUEST, AUTHORIZE_INTERACTION_RESULT
}

func (operation *authorizeInteractionOperation) InputContainer() Pointer {
	return new(AuthorizeInteractionRequest)
}

func (operation *authorizeInteractionOperation) OutputContainer() Pointer {
	return new(AuthorizeInteractionResult)
}

func (operation *authorizeInteractionOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*AuthorizeInteractionRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
