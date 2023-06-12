package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func WebfingerService(context IContext, input *WebfingerRequest) (result *WebfingerResult, err error) {
	source := "webfinger"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.Webfinger(input.Resource, context.Identity())
	if err != nil {
		return nil, err
	}

	outputLinks := make([]*ActivityPubLink, 0)
	for _, link := range commandResult.Links() {
		if link == nil {
			continue
		}

		outputLinks = append(outputLinks, &ActivityPubLink{
			Href:     link.Href(),
			Rel:      link.Rel(),
			Type:     link.Type(),
			Template: link.Template(),
		})
	}

	result = context.ResultContainer().(*WebfingerResult)
	result.Aliases = commandResult.Aliases()
	result.Links = outputLinks
	result.Subject = commandResult.Subject()
	return result, nil
}
