package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func SignupService(context IContext, input *SignupRequest) (result *SignupResult, err error) {
	conductor := core.Conductor
	_ = SIGNUP_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "signup", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "signup", input, result, err) }()

	_result, _err := conductor.Signup(input.Username, input.Email, input.Password, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	context.SetCookie("Diamante", _result.Token())
	result = context.ResultContainer().(*SignupResult)
	result.Token = _result.Token()
	result.Code = _result.Code()
	return result, nil
}
