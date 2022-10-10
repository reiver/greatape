package routes

import (
	"config"
	. "contracts"
	"db/repos"
	"fmt"
	"server/route"
)

// WebFinger	godoc
// @Tags    WebFinger
// @Accept  json
// @Produce json
// @Param   resource query    string true "Resource" default(acct:user@domain.com)
// @Success 200      {object} map[string]interface{}
// @Router  /.well-known/webfinger [get]
func _() {}

var WebFinger = route.New(HttpGet, "/.well-known/webfinger", func(x IContext) error {
	resource := x.Request().Query("resource")
	if !x.StringUtil().Contains(resource, "acct:") {
		return x.BadRequest("Bad request. Please make sure 'acct:user@domain' is what you are sending as the 'resource' query parameter.")
	}

	username := x.StringUtil().Replace(resource, "acct:", "", -1)
	username = x.StringUtil().Replace(username, fmt.Sprintf("@%s", config.DOMAIN), "", -1)

	user, err := repos.Default.FindUserByUsername(username)
	if err != nil {
		return err
	}

	webfinger := createWebfinger(user)
	return x.Json(webfinger)
})
