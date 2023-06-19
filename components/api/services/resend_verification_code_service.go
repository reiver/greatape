package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func ResendVerificationCodeService(context IContext, input *ResendVerificationCodeRequest) (result *ResendVerificationCodeResult, err error) {
	source := "resend_verification_code"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.ResendVerificationCode(input.Email, context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*ResendVerificationCodeResult)
	result.Code = commandResult.Code()
	return result, nil
}
