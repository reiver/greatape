package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
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
			"update_profile_by_user",
			UPDATE_PROFILE_BY_USER_REQUEST,
			UPDATE_PROFILE_BY_USER_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}
