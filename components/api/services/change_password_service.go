package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func ChangePasswordService(context IContext, input *ChangePasswordRequest) (result *ChangePasswordResult, err error) {
	conductor := core.Conductor

	conductor.LogRemoteCall(context, INITIALIZE, "change_password", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "change_password", input, result, err) }()

	_result, _err := conductor.ChangePassword(input.CurrentPassword, input.NewPassword, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*ChangePasswordResult)
	return result, nil
}
