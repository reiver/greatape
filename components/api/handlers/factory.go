package handlers

import . "github.com/xeronith/diamante/contracts/network/http"

type httpHandlerFactory struct{}

func (factory *httpHandlerFactory) Handlers() []IHttpHandler {
	return []IHttpHandler{
		EchoHandler(),                      // │ P . /api/v1/echo
		CheckUsernameAvailabilityHandler(), // │ P . /api/v1/check-username
		SignupHandler(),                    // │ P . /api/v1/signup
		VerifyHandler(),                    // │ P . /api/v1/verify
		LoginHandler(),                     // │ P . /api/v1/login
		GetProfileByUserHandler(),          // │ G . /api/v1/profile
		UpdateProfileByUserHandler(),       // │ P . /api/v1/profile
		LogoutHandler(),                    // │ P . /api/v1/logout
		WebfingerHandler(),                 // │ G . /.well-known/webfinger
		GetPackagesHandler(),               // │ G . /.well-known/packages.txt
		GetActorHandler(),                  // │ G . /users/:username
		FollowActorHandler(),               // │ G . /users/:username/follow
		AuthorizeInteractionHandler(),      // │ G . /authorize_interaction
		GetFollowersHandler(),              // │ G . /users/:username/followers
		GetFollowingHandler(),              // │ G . /users/:username/following
		PostToOutboxHandler(),              // │ P . /users/:username/outbox
		GetOutboxHandler(),                 // │ G . /users/:username/outbox
		PostToInboxHandler(),               // │ P . /users/:username/inbox
		GetInboxHandler(),                  // │ G . /users/:username/inbox
	}
}

func NewFactory() IHttpHandlerFactory {
	return &httpHandlerFactory{}
}
