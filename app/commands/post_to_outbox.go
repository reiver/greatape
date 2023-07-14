package commands

import (
	"fmt"
	"time"

	"github.com/reiver/greatape/app/activitypub"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func PostToOutbox(x IDispatcher, username string, body []byte) (IPostToOutboxResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Username() == username
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_FOUND)
	identity := identities.First()

	item := x.UnmarshalActivityPubObjectOrLink(body)

	actorId := x.GetActorId()

	switch item.GetType() {
	case activitypub.TypeNote:
		{
			note, err := activitypub.UnmarshalNote(body)
			if err != nil {
				return nil, ERROR_INVALID_PARAMETERS
			}

			content := note.Content
			to := note.To[0]
			from := note.AttributedTo

			if from != actorId {
				return nil, ERROR_INVALID_PARAMETERS
			}

			uniqueIdentifier := x.GenerateUUID()

			activity := &activitypub.Activity{
				Context:   activitypub.ActivityStreams,
				Type:      activitypub.TypeCreate,
				ID:        fmt.Sprintf("%s/u/%s/posts/%s", x.PublicUrl(), username, uniqueIdentifier),
				To:        note.To,
				Actor:     fmt.Sprintf("%s/u/%s", x.PublicUrl(), username),
				Published: time.Now(),
				Object:    note,
			}

			if to != ACTIVITY_PUB_PUBLIC {
				recipient := &activitypub.Actor{}
				if err := x.GetActivityStreamSigned(to, nil, recipient); err != nil {
					return nil, err
				}

				to = recipient.ID

				if err := x.PostActivityStreamSigned(recipient.Inbox, activity, nil); err != nil {
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
