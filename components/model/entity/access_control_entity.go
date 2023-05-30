package entity

import (
	"fmt"
	"reflect"
	"time"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
)

var AccessControlEntityType = reflect.TypeOf(accessControlEntity{})

type accessControlEntity struct {
	entity
	KeyField   uint64 `json:"key" storage:"BIGINT" default:"0"`
	ValueField uint64 `json:"value" storage:"BIGINT" default:"0"`
}

func NewAccessControlEntity(id int64, key uint64, value uint64) IAccessControlEntity {
	return &accessControlEntity{
		entity:     entity{IdField: id},
		KeyField:   key,
		ValueField: value,
	}
}

type accessControlPipeEntity struct {
	accessControlEntity
	pipeEntity
}

func NewAccessControlPipeEntity(id int64, key uint64, value uint64, source string, editor int64, payload string) IAccessControlPipeEntity {
	return &accessControlPipeEntity{
		accessControlEntity: accessControlEntity{
			entity:     entity{IdField: id, PayloadField: payload},
			KeyField:   key,
			ValueField: value,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_ACCESS_CONTROL,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *accessControlEntity) Key() uint64 {
	return entity.KeyField
}

func (entity *accessControlEntity) Value() uint64 {
	return entity.ValueField
}

func (entity *accessControlEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *accessControlEntity) String() string {
	return fmt.Sprintf("AccessControl (Id: %d, Key: %v, Value: %v)", entity.Id(), entity.Key(), entity.Value())
}
