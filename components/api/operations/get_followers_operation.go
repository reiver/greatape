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

type getFollowersOperation struct {
	Operation

	run func(IContext, *GetFollowersRequest) (*GetFollowersResult, error)
}

func GetFollowersOperation() IOperation {
	return &getFollowersOperation{
		run: GetFollowersService,
	}
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
	return operation.run(context, payload.(*GetFollowersRequest))
}
