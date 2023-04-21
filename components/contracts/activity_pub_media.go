package contracts

import . "github.com/xeronith/diamante/contracts/security"

var ActivityPubMediaPassThroughFilter = func(IActivityPubMedia) bool { return true }

type (
	ActivityPubMedias               []IActivityPubMedia
	ActivityPubMediaIterator        func(IActivityPubMedia)
	ActivityPubMediaCondition       func(IActivityPubMedia) bool
	ActivityPubMediaFilterPredicate func(IActivityPubMedia) bool
	ActivityPubMediaMapPredicate    func(IActivityPubMedia) IActivityPubMedia
	ActivityPubMediaCacheCallback   func()

	IActivityPubMedia interface {
		// MediaType returns 'MediaType' of this 'ActivityPubMedia' instance.
		MediaType() string
		// SetMediaType sets 'MediaType' in-memory value of this 'ActivityPubMedia' instance.
		// This doesn't affect the persistent data store.
		SetMediaType(mediaType string)
		// Type returns 'Type' of this 'ActivityPubMedia' instance.
		Type() string
		// SetType sets 'Type' in-memory value of this 'ActivityPubMedia' instance.
		// This doesn't affect the persistent data store.
		SetType(type_ string)
		// Url returns 'Url' of this 'ActivityPubMedia' instance.
		Url() string
		// SetUrl sets 'Url' in-memory value of this 'ActivityPubMedia' instance.
		// This doesn't affect the persistent data store.
		SetUrl(url string)
		// Width returns 'Width' of this 'ActivityPubMedia' instance.
		Width() int32
		// SetWidth sets 'Width' in-memory value of this 'ActivityPubMedia' instance.
		// This doesn't affect the persistent data store.
		SetWidth(width int32)
		// Height returns 'Height' of this 'ActivityPubMedia' instance.
		Height() int32
		// SetHeight sets 'Height' in-memory value of this 'ActivityPubMedia' instance.
		// This doesn't affect the persistent data store.
		SetHeight(height int32)
	}

	IActivityPubMediaCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IActivityPubMedia
		Append(activityPubMedia IActivityPubMedia)
		ForEach(ActivityPubMediaIterator)
		Array() ActivityPubMedias
	}

	IActivityPubMediaManager interface {
		ISystemComponent
		OnCacheChanged(ActivityPubMediaCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition ActivityPubMediaCondition) bool
		ListActivityPubMedias(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubMediaCollection
		GetActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error)
		AddActivityPubMedia(editor Identity) (IActivityPubMedia, error)
		AddActivityPubMediaWithCustomId(id int64, editor Identity) (IActivityPubMedia, error)
		AddActivityPubMediaObject(activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error)
		AddActivityPubMediaAtomic(transaction ITransaction, editor Identity) (IActivityPubMedia, error)
		AddActivityPubMediaWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubMedia, error)
		AddActivityPubMediaObjectAtomic(transaction ITransaction, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error)
		Log(source string, editor Identity, payload string)
		UpdateActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error)
		UpdateActivityPubMediaObject(id int64, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error)
		UpdateActivityPubMediaAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubMedia, error)
		UpdateActivityPubMediaObjectAtomic(transaction ITransaction, id int64, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error)
		AddOrUpdateActivityPubMediaObject(id int64, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error)
		AddOrUpdateActivityPubMediaObjectAtomic(transaction ITransaction, id int64, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error)
		RemoveActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error)
		RemoveActivityPubMediaAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubMedia, error)
		Find(id int64) IActivityPubMedia
		ForEach(iterator ActivityPubMediaIterator)
		Filter(predicate ActivityPubMediaFilterPredicate) IActivityPubMediaCollection
		Map(predicate ActivityPubMediaMapPredicate) IActivityPubMediaCollection
	}
)
