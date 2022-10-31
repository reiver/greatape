package entity

import (
	"fmt"
	"reflect"
	"time"

	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts/model"
)

var CategoryEntityType = reflect.TypeOf(categoryEntity{})

// noinspection GoUnusedExportedFunction
func InitializeCategoryEntity() {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
}

type categoryEntity struct {
	entity
	CategoryTypeIdField int64  `json:"category_type_id"`
	CategoryIdField     int64  `json:"category_id"`
	TitleField          string `json:"title" previous:"category_id" storage:"VARCHAR(64)" default:"''"`
	DescriptionField    string `json:"description" previous:"title" storage:"VARCHAR(64)" default:"''"`
}

func NewCategoryEntity(id int64, categoryTypeId int64, categoryId int64, title string, description string) ICategoryEntity {
	return &categoryEntity{
		entity:              entity{IdField: id},
		CategoryTypeIdField: categoryTypeId,
		CategoryIdField:     categoryId,
		TitleField:          title,
		DescriptionField:    description,
	}
}

type categoryPipeEntity struct {
	categoryEntity
	pipeEntity
}

func NewCategoryPipeEntity(id int64, categoryTypeId int64, categoryId int64, title string, description string, source string, editor int64, payload string) ICategoryPipeEntity {
	return &categoryPipeEntity{
		categoryEntity: categoryEntity{
			entity:              entity{IdField: id, PayloadField: payload},
			CategoryTypeIdField: categoryTypeId,
			CategoryIdField:     categoryId,
			TitleField:          title,
			DescriptionField:    description,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_CATEGORY,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *categoryEntity) CategoryTypeId() int64 {
	return entity.CategoryTypeIdField
}

func (entity *categoryEntity) CategoryId() int64 {
	return entity.CategoryIdField
}

func (entity *categoryEntity) Title() string {
	return entity.TitleField
}

func (entity *categoryEntity) Description() string {
	return entity.DescriptionField
}

func (entity *categoryEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *categoryEntity) String() string {
	return fmt.Sprintf("Category (Id: %d, CategoryTypeId: %d, CategoryId: %d, Title: %v, Description: %v)", entity.Id(), entity.CategoryTypeId(), entity.CategoryId(), entity.Title(), entity.Description())
}
