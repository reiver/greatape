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

type getPackagesOperation struct {
	Operation

	run func(IContext, *GetPackagesRequest) (*GetPackagesResult, error)
}

func GetPackagesOperation() IOperation {
	return &getPackagesOperation{
		run: GetPackagesService,
	}
}

func (operation *getPackagesOperation) Id() (ID, ID) {
	return GET_PACKAGES_REQUEST, GET_PACKAGES_RESULT
}

func (operation *getPackagesOperation) InputContainer() Pointer {
	return new(GetPackagesRequest)
}

func (operation *getPackagesOperation) OutputContainer() Pointer {
	return new(GetPackagesResult)
}

func (operation *getPackagesOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	return operation.run(context, payload.(*GetPackagesRequest))
}
