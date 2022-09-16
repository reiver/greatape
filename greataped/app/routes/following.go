package routes

import (
	"activitypub"
	"app/models/domain"
	"app/models/repos"
	"app/models/types"
	"config"
	. "contracts"
	"server/route"
)

var Following = route.New(HttpGet, "/u/:username/following", func(x IContext) error {
	username := domain.Username(x.Request().Params("username"))
	if username.IsEmpty() {
		return x.BadRequest("username required.")
	}

	if username.IsFederated() {
		webfinger, err := x.GetWebFinger(username)
		if err != nil {
			return x.InternalServerError(err)
		}

		actor, err := x.GetActor(webfinger)
		if err != nil {
			return x.InternalServerError(err)
		}

		following, err := x.GetOrderedCollection(actor.Following)
		if err != nil {
			return x.InternalServerError(err)
		}

		return x.Activity(following)
	} else {
		actor := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username)
		id := x.StringUtil().Format("%s://%s/u/%s/following", config.PROTOCOL, config.DOMAIN, username)

		followings := &[]types.FollowerResponse{}
		err := repos.FindFollowing(followings, actor).Error
		if err != nil {
			x.InternalServerError(err)
		}

		items := []string{}
		for _, following := range *followings {
			items = append(items, following.Target)
		}

		result := activitypub.NewOrderedCollection(id, items, len(items))
		return x.Activity(result)
	}
})
