package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/xeronith/diamante/logging"
	. "rail.town/infrastructure/components/contracts/model"
)

// noinspection GoSnakeCaseUsage
const (
	BUFFER_LENGTH              = int64(1000) // Should be set in regard to `max_allowed_packet` size in MySQL
	AUTO_FLUSH_DURATION        = time.Millisecond * 100
	DEFAULT_INPUT_CHANNEL_SIZE = 10000

	SIGNAL_FLUSH = 0
)

type pipe struct {
	repository         IRepository
	descriptorId       int
	input              chan IPipeEntity
	signal             chan int
	semaphore          *sync.Mutex
	query              string
	parametersResolver func(IPipeEntity) Parameters
	cursor             int64
	buffer             [BUFFER_LENGTH]IPipeEntity
	lastFlush          time.Time
}

func NewPipe(descriptor IPipeDescriptor, repository IRepository) IPipe {
	return &pipe{
		repository:         repository,
		descriptorId:       descriptor.Id(),
		input:              make(chan IPipeEntity, DEFAULT_INPUT_CHANNEL_SIZE),
		signal:             make(chan int),
		semaphore:          descriptor.GetSemaphore(),
		query:              descriptor.GetQuery(),
		parametersResolver: descriptor.GetParametersResolver(),
		lastFlush:          time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC),
	}
}

func (pipe *pipe) Input() chan IPipeEntity {
	return pipe.input
}

func (pipe *pipe) Signal() chan int {
	return pipe.signal
}

func (pipe *pipe) flush(cursor int64) int64 {
	if pipe.cursor == 0 {
		return 0
	}

	if cursor < BUFFER_LENGTH && time.Since(pipe.lastFlush) < AUTO_FLUSH_DURATION {
		return cursor
	}

	pipe.semaphore.Lock()
	defer pipe.semaphore.Unlock()

	parameters := make(Parameters, 0)
	for index := int64(0); index < cursor; index++ {
		parameters = append(parameters, pipe.parametersResolver(pipe.buffer[index])...)
	}

	if err := pipe.repository.GetSqlDatabase().InsertAll(pipe.query, cursor, parameters...); err != nil {
		pipe.repository.Serialize(pipe.buffer[:cursor], err)
	}

	pipe.lastFlush = time.Now()

	return 0
}

func (pipe *pipe) OpenValve() {
	defer pipe.catch()

	for {
		select {
		case _entity := <-pipe.Input():
			pipe.buffer[pipe.cursor] = _entity
			pipe.cursor++

			if pipe.cursor == BUFFER_LENGTH {
				pipe.cursor = pipe.flush(pipe.cursor)
			}
		case <-pipe.Signal():
			pipe.cursor = pipe.flush(pipe.cursor)
		}
	}
}

func (pipe *pipe) catch() {
	if reason := recover(); reason != nil {
		logging.GetDefaultLogger().Panic(fmt.Sprintf("PIPE_REPOSITORY: PIPE %d, STORE %s / %s", pipe.descriptorId, pipe.repository.GetSqlDatabase().GetName(), reason))
	}
}
