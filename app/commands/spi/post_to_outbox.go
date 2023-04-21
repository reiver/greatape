package spi

import (
	"encoding/json"

	"github.com/reiver/greatape/app/activitypub"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
)

func PostToOutbox(x IDispatcher,
	username string,
	context string,
	activityType string,
	to string,
	attributedTo string,
	inReplyTo string,
	content string,
) (IPostToOutboxResult, error) {
	identities := x.FilterIdentities(func(identity IIdentity) bool {
		return identity.Username() == username
	})

	x.Assert(identities.HasExactlyOneItem()).Or(ERROR_USER_NOT_FOUND)
	identity := identities.First()

	id := x.Format("%s/u/%s", x.PublicUrl(), identity.Username())
	publicKeyId := x.Format("%s#main-key", id)
	privateKey := identity.PrivateKey()

	_ = publicKeyId

	switch activityType {
	case ACTIVITY_PUB_NOTE:
		{
			uniqueIdentifier := x.GenerateUUID()
			note := activitypub.NewNote(id, to, content)
			activity := note.Wrap(identity.Username(), x.PublicUrl(), uniqueIdentifier)
			to := activity.To.([]string)[0]

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
				note.AttributedTo,
				to,
				note.Content,
				string(raw),
				"PostToOutbox",
				EMPTY_JSON,
			)

			return x.NewPostToOutboxResult(), nil
		}
	default:
		return nil, ERROR_INVALID_PARAMETERS
	}
}
