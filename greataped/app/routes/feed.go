package routes

import (
	"activitypub"
	"config"
	"contracts"
	"db/repos"
	"server/route"
)

// Feed godoc
// @Tags     Feed
// @Accept   json
// @Produce  json
// @Security JWT
// @Param    username path     string true "Username"
// @Success  200      {object} map[string]interface{}
// @Router   /api/v1/users/{username}/feed [get]
func _() {}

var Feed = route.New(contracts.HttpGet, "/api/v1/users/:username/feed", func(x contracts.IContext) error {
	username := x.Request().Params("username")
	actor := x.StringUtil().Format("%s://%s/u/%s", config.PROTOCOL, config.DOMAIN, username)
	id := x.StringUtil().Format("%s://%s/u/%s/outbox", config.PROTOCOL, config.DOMAIN, username)

	messages, err := repos.Default.FindOutgoingActivitiesByUser(actor)
	if err != nil {
		return err
	}

	items := []*activitypub.Activity{}
	for _, message := range messages {
		note := activitypub.NewPublicNote(actor, message.Content)
		activity := note.Wrap(username)
		items = append(items, activity)
	}

	outbox := activitypub.NewOrderedCollection(id, items, len(items))
	return x.Activity(outbox)
})

// FeedTypes godoc
// @Tags    Feed
// @Accept  json
// @Produce json
// @Success 200 {object} []string
// @Router  /api/v1/feed/types [get]
func _() {}

var FeedTypes = route.New(contracts.HttpGet, "/api/v1/feed/types", func(x contracts.IContext) error {
	return x.Json([]string{
		"Most Recent",
		"Most Relevant",
	})
})
