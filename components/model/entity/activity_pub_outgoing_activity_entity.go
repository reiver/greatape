package entity

import (
	"fmt"
	"reflect"
	"time"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
)

var ActivityPubOutgoingActivityEntityType = reflect.TypeOf(activityPubOutgoingActivityEntity{})

type activityPubOutgoingActivityEntity struct {
	entity
	IdentityIdField       int64  `json:"identity_id"`
	UniqueIdentifierField string `json:"unique_identifier" storage:"VARCHAR(128)" default:"''"`
	TimestampField        int64  `json:"timestamp" storage:"BIGINT" default:"0"`
	FromField             string `json:"from" storage:"VARCHAR(256)" default:"''"`
	ToField               string `json:"to" storage:"VARCHAR(256)" default:"''"`
	ContentField          string `json:"content" storage:"VARCHAR(4096)" default:"''"`
	RawField              string `json:"raw" storage:"JSONB" default:"'{}'"`
}

func NewActivityPubOutgoingActivityEntity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubOutgoingActivityEntity {
	return &activityPubOutgoingActivityEntity{
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

type activityPubOutgoingActivityPipeEntity struct {
	activityPubOutgoingActivityEntity
	pipeEntity
}

func NewActivityPubOutgoingActivityPipeEntity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor int64, payload string) IActivityPubOutgoingActivityPipeEntity {
	return &activityPubOutgoingActivityPipeEntity{
		activityPubOutgoingActivityEntity: activityPubOutgoingActivityEntity{
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
			Pipe:           PIPE_ACTIVITY_PUB_OUTGOING_ACTIVITY,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *activityPubOutgoingActivityEntity) IdentityId() int64 {
	return entity.IdentityIdField
}

func (entity *activityPubOutgoingActivityEntity) UniqueIdentifier() string {
	return entity.UniqueIdentifierField
}

func (entity *activityPubOutgoingActivityEntity) Timestamp() int64 {
	return entity.TimestampField
}

func (entity *activityPubOutgoingActivityEntity) From() string {
	return entity.FromField
}

func (entity *activityPubOutgoingActivityEntity) To() string {
	return entity.ToField
}

func (entity *activityPubOutgoingActivityEntity) Content() string {
	return entity.ContentField
}

func (entity *activityPubOutgoingActivityEntity) Raw() string {
	return entity.RawField
}

func (entity *activityPubOutgoingActivityEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *activityPubOutgoingActivityEntity) String() string {
	return fmt.Sprintf("ActivityPubOutgoingActivity (Id: %d, IdentityId: %d, UniqueIdentifier: %v, Timestamp: %v, From: %v, To: %v, Content: %v, Raw: %v)", entity.Id(), entity.IdentityId(), entity.UniqueIdentifier(), entity.Timestamp(), entity.From(), entity.To(), entity.Content(), entity.Raw())
}
