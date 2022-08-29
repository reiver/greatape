package routes

import (
	"activitypub"
	"app/models/repos"
	"app/models/types"
	"config"
	. "contracts"
	"encoding/json"
	"errors"
	"server/mime"
	"server/route"
	"time"

	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

var InboxPost = route.New(HttpPost, "/u/:username/inbox", func(x IContext) error {
	username := x.Request().Params("username")

	object := &activitypub.Object{}
	if err := x.ParseBodyAndValidate(object); err != nil {
		return x.BadRequest("Bad request")
	}

	user := &repos.User{}
	err := repos.FindUserByUsername(user, username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return x.NotFound("No record found for %s.", username)
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
				if err := x.GetActivityStream(url, keyId, user.PrivateKey, nil, actor); err != nil {
					return x.InternalServerError(err.Error())
				}

				inbox = actor.Inbox
			}

			data, err := json.Marshal(activity)
			if err != nil {
				return x.InternalServerError(err.Error())
			}

			follower := &repos.Follower{
				Target:      x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username),
				Handle:      activity.Actor,
				HandleInbox: inbox,
				Activity:    string(data),
				Accepted:    false,
			}

			if err := repos.CreateFollower(follower); err.Error != nil {
				return x.Conflict(err.Error.Error())
			}

			if user.Access == repos.ACCESS_PUBLIC {
				data, _ := json.Marshal(&activitypub.Activity{
					Context: "https://www.w3.org/ns/activitystreams",
					ID:      x.StringUtil().Format("%s://%s/%s", config.PROTOCOL, config.DOMAIN, x.GUID()),
					Type:    activitypub.TypeAccept,
					Actor:   x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username),
					Object:  activity,
				})

				if err := x.PostActivityStream(inbox, keyId, user.PrivateKey, data, nil); err != nil {
					return x.InternalServerError(err.Error())
				}

				err := repos.AcceptFollower(follower.ID).Error
				if err != nil {
					return x.InternalServerError(err.Error())
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
					return x.Conflict(err.Error.Error())
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
		Context:      "https://www.w3.org/ns/activitystreams",
		ID:           id,
		Type:         "OrderedCollection",
		TotalItems:   len(items),
		OrderedItems: items,
	}

	json, _ := outbox.Marshal()
	x.Response().Header("Content-Type", mime.ActivityJsonUtf8)
	return x.WriteString(string(json))
})
