package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type updateProfileByUserHandler struct {
}

func UpdateProfileByUserHandler() IHttpHandler {
	return &updateProfileByUserHandler{}
}

func (handler *updateProfileByUserHandler) Method() string {
	return http.MethodPost
}

func (handler *updateProfileByUserHandler) Path() string {
	return "/api/v1/profile"
}

func (handler *updateProfileByUserHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &UpdateProfileByUserRequest{}
		result := &UpdateProfileByUserResult{}

		onRequestUnmarshalled := func(request *UpdateProfileByUserRequest) {
		}

		return pipeline.Handle(x,
			UPDATE_PROFILE_BY_USER_REQUEST,
			UPDATE_PROFILE_BY_USER_RESULT,
			request, result,
			onRequestUnmarshalled,
			nil,
			false,
		)
	}
}
