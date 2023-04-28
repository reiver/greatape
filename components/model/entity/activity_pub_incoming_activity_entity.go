package entity

import (
	"fmt"
	"reflect"
	"time"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
)

var ActivityPubIncomingActivityEntityType = reflect.TypeOf(activityPubIncomingActivityEntity{})

// noinspection GoUnusedExportedFunction
func InitializeActivityPubIncomingActivityEntity() {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
}

type activityPubIncomingActivityEntity struct {
	entity
	IdentityIdField       int64  `json:"identity_id"`
	UniqueIdentifierField string `json:"unique_identifier" storage:"VARCHAR(128)" default:"''"`
	TimestampField        int64  `json:"timestamp" storage:"BIGINT" default:"0"`
	FromField             string `json:"from" storage:"VARCHAR(256)" default:"''"`
	ToField               string `json:"to" storage:"VARCHAR(256)" default:"''"`
	ContentField          string `json:"content" storage:"VARCHAR(4096)" default:"''"`
	RawField              string `json:"raw" storage:"JSONB" default:"'{}'"`
}

func NewActivityPubIncomingActivityEntity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubIncomingActivityEntity {
	return &activityPubIncomingActivityEntity{
		entity:                entity{IdField: id},
		IdentityIdField:       identityId,
		UniqueIdentifierField: uniqueIdentifier,
		TimestampField:        timestamp,
		FromField:             from,
		ToField:               to,
		ContentField:          content,
		RawField:              raw,
	}
}

type activityPubIncomingActivityPipeEntity struct {
	activityPubIncomingActivityEntity
	pipeEntity
}

func NewActivityPubIncomingActivityPipeEntity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor int64, payload string) IActivityPubIncomingActivityPipeEntity {
	return &activityPubIncomingActivityPipeEntity{
		activityPubIncomingActivityEntity: activityPubIncomingActivityEntity{
			entity:                entity{IdField: id, PayloadField: payload},
			IdentityIdField:       identityId,
			UniqueIdentifierField: uniqueIdentifier,
			TimestampField:        timestamp,
			FromField:             from,
			ToField:               to,
			ContentField:          content,
			RawField:              raw,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_ACTIVITY_PUB_INCOMING_ACTIVITY,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *activityPubIncomingActivityEntity) IdentityId() int64 {
	return entity.IdentityIdField
}

func (entity *activityPubIncomingActivityEntity) UniqueIdentifier() string {
	return entity.UniqueIdentifierField
}

func (entity *activityPubIncomingActivityEntity) Timestamp() int64 {
	return entity.TimestampField
}

func (entity *activityPubIncomingActivityEntity) From() string {
	return entity.FromField
}

func (entity *activityPubIncomingActivityEntity) To() string {
	return entity.ToField
}

func (entity *activityPubIncomingActivityEntity) Content() string {
	return entity.ContentField
}

func (entity *activityPubIncomingActivityEntity) Raw() string {
	return entity.RawField
}

func (entity *activityPubIncomingActivityEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *activityPubIncomingActivityEntity) String() string {
	return fmt.Sprintf("ActivityPubIncomingActivity (Id: %d, IdentityId: %d, UniqueIdentifier: %v, Timestamp: %v, From: %v, To: %v, Content: %v, Raw: %v)", entity.Id(), entity.IdentityId(), entity.UniqueIdentifier(), entity.Timestamp(), entity.From(), entity.To(), entity.Content(), entity.Raw())
}
