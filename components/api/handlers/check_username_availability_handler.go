package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type checkUsernameAvailabilityHandler struct {
}

func CheckUsernameAvailabilityHandler() IHttpHandler {
	return &checkUsernameAvailabilityHandler{}
}

func (handler *checkUsernameAvailabilityHandler) Method() string {
	return http.MethodPost
}

func (handler *checkUsernameAvailabilityHandler) Path() string {
	return "/api/v1/check-username"
}

func (handler *checkUsernameAvailabilityHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &CheckUsernameAvailabilityRequest{}
		result := &CheckUsernameAvailabilityResult{}

		onRequestUnmarshalled := func(request *CheckUsernameAvailabilityRequest) {
		}

		return pipeline.Handle(x,
			CHECK_USERNAME_AVAILABILITY_REQUEST,
			CHECK_USERNAME_AVAILABILITY_RESULT,
			request, result,
			onRequestUnmarshalled,
			nil,
			false,
		)
	}
}
