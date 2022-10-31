package services

import (
	. "github.com/xeronith/diamante/contracts/service"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
)

// noinspection GoUnusedParameter
func ResolveErrorService(context IContext, input *ResolveErrorRequest) (result *ResolveErrorResult, err error) {
	conductor := core.Conductor
	_ = RESOLVE_ERROR_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "resolve_error", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "resolve_error", input, result, err) }()

	var inputDocument IDocument
	if input.Document != nil {
		var err error
		if inputDocument, err = conductor.NewDocument(input.Document.Id, input.Document.Content); err == nil {
		} else {
			return nil, err
		}
	}

	_result, _err := conductor.ResolveError(inputDocument, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*ResolveErrorResult)
	return result, nil
}
