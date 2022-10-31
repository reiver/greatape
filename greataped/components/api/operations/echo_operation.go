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

type echoOperation struct {
	Operation

	run func(IContext, *EchoRequest) (*EchoResult, error)
}

func EchoOperation() IOperation {
	return &echoOperation{
		run: EchoService,
	}
}

func (operation *echoOperation) Id() (ID, ID) {
	return ECHO_REQUEST, ECHO_RESULT
}

func (operation *echoOperation) InputContainer() Pointer {
	return new(EchoRequest)
}

func (operation *echoOperation) OutputContainer() Pointer {
	return new(EchoResult)
}

func (operation *echoOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*EchoRequest))
}

/*
func (operation *echoOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
