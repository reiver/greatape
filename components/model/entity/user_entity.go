package entity

import (
	"fmt"
	"reflect"
	"time"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
)

var UserEntityType = reflect.TypeOf(userEntity{})

type userEntity struct {
	entity
	GithubField string `json:"github" storage:"VARCHAR(512)" default:"''"`
}

func NewUserEntity(id int64, github string) IUserEntity {
	return &userEntity{
		entity:      entity{IdField: id},
		GithubField: github,
	}
}

type userPipeEntity struct {
	userEntity
	pipeEntity
}

func NewUserPipeEntity(id int64, github string, source string, editor int64, payload string) IUserPipeEntity {
	return &userPipeEntity{
		userEntity: userEntity{
			entity:      entity{IdField: id, PayloadField: payload},
			GithubField: github,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_USER,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *userEntity) Github() string {
	return entity.GithubField
}

func (entity *userEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *userEntity) String() string {
	return fmt.Sprintf("User (Id: %d, Github: %v)", entity.Id(), entity.Github())
}
