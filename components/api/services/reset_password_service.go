package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func ResetPasswordService(context IContext, input *ResetPasswordRequest) (result *ResetPasswordResult, err error) {
	source := "reset_password"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	if _, err = Conductor.ResetPassword(input.UsernameOrEmail, context.Identity()); err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*ResetPasswordResult)
	return result, nil
}
