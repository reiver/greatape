package core

import (
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

func (api *api) SystemCall(request *SystemCallRequest) (*SystemCallResult, error) {
	result, err := api.call(SYSTEM_CALL_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*SystemCallResult), nil
	}
}

func (api *api) Echo(request *EchoRequest) (*EchoResult, error) {
	result, err := api.call(ECHO_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*EchoResult), nil
	}
}

func (api *api) ResolveError(request *ResolveErrorRequest) (*ResolveErrorResult, error) {
	result, err := api.call(RESOLVE_ERROR_REQUEST, request)

	if err != nil {
		return nil, err
	} else {
		return result.(*ResolveErrorResult), nil
	}
}

func init() {
	API_RESULT[SYSTEM_CALL_RESULT] = SystemCallResult{}
	API_RESULT[ECHO_RESULT] = EchoResult{}
	API_RESULT[RESOLVE_ERROR_RESULT] = ResolveErrorResult{}
}
