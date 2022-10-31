package entity

import (
	"fmt"
	"reflect"
	"time"

	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts/model"
)

var CategoryTypeEntityType = reflect.TypeOf(categoryTypeEntity{})

// noinspection GoUnusedExportedFunction
func InitializeCategoryTypeEntity() {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
}

type categoryTypeEntity struct {
	entity
	DescriptionField string `json:"description" previous:"id" storage:"VARCHAR(64)" default:"''"`
}

func NewCategoryTypeEntity(id int64, description string) ICategoryTypeEntity {
	return &categoryTypeEntity{
		entity:           entity{IdField: id},
		DescriptionField: description,
	}
}

type categoryTypePipeEntity struct {
	categoryTypeEntity
	pipeEntity
}

func NewCategoryTypePipeEntity(id int64, description string, source string, editor int64, payload string) ICategoryTypePipeEntity {
	return &categoryTypePipeEntity{
		categoryTypeEntity: categoryTypeEntity{
			entity:           entity{IdField: id, PayloadField: payload},
			DescriptionField: description,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_CATEGORY_TYPE,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *categoryTypeEntity) Description() string {
	return entity.DescriptionField
}

func (entity *categoryTypeEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *categoryTypeEntity) String() string {
	return fmt.Sprintf("CategoryType (Id: %d, Description: %v)", entity.Id(), entity.Description())
}
