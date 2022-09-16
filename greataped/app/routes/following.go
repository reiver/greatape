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
		webfinger := activitypub.Webfinger{}
		if err := x.GetActivityStream(username.Webfinger(), nil, &webfinger); err != nil {
			return x.InternalServerError(err)
		}

		actor := activitypub.Actor{}
		if err := x.GetActivityStream(webfinger.Self(), nil, &actor); err != nil {
			return x.InternalServerError(err)
		}

		following := activitypub.OrderedCollection{}
		if err := x.GetActivityStream(actor.Following, nil, &following); err != nil {
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
