package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func ResetPasswordService(context IContext, input *ResetPasswordRequest) (result *ResetPasswordResult, err error) {
	conductor := core.Conductor

	conductor.LogRemoteCall(context, INITIALIZE, "reset_password", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "reset_password", input, result, err) }()

	_result, _err := conductor.ResetPassword(input.UsernameOrEmail, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*ResetPasswordResult)
	return result, nil
}
