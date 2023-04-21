package entity

import (
	"fmt"
	"reflect"
	"time"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
)

var RemoteActivityEntityType = reflect.TypeOf(remoteActivityEntity{})

// noinspection GoUnusedExportedFunction
func InitializeRemoteActivityEntity() {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
}

type remoteActivityEntity struct {
	entity
	EntryPointField    string `json:"entry_point" previous:"id" storage:"VARCHAR(256)" default:"''"`
	DurationField      int64  `json:"duration" previous:"entry_point" storage:"BIGINT" default:"0"`
	SuccessfulField    bool   `json:"successful" previous:"duration" storage:"BIT(1)" default:"FALSE"`
	ErrorMessageField  string `json:"error_message" previous:"successful" storage:"VARCHAR(1024)" default:"''"`
	RemoteAddressField string `json:"remote_address" previous:"error_message" storage:"VARCHAR(128)" default:"''"`
	UserAgentField     string `json:"user_agent" previous:"remote_address" storage:"VARCHAR(512)" default:"''"`
	EventTypeField     uint32 `json:"event_type" previous:"user_agent" storage:"INT UNSIGNED" default:"0"`
	TimestampField     int64  `json:"timestamp" previous:"event_type" storage:"BIGINT" default:"0"`
}

func NewRemoteActivityEntity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) IRemoteActivityEntity {
	return &remoteActivityEntity{
		entity:             entity{IdField: id},
		EntryPointField:    entryPoint,
		DurationField:      duration,
		SuccessfulField:    successful,
		ErrorMessageField:  errorMessage,
		RemoteAddressField: remoteAddress,
		UserAgentField:     userAgent,
		EventTypeField:     eventType,
		TimestampField:     timestamp,
	}
}

type remoteActivityPipeEntity struct {
	remoteActivityEntity
	pipeEntity
}

func NewRemoteActivityPipeEntity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, source string, editor int64, payload string) IRemoteActivityPipeEntity {
	return &remoteActivityPipeEntity{
		remoteActivityEntity: remoteActivityEntity{
			entity:             entity{IdField: id, PayloadField: payload},
			EntryPointField:    entryPoint,
			DurationField:      duration,
			SuccessfulField:    successful,
			ErrorMessageField:  errorMessage,
			RemoteAddressField: remoteAddress,
			UserAgentField:     userAgent,
			EventTypeField:     eventType,
			TimestampField:     timestamp,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_REMOTE_ACTIVITY,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *remoteActivityEntity) EntryPoint() string {
	return entity.EntryPointField
}

func (entity *remoteActivityEntity) Duration() int64 {
	return entity.DurationField
}

func (entity *remoteActivityEntity) Successful() bool {
	return entity.SuccessfulField
}

func (entity *remoteActivityEntity) ErrorMessage() string {
	return entity.ErrorMessageField
}

func (entity *remoteActivityEntity) RemoteAddress() string {
	return entity.RemoteAddressField
}

func (entity *remoteActivityEntity) UserAgent() string {
	return entity.UserAgentField
}

func (entity *remoteActivityEntity) EventType() uint32 {
	return entity.EventTypeField
}

func (entity *remoteActivityEntity) Timestamp() int64 {
	return entity.TimestampField
}

func (entity *remoteActivityEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *remoteActivityEntity) String() string {
	return fmt.Sprintf("RemoteActivity (Id: %d, EntryPoint: %v, Duration: %v, Successful: %v, ErrorMessage: %v, RemoteAddress: %v, UserAgent: %v, EventType: %v, Timestamp: %v)", entity.Id(), entity.EntryPoint(), entity.Duration(), entity.Successful(), entity.ErrorMessage(), entity.RemoteAddress(), entity.UserAgent(), entity.EventType(), entity.Timestamp())
}
