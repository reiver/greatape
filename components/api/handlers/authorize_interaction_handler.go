package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type authorizeInteractionHandler struct {
}

func AuthorizeInteractionHandler() IHttpHandler {
	return &authorizeInteractionHandler{}
}

func (handler *authorizeInteractionHandler) Method() string {
	return http.MethodGet
}

func (handler *authorizeInteractionHandler) Path() string {
	return "/authorize_interaction"
}

func (handler *authorizeInteractionHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &AuthorizeInteractionRequest{}
		result := &AuthorizeInteractionResult{}

		onRequestUnmarshalled := func(request *AuthorizeInteractionRequest) {
			request.Uri = x.Query("uri")
		}

		return pipeline.Handle(x,
			"authorize_interaction",
			AUTHORIZE_INTERACTION_REQUEST,
			AUTHORIZE_INTERACTION_RESULT,
			request, result,
			onRequestUnmarshalled,
			nil,
			false,
		)
	}
}
