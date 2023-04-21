package core

import (
	"fmt"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
)

type activityPubFollower struct {
	object
	handle   string
	inbox    string
	subject  string
	activity string
	accepted bool
}

// noinspection GoUnusedExportedFunction
func InitializeActivityPubFollower() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool) (IActivityPubFollower, error) {
	instance := &activityPubFollower{
		object: object{
			id: id,
		},
		handle:   handle,
		inbox:    inbox,
		subject:  subject,
		activity: activity,
		accepted: accepted,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewActivityPubFollowerFromEntity(entity IActivityPubFollowerEntity) (IActivityPubFollower, error) {
	instance := &activityPubFollower{
		object: object{
			id: entity.Id(),
		},
		handle:   entity.Handle(),
		inbox:    entity.Inbox(),
		subject:  entity.Subject(),
		activity: entity.Activity(),
		accepted: entity.Accepted(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (activityPubFollower *activityPubFollower) Handle() string {
	return activityPubFollower.handle
}

func (activityPubFollower *activityPubFollower) UpdateHandle(handle string, editor Identity) {
	if err := repository.ActivityPubFollowers.UpdateHandle(activityPubFollower.id, handle, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubFollower.handle = handle
}

func (activityPubFollower *activityPubFollower) UpdateHandleAtomic(transaction ITransaction, handle string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubFollower.handle = handle
	})

	if err := repository.ActivityPubFollowers.UpdateHandleAtomic(transaction, activityPubFollower.id, handle, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubFollower *activityPubFollower) Inbox() string {
	return activityPubFollower.inbox
}

func (activityPubFollower *activityPubFollower) UpdateInbox(inbox string, editor Identity) {
	if err := repository.ActivityPubFollowers.UpdateInbox(activityPubFollower.id, inbox, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubFollower.inbox = inbox
}

func (activityPubFollower *activityPubFollower) UpdateInboxAtomic(transaction ITransaction, inbox string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubFollower.inbox = inbox
	})

	if err := repository.ActivityPubFollowers.UpdateInboxAtomic(transaction, activityPubFollower.id, inbox, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubFollower *activityPubFollower) Subject() string {
	return activityPubFollower.subject
}

func (activityPubFollower *activityPubFollower) UpdateSubject(subject string, editor Identity) {
	if err := repository.ActivityPubFollowers.UpdateSubject(activityPubFollower.id, subject, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubFollower.subject = subject
}

func (activityPubFollower *activityPubFollower) UpdateSubjectAtomic(transaction ITransaction, subject string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubFollower.subject = subject
	})

	if err := repository.ActivityPubFollowers.UpdateSubjectAtomic(transaction, activityPubFollower.id, subject, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubFollower *activityPubFollower) Activity() string {
	return activityPubFollower.activity
}

func (activityPubFollower *activityPubFollower) UpdateActivity(activity string, editor Identity) {
	if err := repository.ActivityPubFollowers.UpdateActivity(activityPubFollower.id, activity, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubFollower.activity = activity
}

func (activityPubFollower *activityPubFollower) UpdateActivityAtomic(transaction ITransaction, activity string, editor Identity) {
	transaction.OnCommit(func() {
		activityPubFollower.activity = activity
	})

	if err := repository.ActivityPubFollowers.UpdateActivityAtomic(transaction, activityPubFollower.id, activity, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubFollower *activityPubFollower) Accepted() bool {
	return activityPubFollower.accepted
}

func (activityPubFollower *activityPubFollower) UpdateAccepted(accepted bool, editor Identity) {
	if err := repository.ActivityPubFollowers.UpdateAccepted(activityPubFollower.id, accepted, editor.Id()); err != nil {
		panic(err.Error())
	}

	activityPubFollower.accepted = accepted
}

func (activityPubFollower *activityPubFollower) UpdateAcceptedAtomic(transaction ITransaction, accepted bool, editor Identity) {
	transaction.OnCommit(func() {
		activityPubFollower.accepted = accepted
	})

	if err := repository.ActivityPubFollowers.UpdateAcceptedAtomic(transaction, activityPubFollower.id, accepted, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (activityPubFollower *activityPubFollower) Validate() error {
	return nil
}

func (activityPubFollower *activityPubFollower) String() string {
	return fmt.Sprintf("ActivityPubFollower (Id: %d, Handle: %v, Inbox: %v, Subject: %v, Activity: %v, Accepted: %v)", activityPubFollower.Id(), activityPubFollower.Handle(), activityPubFollower.Inbox(), activityPubFollower.Subject(), activityPubFollower.Activity(), activityPubFollower.Accepted())
}

//------------------------------------------------------------------------------

type activityPubFollowers struct {
	collection ActivityPubFollowers
}

// NewActivityPubFollowers creates an empty collection of 'Activity Pub Follower' which is not thread-safe.
func NewActivityPubFollowers() IActivityPubFollowerCollection {
	return &activityPubFollowers{
		collection: make(ActivityPubFollowers, 0),
	}
}

func (activityPubFollowers *activityPubFollowers) Count() int {
	return len(activityPubFollowers.collection)
}

func (activityPubFollowers *activityPubFollowers) IsEmpty() bool {
	return len(activityPubFollowers.collection) == 0
}

func (activityPubFollowers *activityPubFollowers) IsNotEmpty() bool {
	return len(activityPubFollowers.collection) > 0
}

func (activityPubFollowers *activityPubFollowers) HasExactlyOneItem() bool {
	return len(activityPubFollowers.collection) == 1
}

func (activityPubFollowers *activityPubFollowers) HasAtLeastOneItem() bool {
	return len(activityPubFollowers.collection) >= 1
}

func (activityPubFollowers *activityPubFollowers) First() IActivityPubFollower {
	return activityPubFollowers.collection[0]
}

func (activityPubFollowers *activityPubFollowers) Append(activityPubFollower IActivityPubFollower) {
	activityPubFollowers.collection = append(activityPubFollowers.collection, activityPubFollower)
}

func (activityPubFollowers *activityPubFollowers) ForEach(iterator ActivityPubFollowerIterator) {
	if iterator == nil {
		return
	}

	for _, value := range activityPubFollowers.collection {
		iterator(value)
	}
}

func (activityPubFollowers *activityPubFollowers) Array() ActivityPubFollowers {
	return activityPubFollowers.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) ActivityPubFollowerExists(id int64) bool {
	return dispatcher.conductor.ActivityPubFollowerManager().Exists(id)
}

func (dispatcher *dispatcher) ActivityPubFollowerExistsWhich(condition ActivityPubFollowerCondition) bool {
	return dispatcher.conductor.ActivityPubFollowerManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListActivityPubFollowers() IActivityPubFollowerCollection {
	return dispatcher.conductor.ActivityPubFollowerManager().ListActivityPubFollowers(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubFollower(iterator ActivityPubFollowerIterator) {
	dispatcher.conductor.ActivityPubFollowerManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterActivityPubFollowers(predicate ActivityPubFollowerFilterPredicate) IActivityPubFollowerCollection {
	return dispatcher.conductor.ActivityPubFollowerManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapActivityPubFollowers(predicate ActivityPubFollowerMapPredicate) IActivityPubFollowerCollection {
	return dispatcher.conductor.ActivityPubFollowerManager().Map(predicate)
}

func (dispatcher *dispatcher) GetActivityPubFollower(id int64) IActivityPubFollower {
	if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().GetActivityPubFollower(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return activityPubFollower
	}
}

func (dispatcher *dispatcher) AddActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool) IActivityPubFollower {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().AddActivityPubFollowerAtomic(transaction, handle, inbox, subject, activity, accepted, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	} else {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().AddActivityPubFollower(handle, inbox, subject, activity, accepted, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	}
}

func (dispatcher *dispatcher) AddActivityPubFollowerWithCustomId(id int64, handle string, inbox string, subject string, activity string, accepted bool) IActivityPubFollower {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().AddActivityPubFollowerWithCustomIdAtomic(id, transaction, handle, inbox, subject, activity, accepted, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	} else {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().AddActivityPubFollowerWithCustomId(id, handle, inbox, subject, activity, accepted, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	}
}

func (dispatcher *dispatcher) LogActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool, source string, payload string) {
	dispatcher.conductor.ActivityPubFollowerManager().Log(handle, inbox, subject, activity, accepted, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool) IActivityPubFollower {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().UpdateActivityPubFollowerAtomic(transaction, id, handle, inbox, subject, activity, accepted, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	} else {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().UpdateActivityPubFollower(id, handle, inbox, subject, activity, accepted, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateActivityPubFollowerObject(object IObject, activityPubFollower IActivityPubFollower) IActivityPubFollower {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().UpdateActivityPubFollowerAtomic(transaction, object.Id(), activityPubFollower.Handle(), activityPubFollower.Inbox(), activityPubFollower.Subject(), activityPubFollower.Activity(), activityPubFollower.Accepted(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	} else {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().UpdateActivityPubFollower(object.Id(), activityPubFollower.Handle(), activityPubFollower.Inbox(), activityPubFollower.Subject(), activityPubFollower.Activity(), activityPubFollower.Accepted(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateActivityPubFollowerObject(object IObject, activityPubFollower IActivityPubFollower) IActivityPubFollower {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().AddOrUpdateActivityPubFollowerObjectAtomic(transaction, object.Id(), activityPubFollower, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	} else {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().AddOrUpdateActivityPubFollowerObject(object.Id(), activityPubFollower, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	}
}

func (dispatcher *dispatcher) RemoveActivityPubFollower(id int64) IActivityPubFollower {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().RemoveActivityPubFollowerAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	} else {
		if activityPubFollower, err := dispatcher.conductor.ActivityPubFollowerManager().RemoveActivityPubFollower(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubFollower
		}
	}
}
