package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func VerifyService(context IContext, input *VerifyRequest) (result *VerifyResult, err error) {
	conductor := core.Conductor
	_ = VERIFY_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "verify", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "verify", input, result, err) }()

	_result, _err := conductor.Verify(input.Email, input.Token, input.Code, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	context.SetCookie("Diamante", _result.Token())
	result = context.ResultContainer().(*VerifyResult)
	result.Token = _result.Token()
	return result, nil
}
