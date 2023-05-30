package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func PostToOutboxService(context IContext, input *PostToOutboxRequest) (result *PostToOutboxResult, err error) {
	conductor := core.Conductor

	conductor.LogRemoteCall(context, INITIALIZE, "post_to_outbox", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "post_to_outbox", input, result, err) }()

	_result, _err := conductor.PostToOutbox(input.Username, input.Body, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*PostToOutboxResult)
	result.Body = _result.Body()
	return result, nil
}
