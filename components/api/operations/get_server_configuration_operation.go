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

type getServerConfigurationOperation struct {
	Operation

	run func(IContext, *GetServerConfigurationRequest) (*GetServerConfigurationResult, error)
}

func GetServerConfigurationOperation() IOperation {
	return &getServerConfigurationOperation{
		run: GetServerConfigurationService,
	}
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
	return operation.run(context, payload.(*GetServerConfigurationRequest))
}
