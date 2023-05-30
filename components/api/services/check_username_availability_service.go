package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func CheckUsernameAvailabilityService(context IContext, input *CheckUsernameAvailabilityRequest) (result *CheckUsernameAvailabilityResult, err error) {
	conductor := core.Conductor

	conductor.LogRemoteCall(context, INITIALIZE, "check_username_availability", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "check_username_availability", input, result, err) }()

	_result, _err := conductor.CheckUsernameAvailability(input.Username, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*CheckUsernameAvailabilityResult)
	result.IsAvailable = _result.IsAvailable()
	return result, nil
}
