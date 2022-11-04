package core

import (
	"fmt"

	. "github.com/xeronith/diamante/contracts/security"
	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/contracts/model"
	"rail.town/infrastructure/components/model/repository"
)

type activityPubIncomingActivity struct {
	object
	identityId       int64
	uniqueIdentifier string
	timestamp        int64
	from             string
	to               string
	content          string
	raw              string
}

// noinspection GoUnusedExportedFunction
func InitializeActivityPubIncomingActivity() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) (IActivityPubIncomingActivity, error) {
	instance := &activityPubIncomingActivity{
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

func NewActivityPubIncomingActivityFromEntity(entity IActivityPubIncomingActivityEntity) (IActivityPubIncomingActivity, error) {
	instance := &activityPubIncomingActivity{
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

func (activityPubIncomingActivity *activityPubIncomingActivity) DependenciesAreUnknown() bool {
	// noinspection GoBoolExpressions
	return activityPubIncomingActivity.identityId == 0 || false
}

func (activityPubIncomingActivity *activityPubIncomingActivity) IdentityId() int64 {
	return activityPubIncomingActivity.identityId
}

func (activityPubIncomingActivity *activityPubIncomingActivity) AssertBelongsToIdentity(_identity IIdentity) {
	if activityPubIncomingActivity.identityId != _identity.Id() {
		panic(ERROR_MESSAGE_ACTIVITY_PUB_INCOMING_ACTIVITY_NOT_FOUND)
	}
}

func (activityPubIncomingActivity *activityPubIncomingActivity) IdentityIsUnknown() bool {
	return activityPubIncomingActivity.identityId == 0
}

func (activityPubIncomingActivity *activityPubIncomingActivity) AssertIdentityIsProvided() {
	if activityPubIncomingActivity.identityId == 0 {
		panic(ERROR_MESSAGE_UNKNOWN_IDENTITY)
	}
}

func (activityPubIncomingActivity *activityPubIncomingActivity) AssertIdentity(identityId int64) {
	if activityPubIncomingActivity.identityId != identityId {
		panic(ERROR_MESSAGE_UNKNOWN_IDENTITY)
	}
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UniqueIdentifier() string {
	return activityPubIncomingActivity.uniqueIdentifier
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateUniqueIdentifier(uniqueIdentifier string, editor Identity) {
	if err := repository.ActivityPubIncomingActivities.UpdateUniqueIdentifier(activityPubIncomingActivity.id, uniqueIdentifier, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubIncomingActivity.uniqueIdentifier = uniqueIdentifier
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateUniqueIdentifierAtomic(transaction ITransaction, uniqueIdentifier string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubIncomingActivity.uniqueIdentifier = uniqueIdentifier
	})

	if err := repository.ActivityPubIncomingActivities.UpdateUniqueIdentifierAtomic(transaction, activityPubIncomingActivity.id, uniqueIdentifier, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubIncomingActivity *activityPubIncomingActivity) Timestamp() int64 {
	return activityPubIncomingActivity.timestamp
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateTimestamp(timestamp int64, editor Identity) {
	if err := repository.ActivityPubIncomingActivities.UpdateTimestamp(activityPubIncomingActivity.id, timestamp, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubIncomingActivity.timestamp = timestamp
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateTimestampAtomic(transaction ITransaction, timestamp int64, editor Identity) {
	transaction.OnCommit(func() {
		activityPubIncomingActivity.timestamp = timestamp
	})

	if err := repository.ActivityPubIncomingActivities.UpdateTimestampAtomic(transaction, activityPubIncomingActivity.id, timestamp, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubIncomingActivity *activityPubIncomingActivity) From() string {
	return activityPubIncomingActivity.from
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateFrom(from string, editor Identity) {
	if err := repository.ActivityPubIncomingActivities.UpdateFrom(activityPubIncomingActivity.id, from, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubIncomingActivity.from = from
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateFromAtomic(transaction ITransaction, from string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubIncomingActivity.from = from
	})

	if err := repository.ActivityPubIncomingActivities.UpdateFromAtomic(transaction, activityPubIncomingActivity.id, from, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubIncomingActivity *activityPubIncomingActivity) To() string {
	return activityPubIncomingActivity.to
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateTo(to string, editor Identity) {
	if err := repository.ActivityPubIncomingActivities.UpdateTo(activityPubIncomingActivity.id, to, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubIncomingActivity.to = to
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateToAtomic(transaction ITransaction, to string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubIncomingActivity.to = to
	})

	if err := repository.ActivityPubIncomingActivities.UpdateToAtomic(transaction, activityPubIncomingActivity.id, to, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubIncomingActivity *activityPubIncomingActivity) Content() string {
	return activityPubIncomingActivity.content
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateContent(content string, editor Identity) {
	if err := repository.ActivityPubIncomingActivities.UpdateContent(activityPubIncomingActivity.id, content, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubIncomingActivity.content = content
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateContentAtomic(transaction ITransaction, content string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubIncomingActivity.content = content
	})

	if err := repository.ActivityPubIncomingActivities.UpdateContentAtomic(transaction, activityPubIncomingActivity.id, content, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubIncomingActivity *activityPubIncomingActivity) Raw() string {
	return activityPubIncomingActivity.raw
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateRaw(raw string, editor Identity) {
	if err := repository.ActivityPubIncomingActivities.UpdateRaw(activityPubIncomingActivity.id, raw, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubIncomingActivity.raw = raw
}

func (activityPubIncomingActivity *activityPubIncomingActivity) UpdateRawAtomic(transaction ITransaction, raw string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubIncomingActivity.raw = raw
	})

	if err := repository.ActivityPubIncomingActivities.UpdateRawAtomic(transaction, activityPubIncomingActivity.id, raw, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubIncomingActivity *activityPubIncomingActivity) Validate() error {
	return nil
}

func (activityPubIncomingActivity *activityPubIncomingActivity) String() string {
	return fmt.Sprintf("ActivityPubIncomingActivity (Id: %d, IdentityId: %d, UniqueIdentifier: %v, Timestamp: %v, From: %v, To: %v, Content: %v, Raw: %v)", activityPubIncomingActivity.Id(), activityPubIncomingActivity.IdentityId(), activityPubIncomingActivity.UniqueIdentifier(), activityPubIncomingActivity.Timestamp(), activityPubIncomingActivity.From(), activityPubIncomingActivity.To(), activityPubIncomingActivity.Content(), activityPubIncomingActivity.Raw())
}

//------------------------------------------------------------------------------

type activityPubIncomingActivities struct {
	collection ActivityPubIncomingActivities
}

// NewActivityPubIncomingActivities creates an empty collection of 'Activity Pub Incoming Activity' which is not thread-safe.
func NewActivityPubIncomingActivities() IActivityPubIncomingActivityCollection {
	return &activityPubIncomingActivities{
		collection: make(ActivityPubIncomingActivities, 0),
	}
}

func (activityPubIncomingActivities *activityPubIncomingActivities) Count() int {
	return len(activityPubIncomingActivities.collection)
}

func (activityPubIncomingActivities *activityPubIncomingActivities) IsEmpty() bool {
	return len(activityPubIncomingActivities.collection) == 0
}

func (activityPubIncomingActivities *activityPubIncomingActivities) IsNotEmpty() bool {
	return len(activityPubIncomingActivities.collection) > 0
}

func (activityPubIncomingActivities *activityPubIncomingActivities) HasExactlyOneItem() bool {
	return len(activityPubIncomingActivities.collection) == 1
}

func (activityPubIncomingActivities *activityPubIncomingActivities) HasAtLeastOneItem() bool {
	return len(activityPubIncomingActivities.collection) >= 1
}

func (activityPubIncomingActivities *activityPubIncomingActivities) First() IActivityPubIncomingActivity {
	return activityPubIncomingActivities.collection[0]
}

func (activityPubIncomingActivities *activityPubIncomingActivities) Append(activityPubIncomingActivity IActivityPubIncomingActivity) {
	activityPubIncomingActivities.collection = append(activityPubIncomingActivities.collection, activityPubIncomingActivity)
}

func (activityPubIncomingActivities *activityPubIncomingActivities) ForEach(iterator ActivityPubIncomingActivityIterator) {
	if iterator == nil {
		return
	}

	for _, value := range activityPubIncomingActivities.collection {
		iterator(value)
	}
}

func (activityPubIncomingActivities *activityPubIncomingActivities) Array() ActivityPubIncomingActivities {
	return activityPubIncomingActivities.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) ActivityPubIncomingActivityExists(id int64) bool {
	return dispatcher.conductor.ActivityPubIncomingActivityManager().Exists(id)
}

func (dispatcher *dispatcher) ActivityPubIncomingActivityExistsWhich(condition ActivityPubIncomingActivityCondition) bool {
	return dispatcher.conductor.ActivityPubIncomingActivityManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListActivityPubIncomingActivities() IActivityPubIncomingActivityCollection {
	return dispatcher.conductor.ActivityPubIncomingActivityManager().ListActivityPubIncomingActivities(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubIncomingActivity(iterator ActivityPubIncomingActivityIterator) {
	dispatcher.conductor.ActivityPubIncomingActivityManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterActivityPubIncomingActivities(predicate ActivityPubIncomingActivityFilterPredicate) IActivityPubIncomingActivityCollection {
	return dispatcher.conductor.ActivityPubIncomingActivityManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapActivityPubIncomingActivities(predicate ActivityPubIncomingActivityMapPredicate) IActivityPubIncomingActivityCollection {
	return dispatcher.conductor.ActivityPubIncomingActivityManager().Map(predicate)
}

func (dispatcher *dispatcher) GetActivityPubIncomingActivity(id int64) IActivityPubIncomingActivity {
	if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().GetActivityPubIncomingActivity(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return activityPubIncomingActivity
	}
}

func (dispatcher *dispatcher) AddActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubIncomingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().AddActivityPubIncomingActivityAtomic(transaction, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	} else {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().AddActivityPubIncomingActivity(identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	}
}

func (dispatcher *dispatcher) AddActivityPubIncomingActivityWithCustomId(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubIncomingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().AddActivityPubIncomingActivityWithCustomIdAtomic(id, transaction, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	} else {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().AddActivityPubIncomingActivityWithCustomId(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	}
}

func (dispatcher *dispatcher) LogActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, payload string) {
	dispatcher.conductor.ActivityPubIncomingActivityManager().Log(identityId, uniqueIdentifier, timestamp, from, to, content, raw, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubIncomingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().UpdateActivityPubIncomingActivityAtomic(transaction, id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	} else {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().UpdateActivityPubIncomingActivity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateActivityPubIncomingActivityObject(object IObject, activityPubIncomingActivity IActivityPubIncomingActivity) IActivityPubIncomingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().UpdateActivityPubIncomingActivityAtomic(transaction, object.Id(), activityPubIncomingActivity.IdentityId(), activityPubIncomingActivity.UniqueIdentifier(), activityPubIncomingActivity.Timestamp(), activityPubIncomingActivity.From(), activityPubIncomingActivity.To(), activityPubIncomingActivity.Content(), activityPubIncomingActivity.Raw(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	} else {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().UpdateActivityPubIncomingActivity(object.Id(), activityPubIncomingActivity.IdentityId(), activityPubIncomingActivity.UniqueIdentifier(), activityPubIncomingActivity.Timestamp(), activityPubIncomingActivity.From(), activityPubIncomingActivity.To(), activityPubIncomingActivity.Content(), activityPubIncomingActivity.Raw(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateActivityPubIncomingActivityObject(object IObject, activityPubIncomingActivity IActivityPubIncomingActivity) IActivityPubIncomingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().AddOrUpdateActivityPubIncomingActivityObjectAtomic(transaction, object.Id(), activityPubIncomingActivity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	} else {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().AddOrUpdateActivityPubIncomingActivityObject(object.Id(), activityPubIncomingActivity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	}
}

func (dispatcher *dispatcher) RemoveActivityPubIncomingActivity(id int64) IActivityPubIncomingActivity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().RemoveActivityPubIncomingActivityAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	} else {
		if activityPubIncomingActivity, err := dispatcher.conductor.ActivityPubIncomingActivityManager().RemoveActivityPubIncomingActivity(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubIncomingActivity
		}
	}
}

func (dispatcher *dispatcher) ListActivityPubIncomingActivitiesByIdentity(identity IIdentity) IActivityPubIncomingActivityCollection {
	return dispatcher.conductor.ActivityPubIncomingActivityManager().ListActivityPubIncomingActivitiesByIdentity(identity.Id(), 0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ListActivityPubIncomingActivitiesByIdentityId(identityId int64) IActivityPubIncomingActivityCollection {
	return dispatcher.conductor.ActivityPubIncomingActivityManager().ListActivityPubIncomingActivitiesByIdentity(identityId, 0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubIncomingActivityByIdentity(identity IIdentity, iterator ActivityPubIncomingActivityIterator) {
	dispatcher.conductor.ActivityPubIncomingActivityManager().ForEachByIdentity(identity.Id(), iterator)
}

func (dispatcher *dispatcher) ForEachActivityPubIncomingActivityByIdentityId(identityId int64, iterator ActivityPubIncomingActivityIterator) {
	dispatcher.conductor.ActivityPubIncomingActivityManager().ForEachByIdentity(identityId, iterator)
}
