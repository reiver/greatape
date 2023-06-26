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
	GetServerConfigurationRunner  func(IContext, *GetServerConfigurationRequest) (*GetServerConfigurationResult, error)
	GetServerConfigurationRunners []GetServerConfigurationRunner

	getServerConfigurationOperation struct {
		Operation

		runners GetServerConfigurationRunners
	}
)

func GetServerConfigurationOperation() IOperation {
	return &getServerConfigurationOperation{
		runners: GetServerConfigurationRunners{
			GetServerConfigurationService,
		},
	}
}

func (operation *getServerConfigurationOperation) Tag() string {
	return "GET_SERVER_CONFIGURATION"
}

func (operation *getServerConfigurationOperation) Id() (ID, ID) {
	return GET_SERVER_CONFIGURATION_REQUEST, GET_SERVER_CONFIGURATION_RESULT
}

func (operation *getServerConfigurationOperation) InputContainer() Pointer {
	return new(GetServerConfigurationRequest)
}

func (operation *getServerConfigurationOperation) OutputContainer() Pointer {
	return new(GetServerConfigurationResult)
}

func (operation *getServerConfigurationOperation) Execute(context IContext, payload Pointer) (Pointer, error) {
	if len(operation.runners) <= int(operation.ActiveRunner()) {
		return nil, ERROR_OPERATION_RUNNER_NOT_AVAILABLE
	}

	service := operation.runners[operation.ActiveRunner()]
	if input, valid := payload.(*GetServerConfigurationRequest); valid {
		return service(context, input)
	}

	return nil, ERROR_OPERATION_PAYLOAD_NOT_SUPPORTED
}
