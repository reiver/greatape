package routes

import (
	"app/models/repos"
	. "contracts"
	"server/route"
)

var Message = route.New(HttpGet, "/m/:guid", func(x IContext) error {
	guid := x.Request().Params("guid")
	if guid == "" {
		return x.BadRequest("bad_request")
	}

	response, err := repos.FindOutgoingActivityByGuid(guid)
	if err != nil {
		return err
	}

	return x.String(response.Content)
})
