package core

import (
	"fmt"

	. "github.com/reiver/greatape/components/contracts"
)

type activityPubActivity struct {
	context   string
	id        string
	type_     string
	actor     string
	object    IActivityPubObject
	from      string
	to        []string
	inReplyTo string
	content   string
	published string
}

func NewActivityPubActivity() (IActivityPubActivity, error) {
	instance := &activityPubActivity{}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (activityPubActivity *activityPubActivity) Context() string {
	return activityPubActivity.context
}

func (activityPubActivity *activityPubActivity) SetContext(value string) {
	activityPubActivity.context = value
}

func (activityPubActivity *activityPubActivity) Id() string {
	return activityPubActivity.id
}

func (activityPubActivity *activityPubActivity) SetId(value string) {
	activityPubActivity.id = value
}

func (activityPubActivity *activityPubActivity) Type() string {
	return activityPubActivity.type_
}

func (activityPubActivity *activityPubActivity) SetType(value string) {
	activityPubActivity.type_ = value
}

func (activityPubActivity *activityPubActivity) Actor() string {
	return activityPubActivity.actor
}

func (activityPubActivity *activityPubActivity) SetActor(value string) {
	activityPubActivity.actor = value
}

func (activityPubActivity *activityPubActivity) Object() IActivityPubObject {
	return activityPubActivity.object
}

func (activityPubActivity *activityPubActivity) SetObject(value IActivityPubObject) {
	activityPubActivity.object = value
}

func (activityPubActivity *activityPubActivity) From() string {
	return activityPubActivity.from
}

func (activityPubActivity *activityPubActivity) SetFrom(value string) {
	activityPubActivity.from = value
}

func (activityPubActivity *activityPubActivity) To() []string {
	return activityPubActivity.to
}

func (activityPubActivity *activityPubActivity) SetTo(value []string) {
	activityPubActivity.to = value
}

func (activityPubActivity *activityPubActivity) InReplyTo() string {
	return activityPubActivity.inReplyTo
}

func (activityPubActivity *activityPubActivity) SetInReplyTo(value string) {
	activityPubActivity.inReplyTo = value
}

func (activityPubActivity *activityPubActivity) Content() string {
	return activityPubActivity.content
}

func (activityPubActivity *activityPubActivity) SetContent(value string) {
	activityPubActivity.content = value
}

func (activityPubActivity *activityPubActivity) Published() string {
	return activityPubActivity.published
}

func (activityPubActivity *activityPubActivity) SetPublished(value string) {
	activityPubActivity.published = value
}

func (activityPubActivity *activityPubActivity) Validate() error {
	return nil
}

func (activityPubActivity *activityPubActivity) String() string {
	return fmt.Sprintf("ActivityPubActivity (Id: %d)", 0)
}

//------------------------------------------------------------------------------

type activityPubActivities struct {
	collection ActivityPubActivities
}

// NewActivityPubActivities creates an empty collection of 'Activity Pub Activity' which is not thread-safe.
func NewActivityPubActivities() IActivityPubActivityCollection {
	return &activityPubActivities{
		collection: make(ActivityPubActivities, 0),
	}
}

func (activityPubActivities *activityPubActivities) Count() int {
	return len(activityPubActivities.collection)
}

func (activityPubActivities *activityPubActivities) IsEmpty() bool {
	return len(activityPubActivities.collection) == 0
}

func (activityPubActivities *activityPubActivities) IsNotEmpty() bool {
	return len(activityPubActivities.collection) > 0
}

func (activityPubActivities *activityPubActivities) HasExactlyOneItem() bool {
	return len(activityPubActivities.collection) == 1
}

func (activityPubActivities *activityPubActivities) HasAtLeastOneItem() bool {
	return len(activityPubActivities.collection) >= 1
}

func (activityPubActivities *activityPubActivities) First() IActivityPubActivity {
	return activityPubActivities.collection[0]
}

func (activityPubActivities *activityPubActivities) Append(activityPubActivity IActivityPubActivity) {
	activityPubActivities.collection = append(activityPubActivities.collection, activityPubActivity)
}

func (activityPubActivities *activityPubActivities) Reverse() IActivityPubActivityCollection {
	slice := activityPubActivities.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	activityPubActivities.collection = slice

	return activityPubActivities
}

func (activityPubActivities *activityPubActivities) ForEach(iterator ActivityPubActivityIterator) {
	if iterator == nil {
		return
	}

	for _, value := range activityPubActivities.collection {
		iterator(value)
	}
}

func (activityPubActivities *activityPubActivities) Array() ActivityPubActivities {
	return activityPubActivities.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) ActivityPubActivityExists(id int64) bool {
	return dispatcher.conductor.ActivityPubActivityManager().Exists(id)
}

func (dispatcher *dispatcher) ActivityPubActivityExistsWhich(condition ActivityPubActivityCondition) bool {
	return dispatcher.conductor.ActivityPubActivityManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListActivityPubActivities() IActivityPubActivityCollection {
	return dispatcher.conductor.ActivityPubActivityManager().ListActivityPubActivities(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubActivity(iterator ActivityPubActivityIterator) {
	dispatcher.conductor.ActivityPubActivityManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterActivityPubActivities(predicate ActivityPubActivityFilterPredicate) IActivityPubActivityCollection {
	return dispatcher.conductor.ActivityPubActivityManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapActivityPubActivities(predicate ActivityPubActivityMapPredicate) IActivityPubActivityCollection {
	return dispatcher.conductor.ActivityPubActivityManager().Map(predicate)
}

func (dispatcher *dispatcher) GetActivityPubActivity(id int64) IActivityPubActivity {
	if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().GetActivityPubActivity(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return activityPubActivity
	}
}

func (dispatcher *dispatcher) AddActivityPubActivity() IActivityPubActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().AddActivityPubActivityAtomic(transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	} else {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().AddActivityPubActivity(dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	}
}

func (dispatcher *dispatcher) AddActivityPubActivityWithCustomId(id int64) IActivityPubActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().AddActivityPubActivityWithCustomIdAtomic(id, transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	} else {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().AddActivityPubActivityWithCustomId(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	}
}

func (dispatcher *dispatcher) LogActivityPubActivity(source string, payload string) {
	dispatcher.conductor.ActivityPubActivityManager().Log(source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateActivityPubActivity(id int64) IActivityPubActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().UpdateActivityPubActivityAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	} else {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().UpdateActivityPubActivity(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateActivityPubActivityObject(object IObject, activityPubActivity IActivityPubActivity) IActivityPubActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().UpdateActivityPubActivityAtomic(transaction, object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	} else {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().UpdateActivityPubActivity(object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateActivityPubActivityObject(object IObject, activityPubActivity IActivityPubActivity) IActivityPubActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().AddOrUpdateActivityPubActivityObjectAtomic(transaction, object.Id(), activityPubActivity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	} else {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().AddOrUpdateActivityPubActivityObject(object.Id(), activityPubActivity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	}
}

func (dispatcher *dispatcher) RemoveActivityPubActivity(id int64) IActivityPubActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().RemoveActivityPubActivityAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	} else {
		if activityPubActivity, err := dispatcher.conductor.ActivityPubActivityManager().RemoveActivityPubActivity(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubActivity
		}
	}
}
