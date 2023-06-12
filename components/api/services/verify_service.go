package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func VerifyService(context IContext, input *VerifyRequest) (result *VerifyResult, err error) {
	source := "verify"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.Verify(input.Email, input.Token, input.Code, context.Identity())
	if err != nil {
		return nil, err
	}

	context.SetCookie("Diamante", commandResult.Token())
	result = context.ResultContainer().(*VerifyResult)
	result.Token = commandResult.Token()
	return result, nil
}
