package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type followActorHandler struct {
}

func FollowActorHandler() IHttpHandler {
	return &followActorHandler{}
}

func (handler *followActorHandler) Method() string {
	return http.MethodGet
}

func (handler *followActorHandler) Path() string {
	return "/users/:username/follow"
}

func (handler *followActorHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &FollowActorRequest{}
		result := &FollowActorResult{}

		onRequestUnmarshalled := func(request *FollowActorRequest) {
			request.Username = x.Param("username")
			request.Account = x.Query("account")
		}

		return pipeline.Handle(x,
			FOLLOW_ACTOR_REQUEST,
			FOLLOW_ACTOR_RESULT,
			request, result,
			onRequestUnmarshalled,
			nil,
			false,
		)
	}
}
