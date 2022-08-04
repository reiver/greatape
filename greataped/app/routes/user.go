package routes

import (
	"app/models/repos"
	. "contracts"
	"errors"
	"fmt"
	"server/route"
	"strings"

	"gorm.io/gorm"
)

var User = route.New(HttpGet, "/u/:name", func(x IContext) error {
	name := x.Request().Params("name")
	if name == "" {
		return x.BadRequest("Bad request")
	}

	user := &repos.User{}
	err := repos.FindUserByUsername(user, name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return x.NotFound("No record found for %s.", name)
	}

	actor := createActor(user)
	if strings.Contains(x.Request().Header("Accept"), "application/activity+json") {
		json, _ := actor.Marshal()
		x.Response().Header("Content-Type", "application/activity+json; charset=utf-8")
		return x.WriteString(string(json))
	} else {
		return x.Render("user", ViewData{
			"Actor": actor,
		})
	}
})

var _Followers = route.New(HttpPost, "/u/:name/:followers", func(x IContext) error {
	name := x.Request().Params("name")
	if name == "" {
		return x.BadRequest("Bad request")
	}

	storage := x.Storage()
	domain := x.Config().Get("domain")
	result := storage.Prepare("select followers from accounts where name = ?").Param(fmt.Sprintf("%s@%s", name, domain))
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
	`, domain, name, followers.Length())

	return x.JSON(followersCollection)
})
