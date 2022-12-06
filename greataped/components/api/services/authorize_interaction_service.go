package services

import (
	. "github.com/xeronith/diamante/contracts/service"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
)

// noinspection GoUnusedParameter
func AuthorizeInteractionService(context IContext, input *AuthorizeInteractionRequest) (result *AuthorizeInteractionResult, err error) {
	conductor := core.Conductor
	_ = AUTHORIZE_INTERACTION_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "authorize_interaction", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "authorize_interaction", input, result, err) }()

	_result, _err := conductor.AuthorizeInteraction(input.Uri, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*AuthorizeInteractionResult)
	result.Uri = _result.Uri()
	result.Success = _result.Success()
	return result, nil
}
