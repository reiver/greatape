package routes

import (
	"app/models/repos"
	"config"
	. "contracts"
	"errors"
	"fmt"
	"server/mime"
	"server/route"
	"strings"

	"gorm.io/gorm"
)

var User = route.New(HttpGet, "/u/:username", func(x IContext) error {
	username := x.Request().Params("username")
	if username == "" {
		return x.BadRequest("Bad request")
	}

	user := &repos.User{}
	err := repos.FindUserByUsername(user, username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return x.NotFound("No record found for %s.", username)
	}

	actor := createActor(user)
	if strings.Contains(x.Request().Header("Accept"), mime.ActivityJson) {
		return x.Activity(actor)
	} else if config.DOMAIN == config.CLIENT_DOMAIN {
		return x.Render("user", ViewData{
			"Title":    fmt.Sprintf("%s's Public Profile", user.DisplayName),
			"Username": user.Username,
			"Actor":    actor,
		})
	} else {
		client := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.CLIENT_DOMAIN, user.Username)
		return x.Redirect(client)
	}
})
