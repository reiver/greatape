package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
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
	return "/users/:username/inbox"
}

func (handler *getInboxHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &GetInboxRequest{}
		result := &GetInboxResult{}

		onRequestUnmarshalled := func(request *GetInboxRequest) {
			request.Username = x.Param("username")
		}

		return pipeline.Handle(x,
			GET_INBOX_REQUEST,
			GET_INBOX_RESULT,
			request, result,
			onRequestUnmarshalled,
			nil,
			false,
		)
	}
}
