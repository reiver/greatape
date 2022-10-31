package contracts

import . "rail.town/infrastructure/components/api/protobuf"

type IApi interface {
	SetToken(string)
	SetDebugMode(bool)
	//API Methods
	SystemCall(*SystemCallRequest) (*SystemCallResult, error)
	Echo(*EchoRequest) (*EchoResult, error)
	ResolveError(*ResolveErrorRequest) (*ResolveErrorResult, error)
}