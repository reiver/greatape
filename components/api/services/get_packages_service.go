package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func GetPackagesService(context IContext, input *GetPackagesRequest) (result *GetPackagesResult, err error) {
	conductor := core.Conductor
	_ = GET_PACKAGES_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "get_packages", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "get_packages", input, result, err) }()

	_result, _err := conductor.GetPackages(context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*GetPackagesResult)
	result.Body = _result.Body()
	return result, nil
}
