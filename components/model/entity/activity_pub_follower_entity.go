package entity

import (
	"fmt"
	"reflect"
	"time"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts/model"
)

var ActivityPubFollowerEntityType = reflect.TypeOf(activityPubFollowerEntity{})

// noinspection GoUnusedExportedFunction
func InitializeActivityPubFollowerEntity() {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
}

type activityPubFollowerEntity struct {
	entity
	HandleField   string `json:"handle" previous:"id" storage:"VARCHAR(256)" default:"''"`
	InboxField    string `json:"inbox" previous:"handle" storage:"VARCHAR(256)" default:"''"`
	SubjectField  string `json:"subject" previous:"inbox" storage:"VARCHAR(256)" default:"''"`
	ActivityField string `json:"activity" previous:"subject" storage:"VARCHAR(4096)" default:"''"`
	AcceptedField bool   `json:"accepted" previous:"activity" storage:"BIT(1)" default:"FALSE"`
}

func NewActivityPubFollowerEntity(id int64, handle string, inbox string, subject string, activity string, accepted bool) IActivityPubFollowerEntity {
	return &activityPubFollowerEntity{
		entity:        entity{IdField: id},
		HandleField:   handle,
		InboxField:    inbox,
		SubjectField:  subject,
		ActivityField: activity,
		AcceptedField: accepted,
	}
}

type activityPubFollowerPipeEntity struct {
	activityPubFollowerEntity
	pipeEntity
}

func NewActivityPubFollowerPipeEntity(id int64, handle string, inbox string, subject string, activity string, accepted bool, source string, editor int64, payload string) IActivityPubFollowerPipeEntity {
	return &activityPubFollowerPipeEntity{
		activityPubFollowerEntity: activityPubFollowerEntity{
			entity:        entity{IdField: id, PayloadField: payload},
			HandleField:   handle,
			InboxField:    inbox,
			SubjectField:  subject,
			ActivityField: activity,
			AcceptedField: accepted,
		},
		pipeEntity: pipeEntity{
			Pipe:           PIPE_ACTIVITY_PUB_FOLLOWER,
			Source:         source,
			Editor:         editor,
			QueueTimestamp: time.Now(),
		},
	}
}

func (entity *activityPubFollowerEntity) Handle() string {
	return entity.HandleField
}

func (entity *activityPubFollowerEntity) Inbox() string {
	return entity.InboxField
}

func (entity *activityPubFollowerEntity) Subject() string {
	return entity.SubjectField
}

func (entity *activityPubFollowerEntity) Activity() string {
	return entity.ActivityField
}

func (entity *activityPubFollowerEntity) Accepted() bool {
	return entity.AcceptedField
}

func (entity *activityPubFollowerEntity) Validate() error {
	if entity.IdField <= 0 {
		return ERROR_INVALID_ID
	}

	return nil
}

func (entity *activityPubFollowerEntity) String() string {
	return fmt.Sprintf("ActivityPubFollower (Id: %d, Handle: %v, Inbox: %v, Subject: %v, Activity: %v, Accepted: %v)", entity.Id(), entity.Handle(), entity.Inbox(), entity.Subject(), entity.Activity(), entity.Accepted())
}
