package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type signupHandler struct {
}

func SignupHandler() IHttpHandler {
	return &signupHandler{}
}

func (handler *signupHandler) Method() string {
	return http.MethodPost
}

func (handler *signupHandler) Path() string {
	return "/api/v1/signup"
}

func (handler *signupHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &SignupRequest{}
		result := &SignupResult{}

		onRequestUnmarshalled := func(request *SignupRequest) {
		}

		return pipeline.Handle(x,
			SIGNUP_REQUEST,
			SIGNUP_RESULT,
			request, result,
			onRequestUnmarshalled,
			nil,
			false,
		)
	}
}
