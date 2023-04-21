package operations

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/api/services"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/operation"
	. "github.com/xeronith/diamante/contracts/service"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/operation"
)

type loginOperation struct {
	Operation

	run func(IContext, *LoginRequest) (*LoginResult, error)
}

func LoginOperation() IOperation {
	return &loginOperation{
		run: LoginService,
	}
}

func (operation *loginOperation) Id() (ID, ID) {
	return LOGIN_REQUEST, LOGIN_RESULT
}

func (operation *loginOperation) InputContainer() Pointer {
	return new(LoginRequest)
}

func (operation *loginOperation) OutputContainer() Pointer {
	return new(LoginResult)
}

func (operation *loginOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*LoginRequest))
}

/*
func (operation *loginOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
