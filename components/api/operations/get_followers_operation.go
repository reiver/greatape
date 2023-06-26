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
	GetFollowersRunner  func(IContext, *GetFollowersRequest) (*GetFollowersResult, error)
	GetFollowersRunners []GetFollowersRunner

	getFollowersOperation struct {
		Operation

		runners GetFollowersRunners
	}
)

func GetFollowersOperation() IOperation {
	return &getFollowersOperation{
		runners: GetFollowersRunners{
			GetFollowersService,
		},
	}
}

func (operation *getFollowersOperation) Tag() string {
	return "GET_FOLLOWERS"
}

func (operation *getFollowersOperation) Id() (ID, ID) {
	return GET_FOLLOWERS_REQUEST, GET_FOLLOWERS_RESULT
}

func (operation *getFollowersOperation) InputContainer() Pointer {
	return new(GetFollowersRequest)
}

func (operation *getFollowersOperation) OutputContainer() Pointer {
	return new(GetFollowersResult)
}

func (operation *getFollowersOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*GetFollowersRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
