package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func CheckUsernameAvailabilityService(context IContext, input *CheckUsernameAvailabilityRequest) (result *CheckUsernameAvailabilityResult, err error) {
	source := "check_username_availability"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.CheckUsernameAvailability(input.Username, context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*CheckUsernameAvailabilityResult)
	result.IsAvailable = commandResult.IsAvailable()
	return result, nil
}
