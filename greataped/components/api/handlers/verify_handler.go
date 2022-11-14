package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

type verifyHandler struct {
}

func VerifyHandler() IHttpHandler {
	return &verifyHandler{}
}

func (handler *verifyHandler) Method() string {
	return http.MethodPost
}

func (handler *verifyHandler) Path() string {
	return "/api/v1/verify"
}

func (handler *verifyHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &VerifyRequest{}
		result := &VerifyResult{}

		onRequestUnmarshalled := func(request *VerifyRequest) {
		}

		return pipeline.Handle(x,
			"verify",
			VERIFY_REQUEST,
			VERIFY_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}
