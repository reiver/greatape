package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func ChangePasswordService(context IContext, input *ChangePasswordRequest) (result *ChangePasswordResult, err error) {
	source := "change_password"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	if _, err = Conductor.ChangePassword(input.CurrentPassword, input.NewPassword, context.Identity()); err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*ChangePasswordResult)
	return result, nil
}
