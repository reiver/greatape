package routes

import (
	"app/models/domain"
	"config"
	"contracts"
	"db/repos"
	"server/route"
)

// User		godoc
// @Tags	ActivityPub
// @Accept	json
// @Produce	json
// @Param	username path string true "Username"
// @Success	200 {object} map[string]interface{}
// @Router	/u/{username} [get]
func _() {}

var User = route.New(contracts.HttpGet, "/u/:username", func(x contracts.IContext) error {
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

		return x.Activity(actor)
	} else {
		user, err := repos.Default.FindUserByUsername(username.String())
		if err != nil {
			return err
		}

		str := x.StringUtil()
		actor := createActor(user)
		if x.Request().AcceptsActivityStream() || x.Request().AcceptsJSON() {
			return x.Activity(actor)
		} else if config.ExternalClient() {
			return x.Redirect(str.Format("%s://%s/u/%s", config.PROTOCOL, config.CLIENT_DOMAIN, user.Username))
		} else {
			return x.Render("user", contracts.ViewData{
				"Title":    str.Format("%s's Public Profile", user.DisplayName),
				"Username": user.Username,
				"Actor":    actor,
			})
		}
	}
})
