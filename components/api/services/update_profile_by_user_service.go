package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func UpdateProfileByUserService(context IContext, input *UpdateProfileByUserRequest) (result *UpdateProfileByUserResult, err error) {
	conductor := core.Conductor

	conductor.LogRemoteCall(context, INITIALIZE, "update_profile_by_user", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "update_profile_by_user", input, result, err) }()

	_result, _err := conductor.UpdateProfileByUser(input.DisplayName, input.Avatar, input.Banner, input.Summary, input.Github, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*UpdateProfileByUserResult)
	result.DisplayName = _result.DisplayName()
	result.Avatar = _result.Avatar()
	result.Banner = _result.Banner()
	result.Summary = _result.Summary()
	result.Github = _result.Github()
	return result, nil
}
