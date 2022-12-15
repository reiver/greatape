package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

type getFollowingHandler struct {
}

func GetFollowingHandler() IHttpHandler {
	return &getFollowingHandler{}
}

func (handler *getFollowingHandler) Method() string {
	return http.MethodGet
}

func (handler *getFollowingHandler) Path() string {
	return "/u/:username/following"
}

func (handler *getFollowingHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &GetFollowingRequest{}
		result := &GetFollowingResult{}

		onRequestUnmarshalled := func(request *GetFollowingRequest) {
			request.Username = x.Param("username")
		}

		return pipeline.Handle(x,
			"get_following",
			GET_FOLLOWING_REQUEST,
			GET_FOLLOWING_RESULT,
			request, result,
			onRequestUnmarshalled,
			false,
		)
	}
}