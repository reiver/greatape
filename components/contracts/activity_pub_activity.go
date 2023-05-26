package contracts

import . "github.com/xeronith/diamante/contracts/security"

var ActivityPubActivityPassThroughFilter = func(IActivityPubActivity) bool { return true }

type (
	ActivityPubActivities              []IActivityPubActivity
	ActivityPubActivityIterator        func(IActivityPubActivity)
	ActivityPubActivityCondition       func(IActivityPubActivity) bool
	ActivityPubActivityFilterPredicate func(IActivityPubActivity) bool
	ActivityPubActivityMapPredicate    func(IActivityPubActivity) IActivityPubActivity
	ActivityPubActivityCacheCallback   func()

	IActivityPubActivity interface {
		// Context returns 'Context' of this 'ActivityPubActivity' instance.
		Context() string
		// SetContext sets 'Context' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetContext(context string)
		// Id returns 'Id' of this 'ActivityPubActivity' instance.
		Id() string
		// SetId sets 'Id' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetId(id string)
		// Type returns 'Type' of this 'ActivityPubActivity' instance.
		Type() string
		// SetType sets 'Type' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetType(type_ string)
		// Actor returns 'Actor' of this 'ActivityPubActivity' instance.
		Actor() string
		// SetActor sets 'Actor' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetActor(actor string)
		// Object returns 'Object' of this 'ActivityPubActivity' instance.
		Object() IActivityPubObject
		// SetObject sets 'Object' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetObject(object IActivityPubObject)
		// From returns 'From' of this 'ActivityPubActivity' instance.
		From() string
		// SetFrom sets 'From' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetFrom(from string)
		// To returns 'To' of this 'ActivityPubActivity' instance.
		To() []string
		// SetTo sets 'To' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetTo(to []string)
		// InReplyTo returns 'InReplyTo' of this 'ActivityPubActivity' instance.
		InReplyTo() string
		// SetInReplyTo sets 'InReplyTo' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetInReplyTo(inReplyTo string)
		// Content returns 'Content' of this 'ActivityPubActivity' instance.
		Content() string
		// SetContent sets 'Content' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetContent(content string)
		// Published returns 'Published' of this 'ActivityPubActivity' instance.
		Published() string
		// SetPublished sets 'Published' in-memory value of this 'ActivityPubActivity' instance.
		// This doesn't affect the persistent data store.
		SetPublished(published string)
	}

	IActivityPubActivityCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IActivityPubActivity
		Append(activityPubActivity IActivityPubActivity)
		ForEach(ActivityPubActivityIterator)
		Reverse() IActivityPubActivityCollection
		Array() ActivityPubActivities
	}

	IActivityPubActivityManager interface {
		ISystemComponent
		OnCacheChanged(ActivityPubActivityCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition ActivityPubActivityCondition) bool
		ListActivityPubActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubActivityCollection
		GetActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error)
		AddActivityPubActivity(editor Identity) (IActivityPubActivity, error)
		AddActivityPubActivityWithCustomId(id int64, editor Identity) (IActivityPubActivity, error)
		AddActivityPubActivityObject(activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error)
		AddActivityPubActivityAtomic(transaction ITransaction, editor Identity) (IActivityPubActivity, error)
		AddActivityPubActivityWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubActivity, error)
		AddActivityPubActivityObjectAtomic(transaction ITransaction, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error)
		Log(source string, editor Identity, payload string)
		UpdateActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error)
		UpdateActivityPubActivityObject(id int64, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error)
		UpdateActivityPubActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubActivity, error)
		UpdateActivityPubActivityObjectAtomic(transaction ITransaction, id int64, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error)
		AddOrUpdateActivityPubActivityObject(id int64, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error)
		AddOrUpdateActivityPubActivityObjectAtomic(transaction ITransaction, id int64, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error)
		RemoveActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error)
		RemoveActivityPubActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubActivity, error)
		Find(id int64) IActivityPubActivity
		ForEach(iterator ActivityPubActivityIterator)
		Filter(predicate ActivityPubActivityFilterPredicate) IActivityPubActivityCollection
		Map(predicate ActivityPubActivityMapPredicate) IActivityPubActivityCollection
	}
)
