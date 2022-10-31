package services

import (
	. "github.com/xeronith/diamante/contracts/service"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
)

// noinspection GoUnusedParameter
func EchoService(context IContext, input *EchoRequest) (result *EchoResult, err error) {
	conductor := core.Conductor
	_ = ECHO_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "echo", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "echo", input, result, err) }()

	var inputDocument IDocument
	if input.Document != nil {
		var err error
		if inputDocument, err = conductor.NewDocument(input.Document.Id, input.Document.Content); err == nil {
		} else {
			return nil, err
		}
	}

	_result, _err := conductor.Echo(inputDocument, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	var outputDocument *Document = nil
	if _result.Document() != nil {
		outputDocument = &Document{
			Id:      _result.Document().Id(),
			Content: _result.Document().Content(),
		}
	}

	result = context.ResultContainer().(*EchoResult)
	result.Document = outputDocument
	return result, nil
}
