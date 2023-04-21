package entity

import (
	"fmt"
	"reflect"
	"time"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
)

var DocumentEntityType = reflect.TypeOf(documentEntity{})

// noinspection GoUnusedExportedFunction
func InitializeDocumentEntity() {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
}

type documentEntity struct {
	entity
	ContentField string `json:"content" previous:"id" storage:"JSON" default:"'{}'"`
}

func NewDocumentEntity(id int64, content string) IDocumentEntity {
	return &documentEntity{
		entity:       entity{IdField: id},
		ContentField: content,
	}
}

type documentPipeEntity struct {
	documentEntity
	pipeEntity
}

func NewDocumentPipeEntity(id int64, content string, source string, editor int64, payload string) IDocumentPipeEntity {
	return &documentPipeEntity{
		documentEntity: documentEntity{
			entity:       entity{IdField: id, PayloadField: payload},
			ContentField: content,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_DOCUMENT,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *documentEntity) Content() string {
	return entity.ContentField
}

func (entity *documentEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *documentEntity) String() string {
	return fmt.Sprintf("Document (Id: %d, Content: %v)", entity.Id(), entity.Content())
}
