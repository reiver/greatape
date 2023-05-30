package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func GetProfileByUserService(context IContext, input *GetProfileByUserRequest) (result *GetProfileByUserResult, err error) {
	conductor := core.Conductor

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
