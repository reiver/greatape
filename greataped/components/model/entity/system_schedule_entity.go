package entity

import (
	"fmt"
	"reflect"
	"time"

	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts/model"
)

var SystemScheduleEntityType = reflect.TypeOf(systemScheduleEntity{})

// noinspection GoUnusedExportedFunction
func InitializeSystemScheduleEntity() {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
}

type systemScheduleEntity struct {
	entity
	EnabledField bool   `json:"enabled" previous:"id" storage:"BIT(1)" default:"FALSE"`
	ConfigField  string `json:"config" previous:"enabled" storage:"VARCHAR(1024)" default:"''"`
}

func NewSystemScheduleEntity(id int64, enabled bool, config string) ISystemScheduleEntity {
	return &systemScheduleEntity{
		entity:       entity{IdField: id},
		EnabledField: enabled,
		ConfigField:  config,
	}
}

type systemSchedulePipeEntity struct {
	systemScheduleEntity
	pipeEntity
}

func NewSystemSchedulePipeEntity(id int64, enabled bool, config string, source string, editor int64, payload string) ISystemSchedulePipeEntity {
	return &systemSchedulePipeEntity{
		systemScheduleEntity: systemScheduleEntity{
			entity:       entity{IdField: id, PayloadField: payload},
			EnabledField: enabled,
			ConfigField:  config,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_SYSTEM_SCHEDULE,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *systemScheduleEntity) Enabled() bool {
	return entity.EnabledField
}

func (entity *systemScheduleEntity) Config() string {
	return entity.ConfigField
}

func (entity *systemScheduleEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *systemScheduleEntity) String() string {
	return fmt.Sprintf("SystemSchedule (Id: %d, Enabled: %v, Config: %v)", entity.Id(), entity.Enabled(), entity.Config())
}
