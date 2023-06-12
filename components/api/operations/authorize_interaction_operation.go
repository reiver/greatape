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

type authorizeInteractionOperation struct {
	Operation

	run func(IContext, *AuthorizeInteractionRequest) (*AuthorizeInteractionResult, error)
}

func AuthorizeInteractionOperation() IOperation {
	return &authorizeInteractionOperation{
		run: AuthorizeInteractionService,
	}
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
	return operation.run(context, payload.(*AuthorizeInteractionRequest))
}
