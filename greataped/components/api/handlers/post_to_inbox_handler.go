package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

type postToInboxHandler struct {
}

func PostToInboxHandler() IHttpHandler {
	return &postToInboxHandler{}
}

func (handler *postToInboxHandler) Method() string {
	return http.MethodPost
}

func (handler *postToInboxHandler) Path() string {
	return "/u/:username/inbox"
}

func (handler *postToInboxHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &PostToInboxRequest{}
		result := &PostToInboxResult{}

		onRequestUnmarshalled := func(request *PostToInboxRequest) {
			request.Username = x.Param("username")
		}

		return pipeline.Handle(x,
			"post_to_inbox",
			POST_TO_INBOX_REQUEST,
			POST_TO_INBOX_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}
