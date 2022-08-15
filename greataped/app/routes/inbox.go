package routes

import (
	"app/activitypub"
	"app/models/repos"
	"app/models/types"
	"config"
	. "contracts"
	"encoding/json"
	"errors"
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

	key := &types.KeyResponse{}
	err := repos.FindUserByUsername(key, username).Error
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
			follower := activity.Actor
			var inbox string

			{
				actor := &activitypub.Actor{}
				if err := x.GetActivityStream(url, keyId, key.PrivateKey, nil, actor); err != nil {
					return x.InternalServerError(err.Error())
				}

				inbox = actor.Inbox
			}

			{
				data, _ := json.Marshal(&activitypub.Activity{
					Context: "https://www.w3.org/ns/activitystreams",
					ID:      x.StringUtil().Format("%s://%s/%s", config.PROTOCOL, config.DOMAIN, x.GUID()),
					Type:    activitypub.TypeAccept,
					Actor:   x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username),
					Object:  activity,
				})

				if err := x.PostActivityStream(inbox, keyId, key.PrivateKey, data, nil); err != nil {
					return x.InternalServerError(err.Error())
				}

				if err := repos.CreateFollower(&repos.Follower{
					Target: x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username),
					Handle: follower,
				}); err.Error != nil {
					return x.Conflict(err.Error.Error())
				}
			}

			return x.Nothing()
		}
	case activitypub.TypeCreate:
		{
			activity := &activitypub.Activity{}
			if err := x.ParseBodyAndValidate(activity); err != nil {
				return x.BadRequest("bar_request")
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
	user := x.Request().Params("username")
	actor := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, user)

	messages := &[]types.MessageResponse{}
	err := repos.FindIncomingActivitiesForUser(messages, actor).Error
	if err != nil {
		x.InternalServerError("internal server error")
	}

	return x.JSON(messages)
})
