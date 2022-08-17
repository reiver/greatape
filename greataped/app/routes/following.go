package routes

import (
	"activitypub"
	"app/models/repos"
	"app/models/types"
	"config"
	. "contracts"
	"server/route"
)

var Following = route.New(HttpGet, "/u/:username/following", func(x IContext) error {
	username := x.Request().Params("username")
	actor := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username)
	id := x.StringUtil().Format("%s://%s/u/%s/following", config.PROTOCOL, config.DOMAIN, username)

	followings := &[]types.FollowerResponse{}
	err := repos.FindFollowing(followings, actor).Error
	if err != nil {
		x.InternalServerError(err.Error())
	}

	items := []string{}
	for _, following := range *followings {
		items = append(items, following.Target)
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
