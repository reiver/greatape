package contracts

import . "github.com/xeronith/diamante/contracts/security"

var ActivityPubObjectPassThroughFilter = func(IActivityPubObject) bool { return true }

type (
	ActivityPubObjects               []IActivityPubObject
	ActivityPubObjectIterator        func(IActivityPubObject)
	ActivityPubObjectCondition       func(IActivityPubObject) bool
	ActivityPubObjectFilterPredicate func(IActivityPubObject) bool
	ActivityPubObjectMapPredicate    func(IActivityPubObject) IActivityPubObject
	ActivityPubObjectCacheCallback   func()

	IActivityPubObject interface {
		// Context returns 'Context' of this 'ActivityPubObject' instance.
		Context() string
		// SetContext sets 'Context' in-memory value of this 'ActivityPubObject' instance.
		// This doesn't affect the persistent data store.
		SetContext(context string)
		// Id returns 'Id' of this 'ActivityPubObject' instance.
		Id() string
		// SetId sets 'Id' in-memory value of this 'ActivityPubObject' instance.
		// This doesn't affect the persistent data store.
		SetId(id string)
		// Type returns 'Type' of this 'ActivityPubObject' instance.
		Type() string
		// SetType sets 'Type' in-memory value of this 'ActivityPubObject' instance.
		// This doesn't affect the persistent data store.
		SetType(type_ string)
		// Actor returns 'Actor' of this 'ActivityPubObject' instance.
		Actor() string
		// SetActor sets 'Actor' in-memory value of this 'ActivityPubObject' instance.
		// This doesn't affect the persistent data store.
		SetActor(actor string)
		// From returns 'From' of this 'ActivityPubObject' instance.
		From() string
		// SetFrom sets 'From' in-memory value of this 'ActivityPubObject' instance.
		// This doesn't affect the persistent data store.
		SetFrom(from string)
		// To returns 'To' of this 'ActivityPubObject' instance.
		To() []string
		// SetTo sets 'To' in-memory value of this 'ActivityPubObject' instance.
		// This doesn't affect the persistent data store.
		SetTo(to []string)
		// InReplyTo returns 'InReplyTo' of this 'ActivityPubObject' instance.
		InReplyTo() string
		// SetInReplyTo sets 'InReplyTo' in-memory value of this 'ActivityPubObject' instance.
		// This doesn't affect the persistent data store.
		SetInReplyTo(inReplyTo string)
		// Content returns 'Content' of this 'ActivityPubObject' instance.
		Content() string
		// SetContent sets 'Content' in-memory value of this 'ActivityPubObject' instance.
		// This doesn't affect the persistent data store.
		SetContent(content string)
		// Published returns 'Published' of this 'ActivityPubObject' instance.
		Published() string
		// SetPublished sets 'Published' in-memory value of this 'ActivityPubObject' instance.
		// This doesn't affect the persistent data store.
		SetPublished(published string)
	}

	IActivityPubObjectCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IActivityPubObject
		Append(activityPubObject IActivityPubObject)
		ForEach(ActivityPubObjectIterator)
		Reverse() IActivityPubObjectCollection
		Array() ActivityPubObjects
	}

	IActivityPubObjectManager interface {
		ISystemComponent
		OnCacheChanged(ActivityPubObjectCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition ActivityPubObjectCondition) bool
		ListActivityPubObjects(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubObjectCollection
		GetActivityPubObject(id int64, editor Identity) (IActivityPubObject, error)
		AddActivityPubObject(editor Identity) (IActivityPubObject, error)
		AddActivityPubObjectWithCustomId(id int64, editor Identity) (IActivityPubObject, error)
		AddActivityPubObjectObject(activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error)
		AddActivityPubObjectAtomic(transaction ITransaction, editor Identity) (IActivityPubObject, error)
		AddActivityPubObjectWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubObject, error)
		AddActivityPubObjectObjectAtomic(transaction ITransaction, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error)
		Log(source string, editor Identity, payload string)
		UpdateActivityPubObject(id int64, editor Identity) (IActivityPubObject, error)
		UpdateActivityPubObjectObject(id int64, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error)
		UpdateActivityPubObjectAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubObject, error)
		UpdateActivityPubObjectObjectAtomic(transaction ITransaction, id int64, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error)
		AddOrUpdateActivityPubObjectObject(id int64, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error)
		AddOrUpdateActivityPubObjectObjectAtomic(transaction ITransaction, id int64, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error)
		RemoveActivityPubObject(id int64, editor Identity) (IActivityPubObject, error)
		RemoveActivityPubObjectAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubObject, error)
		Find(id int64) IActivityPubObject
		ForEach(iterator ActivityPubObjectIterator)
		Filter(predicate ActivityPubObjectFilterPredicate) IActivityPubObjectCollection
		Map(predicate ActivityPubObjectMapPredicate) IActivityPubObjectCollection
	}
)
