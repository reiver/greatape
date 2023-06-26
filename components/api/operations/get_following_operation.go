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
	GetFollowingRunner  func(IContext, *GetFollowingRequest) (*GetFollowingResult, error)
	GetFollowingRunners []GetFollowingRunner

	getFollowingOperation struct {
		Operation

		runners GetFollowingRunners
	}
)

func GetFollowingOperation() IOperation {
	return &getFollowingOperation{
		runners: GetFollowingRunners{
			GetFollowingService,
		},
	}
}

func (operation *getFollowingOperation) Tag() string {
	return "GET_FOLLOWING"
}

func (operation *getFollowingOperation) Id() (ID, ID) {
	return GET_FOLLOWING_REQUEST, GET_FOLLOWING_RESULT
}

func (operation *getFollowingOperation) InputContainer() Pointer {
	return new(GetFollowingRequest)
}

func (operation *getFollowingOperation) OutputContainer() Pointer {
	return new(GetFollowingResult)
}

func (operation *getFollowingOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*GetFollowingRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
