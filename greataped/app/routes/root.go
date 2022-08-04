package routes

import (
	. "contracts"
	"server/route"
)

var Root = route.New(HttpGet, "/", func(x IContext) error {
	return x.Render("index", ViewData{
		"Title": "Great Ape",
	})
})
