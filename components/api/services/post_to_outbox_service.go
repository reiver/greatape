package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func PostToOutboxService(context IContext, input *PostToOutboxRequest) (result *PostToOutboxResult, err error) {
	source := "post_to_outbox"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.PostToOutbox(input.Username, input.Body, context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*PostToOutboxResult)
	result.Body = commandResult.Body()
	return result, nil
}
