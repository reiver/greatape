package commands

import (
	"github.com/reiver/greatape/app/activitypub"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func PostToInbox(x IDispatcher, username string, body []byte) (IPostToInboxResult, error) {
	identity := x.GetIdentityByUsername(username)
	object := x.UnmarshalActivityPubObjectOrLink(body)

	switch object.GetType() {
	case activitypub.TypeAccept:
		{
			activity := &activitypub.Activity{}
			x.UnmarshalJson(body, activity)

			switch activity.Object.(map[string]interface{})["type"] {
			case activitypub.TypeFollow:
				follow := &activitypub.Follow{}
				x.DecodeMapStructure(activity.Object, follow)

				x.Atomic(func() error {
					x.ForEachActivityPubFollower(func(record IActivityPubFollower) {
						if record.Handle() == follow.Actor && record.Subject() == follow.Object {
							record.UpdateAcceptedAtomic(x.Transaction(), true, x.Identity())
						}
					})

					x.AddActivityPubIncomingActivity(
						identity.Id(),
						x.GenerateUUID(),
						x.UnixNano(),
						follow.Object,
						follow.Actor,
						activitypub.TypeAccept,
						string(body),
					)

					return nil
				})

			default:
				return nil, ERROR_INVALID_PARAMETERS
			}
		}
	case activitypub.TypeFollow:
		{
			follow := &activitypub.Follow{}
			x.UnmarshalJson(body, follow)

			actorId := x.GetActorId(identity)

			actor := &activitypub.Actor{}
			if err := x.GetSignedActivityStream(follow.Actor, actor, identity); err != nil {
				return nil, err
			}

			follower := x.AddActivityPubFollower(
				follow.Actor,
				actor.Inbox,
				actorId,
				x.MarshalJson(follow),
				false,
			)

			activity := &activitypub.Activity{
				Context: activitypub.ActivityStreams,
				ID:      x.Format("%s/%s", x.PublicUrl(), x.GenerateUUID()),
				Type:    activitypub.TypeAccept,
				Actor:   actorId,
				Object:  follow,
			}

			if err := x.PostSignedActivityStream(actor.Inbox, activity, identity); err != nil {
				return nil, err
			}

			follower.UpdateAccepted(true, x.Identity())
		}
	case activitypub.TypeCreate:
		{
			activity := &activitypub.Activity{}
			x.UnmarshalJson(body, activity)

			switch activity.Object.(map[string]interface{})["type"] {
			case activitypub.TypeNote:
				note := &activitypub.Note{}
				x.DecodeMapStructure(activity.Object, note)

				x.AddActivityPubIncomingActivity(
					identity.Id(),
					x.GenerateUUID(),
					x.UnixNano(),
					note.AttributedTo,
					note.To[0],
					note.Content,
					string(body),
				)
			default:
				return nil, ERROR_INVALID_PARAMETERS
			}
		}
	default:
		return nil, ERROR_INVALID_PARAMETERS
	}

	return x.NewPostToInboxResult(body), nil
}
