package routes

import (
	. "contracts"
	"server/route"
)

var Authorize = route.New(HttpGet, "/authorize_interaction", func(x IContext) error {
	uri := x.Request().Query("uri")
	return x.Json(struct {
		Uri     string
		Success bool `json:"success"`
	}{
		Uri:     uri,
		Success: true,
	})
})
