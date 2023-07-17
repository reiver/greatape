package commands

import (
	"time"

	ap "github.com/go-ap/activitypub"
	"github.com/reiver/greatape/app/activitypub"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func PostToOutbox(x IDispatcher, username string, body []byte) (IPostToOutboxResult, error) {
	identity := x.GetIdentityByUsername(username)
	object := x.UnmarshalActivityPubObjectOrLink(body)

	switch object.GetType() {
	case activitypub.TypeNote:
		{
			note := object.(*ap.Note)

			actorId := x.GetActorId(identity)
			content := note.Content.First().String()
			to := note.To[0].GetID().String()
			from := note.AttributedTo.GetID().String()

			if from != actorId {
				return nil, ERROR_INVALID_PARAMETERS
			}

			uniqueIdentifier := x.GenerateUUID()

			activity := &activitypub.Activity{
				Context:   activitypub.ActivityStreams,
				Type:      activitypub.TypeCreate,
				ID:        x.Format("%s/posts/%s", actorId, uniqueIdentifier),
				To:        note.To,
				Actor:     actorId,
				Published: time.Now(),
				Object:    note,
			}

			if to != ACTIVITY_PUB_PUBLIC {
				recipient := &activitypub.Actor{}
				if err := x.GetSignedActivityStream(to, recipient, identity); err != nil {
					return nil, err
				}

				to = recipient.Id

				if err := x.PostSignedActivityStream(recipient.Inbox, activity, identity); err != nil {
					return nil, err
				}
			}

			x.LogActivityPubOutgoingActivity(
				identity.Id(),
				uniqueIdentifier,
				x.UnixNano(),
				from,
				to,
				content,
				string(body),
				"PostToOutbox",
				EMPTY_JSON,
			)

			return x.NewPostToOutboxResult(nil), nil
		}
	default:
		return nil, ERROR_INVALID_PARAMETERS
	}
}
