package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func LogoutService(context IContext, input *LogoutRequest) (result *LogoutResult, err error) {
	conductor := core.Conductor
	_ = LOGOUT_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "logout", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "logout", input, result, err) }()

	_result, _err := conductor.Logout(context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*LogoutResult)
	return result, nil
}
