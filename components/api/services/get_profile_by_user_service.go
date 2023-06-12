package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func GetProfileByUserService(context IContext, input *GetProfileByUserRequest) (result *GetProfileByUserResult, err error) {
	source := "get_profile_by_user"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.GetProfileByUser(context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*GetProfileByUserResult)
	result.Username = commandResult.Username()
	result.DisplayName = commandResult.DisplayName()
	result.Avatar = commandResult.Avatar()
	result.Banner = commandResult.Banner()
	result.Summary = commandResult.Summary()
	result.Github = commandResult.Github()
	return result, nil
}
