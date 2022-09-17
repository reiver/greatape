package routes

import (
	"activitypub"
	"app/models/repos"
	"app/models/types"
	"config"
	. "contracts"
	"encoding/json"
	"server/route"
	"time"

	"github.com/mitchellh/mapstructure"
)

var InboxPost = route.New(HttpPost, "/u/:username/inbox", func(x IContext) error {
	username := x.Request().Params("username")

	object := &activitypub.Object{}
	if err := x.ParseBodyAndValidate(object); err != nil {
		return x.BadRequest("Bad request")
	}

	user, err := repos.FindUserByUsername(username)
	if err != nil {
		return err
	}

	keyId := x.StringUtil().Format("%s://%s/u/%s#main-key", config.PROTOCOL, config.DOMAIN, username)

	switch object.Type {
	case activitypub.TypeFollow:
		{
			activity := &activitypub.Activity{}
			if err := x.ParseBodyAndValidate(activity); err != nil {
				return x.BadRequest("Bad request")
			}

			url := activity.Actor
			var inbox string

			{
				actor := &activitypub.Actor{}
				if err := x.GetActivityStreamSigned(url, keyId, user.PrivateKey, nil, actor); err != nil {
					return err
				}

				inbox = actor.Inbox
			}

			data, err := json.Marshal(activity)
			if err != nil {
				return err
			}

			follower := &repos.Follower{
				Target:      x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username),
				Handle:      activity.Actor,
				HandleInbox: inbox,
				Activity:    string(data),
				Accepted:    false,
			}

			if err := repos.CreateFollower(follower); err.Error != nil {
				return x.Conflict(err.Error)
			}

			if user.Access == repos.ACCESS_PUBLIC {
				data, _ := json.Marshal(&activitypub.Activity{
					Context: activitypub.ActivityStreams,
					ID:      x.StringUtil().Format("%s://%s/%s", config.PROTOCOL, config.DOMAIN, x.GUID()),
					Type:    activitypub.TypeAccept,
					Actor:   x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username),
					Object:  activity,
				})

				if err := x.PostActivityStreamSigned(inbox, keyId, user.PrivateKey, data, nil); err != nil {
					return err
				}

				err := repos.AcceptFollower(follower.ID).Error
				if err != nil {
					return err
				}
			}

			return x.Nothing()
		}
	case activitypub.TypeCreate:
		{
			activity := &activitypub.Activity{}
			if err := x.ParseBodyAndValidate(activity); err != nil {
				return x.BadRequest("bad_request")
			}

			switch activity.Object.(map[string]interface{})["type"] {
			case activitypub.TypeNote:
				note := &activitypub.Note{}
				if err := mapstructure.Decode(activity.Object, note); err != nil {
					return x.InternalServerError("decode_failed")
				}

				message := &repos.IncomingActivity{
					Timestamp: time.Now().UnixNano(),
					From:      note.AttributedTo,
					To:        note.To[0],
					Guid:      x.GUID(),
					Content:   note.Content,
				}

				if err := repos.CreateIncomingActivity(message); err.Error != nil {
					return x.Conflict(err.Error)
				}

				return x.Nothing()
			default:
				return x.BadRequest("")
			}
		}
	default:
		{
			return x.BadRequest("")
		}
	}
})

var InboxGet = route.New(HttpGet, "/u/:username/inbox", func(x IContext) error {
	username := x.Request().Params("username")
	actor := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username)
	id := x.StringUtil().Format("%s://%s/u/%s/inbox", config.PROTOCOL, config.DOMAIN, username)

	messages := &[]types.MessageResponse{}
	err := repos.FindIncomingActivitiesForUser(messages, actor).Error
	if err != nil {
		x.InternalServerError("internal_server_error")
	}

	items := []*activitypub.Activity{}
	for _, message := range *messages {
		note := activitypub.NewPublicNote(message.From, message.Content)
		activity := note.Wrap(username)
		items = append(items, activity)
	}

	outbox := &activitypub.Outbox{
		Context:      activitypub.ActivityStreams,
		ID:           id,
		Type:         activitypub.TypeOrderedCollection,
		TotalItems:   len(items),
		OrderedItems: items,
	}

	return x.Activity(outbox)
})
