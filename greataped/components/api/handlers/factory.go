package handlers

import . "github.com/xeronith/diamante/contracts/network/http"

type httpHandlerFactory struct{}

func (factory *httpHandlerFactory) Handlers() []IHttpHandler {
	return []IHttpHandler{
		EchoHandler(),                 // │ P . /api/v1/echo
		SignupHandler(),               // │ P . /api/v1/signup
		VerifyHandler(),               // │ P . /api/v1/verify
		LoginHandler(),                // │ P . /api/v1/login
		GetProfileByUserHandler(),     // │ G . /api/v1/profile
		UpdateProfileByUserHandler(),  // │ P . /api/v1/profile
		LogoutHandler(),               // │ P . /api/v1/logout
		WebfingerHandler(),            // │ G . /.well-known/webfinger
		GetActorHandler(),             // │ G . /u/:username
		FollowActorHandler(),          // │ G . /u/:username/follow
		AuthorizeInteractionHandler(), // │ G . /authorize_interaction
		GetFollowersHandler(),         // │ G . /u/:username/followers
		GetFollowingHandler(),         // │ G . /u/:username/following
		PostToOutboxHandler(),         // │ P . /u/:username/outbox
		GetOutboxHandler(),            // │ G . /u/:username/outbox
	}
}

func NewFactory() IHttpHandlerFactory {
	return &httpHandlerFactory{}
}
