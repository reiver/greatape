package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type logoutHandler struct {
}

func LogoutHandler() IHttpHandler {
	return &logoutHandler{}
}

func (handler *logoutHandler) Method() string {
	return http.MethodPost
}

func (handler *logoutHandler) Path() string {
	return "/api/v1/logout"
}

func (handler *logoutHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &LogoutRequest{}
		result := &LogoutResult{}

		onRequestUnmarshalled := func(request *LogoutRequest) {
		}

		return pipeline.Handle(x,
			LOGOUT_REQUEST,
			LOGOUT_RESULT,
			request, result,
			onRequestUnmarshalled,
			nil,
			false,
		)
	}
}
