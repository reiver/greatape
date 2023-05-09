package commands

import (
	"encoding/json"
	"fmt"
	"time"

	ap "github.com/go-ap/activitypub"
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

	id := x.Format("%s/u/%s", x.PublicUrl(), identity.Username())

	publicKeyId := x.Format("%s#main-key", id)
	privateKey := identity.PrivateKey()

	switch item.GetType() {
	case ap.NoteType:
		{
			note := x.UnmarshalActivityPubNote(body)

			content := note.Content.First().Value.String()
			to := note.To.First().GetID().String()
			from := note.AttributedTo.GetID().String()

			if from != id {
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

			if to != activitypub.Public {
				recipient := &activitypub.Actor{}
				if err := x.GetActivityStreamSigned(to, publicKeyId, privateKey, nil, recipient); err != nil {
					return nil, err
				}

				to = recipient.ID

				data, _ := json.Marshal(activity)
				output := &struct{}{}
				if err := x.PostActivityStreamSigned(recipient.Inbox, publicKeyId, privateKey, data, output); err != nil {
					return nil, err
				}
			}

			raw, _ := json.Marshal(note)

			x.LogActivityPubOutgoingActivity(
				identity.Id(),
				uniqueIdentifier,
				x.UnixNano(),
				from,
				to,
				content,
				string(raw),
				"PostToOutbox",
				EMPTY_JSON,
			)

			return x.NewPostToOutboxResult(nil), nil
		}
	default:
		return nil, ERROR_INVALID_PARAMETERS
	}
}
