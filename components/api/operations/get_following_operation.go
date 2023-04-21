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

type getFollowingOperation struct {
	Operation

	run func(IContext, *GetFollowingRequest) (*GetFollowingResult, error)
}

func GetFollowingOperation() IOperation {
	return &getFollowingOperation{
		run: GetFollowingService,
	}
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
	return operation.run(context, payload.(*GetFollowingRequest))
}

/*
func (operation *getFollowingOperation) ExecutionTimeLimits() (Duration, Duration, Duration) {
	var (
		TIME_LIMIT_WARNING  Duration = 20_000_000
		TIME_LIMIT_ALERT    Duration = 35_000_000
		TIME_LIMIT_CRITICAL Duration = 50_000_000
	)

	return TIME_LIMIT_WARNING, TIME_LIMIT_ALERT, TIME_LIMIT_CRITICAL
}
*/
