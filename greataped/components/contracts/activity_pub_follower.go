package contracts

import . "github.com/xeronith/diamante/contracts/security"

var ActivityPubFollowerPassThroughFilter = func(IActivityPubFollower) bool { return true }

type (
	ActivityPubFollowers               []IActivityPubFollower
	ActivityPubFollowerIterator        func(IActivityPubFollower)
	ActivityPubFollowerCondition       func(IActivityPubFollower) bool
	ActivityPubFollowerFilterPredicate func(IActivityPubFollower) bool
	ActivityPubFollowerMapPredicate    func(IActivityPubFollower) IActivityPubFollower
	ActivityPubFollowerCacheCallback   func()

	IActivityPubFollower interface {
		IObject
		// Handle returns 'Handle' of this 'ActivityPubFollower' instance.
		Handle() string
		// UpdateHandle directly updates 'Handle' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateHandle(handle string, editor Identity)
		// UpdateHandleAtomic updates 'Handle' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateHandleAtomic(transaction ITransaction, handle string, editor Identity)
		// Inbox returns 'Inbox' of this 'ActivityPubFollower' instance.
		Inbox() string
		// UpdateInbox directly updates 'Inbox' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateInbox(inbox string, editor Identity)
		// UpdateInboxAtomic updates 'Inbox' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateInboxAtomic(transaction ITransaction, inbox string, editor Identity)
		// Subject returns 'Subject' of this 'ActivityPubFollower' instance.
		Subject() string
		// UpdateSubject directly updates 'Subject' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateSubject(subject string, editor Identity)
		// UpdateSubjectAtomic updates 'Subject' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateSubjectAtomic(transaction ITransaction, subject string, editor Identity)
		// Activity returns 'Activity' of this 'ActivityPubFollower' instance.
		Activity() string
		// UpdateActivity directly updates 'Activity' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateActivity(activity string, editor Identity)
		// UpdateActivityAtomic updates 'Activity' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateActivityAtomic(transaction ITransaction, activity string, editor Identity)
		// Accepted returns 'Accepted' of this 'ActivityPubFollower' instance.
		Accepted() bool
		// UpdateAccepted directly updates 'Accepted' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateAccepted(accepted bool, editor Identity)
		// UpdateAcceptedAtomic updates 'Accepted' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateAcceptedAtomic(transaction ITransaction, accepted bool, editor Identity)
	}

	IActivityPubFollowerCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IActivityPubFollower
		Append(activityPubFollower IActivityPubFollower)
		ForEach(ActivityPubFollowerIterator)
		Array() ActivityPubFollowers
	}

	IActivityPubFollowerManager interface {
		ISystemComponent
		OnCacheChanged(ActivityPubFollowerCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition ActivityPubFollowerCondition) bool
		ListActivityPubFollowers(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubFollowerCollection
		GetActivityPubFollower(id int64, editor Identity) (IActivityPubFollower, error)
		AddActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		AddActivityPubFollowerWithCustomId(id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		AddActivityPubFollowerObject(activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error)
		AddActivityPubFollowerAtomic(transaction ITransaction, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		AddActivityPubFollowerWithCustomIdAtomic(id int64, transaction ITransaction, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		AddActivityPubFollowerObjectAtomic(transaction ITransaction, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error)
		Log(handle string, inbox string, subject string, activity string, accepted bool, source string, editor Identity, payload string)
		UpdateActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		UpdateActivityPubFollowerObject(id int64, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error)
		UpdateActivityPubFollowerAtomic(transaction ITransaction, id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		UpdateActivityPubFollowerObjectAtomic(transaction ITransaction, id int64, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error)
		AddOrUpdateActivityPubFollowerObject(id int64, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error)
		AddOrUpdateActivityPubFollowerObjectAtomic(transaction ITransaction, id int64, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error)
		RemoveActivityPubFollower(id int64, editor Identity) (IActivityPubFollower, error)
		RemoveActivityPubFollowerAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubFollower, error)
		Find(id int64) IActivityPubFollower
		ForEach(iterator ActivityPubFollowerIterator)
		Filter(predicate ActivityPubFollowerFilterPredicate) IActivityPubFollowerCollection
		Map(predicate ActivityPubFollowerMapPredicate) IActivityPubFollowerCollection
	}
)
