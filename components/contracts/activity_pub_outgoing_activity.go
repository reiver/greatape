package contracts

import . "github.com/xeronith/diamante/contracts/security"

var ActivityPubOutgoingActivityPassThroughFilter = func(IActivityPubOutgoingActivity) bool { return true }

type (
	ActivityPubOutgoingActivities              []IActivityPubOutgoingActivity
	ActivityPubOutgoingActivityIterator        func(IActivityPubOutgoingActivity)
	ActivityPubOutgoingActivityCondition       func(IActivityPubOutgoingActivity) bool
	ActivityPubOutgoingActivityFilterPredicate func(IActivityPubOutgoingActivity) bool
	ActivityPubOutgoingActivityMapPredicate    func(IActivityPubOutgoingActivity) IActivityPubOutgoingActivity
	ActivityPubOutgoingActivityCacheCallback   func()

	IActivityPubOutgoingActivity interface {
		IObject
		// DependenciesAreUnknown scans all dependencies to make sure a valid parent is set for all of them.
		DependenciesAreUnknown() bool
		// IdentityId returns parent 'IdentityId' of this 'ActivityPubOutgoingActivity' instance.
		IdentityId() int64
		// AssertBelongsToIdentity checks whether this 'ActivityPubOutgoingActivity' instance is a child of the specified 'Identity'
		AssertBelongsToIdentity(identity IIdentity)
		// IdentityIsUnknown checks whether a valid parent 'IdentityId' is provided for this 'ActivityPubOutgoingActivity' instance or not.
		IdentityIsUnknown() bool
		// AssertIdentityIsProvided asserts that a valid 'IdentityId' is provided for this 'ActivityPubOutgoingActivity' instance. A panic will occur if the assertion is not valid.
		AssertIdentityIsProvided()
		// AssertIdentity asserts the given 'IdentityId' is in fact the parent of this 'ActivityPubOutgoingActivity' instance. A panic will occur if the assertion is not valid.
		AssertIdentity(identityId int64)
		// UniqueIdentifier returns 'UniqueIdentifier' of this 'ActivityPubOutgoingActivity' instance.
		UniqueIdentifier() string
		// UpdateUniqueIdentifier directly updates 'UniqueIdentifier' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateUniqueIdentifier(uniqueIdentifier string, editor Identity)
		// UpdateUniqueIdentifierAtomic updates 'UniqueIdentifier' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateUniqueIdentifierAtomic(transaction ITransaction, uniqueIdentifier string, editor Identity)
		// Timestamp returns 'Timestamp' of this 'ActivityPubOutgoingActivity' instance.
		Timestamp() int64
		// UpdateTimestamp directly updates 'Timestamp' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateTimestamp(timestamp int64, editor Identity)
		// UpdateTimestampAtomic updates 'Timestamp' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateTimestampAtomic(transaction ITransaction, timestamp int64, editor Identity)
		// From returns 'From' of this 'ActivityPubOutgoingActivity' instance.
		From() string
		// UpdateFrom directly updates 'From' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateFrom(from string, editor Identity)
		// UpdateFromAtomic updates 'From' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateFromAtomic(transaction ITransaction, from string, editor Identity)
		// To returns 'To' of this 'ActivityPubOutgoingActivity' instance.
		To() string
		// UpdateTo directly updates 'To' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateTo(to string, editor Identity)
		// UpdateToAtomic updates 'To' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateToAtomic(transaction ITransaction, to string, editor Identity)
		// Content returns 'Content' of this 'ActivityPubOutgoingActivity' instance.
		Content() string
		// UpdateContent directly updates 'Content' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateContent(content string, editor Identity)
		// UpdateContentAtomic updates 'Content' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateContentAtomic(transaction ITransaction, content string, editor Identity)
		// Raw returns 'Raw' of this 'ActivityPubOutgoingActivity' instance.
		Raw() string
		// UpdateRaw directly updates 'Raw' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateRaw(raw string, editor Identity)
		// UpdateRawAtomic updates 'Raw' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateRawAtomic(transaction ITransaction, raw string, editor Identity)
	}

	IActivityPubOutgoingActivityCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IActivityPubOutgoingActivity
		Append(activityPubOutgoingActivity IActivityPubOutgoingActivity)
		ForEach(ActivityPubOutgoingActivityIterator)
		Array() ActivityPubOutgoingActivities
	}

	IActivityPubOutgoingActivityManager interface {
		ISystemComponent
		OnCacheChanged(ActivityPubOutgoingActivityCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition ActivityPubOutgoingActivityCondition) bool
		ListActivityPubOutgoingActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubOutgoingActivityCollection
		GetActivityPubOutgoingActivity(id int64, editor Identity) (IActivityPubOutgoingActivity, error)
		AddActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		AddActivityPubOutgoingActivityWithCustomId(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		AddActivityPubOutgoingActivityObject(activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error)
		AddActivityPubOutgoingActivityAtomic(transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		AddActivityPubOutgoingActivityWithCustomIdAtomic(id int64, transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		AddActivityPubOutgoingActivityObjectAtomic(transaction ITransaction, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error)
		Log(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor Identity, payload string)
		UpdateActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		UpdateActivityPubOutgoingActivityObject(id int64, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error)
		UpdateActivityPubOutgoingActivityAtomic(transaction ITransaction, id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		UpdateActivityPubOutgoingActivityObjectAtomic(transaction ITransaction, id int64, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error)
		AddOrUpdateActivityPubOutgoingActivityObject(id int64, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error)
		AddOrUpdateActivityPubOutgoingActivityObjectAtomic(transaction ITransaction, id int64, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error)
		RemoveActivityPubOutgoingActivity(id int64, editor Identity) (IActivityPubOutgoingActivity, error)
		RemoveActivityPubOutgoingActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubOutgoingActivity, error)
		Find(id int64) IActivityPubOutgoingActivity
		ForEach(iterator ActivityPubOutgoingActivityIterator)
		Filter(predicate ActivityPubOutgoingActivityFilterPredicate) IActivityPubOutgoingActivityCollection
		Map(predicate ActivityPubOutgoingActivityMapPredicate) IActivityPubOutgoingActivityCollection
		ListActivityPubOutgoingActivitiesByIdentity(identityId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubOutgoingActivityCollection
		ForEachByIdentity(identityId int64, iterator ActivityPubOutgoingActivityIterator)
	}
)
