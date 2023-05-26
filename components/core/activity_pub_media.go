package core

import (
	"fmt"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
)

type activityPubMedia struct {
	mediaType string
	type_     string
	url       string
	width     int32
	height    int32
}

// noinspection GoUnusedExportedFunction
func InitializeActivityPubMedia() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewActivityPubMedia() (IActivityPubMedia, error) {
	instance := &activityPubMedia{}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (activityPubMedia *activityPubMedia) MediaType() string {
	return activityPubMedia.mediaType
}

func (activityPubMedia *activityPubMedia) SetMediaType(value string) {
	activityPubMedia.mediaType = value
}

func (activityPubMedia *activityPubMedia) Type() string {
	return activityPubMedia.type_
}

func (activityPubMedia *activityPubMedia) SetType(value string) {
	activityPubMedia.type_ = value
}

func (activityPubMedia *activityPubMedia) Url() string {
	return activityPubMedia.url
}

func (activityPubMedia *activityPubMedia) SetUrl(value string) {
	activityPubMedia.url = value
}

func (activityPubMedia *activityPubMedia) Width() int32 {
	return activityPubMedia.width
}

func (activityPubMedia *activityPubMedia) SetWidth(value int32) {
	activityPubMedia.width = value
}

func (activityPubMedia *activityPubMedia) Height() int32 {
	return activityPubMedia.height
}

func (activityPubMedia *activityPubMedia) SetHeight(value int32) {
	activityPubMedia.height = value
}

func (activityPubMedia *activityPubMedia) Validate() error {
	return nil
}

func (activityPubMedia *activityPubMedia) String() string {
	return fmt.Sprintf("ActivityPubMedia (Id: %d)", 0)
}

//------------------------------------------------------------------------------

type activityPubMedias struct {
	collection ActivityPubMedias
}

// NewActivityPubMedias creates an empty collection of 'Activity Pub Media' which is not thread-safe.
func NewActivityPubMedias() IActivityPubMediaCollection {
	return &activityPubMedias{
		collection: make(ActivityPubMedias, 0),
	}
}

func (activityPubMedias *activityPubMedias) Count() int {
	return len(activityPubMedias.collection)
}

func (activityPubMedias *activityPubMedias) IsEmpty() bool {
	return len(activityPubMedias.collection) == 0
}

func (activityPubMedias *activityPubMedias) IsNotEmpty() bool {
	return len(activityPubMedias.collection) > 0
}

func (activityPubMedias *activityPubMedias) HasExactlyOneItem() bool {
	return len(activityPubMedias.collection) == 1
}

func (activityPubMedias *activityPubMedias) HasAtLeastOneItem() bool {
	return len(activityPubMedias.collection) >= 1
}

func (activityPubMedias *activityPubMedias) First() IActivityPubMedia {
	return activityPubMedias.collection[0]
}

func (activityPubMedias *activityPubMedias) Append(activityPubMedia IActivityPubMedia) {
	activityPubMedias.collection = append(activityPubMedias.collection, activityPubMedia)
}

func (activityPubMedias *activityPubMedias) Reverse() IActivityPubMediaCollection {
	slice := activityPubMedias.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	activityPubMedias.collection = slice

	return activityPubMedias
}

func (activityPubMedias *activityPubMedias) ForEach(iterator ActivityPubMediaIterator) {
	if iterator == nil {
		return
	}

	for _, value := range activityPubMedias.collection {
		iterator(value)
	}
}

func (activityPubMedias *activityPubMedias) Array() ActivityPubMedias {
	return activityPubMedias.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) ActivityPubMediaExists(id int64) bool {
	return dispatcher.conductor.ActivityPubMediaManager().Exists(id)
}

func (dispatcher *dispatcher) ActivityPubMediaExistsWhich(condition ActivityPubMediaCondition) bool {
	return dispatcher.conductor.ActivityPubMediaManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListActivityPubMedias() IActivityPubMediaCollection {
	return dispatcher.conductor.ActivityPubMediaManager().ListActivityPubMedias(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubMedia(iterator ActivityPubMediaIterator) {
	dispatcher.conductor.ActivityPubMediaManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterActivityPubMedias(predicate ActivityPubMediaFilterPredicate) IActivityPubMediaCollection {
	return dispatcher.conductor.ActivityPubMediaManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapActivityPubMedias(predicate ActivityPubMediaMapPredicate) IActivityPubMediaCollection {
	return dispatcher.conductor.ActivityPubMediaManager().Map(predicate)
}

func (dispatcher *dispatcher) GetActivityPubMedia(id int64) IActivityPubMedia {
	if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().GetActivityPubMedia(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return activityPubMedia
	}
}

func (dispatcher *dispatcher) AddActivityPubMedia() IActivityPubMedia {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().AddActivityPubMediaAtomic(transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	} else {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().AddActivityPubMedia(dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	}
}

func (dispatcher *dispatcher) AddActivityPubMediaWithCustomId(id int64) IActivityPubMedia {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().AddActivityPubMediaWithCustomIdAtomic(id, transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	} else {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().AddActivityPubMediaWithCustomId(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	}
}

func (dispatcher *dispatcher) LogActivityPubMedia(source string, payload string) {
	dispatcher.conductor.ActivityPubMediaManager().Log(source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateActivityPubMedia(id int64) IActivityPubMedia {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().UpdateActivityPubMediaAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	} else {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().UpdateActivityPubMedia(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateActivityPubMediaObject(object IObject, activityPubMedia IActivityPubMedia) IActivityPubMedia {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().UpdateActivityPubMediaAtomic(transaction, object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	} else {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().UpdateActivityPubMedia(object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateActivityPubMediaObject(object IObject, activityPubMedia IActivityPubMedia) IActivityPubMedia {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().AddOrUpdateActivityPubMediaObjectAtomic(transaction, object.Id(), activityPubMedia, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	} else {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().AddOrUpdateActivityPubMediaObject(object.Id(), activityPubMedia, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	}
}

func (dispatcher *dispatcher) RemoveActivityPubMedia(id int64) IActivityPubMedia {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().RemoveActivityPubMediaAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	} else {
		if activityPubMedia, err := dispatcher.conductor.ActivityPubMediaManager().RemoveActivityPubMedia(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubMedia
		}
	}
}
