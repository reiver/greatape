package contracts

import . "github.com/xeronith/diamante/contracts/security"

var ActivityPubIncomingActivityPassThroughFilter = func(IActivityPubIncomingActivity) bool { return true }

type (
	ActivityPubIncomingActivities              []IActivityPubIncomingActivity
	ActivityPubIncomingActivityIterator        func(IActivityPubIncomingActivity)
	ActivityPubIncomingActivityCondition       func(IActivityPubIncomingActivity) bool
	ActivityPubIncomingActivityFilterPredicate func(IActivityPubIncomingActivity) bool
	ActivityPubIncomingActivityMapPredicate    func(IActivityPubIncomingActivity) IActivityPubIncomingActivity
	ActivityPubIncomingActivityCacheCallback   func()

	IActivityPubIncomingActivity interface {
		IObject
		// DependenciesAreUnknown scans all dependencies to make sure a valid parent is set for all of them.
		DependenciesAreUnknown() bool
		// IdentityId returns parent 'IdentityId' of this 'ActivityPubIncomingActivity' instance.
		IdentityId() int64
		// AssertBelongsToIdentity checks whether this 'ActivityPubIncomingActivity' instance is a child of the specified 'Identity'
		AssertBelongsToIdentity(identity IIdentity)
		// IdentityIsUnknown checks whether a valid parent 'IdentityId' is provided for this 'ActivityPubIncomingActivity' instance or not.
		IdentityIsUnknown() bool
		// AssertIdentityIsProvided asserts that a valid 'IdentityId' is provided for this 'ActivityPubIncomingActivity' instance. A panic will occur if the assertion is not valid.
		AssertIdentityIsProvided()
		// AssertIdentity asserts the given 'IdentityId' is in fact the parent of this 'ActivityPubIncomingActivity' instance. A panic will occur if the assertion is not valid.
		AssertIdentity(identityId int64)
		// UniqueIdentifier returns 'UniqueIdentifier' of this 'ActivityPubIncomingActivity' instance.
		UniqueIdentifier() string
		// UpdateUniqueIdentifier directly updates 'UniqueIdentifier' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateUniqueIdentifier(uniqueIdentifier string, editor Identity)
		// UpdateUniqueIdentifierAtomic updates 'UniqueIdentifier' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateUniqueIdentifierAtomic(transaction ITransaction, uniqueIdentifier string, editor Identity)
		// Timestamp returns 'Timestamp' of this 'ActivityPubIncomingActivity' instance.
		Timestamp() int64
		// UpdateTimestamp directly updates 'Timestamp' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateTimestamp(timestamp int64, editor Identity)
		// UpdateTimestampAtomic updates 'Timestamp' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateTimestampAtomic(transaction ITransaction, timestamp int64, editor Identity)
		// From returns 'From' of this 'ActivityPubIncomingActivity' instance.
		From() string
		// UpdateFrom directly updates 'From' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateFrom(from string, editor Identity)
		// UpdateFromAtomic updates 'From' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateFromAtomic(transaction ITransaction, from string, editor Identity)
		// To returns 'To' of this 'ActivityPubIncomingActivity' instance.
		To() string
		// UpdateTo directly updates 'To' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateTo(to string, editor Identity)
		// UpdateToAtomic updates 'To' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateToAtomic(transaction ITransaction, to string, editor Identity)
		// Content returns 'Content' of this 'ActivityPubIncomingActivity' instance.
		Content() string
		// UpdateContent directly updates 'Content' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateContent(content string, editor Identity)
		// UpdateContentAtomic updates 'Content' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateContentAtomic(transaction ITransaction, content string, editor Identity)
		// Raw returns 'Raw' of this 'ActivityPubIncomingActivity' instance.
		Raw() string
		// UpdateRaw directly updates 'Raw' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateRaw(raw string, editor Identity)
		// UpdateRawAtomic updates 'Raw' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateRawAtomic(transaction ITransaction, raw string, editor Identity)
	}

	IActivityPubIncomingActivityCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IActivityPubIncomingActivity
		Append(activityPubIncomingActivity IActivityPubIncomingActivity)
		ForEach(ActivityPubIncomingActivityIterator)
		Array() ActivityPubIncomingActivities
	}

	IActivityPubIncomingActivityManager interface {
		ISystemComponent
		OnCacheChanged(ActivityPubIncomingActivityCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition ActivityPubIncomingActivityCondition) bool
		ListActivityPubIncomingActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubIncomingActivityCollection
		GetActivityPubIncomingActivity(id int64, editor Identity) (IActivityPubIncomingActivity, error)
		AddActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		AddActivityPubIncomingActivityWithCustomId(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		AddActivityPubIncomingActivityObject(activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error)
		AddActivityPubIncomingActivityAtomic(transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		AddActivityPubIncomingActivityWithCustomIdAtomic(id int64, transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		AddActivityPubIncomingActivityObjectAtomic(transaction ITransaction, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error)
		Log(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor Identity, payload string)
		UpdateActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		UpdateActivityPubIncomingActivityObject(id int64, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error)
		UpdateActivityPubIncomingActivityAtomic(transaction ITransaction, id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		UpdateActivityPubIncomingActivityObjectAtomic(transaction ITransaction, id int64, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error)
		AddOrUpdateActivityPubIncomingActivityObject(id int64, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error)
		AddOrUpdateActivityPubIncomingActivityObjectAtomic(transaction ITransaction, id int64, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error)
		RemoveActivityPubIncomingActivity(id int64, editor Identity) (IActivityPubIncomingActivity, error)
		RemoveActivityPubIncomingActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubIncomingActivity, error)
		Find(id int64) IActivityPubIncomingActivity
		ForEach(iterator ActivityPubIncomingActivityIterator)
		Filter(predicate ActivityPubIncomingActivityFilterPredicate) IActivityPubIncomingActivityCollection
		Map(predicate ActivityPubIncomingActivityMapPredicate) IActivityPubIncomingActivityCollection
		ListActivityPubIncomingActivitiesByIdentity(identityId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubIncomingActivityCollection
		ForEachByIdentity(identityId int64, iterator ActivityPubIncomingActivityIterator)
	}
)
