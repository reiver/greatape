package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func EchoService(context IContext, input *EchoRequest) (result *EchoResult, err error) {
	source := "echo"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	var inputDocument IDocument
	if input.Document != nil {
		var err error
		if inputDocument, err = Conductor.NewDocument(input.Document.Id, input.Document.Content); err == nil {
		} else {
			return nil, err
		}
	}

	commandResult, err := Conductor.Echo(inputDocument, context.Identity())
	if err != nil {
		return nil, err
	}

	var outputDocument *Document = nil
	if commandResult.Document() != nil {
		outputDocument = &Document{
			Id:      commandResult.Document().Id(),
			Content: commandResult.Document().Content(),
		}
	}

	result = context.ResultContainer().(*EchoResult)
	result.Document = outputDocument
	return result, nil
}
