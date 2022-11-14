package services

import (
	. "github.com/xeronith/diamante/contracts/service"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/core"
)

// noinspection GoUnusedParameter
func LoginService(context IContext, input *LoginRequest) (result *LoginResult, err error) {
	conductor := core.Conductor
	_ = LOGIN_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "login", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "login", input, result, err) }()

	_result, _err := conductor.Login(input.Email, input.Password, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	context.SetCookie("Diamante", _result.Token())
	result = context.ResultContainer().(*LoginResult)
	result.Username = _result.Username()
	result.Token = _result.Token()
	return result, nil
}
