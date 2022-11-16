package services

import (
	. "github.com/xeronith/diamante/contracts/service"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
)

// noinspection GoUnusedParameter
func GetProfileByUserService(context IContext, input *GetProfileByUserRequest) (result *GetProfileByUserResult, err error) {
	conductor := core.Conductor
	_ = GET_PROFILE_BY_USER_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "get_profile_by_user", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "get_profile_by_user", input, result, err) }()

	_result, _err := conductor.GetProfileByUser(context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*GetProfileByUserResult)
	result.Username = _result.Username()
	result.DisplayName = _result.DisplayName()
	result.Avatar = _result.Avatar()
	result.Banner = _result.Banner()
	result.Summary = _result.Summary()
	result.Github = _result.Github()
	return result, nil
}
