package routes

import (
	. "contracts"
	"server/route"
)

var Authorize = route.New(HttpGet, "/authorize_interaction", func(x IContext) error {
	uri := x.Request().Query("uri")
	return x.JSON(struct {
		Uri     string
		Success bool `json:"success"`
	}{
		Uri:     uri,
		Success: true,
	})
})
