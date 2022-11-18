package operations

import (
	. "github.com/xeronith/diamante/contracts/operation"
	. "github.com/xeronith/diamante/contracts/service"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/operation"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/api/services"
	. "rail.town/infrastructure/components/contracts"
)

type logoutOperation struct {
	SecureOperation

	run func(IContext, *LogoutRequest) (*LogoutResult, error)
}

func LogoutOperation() IOperation {
	return &logoutOperation{
		run: LogoutService,
	}
}

func (operation *logoutOperation) Id() (ID, ID) {
	return LOGOUT_REQUEST, LOGOUT_RESULT
}

func (operation *logoutOperation) InputContainer() Pointer {
	return new(LogoutRequest)
}

func (operation *logoutOperation) OutputContainer() Pointer {
	return new(LogoutResult)
}

func (operation *logoutOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*LogoutRequest))
}

/*
func (operation *logoutOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
