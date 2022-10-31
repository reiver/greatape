package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

type echoHandler struct {
}

func EchoHandler() IHttpHandler {
	return &echoHandler{}
}

func (handler *echoHandler) Method() string {
	return http.MethodPost
}

func (handler *echoHandler) Path() string {
	return "/api/v1/echo"
}

func (handler *echoHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &EchoRequest{}
		result := &EchoResult{}

		onRequestUnmarshalled := func(request *EchoRequest) {
		}

		return pipeline.Handle(x,
			"echo",
			ECHO_REQUEST,
			ECHO_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}
