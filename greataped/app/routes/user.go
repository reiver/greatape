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

var _ = route.New(HttpPost, "/u/:username/:followers", func(x IContext) error {
	username := x.Request().Params("username")
	if username == "" {
		return x.BadRequest("Bad request")
	}

	storage := x.Storage()
	domain := x.Config().Get("domain")
	result := storage.Prepare("select followers from accounts where name = ?").Param(fmt.Sprintf("%s@%s", username, domain))
	if result.Get("followers") == nil {
		result.Set("followers", "[]")
	}

	followers := x.ParseJson(result.Get("followers"))
	followersCollection := fmt.Sprintf(`
	{
		"type":"OrderedCollection",
		"totalItems":followers.length,
		"id": "https://%[1]s/u/%[2]s/followers",
		"first": {
		  "type":"OrderedCollectionPage",
		  "totalItems":%[3]d,
		  "partOf": "https://%[1]s/u/%[2]s/followers",
		  "orderedItems": followers,
		  "id": "https://%[1]s/u/%[2]s/followers?page=1"
		},
		"@context":["https://www.w3.org/ns/activitystreams"]
	}
	`, domain, username, followers.Length())

	return x.Json(followersCollection)
})
