package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type postToOutboxHandler struct {
}

func PostToOutboxHandler() IHttpHandler {
	return &postToOutboxHandler{}
}

func (handler *postToOutboxHandler) Method() string {
	return http.MethodPost
}

func (handler *postToOutboxHandler) Path() string {
	return "/u/:username/outbox"
}

func (handler *postToOutboxHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &PostToOutboxRequest{}
		result := &PostToOutboxResult{}

		onRequestUnmarshalled := func(request *PostToOutboxRequest) {
			request.Username = x.Param("username")
		}

		onRequestProcessed := func(output *PostToOutboxResult) (string, []byte) {
			return "application/activity+json; charset=utf-8", output.Body
		}

		return pipeline.Handle(x,
			"post_to_outbox",
			POST_TO_OUTBOX_REQUEST,
			POST_TO_OUTBOX_RESULT,
			request, result,
			onRequestUnmarshalled,
			onRequestProcessed,
			false,
		)
	}
}
