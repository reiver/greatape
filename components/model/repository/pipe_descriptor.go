package repository

import (
	"sync"

	. "github.com/reiver/greatape/components/contracts/model"
)

type pipeDescriptor struct {
	id                 int
	semaphore          *sync.Mutex
	query              string
	parametersResolver func(entity IPipeEntity) Parameters
}

func (command *pipeDescriptor) Id() int {
	return command.id
}

func (command *pipeDescriptor) GetSemaphore() *sync.Mutex {
	return command.semaphore
}

func (command *pipeDescriptor) GetQuery() string {
	return command.query
}

func (command *pipeDescriptor) GetParametersResolver() func(IPipeEntity) Parameters {
	return command.parametersResolver
}
