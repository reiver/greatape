package services

import (
	. "github.com/xeronith/diamante/contracts/service"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
)

// noinspection GoUnusedParameter
func GetFollowingService(context IContext, input *GetFollowingRequest) (result *GetFollowingResult, err error) {
	conductor := core.Conductor
	_ = GET_FOLLOWING_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "get_following", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "get_following", input, result, err) }()

	_result, _err := conductor.GetFollowing(input.Username, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*GetFollowingResult)
	result.Context = _result.Context()
	result.Id = _result.Id()
	result.Type = _result.Type()
	result.TotalItems = _result.TotalItems()
	result.OrderedItems = _result.OrderedItems()
	result.First = _result.First()
	return result, nil
}
