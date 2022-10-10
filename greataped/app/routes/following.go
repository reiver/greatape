package routes

import (
	"activitypub"
	"app/models/domain"
	"config"
	. "contracts"
	"db/repos"
	"server/route"
)

// Following	godoc
// @Tags    ActivityPub
// @Accept  json
// @Produce json
// @Param   username path     string true "Username"
// @Success 200      {object} map[string]interface{}
// @Router  /u/{username}/following [get]
func _() {}

var Following = route.New(HttpGet, "/u/:username/following", func(x IContext) error {
	username := domain.Username(x.Request().Params("username"))
	if username.IsEmpty() {
		return x.BadRequest("username required.")
	}

	if username.IsFederated() {
		webfinger, err := x.GetWebFinger(username)
		if err != nil {
			return err
		}

		actor, err := x.GetActor(webfinger)
		if err != nil {
			return err
		}

		following, err := x.GetOrderedCollection(actor.Following)
		if err != nil {
			return err
		}

		return x.Activity(following)
	} else {
		actor := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username)
		id := x.StringUtil().Format("%s://%s/u/%s/following", config.PROTOCOL, config.DOMAIN, username)

		followings, err := repos.Default.FindFollowing(actor)
		if err != nil {
			return err
		}

		items := []string{}
		for _, following := range followings {
			items = append(items, following.Target)
		}

		result := activitypub.NewOrderedCollection(id, items, len(items))
		return x.Activity(result)
	}
})
