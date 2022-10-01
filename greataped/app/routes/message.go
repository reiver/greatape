package routes

import (
	. "contracts"
	"db/repos"
	"server/route"
)

var Message = route.New(HttpGet, "/m/:guid", func(x IContext) error {
	guid := x.Request().Params("guid")
	if guid == "" {
		return x.BadRequest("bad_request")
	}

	response, err := repos.Default.FindOutgoingActivityByGuid(guid)
	if err != nil {
		return err
	}

	return x.String(response.Content)
})
