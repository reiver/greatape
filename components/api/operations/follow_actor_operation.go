package operations

import (
	. "github.com/reiver/greatape/components/api/protobuf"
	. "github.com/reiver/greatape/components/api/services"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/operation"
	. "github.com/xeronith/diamante/contracts/service"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/operation"
)

type (
	FollowActorRunner  func(IContext, *FollowActorRequest) (*FollowActorResult, error)
	FollowActorRunners []FollowActorRunner

	followActorOperation struct {
		Operation

		runners FollowActorRunners
	}
)

func FollowActorOperation() IOperation {
	return &followActorOperation{
		runners: FollowActorRunners{
			FollowActorService,
		},
	}
}

func (operation *followActorOperation) Tag() string {
	return "FOLLOW_ACTOR"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*FollowActorRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
