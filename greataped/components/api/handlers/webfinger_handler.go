package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

type webfingerHandler struct {
}

func WebfingerHandler() IHttpHandler {
	return &webfingerHandler{}
}

func (handler *webfingerHandler) Method() string {
	return http.MethodGet
}

func (handler *webfingerHandler) Path() string {
	return "/.well-known/webfinger"
}

func (handler *webfingerHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &WebfingerRequest{}
		result := &WebfingerResult{}

		onRequestUnmarshalled := func(request *WebfingerRequest) {
			request.Resource = x.Query("resource")
		}

		return pipeline.Handle(x,
			"webfinger",
			WEBFINGER_REQUEST,
			WEBFINGER_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}
