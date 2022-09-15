package routes

import (
	"activitypub"
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

	if x.StringUtil().Contains(username, "@") {
		parts := x.StringUtil().Split(username, "@")
		url := x.StringUtil().Format("https://%s/.well-known/webfinger?resource=acct:%s", parts[1], username)

		webfinger := activitypub.Webfinger{}
		if err := x.GetActivityStream(url, "", "", nil, &webfinger); err != nil {
			return x.InternalServerError(err.Error())
		}

		actor := activitypub.Actor{}
		if err := x.GetActivityStream(webfinger.Aliases[0], "activitystream", "", nil, &actor); err != nil {
			return x.InternalServerError(err.Error())
		}

		return x.Activity(actor)
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
