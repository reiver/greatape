package routes

import (
	"app/models/domain"
	"app/models/repos"
	"config"
	"contracts"
	"errors"
	"server/route"

	"gorm.io/gorm"
)

var User = route.New(contracts.HttpGet, "/u/:username", func(x contracts.IContext) error {
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

		return x.Activity(actor)
	} else {
		user := &repos.User{}
		err := repos.FindUserByUsername(user, username.String()).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return x.NotFound("No record found for %s.", username)
		}

		str := x.StringUtil()
		actor := createActor(user)
		if x.Request().AcceptsActivityStream() {
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
