package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func GetServerConfigurationService(context IContext, input *GetServerConfigurationRequest) (result *GetServerConfigurationResult, err error) {
	source := "get_server_configuration"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.GetServerConfiguration(context.Identity())
	if err != nil {
		return nil, err
	}

	result = context.ResultContainer().(*GetServerConfigurationResult)
	result.Product = commandResult.Product()
	result.Environment = commandResult.Environment()
	result.Fqdn = commandResult.Fqdn()
	return result, nil
}
