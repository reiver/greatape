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

type verifyOperation struct {
	Operation

	run func(IContext, *VerifyRequest) (*VerifyResult, error)
}

func VerifyOperation() IOperation {
	return &verifyOperation{
		run: VerifyService,
	}
}

func (operation *verifyOperation) Id() (ID, ID) {
	return VERIFY_REQUEST, VERIFY_RESULT
}

func (operation *verifyOperation) InputContainer() Pointer {
	return new(VerifyRequest)
}

func (operation *verifyOperation) OutputContainer() Pointer {
	return new(VerifyResult)
}

func (operation *verifyOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*VerifyRequest))
}

/*
func (operation *verifyOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
