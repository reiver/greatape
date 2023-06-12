package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func UpdateProfileByUserService(context IContext, input *UpdateProfileByUserRequest) (result *UpdateProfileByUserResult, err error) {
	source := "update_profile_by_user"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.UpdateProfileByUser(input.DisplayName, input.Avatar, input.Banner, input.Summary, input.Github, context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*UpdateProfileByUserResult)
	result.DisplayName = commandResult.DisplayName()
	result.Avatar = commandResult.Avatar()
	result.Banner = commandResult.Banner()
	result.Summary = commandResult.Summary()
	result.Github = commandResult.Github()
	return result, nil
}
