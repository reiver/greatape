package handlers

import (
	"net/http"

	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/network/http"
	pipeline "github.com/xeronith/diamante/network/http"
)

type getPackagesHandler struct {
}

func GetPackagesHandler() IHttpHandler {
	return &getPackagesHandler{}
}

func (handler *getPackagesHandler) Method() string {
	return http.MethodGet
}

func (handler *getPackagesHandler) Path() string {
	return "/.well-known/packages.txt"
}

func (handler *getPackagesHandler) HandlerFunc() HttpHandlerFunc {
	return func(x IServerDispatcher) error {
		request := &GetPackagesRequest{}
		result := &GetPackagesResult{}

		onRequestUnmarshalled := func(request *GetPackagesRequest) {
		}

		onRequestProcessed := func(output *GetPackagesResult) (string, []byte) {
			return "text/plain", output.Body
		}

		return pipeline.Handle(x,
			"get_packages",
			GET_PACKAGES_REQUEST,
			GET_PACKAGES_RESULT,
			request, result,
			onRequestUnmarshalled,
			onRequestProcessed,
			false,
		)
	}
}
