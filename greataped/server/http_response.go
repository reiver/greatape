package server

import (
	"contracts"

	"github.com/gofiber/fiber/v2"
)

type httpResponse struct {
	context *fiber.Ctx
}

func newResponse(context *fiber.Ctx) contracts.IResponse {
	return &httpResponse{
		context: context,
	}
}

func (response *httpResponse) Header(key, value string) {
	response.context.Response().Header.Add(key, value)
}
