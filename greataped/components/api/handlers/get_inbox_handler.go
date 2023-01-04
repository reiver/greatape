package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

type getInboxHandler struct {
}

func GetInboxHandler() IHttpHandler {
	return &getInboxHandler{}
}

func (handler *getInboxHandler) Method() string {
	return http.MethodGet
}

func (handler *getInboxHandler) Path() string {
	return "/u/:username/inbox"
}

func (handler *getInboxHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &GetInboxRequest{}
		result := &GetInboxResult{}

		onRequestUnmarshalled := func(request *GetInboxRequest) {
			request.Username = x.Param("username")
		}

		return pipeline.Handle(x,
			"get_inbox",
			GET_INBOX_REQUEST,
			GET_INBOX_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}
