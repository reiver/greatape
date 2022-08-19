package routes

import (
	"activitypub"
	"app/models/repos"
	"app/models/types"
	"config"
	. "contracts"
	"server/route"
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

	json, _ := result.Marshal()
	x.Response().Header("Content-Type", "application/activity+json; charset=utf-8")
	return x.WriteString(string(json))
})

var AcceptFollowRequest = route.New(HttpPut, "/u/:username/followers/:id/accept", func(x IContext) error {
	id := x.Request().Params("id")

	err := repos.AcceptFollower(id).Error
	if err != nil {
		return x.InternalServerError(err.Error())
	}

	return x.Nothing()
})
