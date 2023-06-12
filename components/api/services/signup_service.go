package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func SignupService(context IContext, input *SignupRequest) (result *SignupResult, err error) {
	source := "signup"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.Signup(input.Username, input.Email, input.Password, context.Identity())
	if err != nil {
		return nil, err
	}

	context.SetCookie("Diamante", commandResult.Token())
	result = context.ResultContainer().(*SignupResult)
	result.Token = commandResult.Token()
	result.Code = commandResult.Code()
	return result, nil
}
