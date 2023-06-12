package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func GetFollowingService(context IContext, input *GetFollowingRequest) (result *GetFollowingResult, err error) {
	source := "get_following"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.GetFollowing(input.Username, context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*GetFollowingResult)
	result.Context = commandResult.Context()
	result.Id = commandResult.Id()
	result.Type = commandResult.Type()
	result.TotalItems = commandResult.TotalItems()
	result.OrderedItems = commandResult.OrderedItems()
	result.First = commandResult.First()
	return result, nil
}
