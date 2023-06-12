package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func LogoutService(context IContext, input *LogoutRequest) (result *LogoutResult, err error) {
	source := "logout"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	if _, err = Conductor.Logout(context.Identity()); err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*LogoutResult)
	return result, nil
}
