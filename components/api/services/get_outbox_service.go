package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func GetOutboxService(context IContext, input *GetOutboxRequest) (result *GetOutboxResult, err error) {
	source := "get_outbox"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.GetOutbox(input.Username, context.Identity())
	if err != nil {
		return nil, err
	}

	outputOrderedItems := make([]*ActivityPubActivity, 0)
	for _, orderedItem := range commandResult.OrderedItems() {
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

	result = context.ResultContainer().(*GetOutboxResult)
	result.Context = commandResult.Context()
	result.Id = commandResult.Id()
	result.Type = commandResult.Type()
	result.TotalItems = commandResult.TotalItems()
	result.OrderedItems = outputOrderedItems
	result.First = commandResult.First()
	return result, nil
}
