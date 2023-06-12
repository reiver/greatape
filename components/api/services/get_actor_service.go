package services

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/core"
	. "github.com/xeronith/diamante/contracts/service"
)

func GetActorService(context IContext, input *GetActorRequest) (result *GetActorResult, err error) {
	source := "get_actor"
	/* //////// */ Conductor.LogRemoteCall(context, INIT, source, input, result, err)
	defer func() { Conductor.LogRemoteCall(context, DONE, source, input, result, err) }()

	commandResult, err := Conductor.GetActor(input.Username, context.Identity())
	if err != nil {
		return nil, err
	}

	var outputIcon *ActivityPubMedia = nil
	if commandResult.Icon() != nil {
		outputIcon = &ActivityPubMedia{
			MediaType: commandResult.Icon().MediaType(),
			Type:      commandResult.Icon().Type(),
			Url:       commandResult.Icon().Url(),
			Width:     commandResult.Icon().Width(),
			Height:    commandResult.Icon().Height(),
		}
	}

	var outputImage *ActivityPubMedia = nil
	if commandResult.Image() != nil {
		outputImage = &ActivityPubMedia{
			MediaType: commandResult.Image().MediaType(),
			Type:      commandResult.Image().Type(),
			Url:       commandResult.Image().Url(),
			Width:     commandResult.Image().Width(),
			Height:    commandResult.Image().Height(),
		}
	}

	var outputPublicKey *ActivityPubPublicKey = nil
	if commandResult.PublicKey() != nil {
		outputPublicKey = &ActivityPubPublicKey{
			Id:           commandResult.PublicKey().Id(),
			Owner:        commandResult.PublicKey().Owner(),
			PublicKeyPem: commandResult.PublicKey().PublicKeyPem(),
		}
	}

	result = context.ResultContainer().(*GetActorResult)
	result.Context = commandResult.Context()
	result.Id = commandResult.Id()
	result.Followers = commandResult.Followers()
	result.Following = commandResult.Following()
	result.Inbox = commandResult.Inbox()
	result.Outbox = commandResult.Outbox()
	result.Name = commandResult.Name()
	result.PreferredUsername = commandResult.PreferredUsername()
	result.Type = commandResult.Type()
	result.Url = commandResult.Url()
	result.Icon = outputIcon
	result.Image = outputImage
	result.PublicKey = outputPublicKey
	result.Summary = commandResult.Summary()
	result.Published = commandResult.Published()
	return result, nil
}
