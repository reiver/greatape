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
	GetPackagesRunner  func(IContext, *GetPackagesRequest) (*GetPackagesResult, error)
	GetPackagesRunners []GetPackagesRunner

	getPackagesOperation struct {
		Operation

		runners GetPackagesRunners
	}
)

func GetPackagesOperation() IOperation {
	return &getPackagesOperation{
		runners: GetPackagesRunners{
			GetPackagesService,
		},
	}
}

func (operation *getPackagesOperation) Tag() string {
	return "GET_PACKAGES"
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
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*GetPackagesRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
