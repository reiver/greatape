package core

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	. "github.com/xeronith/diamante/client"
	. "github.com/xeronith/diamante/contracts/client"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/operation"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/server"
	. "rail.town/infrastructure/components/api/protobuf"
	. "rail.town/infrastructure/components/contracts"
)

// noinspection GoSnakeCaseUsage
var API_RESULT = make(map[uint64]Pointer)

//--------------------------------------------------------------------------------------

type api struct {
	client    IClient
	output    chan *output
	logger    ILogger
	debugMode bool
}

func NewApi(endpoint string, logger ILogger) IApi {
	initialized := make(chan bool, 1)

	httpClient := NewHttpClient()
	httpClient.SetName("API-CLIENT")
	httpClient.SetVersion(1)
	httpClient.SetApiVersion(1)
	httpClient.OnConnectionEstablished(func(client IClient) {
		initialized <- true
	})

	_ = httpClient.Connect(endpoint, "AUTH-TOKEN")

	<-initialized
	api := &api{
		client: httpClient,
		output: make(chan *output, 1),
		logger: logger,
	}

	api.client.SetBinaryOperationResultListener(api.handler)
	return api
}

func (api *api) SetToken(token string) {
	api.client.SetToken(token)
}

func (api *api) SetDebugMode(enabled bool) {
	api.debugMode = enabled
}

func (api *api) call(operation uint64, payload Pointer) (Pointer, error) {
	requestId := uint64(time.Now().UnixNano())
	if api.debugMode {
		api.logger.Debug(fmt.Sprintf("REQ { ID: %d, OP: %s }", requestId, OPCODES[operation]))
	}

	if err := api.client.Send(requestId, operation, payload); err != nil {
		return nil, err
	}

	result := <-api.output
	if result.isError {
		return nil, errors.New(result.payload.(*Error).Message)
	}

	return result.payload, nil
}

func (api *api) handler(bundle IBinaryOperationResult) {
	isError := false
	var result Pointer

	_type, exists := API_RESULT[bundle.Type()]
	if exists {
		resultType := reflect.TypeOf(_type)
		result = reflect.New(resultType).Interface()
	} else {
		switch bundle.Type() {
		case ERROR:
			result = new(Error)
			isError = true
		default:
			api.logger.Fatal("unregistered_result_type")
		}
	}

	if err := DefaultBinarySerializer.Deserialize(bundle.Payload(), result); err != nil {
		api.logger.Fatal(err)
	}

	api.output <- &output{payload: result, isError: isError}
}

//--------------------------------------------------------------------------------------

type output struct {
	payload Pointer
	isError bool
}

//--------------------------------------------------------------------------------------
