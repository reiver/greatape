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

	"gorm.io/gorm"
)

var OutboxPost = route.New(HttpPost, "/u/:username/outbox", func(x IContext) error {
	username := x.Request().Params("username")

	object := &activitypub.Object{}
	if err := x.ParseBodyAndValidate(object); err != nil {
		return x.BadRequest(err.Error())
	}

	key := &types.KeyResponse{}
	err := repos.FindUserByUsername(key, username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return x.NotFound("No record found for %s.", username)
	}

	keyId := x.StringUtil().Format("%s://%s/u/%s#main-key", config.PROTOCOL, config.DOMAIN, username)

	switch object.Type {
	case activitypub.TypeNote:
		{
			note := &activitypub.Note{}
			if err := x.ParseBodyAndValidate(note); err != nil {
				return x.BadRequest(err.Error())
			}

			activity := note.Wrap(username)

			recipient := &activitypub.Actor{}
			if err := x.GetActivityStream(activity.To.([]string)[0], keyId, key.PrivateKey, nil, recipient); err != nil {
				return x.InternalServerError(err.Error())
			}

			data, _ := json.Marshal(activity)
			output := &struct{}{}
			if err := x.PostActivityStream(recipient.Inbox, keyId, key.PrivateKey, data, output); err != nil {
				return x.InternalServerError(err.Error())
			}

			message := &repos.OutgoingActivity{
				Timestamp: time.Now().UnixNano(),
				From:      note.AttributedTo,
				To:        recipient.ID,
				Guid:      x.GUID(),
				Content:   note.Content,
			}

			if err := repos.CreateOutgoingActivity(message); err.Error != nil {
				return x.Conflict(err.Error.Error())
			}

			return x.Nothing()
		}
	default:
		return x.BadRequest("")
	}
})

var OutboxGet = route.New(HttpGet, "/u/:username/outbox", func(x IContext) error {
	user := x.Request().Params("username")
	actor := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, user)

	messages := &[]types.MessageResponse{}
	err := repos.FindOutgoingActivitiesByUser(messages, actor).Error
	if err != nil {
		x.InternalServerError("internal server error")
	}

	return x.JSON(messages)
})
