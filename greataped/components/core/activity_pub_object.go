package core

import (
	"fmt"

	. "github.com/xeronith/diamante/contracts/security"
	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/model/repository"
)

type activityPubObject struct {
	context   string
	id        string
	type_     string
	actor     string
	from      string
	to        []string
	inReplyTo string
	content   string
	published string
}

// noinspection GoUnusedExportedFunction
func InitializeActivityPubObject() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewActivityPubObject() (IActivityPubObject, error) {
	instance := &activityPubObject{}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (activityPubObject *activityPubObject) Context() string {
	return activityPubObject.context
}

func (activityPubObject *activityPubObject) SetContext(value string) {
	activityPubObject.context = value
}

func (activityPubObject *activityPubObject) Id() string {
	return activityPubObject.id
}

func (activityPubObject *activityPubObject) SetId(value string) {
	activityPubObject.id = value
}

func (activityPubObject *activityPubObject) Type() string {
	return activityPubObject.type_
}

func (activityPubObject *activityPubObject) SetType(value string) {
	activityPubObject.type_ = value
}

func (activityPubObject *activityPubObject) Actor() string {
	return activityPubObject.actor
}

func (activityPubObject *activityPubObject) SetActor(value string) {
	activityPubObject.actor = value
}

func (activityPubObject *activityPubObject) From() string {
	return activityPubObject.from
}

func (activityPubObject *activityPubObject) SetFrom(value string) {
	activityPubObject.from = value
}

func (activityPubObject *activityPubObject) To() []string {
	return activityPubObject.to
}

func (activityPubObject *activityPubObject) SetTo(value []string) {
	activityPubObject.to = value
}

func (activityPubObject *activityPubObject) InReplyTo() string {
	return activityPubObject.inReplyTo
}

func (activityPubObject *activityPubObject) SetInReplyTo(value string) {
	activityPubObject.inReplyTo = value
}

func (activityPubObject *activityPubObject) Content() string {
	return activityPubObject.content
}

func (activityPubObject *activityPubObject) SetContent(value string) {
	activityPubObject.content = value
}

func (activityPubObject *activityPubObject) Published() string {
	return activityPubObject.published
}

func (activityPubObject *activityPubObject) SetPublished(value string) {
	activityPubObject.published = value
}

func (activityPubObject *activityPubObject) Validate() error {
	return nil
}

func (activityPubObject *activityPubObject) String() string {
	return fmt.Sprintf("ActivityPubObject (Id: %d)", 0)
}

//------------------------------------------------------------------------------

type activityPubObjects struct {
	collection ActivityPubObjects
}

// NewActivityPubObjects creates an empty collection of 'Activity Pub Object' which is not thread-safe.
func NewActivityPubObjects() IActivityPubObjectCollection {
	return &activityPubObjects{
		collection: make(ActivityPubObjects, 0),
	}
}

func (activityPubObjects *activityPubObjects) Count() int {
	return len(activityPubObjects.collection)
}

func (activityPubObjects *activityPubObjects) IsEmpty() bool {
	return len(activityPubObjects.collection) == 0
}

func (activityPubObjects *activityPubObjects) IsNotEmpty() bool {
	return len(activityPubObjects.collection) > 0
}

func (activityPubObjects *activityPubObjects) HasExactlyOneItem() bool {
	return len(activityPubObjects.collection) == 1
}

func (activityPubObjects *activityPubObjects) HasAtLeastOneItem() bool {
	return len(activityPubObjects.collection) >= 1
}

func (activityPubObjects *activityPubObjects) First() IActivityPubObject {
	return activityPubObjects.collection[0]
}

func (activityPubObjects *activityPubObjects) Append(activityPubObject IActivityPubObject) {
	activityPubObjects.collection = append(activityPubObjects.collection, activityPubObject)
}

func (activityPubObjects *activityPubObjects) ForEach(iterator ActivityPubObjectIterator) {
	if iterator == nil {
		return
	}

	for _, value := range activityPubObjects.collection {
		iterator(value)
	}
}

func (activityPubObjects *activityPubObjects) Array() ActivityPubObjects {
	return activityPubObjects.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) ActivityPubObjectExists(id int64) bool {
	return dispatcher.conductor.ActivityPubObjectManager().Exists(id)
}

func (dispatcher *dispatcher) ActivityPubObjectExistsWhich(condition ActivityPubObjectCondition) bool {
	return dispatcher.conductor.ActivityPubObjectManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListActivityPubObjects() IActivityPubObjectCollection {
	return dispatcher.conductor.ActivityPubObjectManager().ListActivityPubObjects(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubObject(iterator ActivityPubObjectIterator) {
	dispatcher.conductor.ActivityPubObjectManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterActivityPubObjects(predicate ActivityPubObjectFilterPredicate) IActivityPubObjectCollection {
	return dispatcher.conductor.ActivityPubObjectManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapActivityPubObjects(predicate ActivityPubObjectMapPredicate) IActivityPubObjectCollection {
	return dispatcher.conductor.ActivityPubObjectManager().Map(predicate)
}

func (dispatcher *dispatcher) GetActivityPubObject(id int64) IActivityPubObject {
	if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().GetActivityPubObject(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return activityPubObject
	}
}

func (dispatcher *dispatcher) AddActivityPubObject() IActivityPubObject {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().AddActivityPubObjectAtomic(transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	} else {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().AddActivityPubObject(dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	}
}

func (dispatcher *dispatcher) AddActivityPubObjectWithCustomId(id int64) IActivityPubObject {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().AddActivityPubObjectWithCustomIdAtomic(id, transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	} else {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().AddActivityPubObjectWithCustomId(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	}
}

func (dispatcher *dispatcher) LogActivityPubObject(source string, payload string) {
	dispatcher.conductor.ActivityPubObjectManager().Log(source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateActivityPubObject(id int64) IActivityPubObject {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().UpdateActivityPubObjectAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	} else {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().UpdateActivityPubObject(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateActivityPubObjectObject(object IObject, activityPubObject IActivityPubObject) IActivityPubObject {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().UpdateActivityPubObjectAtomic(transaction, object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	} else {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().UpdateActivityPubObject(object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateActivityPubObjectObject(object IObject, activityPubObject IActivityPubObject) IActivityPubObject {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().AddOrUpdateActivityPubObjectObjectAtomic(transaction, object.Id(), activityPubObject, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	} else {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().AddOrUpdateActivityPubObjectObject(object.Id(), activityPubObject, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	}
}

func (dispatcher *dispatcher) RemoveActivityPubObject(id int64) IActivityPubObject {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().RemoveActivityPubObjectAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	} else {
		if activityPubObject, err := dispatcher.conductor.ActivityPubObjectManager().RemoveActivityPubObject(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubObject
		}
	}
}
