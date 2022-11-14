package contracts

import . "rail.town/infrastructure/components/api/protobuf"

type IApi interface {
	SetToken(string)
	SetDebugMode(bool)
	//API Methods
	SystemCall(*SystemCallRequest) (*SystemCallResult, error)
	Echo(*EchoRequest) (*EchoResult, error)
	Signup(*SignupRequest) (*SignupResult, error)
	Verify(*VerifyRequest) (*VerifyResult, error)
	Login(*LoginRequest) (*LoginResult, error)
}
