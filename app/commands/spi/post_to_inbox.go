package spi

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
	"github.com/reiver/greatape/app/activitypub"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func PostToInbox(x IDispatcher, username string, body []byte) (IPostToInboxResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Username() == username
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_FOUND)
	identity := identities.First()

	object := &activitypub.Object{}
	if err := json.Unmarshal(body, object); err != nil {
		return nil, ERROR_UNKNOWN_ACTIVITY_PUB_OBJECT
	}

	keyId := x.Format("%s/u/%s#main-key", x.PublicUrl(), username)

	switch object.Type {
	case activitypub.TypeFollow:
		{
			activity := &activitypub.Activity{}
			if err := json.Unmarshal(body, activity); err != nil {
				return nil, ERROR_UNKNOWN_ACTIVITY_PUB_ACTIVITY
			}

			url := activity.Actor
			var inbox string

			{
				actor := &activitypub.Actor{}
				if err := x.GetActivityStreamSigned(url, keyId, identity.PrivateKey(), nil, actor); err != nil {
					return nil, err
				}

				inbox = actor.Inbox
			}

			data, err := json.Marshal(activity)
			if err != nil {
				return nil, err
			}

			follower := x.AddActivityPubFollower(
				activity.Actor,
				inbox,
				x.Format("%s/u/%s", x.PublicUrl(), username),
				string(data),
				false,
			)

			data, _ = json.Marshal(&activitypub.Activity{
				Context: activitypub.ActivityStreams,
				ID:      x.Format("%s/%s", x.PublicUrl(), x.GenerateUUID()),
				Type:    activitypub.TypeAccept,
				Actor:   x.Format("%s/u/%s", x.PublicUrl(), username),
				Object:  activity,
			})

			if err := x.PostActivityStreamSigned(inbox, keyId, identity.PrivateKey(), data, nil); err != nil {
				return nil, err
			}

			follower.UpdateAccepted(true, x.Identity())
		}
	case activitypub.TypeCreate:
		{
			activity := &activitypub.Activity{}
			if err := json.Unmarshal(body, activity); err != nil {
				return nil, ERROR_UNKNOWN_ACTIVITY_PUB_ACTIVITY
			}

			switch activity.Object.(map[string]interface{})["type"] {
			case activitypub.TypeNote:
				note := &activitypub.Note{}
				if err := mapstructure.Decode(activity.Object, note); err != nil {
					return nil, ERROR_UNKNOWN_ACTIVITY_PUB_ACTIVITY
				}

				raw, _ := json.Marshal(note)

				x.AddActivityPubIncomingActivity(
					identity.Id(),
					x.GenerateUUID(),
					x.UnixNano(),
					note.AttributedTo,
					note.To[0],
					note.Content,
					string(raw),
				)
			default:
				return nil, ERROR_INVALID_PARAMETERS
			}
		}
	default:
		{
			return nil, ERROR_INVALID_PARAMETERS
		}
	}

	return x.NewPostToInboxResult(body), nil
}
