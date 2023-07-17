package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type getActorHandler struct {
}

func GetActorHandler() IHttpHandler {
	return &getActorHandler{}
}

func (handler *getActorHandler) Method() string {
	return http.MethodGet
}

func (handler *getActorHandler) Path() string {
	return "/users/:username"
}

func (handler *getActorHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &GetActorRequest{}
		result := &GetActorResult{}

		onRequestUnmarshalled := func(request *GetActorRequest) {
			request.Username = x.Param("username")
		}

		return pipeline.Handle(x,
			GET_ACTOR_REQUEST,
			GET_ACTOR_RESULT,
			request, result,
			onRequestUnmarshalled,
			nil,
			false,
		)
	}
}
