package routes

import (
	. "contracts"
	"server/route"
)

var Followers = route.New(HttpGet, "/u/:name/followers", func(x IContext) error {
	return x.JSON(struct{}{})
})
