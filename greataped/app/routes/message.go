package routes

import (
	"app/models/repos"
	"app/models/types"
	. "contracts"
	"errors"
	"server/route"

	"gorm.io/gorm"
)

var Message = route.New(HttpGet, "/m/:guid", func(x IContext) error {
	guid := x.Request().Params("guid")
	if guid == "" {
		return x.BadRequest("Bad request.")
	}

	response := &types.MessageResponse{}
	err := repos.FindOutgoingActivityByGuid(response, guid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return x.NotFound("Message not found")
	}

	return x.String(response.Content)
})
