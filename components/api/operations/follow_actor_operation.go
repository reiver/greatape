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
