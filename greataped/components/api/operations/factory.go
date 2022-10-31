package operations

import . "github.com/xeronith/diamante/contracts/operation"

type operationFactory struct{}

func (factory *operationFactory) Operations() []IOperation {
	return []IOperation{
		SystemCallOperation(),
		EchoOperation(),
		ResolveErrorOperation(),
	}
}

func NewFactory() IOperationFactory {
	return &operationFactory{}
}