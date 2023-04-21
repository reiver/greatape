package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func WebfingerService(context IContext, input *WebfingerRequest) (result *WebfingerResult, err error) {
	conductor := core.Conductor
	_ = WEBFINGER_REQUEST

	conductor.LogRemoteCall(context, INITIALIZE, "webfinger", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "webfinger", input, result, err) }()

	_result, _err := conductor.Webfinger(input.Resource, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	outputLinks := make([]*ActivityPubLink, 0)
	for _, link := range _result.Links() {
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
	result.Aliases = _result.Aliases()
	result.Links = outputLinks
	result.Subject = _result.Subject()
	return result, nil
}
