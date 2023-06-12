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

type getActorOperation struct {
	Operation

	run func(IContext, *GetActorRequest) (*GetActorResult, error)
}

func GetActorOperation() IOperation {
	return &getActorOperation{
		run: GetActorService,
	}
}

func (operation *getActorOperation) Id() (ID, ID) {
	return GET_ACTOR_REQUEST, GET_ACTOR_RESULT
}

func (operation *getActorOperation) InputContainer() Pointer {
	return new(GetActorRequest)
}

func (operation *getActorOperation) OutputContainer() Pointer {
	return new(GetActorResult)
}

func (operation *getActorOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*GetActorRequest))
}
