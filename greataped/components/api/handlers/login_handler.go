package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

type loginHandler struct {
}

func LoginHandler() IHttpHandler {
	return &loginHandler{}
}

func (handler *loginHandler) Method() string {
	return http.MethodPost
}

func (handler *loginHandler) Path() string {
	return "/api/v1/login"
}

func (handler *loginHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &LoginRequest{}
		result := &LoginResult{}

		onRequestUnmarshalled := func(request *LoginRequest) {
		}

		return pipeline.Handle(x,
			"login",
			LOGIN_REQUEST,
			LOGIN_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}