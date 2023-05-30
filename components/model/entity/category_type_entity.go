package entity

import (
	"fmt"
	"reflect"
	"time"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
)

var CategoryTypeEntityType = reflect.TypeOf(categoryTypeEntity{})

type categoryTypeEntity struct {
	entity
	DescriptionField string `json:"description" storage:"VARCHAR(64)" default:"''"`
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
