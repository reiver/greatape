package routes

import (
	"app/models/repos"
	"config"
	. "contracts"
	"errors"
	"fmt"
	"server/route"

	"gorm.io/gorm"
)

var WebFinger = route.New(HttpGet, "/.well-known/webfinger", func(x IContext) error {
	resource := x.Request().Query("resource")
	if !x.StringUtil().Contains(resource, "acct:") {
		return x.BadRequest("Bad request. Please make sure 'acct:user@domain' is what you are sending as the 'resource' query parameter.")
	}

	name := x.StringUtil().Replace(resource, "acct:", "", -1)
	name = x.StringUtil().Replace(name, fmt.Sprintf("@%s", config.DOMAIN), "", -1)

	user := &repos.User{}
	if err := repos.FindUserByUsername(user, name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return x.NotFound("No record found for %s.", name)
		} else {
			return x.InternalServerError(err.Error())
		}
	}

	webfinger := createWebfinger(user)
	return x.Json(webfinger)
})
