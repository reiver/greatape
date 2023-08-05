package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func LoginService(context IContext, input *LoginRequest) (result *LoginResult, err error) {
	source := "login"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.Login(input.Email, input.Password, context.Identity())
	if err != nil {
		return nil, err
	}

	context.SetAuthCookie(commandResult.Token())

	result = context.ResultContainer().(*LoginResult)
	result.Username = commandResult.Username()
	result.Token = "Automatic"
	return result, nil
}
