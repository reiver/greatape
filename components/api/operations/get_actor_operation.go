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
	GetActorRunner  func(IContext, *GetActorRequest) (*GetActorResult, error)
	GetActorRunners []GetActorRunner

	getActorOperation struct {
		Operation

		runners GetActorRunners
	}
)

func GetActorOperation() IOperation {
	return &getActorOperation{
		runners: GetActorRunners{
			GetActorService,
		},
	}
}

func (operation *getActorOperation) Tag() string {
	return "GET_ACTOR"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*GetActorRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
