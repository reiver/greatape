package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type getFollowersHandler struct {
}

func GetFollowersHandler() IHttpHandler {
	return &getFollowersHandler{}
}

func (handler *getFollowersHandler) Method() string {
	return http.MethodGet
}

func (handler *getFollowersHandler) Path() string {
	return "/users/:username/followers"
}

func (handler *getFollowersHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &GetFollowersRequest{}
		result := &GetFollowersResult{}

		onRequestUnmarshalled := func(request *GetFollowersRequest) {
			request.Username = x.Param("username")
		}

		return pipeline.Handle(x,
			GET_FOLLOWERS_REQUEST,
			GET_FOLLOWERS_RESULT,
			request, result,
			onRequestUnmarshalled,
			nil,
			false,
		)
	}
}
