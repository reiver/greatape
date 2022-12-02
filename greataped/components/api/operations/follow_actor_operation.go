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

type followActorOperation struct {
	Operation

	run func(IContext, *FollowActorRequest) (*FollowActorResult, error)
}

func FollowActorOperation() IOperation {
	return &followActorOperation{
		run: FollowActorService,
	}
}

func (operation *followActorOperation) Id() (ID, ID) {
	return FOLLOW_ACTOR_REQUEST, FOLLOW_ACTOR_RESULT
}

func (operation *followActorOperation) InputContainer() Pointer {
	return new(FollowActorRequest)
}

func (operation *followActorOperation) OutputContainer() Pointer {
	return new(FollowActorResult)
}

func (operation *followActorOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*FollowActorRequest))
}

/*
func (operation *followActorOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
