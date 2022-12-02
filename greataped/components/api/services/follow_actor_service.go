package services

import (
	. "github.com/xeronith/diamante/contracts/service"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
)

// noinspection GoUnusedParameter
func FollowActorService(context IContext, input *FollowActorRequest) (result *FollowActorResult, err error) {
	conductor := core.Conductor
	_ = FOLLOW_ACTOR_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "follow_actor", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "follow_actor", input, result, err) }()

	_result, _err := conductor.FollowActor(input.Username, input.Acct, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*FollowActorResult)
	result.Url = _result.Url()
	return result, nil
}
