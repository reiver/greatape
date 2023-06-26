package core

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/client"
	. "github.com/xeronith/diamante/contracts/client"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/operation"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/protobuf"
	. "github.com/xeronith/diamante/serialization"
	. "github.com/xeronith/diamante/server"
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

	api.client.SetOperationResultListener(api.handler)
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
		api.logger.Debug(fmt.Sprintf("REQ { ID: %d, OP: %d }", requestId, operation))
	}

	if err := api.client.Send(requestId, operation, payload); err != nil {
		return nil, err
	}

	result := <-api.output
	if result.isError {
		return nil, errors.New(result.payload.(*ServerError).Message)
	}

	return result.payload, nil
}

func (api *api) handler(bundle IOperationResult) {
	isError := false
	var result Pointer

	_type, exists := API_RESULT[bundle.Type()]
	if exists {
		resultType := reflect.TypeOf(_type)
		result = reflect.New(resultType).Interface()
	} else {
		switch bundle.Type() {
		case ERROR:
			result = new(ServerError)
			isError = true
		default:
			api.logger.Fatal("unregistered_result_type")
		}
	}

	if err := NewProtobufSerializer().Deserialize(bundle.Payload(), result); err != nil {
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
