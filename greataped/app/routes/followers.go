package routes

import (
	"activitypub"
	"app/models/repos"
	"app/models/types"
	"config"
	. "contracts"
	"encoding/json"
	"errors"
	"server/route"
	"strconv"

	"gorm.io/gorm"
)

var Followers = route.New(HttpGet, "/u/:username/followers", func(x IContext) error {
	username := x.Request().Params("username")
	actor := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username)
	id := x.StringUtil().Format("%s://%s/u/%s/followers", config.PROTOCOL, config.DOMAIN, username)

	followers := &[]types.FollowerResponse{}
	err := repos.FindFollowers(followers, actor).Error
	if err != nil {
		x.InternalServerError(err.Error())
	}

	items := []string{}
	for _, follower := range *followers {
		items = append(items, follower.Handle)
	}

	result := &activitypub.Followers{
		Context:      "https://www.w3.org/ns/activitystreams",
		ID:           id,
		Type:         "OrderedCollection",
		TotalItems:   len(items),
		OrderedItems: items,
	}

	return x.Activity(result)
})

var AcceptFollowRequest = route.New(HttpPut, "/u/:username/followers/:id/accept", func(x IContext) error {
	username := x.Request().Params("username")
	id := x.Request().Params("id")
	followerId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return x.BadRequest("invalid_id")
	}

	follower := &repos.Follower{}
	if err := repos.FindFollowerById(follower, followerId).Error; err != nil {
		return x.InternalServerError(err.Error())
	}

	data, _ := json.Marshal(&activitypub.Activity{
		Context: "https://www.w3.org/ns/activitystreams",
		ID:      x.StringUtil().Format("%s://%s/%s", config.PROTOCOL, config.DOMAIN, x.GUID()),
		Type:    activitypub.TypeAccept,
		Actor:   x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username),
		Object:  follower.Activity,
	})

	user := &repos.User{}
	err = repos.FindUserByUsername(user, username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return x.NotFound("No record found for %s.", username)
	}

	keyId := x.StringUtil().Format("%s://%s/u/%s#main-key", config.PROTOCOL, config.DOMAIN, username)

	if err := x.PostActivityStream(follower.HandleInbox, keyId, user.PrivateKey, data, nil); err != nil {
		return x.InternalServerError(err.Error())
	}

	if err := repos.AcceptFollower(follower.ID).Error; err != nil {
		return x.InternalServerError(err.Error())
	}

	return x.Nothing()
})
