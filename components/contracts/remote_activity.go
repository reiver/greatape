package contracts

import . "github.com/xeronith/diamante/contracts/security"

var RemoteActivityPassThroughFilter = func(IRemoteActivity) bool { return true }

type (
	RemoteActivities              []IRemoteActivity
	RemoteActivityIterator        func(IRemoteActivity)
	RemoteActivityCondition       func(IRemoteActivity) bool
	RemoteActivityFilterPredicate func(IRemoteActivity) bool
	RemoteActivityMapPredicate    func(IRemoteActivity) IRemoteActivity
	RemoteActivityCacheCallback   func()

	IRemoteActivity interface {
		IObject
		// EntryPoint returns 'EntryPoint' of this 'RemoteActivity' instance.
		EntryPoint() string
		// UpdateEntryPoint directly updates 'EntryPoint' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateEntryPoint(entryPoint string, editor Identity)
		// UpdateEntryPointAtomic updates 'EntryPoint' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateEntryPointAtomic(transaction ITransaction, entryPoint string, editor Identity)
		// Duration returns 'Duration' of this 'RemoteActivity' instance.
		Duration() int64
		// UpdateDuration directly updates 'Duration' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateDuration(duration int64, editor Identity)
		// UpdateDurationAtomic updates 'Duration' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateDurationAtomic(transaction ITransaction, duration int64, editor Identity)
		// Successful returns 'Successful' of this 'RemoteActivity' instance.
		Successful() bool
		// UpdateSuccessful directly updates 'Successful' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateSuccessful(successful bool, editor Identity)
		// UpdateSuccessfulAtomic updates 'Successful' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateSuccessfulAtomic(transaction ITransaction, successful bool, editor Identity)
		// ErrorMessage returns 'ErrorMessage' of this 'RemoteActivity' instance.
		ErrorMessage() string
		// UpdateErrorMessage directly updates 'ErrorMessage' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateErrorMessage(errorMessage string, editor Identity)
		// UpdateErrorMessageAtomic updates 'ErrorMessage' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateErrorMessageAtomic(transaction ITransaction, errorMessage string, editor Identity)
		// RemoteAddress returns 'RemoteAddress' of this 'RemoteActivity' instance.
		RemoteAddress() string
		// UpdateRemoteAddress directly updates 'RemoteAddress' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateRemoteAddress(remoteAddress string, editor Identity)
		// UpdateRemoteAddressAtomic updates 'RemoteAddress' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateRemoteAddressAtomic(transaction ITransaction, remoteAddress string, editor Identity)
		// UserAgent returns 'UserAgent' of this 'RemoteActivity' instance.
		UserAgent() string
		// UpdateUserAgent directly updates 'UserAgent' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateUserAgent(userAgent string, editor Identity)
		// UpdateUserAgentAtomic updates 'UserAgent' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateUserAgentAtomic(transaction ITransaction, userAgent string, editor Identity)
		// EventType returns 'EventType' of this 'RemoteActivity' instance.
		EventType() uint32
		// UpdateEventType directly updates 'EventType' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateEventType(eventType uint32, editor Identity)
		// UpdateEventTypeAtomic updates 'EventType' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateEventTypeAtomic(transaction ITransaction, eventType uint32, editor Identity)
		// Timestamp returns 'Timestamp' of this 'RemoteActivity' instance.
		Timestamp() int64
		// UpdateTimestamp directly updates 'Timestamp' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateTimestamp(timestamp int64, editor Identity)
		// UpdateTimestampAtomic updates 'Timestamp' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateTimestampAtomic(transaction ITransaction, timestamp int64, editor Identity)
	}

	IRemoteActivityCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IRemoteActivity
		Append(remoteActivity IRemoteActivity)
		ForEach(RemoteActivityIterator)
		Reverse() IRemoteActivityCollection
		Array() RemoteActivities
	}

	IRemoteActivityManager interface {
		ISystemComponent
		OnCacheChanged(RemoteActivityCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition RemoteActivityCondition) bool
		ListRemoteActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IRemoteActivityCollection
		GetRemoteActivity(id int64, editor Identity) (IRemoteActivity, error)
		AddRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		AddRemoteActivityWithCustomId(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		AddRemoteActivityObject(remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error)
		AddRemoteActivityAtomic(transaction ITransaction, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		AddRemoteActivityWithCustomIdAtomic(id int64, transaction ITransaction, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		AddRemoteActivityObjectAtomic(transaction ITransaction, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error)
		Log(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, source string, editor Identity, payload string)
		UpdateRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		UpdateRemoteActivityObject(id int64, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error)
		UpdateRemoteActivityAtomic(transaction ITransaction, id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		UpdateRemoteActivityObjectAtomic(transaction ITransaction, id int64, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error)
		AddOrUpdateRemoteActivityObject(id int64, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error)
		AddOrUpdateRemoteActivityObjectAtomic(transaction ITransaction, id int64, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error)
		RemoveRemoteActivity(id int64, editor Identity) (IRemoteActivity, error)
		RemoveRemoteActivityAtomic(transaction ITransaction, id int64, editor Identity) (IRemoteActivity, error)
		Find(id int64) IRemoteActivity
		ForEach(iterator RemoteActivityIterator)
		Filter(predicate RemoteActivityFilterPredicate) IRemoteActivityCollection
		Map(predicate RemoteActivityMapPredicate) IRemoteActivityCollection
	}
)
