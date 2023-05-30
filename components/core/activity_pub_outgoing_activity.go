package core

import (
	"fmt"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
)

type activityPubOutgoingActivity struct {
	object
	identityId       int64
	uniqueIdentifier string
	timestamp        int64
	from             string
	to               string
	content          string
	raw              string
}

func NewActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) (IActivityPubOutgoingActivity, error) {
	instance := &activityPubOutgoingActivity{
		object: object{
			id: id,
		},
		identityId:       identityId,
		uniqueIdentifier: uniqueIdentifier,
		timestamp:        timestamp,
		from:             from,
		to:               to,
		content:          content,
		raw:              raw,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewActivityPubOutgoingActivityFromEntity(entity IActivityPubOutgoingActivityEntity) (IActivityPubOutgoingActivity, error) {
	instance := &activityPubOutgoingActivity{
		object: object{
			id: entity.Id(),
		},
		identityId:       entity.IdentityId(),
		uniqueIdentifier: entity.UniqueIdentifier(),
		timestamp:        entity.Timestamp(),
		from:             entity.From(),
		to:               entity.To(),
		content:          entity.Content(),
		raw:              entity.Raw(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) DependenciesAreUnknown() bool {
	// noinspection GoBoolExpressions
	return activityPubOutgoingActivity.identityId == 0 || false
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) IdentityId() int64 {
	return activityPubOutgoingActivity.identityId
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) AssertBelongsToIdentity(_identity IIdentity) {
	if activityPubOutgoingActivity.identityId != _identity.Id() {
		panic(ERROR_MESSAGE_ACTIVITY_PUB_OUTGOING_ACTIVITY_NOT_FOUND)
	}
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) IdentityIsUnknown() bool {
	return activityPubOutgoingActivity.identityId == 0
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) AssertIdentityIsProvided() {
	if activityPubOutgoingActivity.identityId == 0 {
		panic(ERROR_MESSAGE_UNKNOWN_IDENTITY)
	}
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) AssertIdentity(identityId int64) {
	if activityPubOutgoingActivity.identityId != identityId {
		panic(ERROR_MESSAGE_UNKNOWN_IDENTITY)
	}
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UniqueIdentifier() string {
	return activityPubOutgoingActivity.uniqueIdentifier
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateUniqueIdentifier(uniqueIdentifier string, editor Identity) {
	if err := repository.ActivityPubOutgoingActivities.UpdateUniqueIdentifier(activityPubOutgoingActivity.id, uniqueIdentifier, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubOutgoingActivity.uniqueIdentifier = uniqueIdentifier
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateUniqueIdentifierAtomic(transaction ITransaction, uniqueIdentifier string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubOutgoingActivity.uniqueIdentifier = uniqueIdentifier
	})

	if err := repository.ActivityPubOutgoingActivities.UpdateUniqueIdentifierAtomic(transaction, activityPubOutgoingActivity.id, uniqueIdentifier, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) Timestamp() int64 {
	return activityPubOutgoingActivity.timestamp
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateTimestamp(timestamp int64, editor Identity) {
	if err := repository.ActivityPubOutgoingActivities.UpdateTimestamp(activityPubOutgoingActivity.id, timestamp, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubOutgoingActivity.timestamp = timestamp
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateTimestampAtomic(transaction ITransaction, timestamp int64, editor Identity) {
	transaction.OnCommit(func() {
		activityPubOutgoingActivity.timestamp = timestamp
	})

	if err := repository.ActivityPubOutgoingActivities.UpdateTimestampAtomic(transaction, activityPubOutgoingActivity.id, timestamp, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) From() string {
	return activityPubOutgoingActivity.from
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateFrom(from string, editor Identity) {
	if err := repository.ActivityPubOutgoingActivities.UpdateFrom(activityPubOutgoingActivity.id, from, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubOutgoingActivity.from = from
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateFromAtomic(transaction ITransaction, from string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubOutgoingActivity.from = from
	})

	if err := repository.ActivityPubOutgoingActivities.UpdateFromAtomic(transaction, activityPubOutgoingActivity.id, from, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) To() string {
	return activityPubOutgoingActivity.to
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateTo(to string, editor Identity) {
	if err := repository.ActivityPubOutgoingActivities.UpdateTo(activityPubOutgoingActivity.id, to, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubOutgoingActivity.to = to
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateToAtomic(transaction ITransaction, to string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubOutgoingActivity.to = to
	})

	if err := repository.ActivityPubOutgoingActivities.UpdateToAtomic(transaction, activityPubOutgoingActivity.id, to, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) Content() string {
	return activityPubOutgoingActivity.content
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateContent(content string, editor Identity) {
	if err := repository.ActivityPubOutgoingActivities.UpdateContent(activityPubOutgoingActivity.id, content, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubOutgoingActivity.content = content
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateContentAtomic(transaction ITransaction, content string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubOutgoingActivity.content = content
	})

	if err := repository.ActivityPubOutgoingActivities.UpdateContentAtomic(transaction, activityPubOutgoingActivity.id, content, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) Raw() string {
	return activityPubOutgoingActivity.raw
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateRaw(raw string, editor Identity) {
	if err := repository.ActivityPubOutgoingActivities.UpdateRaw(activityPubOutgoingActivity.id, raw, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubOutgoingActivity.raw = raw
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) UpdateRawAtomic(transaction ITransaction, raw string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubOutgoingActivity.raw = raw
	})

	if err := repository.ActivityPubOutgoingActivities.UpdateRawAtomic(transaction, activityPubOutgoingActivity.id, raw, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) Validate() error {
	return nil
}

func (activityPubOutgoingActivity *activityPubOutgoingActivity) String() string {
	return fmt.Sprintf("ActivityPubOutgoingActivity (Id: %d, IdentityId: %d, UniqueIdentifier: %v, Timestamp: %v, From: %v, To: %v, Content: %v, Raw: %v)", activityPubOutgoingActivity.Id(), activityPubOutgoingActivity.IdentityId(), activityPubOutgoingActivity.UniqueIdentifier(), activityPubOutgoingActivity.Timestamp(), activityPubOutgoingActivity.From(), activityPubOutgoingActivity.To(), activityPubOutgoingActivity.Content(), activityPubOutgoingActivity.Raw())
}

//------------------------------------------------------------------------------

type activityPubOutgoingActivities struct {
	collection ActivityPubOutgoingActivities
}

// NewActivityPubOutgoingActivities creates an empty collection of 'Activity Pub Outgoing Activity' which is not thread-safe.
func NewActivityPubOutgoingActivities() IActivityPubOutgoingActivityCollection {
	return &activityPubOutgoingActivities{
		collection: make(ActivityPubOutgoingActivities, 0),
	}
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) Count() int {
	return len(activityPubOutgoingActivities.collection)
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) IsEmpty() bool {
	return len(activityPubOutgoingActivities.collection) == 0
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) IsNotEmpty() bool {
	return len(activityPubOutgoingActivities.collection) > 0
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) HasExactlyOneItem() bool {
	return len(activityPubOutgoingActivities.collection) == 1
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) HasAtLeastOneItem() bool {
	return len(activityPubOutgoingActivities.collection) >= 1
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) First() IActivityPubOutgoingActivity {
	return activityPubOutgoingActivities.collection[0]
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) Append(activityPubOutgoingActivity IActivityPubOutgoingActivity) {
	activityPubOutgoingActivities.collection = append(activityPubOutgoingActivities.collection, activityPubOutgoingActivity)
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) Reverse() IActivityPubOutgoingActivityCollection {
	slice := activityPubOutgoingActivities.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	activityPubOutgoingActivities.collection = slice

	return activityPubOutgoingActivities
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) ForEach(iterator ActivityPubOutgoingActivityIterator) {
	if iterator == nil {
		return
	}

	for _, value := range activityPubOutgoingActivities.collection {
		iterator(value)
	}
}

func (activityPubOutgoingActivities *activityPubOutgoingActivities) Array() ActivityPubOutgoingActivities {
	return activityPubOutgoingActivities.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) ActivityPubOutgoingActivityExists(id int64) bool {
	return dispatcher.conductor.ActivityPubOutgoingActivityManager().Exists(id)
}

func (dispatcher *dispatcher) ActivityPubOutgoingActivityExistsWhich(condition ActivityPubOutgoingActivityCondition) bool {
	return dispatcher.conductor.ActivityPubOutgoingActivityManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListActivityPubOutgoingActivities() IActivityPubOutgoingActivityCollection {
	return dispatcher.conductor.ActivityPubOutgoingActivityManager().ListActivityPubOutgoingActivities(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubOutgoingActivity(iterator ActivityPubOutgoingActivityIterator) {
	dispatcher.conductor.ActivityPubOutgoingActivityManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterActivityPubOutgoingActivities(predicate ActivityPubOutgoingActivityFilterPredicate) IActivityPubOutgoingActivityCollection {
	return dispatcher.conductor.ActivityPubOutgoingActivityManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapActivityPubOutgoingActivities(predicate ActivityPubOutgoingActivityMapPredicate) IActivityPubOutgoingActivityCollection {
	return dispatcher.conductor.ActivityPubOutgoingActivityManager().Map(predicate)
}

func (dispatcher *dispatcher) GetActivityPubOutgoingActivity(id int64) IActivityPubOutgoingActivity {
	if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().GetActivityPubOutgoingActivity(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return activityPubOutgoingActivity
	}
}

func (dispatcher *dispatcher) AddActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubOutgoingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().AddActivityPubOutgoingActivityAtomic(transaction, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	} else {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().AddActivityPubOutgoingActivity(identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	}
}

func (dispatcher *dispatcher) AddActivityPubOutgoingActivityWithCustomId(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubOutgoingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().AddActivityPubOutgoingActivityWithCustomIdAtomic(id, transaction, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	} else {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().AddActivityPubOutgoingActivityWithCustomId(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	}
}

func (dispatcher *dispatcher) LogActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, payload string) {
	dispatcher.conductor.ActivityPubOutgoingActivityManager().Log(identityId, uniqueIdentifier, timestamp, from, to, content, raw, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubOutgoingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().UpdateActivityPubOutgoingActivityAtomic(transaction, id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	} else {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().UpdateActivityPubOutgoingActivity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateActivityPubOutgoingActivityObject(object IObject, activityPubOutgoingActivity IActivityPubOutgoingActivity) IActivityPubOutgoingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().UpdateActivityPubOutgoingActivityAtomic(transaction, object.Id(), activityPubOutgoingActivity.IdentityId(), activityPubOutgoingActivity.UniqueIdentifier(), activityPubOutgoingActivity.Timestamp(), activityPubOutgoingActivity.From(), activityPubOutgoingActivity.To(), activityPubOutgoingActivity.Content(), activityPubOutgoingActivity.Raw(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	} else {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().UpdateActivityPubOutgoingActivity(object.Id(), activityPubOutgoingActivity.IdentityId(), activityPubOutgoingActivity.UniqueIdentifier(), activityPubOutgoingActivity.Timestamp(), activityPubOutgoingActivity.From(), activityPubOutgoingActivity.To(), activityPubOutgoingActivity.Content(), activityPubOutgoingActivity.Raw(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateActivityPubOutgoingActivityObject(object IObject, activityPubOutgoingActivity IActivityPubOutgoingActivity) IActivityPubOutgoingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().AddOrUpdateActivityPubOutgoingActivityObjectAtomic(transaction, object.Id(), activityPubOutgoingActivity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	} else {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().AddOrUpdateActivityPubOutgoingActivityObject(object.Id(), activityPubOutgoingActivity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	}
}

func (dispatcher *dispatcher) RemoveActivityPubOutgoingActivity(id int64) IActivityPubOutgoingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().RemoveActivityPubOutgoingActivityAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	} else {
		if activityPubOutgoingActivity, err := dispatcher.conductor.ActivityPubOutgoingActivityManager().RemoveActivityPubOutgoingActivity(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubOutgoingActivity
		}
	}
}

func (dispatcher *dispatcher) ListActivityPubOutgoingActivitiesByIdentity(identity IIdentity) IActivityPubOutgoingActivityCollection {
	return dispatcher.conductor.ActivityPubOutgoingActivityManager().ListActivityPubOutgoingActivitiesByIdentity(identity.Id(), 0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ListActivityPubOutgoingActivitiesByIdentityId(identityId int64) IActivityPubOutgoingActivityCollection {
	return dispatcher.conductor.ActivityPubOutgoingActivityManager().ListActivityPubOutgoingActivitiesByIdentity(identityId, 0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubOutgoingActivityByIdentity(identity IIdentity, iterator ActivityPubOutgoingActivityIterator) {
	dispatcher.conductor.ActivityPubOutgoingActivityManager().ForEachByIdentity(identity.Id(), iterator)
}

func (dispatcher *dispatcher) ForEachActivityPubOutgoingActivityByIdentityId(identityId int64, iterator ActivityPubOutgoingActivityIterator) {
	dispatcher.conductor.ActivityPubOutgoingActivityManager().ForEachByIdentity(identityId, iterator)
}
