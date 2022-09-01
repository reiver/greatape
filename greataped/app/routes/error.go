package routes

import (
	"contracts"
	"server/route"
)

var Error = route.New(contracts.HttpGet, "/error/v1/:code/:message", func(x contracts.IContext) error {

	errorCode := x.Request().Params("code")
	errorMessage := x.Request().Params("message")

	switch errorCode {
	case "400":
		return x.BadRequest(errorMessage)
	case "401":
		return x.Unauthorized(errorMessage)
	case "500":
		return x.InternalServerError(errorMessage)
	}

	return x.NotFound(errorMessage)
})
