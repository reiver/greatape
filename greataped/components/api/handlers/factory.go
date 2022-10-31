package handlers

import . "github.com/xeronith/diamante/contracts/network/http"

type httpHandlerFactory struct{}

func (factory *httpHandlerFactory) Handlers() []IHttpHandler {
	return []IHttpHandler{
		EchoHandler(), // │ P . /api/v1/echo
	}
}

func NewFactory() IHttpHandlerFactory {
	return &httpHandlerFactory{}
}
