package handlers

import (
	"net/http"

	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
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
	return "/u/:username/follow"
}

func (handler *followActorHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &FollowActorRequest{}
		result := &FollowActorResult{}

		onRequestUnmarshalled := func(request *FollowActorRequest) {
			request.Username = x.Param("username")
			request.Acct = x.Query("acct")
		}

		if err := pipeline.Handle(x,
			"follow_actor",
			FOLLOW_ACTOR_REQUEST,
			FOLLOW_ACTOR_RESULT,
			request, result,
			onRequestUnmarshalled,
			true,
		); err != nil {
			return err
		}

		x.Redirect(result.Url)
		return nil
	}
}
