package route

import "contracts"

type httpRoute struct {
	method  contracts.HttpMethod
	path    string
	handler contracts.IHandler
}

func New(method contracts.HttpMethod, path string, handler contracts.IHandler) contracts.IRoute {
	return &httpRoute{
		method:  method,
		path:    path,
		handler: handler,
	}
}

func (route *httpRoute) Method() contracts.HttpMethod {
	return route.method
}

func (route *httpRoute) Path() string {
	return route.path
}

func (route *httpRoute) Handler() contracts.IHandler {
	return route.handler
}
