package commands

import . "github.com/reiver/greatape/components/contracts"

const packages = `
	google.golang.org/protobuf version=1.28.1 url=https://github.com/protocolbuffers/protobuf-go
	github.com/gorilla/websocket version=1.5.0 url=https://github.com/gorilla/websocket
	github.com/gorilla/securecookie version=1.1.1 url=https://github.com/gorilla/securecookie
	gopkg.in/yaml.v2 version=2.4.0 url=https://github.com/go-yaml/yaml
`

func GetPackages(x IDispatcher) (IGetPackagesResult, error) {
	return x.NewGetPackagesResult([]byte(packages)), nil
}
