package routes

import (
	"app/models/repos"
	"config"
	. "contracts"
	"fmt"
	"server/route"
)

var WebFinger = route.New(HttpGet, "/.well-known/webfinger", func(x IContext) error {
	resource := x.Request().Query("resource")
	if !x.StringUtil().Contains(resource, "acct:") {
		return x.BadRequest("Bad request. Please make sure 'acct:user@domain' is what you are sending as the 'resource' query parameter.")
	}

	username := x.StringUtil().Replace(resource, "acct:", "", -1)
	username = x.StringUtil().Replace(username, fmt.Sprintf("@%s", config.DOMAIN), "", -1)

	user, err := repos.FindUserByUsername(username)
	if err != nil {
		return err
	}

	webfinger := createWebfinger(user)
	return x.Json(webfinger)
})
