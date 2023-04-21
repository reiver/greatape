package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func GetInboxService(context IContext, input *GetInboxRequest) (result *GetInboxResult, err error) {
	conductor := core.Conductor
	_ = GET_INBOX_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "get_inbox", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "get_inbox", input, result, err) }()

	_result, _err := conductor.GetInbox(input.Username, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	outputOrderedItems := make([]*ActivityPubActivity, 0)
	for _, orderedItem := range _result.OrderedItems() {
		if orderedItem == nil {
			continue
		}

		var object *ActivityPubObject
		if orderedItem.Object() != nil {
			object = &ActivityPubObject{
				Context:   orderedItem.Object().Context(),
				Id:        orderedItem.Object().Id(),
				Type:      orderedItem.Object().Type(),
				Actor:     orderedItem.Object().Actor(),
				From:      orderedItem.Object().From(),
				To:        orderedItem.Object().To(),
				InReplyTo: orderedItem.Object().InReplyTo(),
				Content:   orderedItem.Object().Content(),
				Published: orderedItem.Object().Published(),
			}
		}

		outputOrderedItems = append(outputOrderedItems, &ActivityPubActivity{
			Context:   orderedItem.Context(),
			Id:        orderedItem.Id(),
			Type:      orderedItem.Type(),
			Actor:     orderedItem.Actor(),
			Object:    object,
			From:      orderedItem.From(),
			To:        orderedItem.To(),
			InReplyTo: orderedItem.InReplyTo(),
			Content:   orderedItem.Content(),
			Published: orderedItem.Published(),
		})
	}

	result = context.ResultContainer().(*GetInboxResult)
	result.Context = _result.Context()
	result.Id = _result.Id()
	result.Type = _result.Type()
	result.TotalItems = _result.TotalItems()
	result.OrderedItems = outputOrderedItems
	result.First = _result.First()
	return result, nil
}
