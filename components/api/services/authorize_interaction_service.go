package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func AuthorizeInteractionService(context IContext, input *AuthorizeInteractionRequest) (result *AuthorizeInteractionResult, err error) {
	source := "authorize_interaction"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.AuthorizeInteraction(input.Uri, context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*AuthorizeInteractionResult)
	result.Uri = commandResult.Uri()
	result.Success = commandResult.Success()
	return result, nil
}
