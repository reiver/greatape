package server

import (
	"contracts"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type httpRequest struct {
	context *fiber.Ctx
}

func newRequest(context *fiber.Ctx) contracts.IRequest {
	return &httpRequest{
		context: context,
	}
}

func (request *httpRequest) Body(key string) string {
	// request.context.BodyParser()
	return string(request.context.Body())
}

func (request *httpRequest) Query(key string) string {
	return request.context.Query(key, "")
}

func (request *httpRequest) Header(key string) string {
	return string(request.context.Request().Header.Peek(key))
}

func (request *httpRequest) Params(key string) string {
	return request.context.Params(key, "")
}

func (request *httpRequest) FormValue(key string) string {
	return request.context.FormValue(key, "")
}

func (request *httpRequest) FormFile(key string) (*multipart.FileHeader, error) {
	return request.context.FormFile(key)
}
