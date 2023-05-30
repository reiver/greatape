package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

// noinspection GoUnusedParameter
func GetActorService(context IContext, input *GetActorRequest) (result *GetActorResult, err error) {
	conductor := core.Conductor

	conductor.LogRemoteCall(context, INITIALIZE, "get_actor", input, result, err)
	defer func() { conductor.LogRemoteCall(context, FINALIZE, "get_actor", input, result, err) }()

	_result, _err := conductor.GetActor(input.Username, context.Identity())
	if _err != nil {
		err = _err
		return nil, err
	}

	_ = _result

	var outputIcon *ActivityPubMedia = nil
	if _result.Icon() != nil {
		outputIcon = &ActivityPubMedia{
			MediaType: _result.Icon().MediaType(),
			Type:      _result.Icon().Type(),
			Url:       _result.Icon().Url(),
			Width:     _result.Icon().Width(),
			Height:    _result.Icon().Height(),
		}
	}

	var outputImage *ActivityPubMedia = nil
	if _result.Image() != nil {
		outputImage = &ActivityPubMedia{
			MediaType: _result.Image().MediaType(),
			Type:      _result.Image().Type(),
			Url:       _result.Image().Url(),
			Width:     _result.Image().Width(),
			Height:    _result.Image().Height(),
		}
	}

	var outputPublicKey *ActivityPubPublicKey = nil
	if _result.PublicKey() != nil {
		outputPublicKey = &ActivityPubPublicKey{
			Id:           _result.PublicKey().Id(),
			Owner:        _result.PublicKey().Owner(),
			PublicKeyPem: _result.PublicKey().PublicKeyPem(),
		}
	}

	result = context.ResultContainer().(*GetActorResult)
	result.Context = _result.Context()
	result.Id = _result.Id()
	result.Followers = _result.Followers()
	result.Following = _result.Following()
	result.Inbox = _result.Inbox()
	result.Outbox = _result.Outbox()
	result.Name = _result.Name()
	result.PreferredUsername = _result.PreferredUsername()
	result.Type = _result.Type()
	result.Url = _result.Url()
	result.Icon = outputIcon
	result.Image = outputImage
	result.PublicKey = outputPublicKey
	result.Summary = _result.Summary()
	result.Published = _result.Published()
	return result, nil
}
