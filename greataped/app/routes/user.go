package routes

import (
	"app/models/repos"
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
	} else {
		return x.Render("user", ViewData{
			"Title":    fmt.Sprintf("%s's Public Profile", user.DisplayName),
			"Username": user.Username,
			"Actor":    actor,
		})
	}
})
