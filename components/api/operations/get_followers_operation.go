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

/*
func (operation *getFollowersOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
