package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func GetPackagesService(context IContext, input *GetPackagesRequest) (result *GetPackagesResult, err error) {
	source := "get_packages"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.GetPackages(context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*GetPackagesResult)
	result.Body = commandResult.Body()
	return result, nil
}
