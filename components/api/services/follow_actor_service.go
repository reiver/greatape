package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func FollowActorService(context IContext, input *FollowActorRequest) (result *FollowActorResult, err error) {
	source := "follow_actor"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.FollowActor(input.Username, input.Acct, context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*FollowActorResult)
	result.Url = commandResult.Url()
	return result, nil
}
