package model

import (
	"sync"
	. "time"
)

// noinspection GoUnusedConst,GoSnakeCaseUsage
const (
	PIPE_DOCUMENT                       = 0x00000001
	PIPE_SYSTEM_SCHEDULE                = 0x00000002
	PIPE_IDENTITY                       = 0x00000003
	PIPE_ACCESS_CONTROL                 = 0x00000004
	PIPE_REMOTE_ACTIVITY                = 0x00000005
	PIPE_CATEGORY_TYPE                  = 0x00000006
	PIPE_CATEGORY                       = 0x00000007
	PIPE_USER                           = 0x00000008
	PIPE_ACTIVITY_PUB_INCOMING_ACTIVITY = 0x0000000E
	PIPE_ACTIVITY_PUB_OUTGOING_ACTIVITY = 0x0000000F
)

type (
	Parameters []interface{}

	IPipe interface {
		Input() chan IPipeEntity
		Signal() chan int
		OpenValve()
	}

	IPipeDescriptor interface {
		Id() int
		GetSemaphore() *sync.Mutex
		GetQuery() string
		GetParametersResolver() func(IPipeEntity) Parameters
	}

	IPipeEntity interface {
		GetPipe() int
		GetSource() string
		GetEditor() int64
		GetQueueTimestamp() Time
	}

	IPipeRepository interface {
		IRepository
		Insert(...IPipeEntity)
	}
)
