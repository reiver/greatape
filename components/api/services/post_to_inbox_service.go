package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func PostToInboxService(context IContext, input *PostToInboxRequest) (result *PostToInboxResult, err error) {
	conductor := core.Conductor
	_ = POST_TO_INBOX_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "post_to_inbox", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "post_to_inbox", input, result, err) }()

	_result, _err := conductor.PostToInbox(input.Username, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	result = context.ResultContainer().(*PostToInboxResult)
	return result, nil
}
