package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func PostToInboxService(context IContext, input *PostToInboxRequest) (result *PostToInboxResult, err error) {
	source := "post_to_inbox"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.PostToInbox(input.Username, input.Body, context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*PostToInboxResult)
	result.Body = commandResult.Body()
	return result, nil
}
