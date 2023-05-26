package contracts

import . "github.com/xeronith/diamante/contracts/security"

var ActivityPubLinkPassThroughFilter = func(IActivityPubLink) bool { return true }

type (
	ActivityPubLinks               []IActivityPubLink
	ActivityPubLinkIterator        func(IActivityPubLink)
	ActivityPubLinkCondition       func(IActivityPubLink) bool
	ActivityPubLinkFilterPredicate func(IActivityPubLink) bool
	ActivityPubLinkMapPredicate    func(IActivityPubLink) IActivityPubLink
	ActivityPubLinkCacheCallback   func()

	IActivityPubLink interface {
		// Href returns 'Href' of this 'ActivityPubLink' instance.
		Href() string
		// SetHref sets 'Href' in-memory value of this 'ActivityPubLink' instance.
		// This doesn't affect the persistent data store.
		SetHref(href string)
		// Rel returns 'Rel' of this 'ActivityPubLink' instance.
		Rel() string
		// SetRel sets 'Rel' in-memory value of this 'ActivityPubLink' instance.
		// This doesn't affect the persistent data store.
		SetRel(rel string)
		// Type returns 'Type' of this 'ActivityPubLink' instance.
		Type() string
		// SetType sets 'Type' in-memory value of this 'ActivityPubLink' instance.
		// This doesn't affect the persistent data store.
		SetType(type_ string)
		// Template returns 'Template' of this 'ActivityPubLink' instance.
		Template() string
		// SetTemplate sets 'Template' in-memory value of this 'ActivityPubLink' instance.
		// This doesn't affect the persistent data store.
		SetTemplate(template string)
	}

	IActivityPubLinkCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IActivityPubLink
		Append(activityPubLink IActivityPubLink)
		ForEach(ActivityPubLinkIterator)
		Reverse() IActivityPubLinkCollection
		Array() ActivityPubLinks
	}

	IActivityPubLinkManager interface {
		ISystemComponent
		OnCacheChanged(ActivityPubLinkCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition ActivityPubLinkCondition) bool
		ListActivityPubLinks(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubLinkCollection
		GetActivityPubLink(id int64, editor Identity) (IActivityPubLink, error)
		AddActivityPubLink(editor Identity) (IActivityPubLink, error)
		AddActivityPubLinkWithCustomId(id int64, editor Identity) (IActivityPubLink, error)
		AddActivityPubLinkObject(activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error)
		AddActivityPubLinkAtomic(transaction ITransaction, editor Identity) (IActivityPubLink, error)
		AddActivityPubLinkWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubLink, error)
		AddActivityPubLinkObjectAtomic(transaction ITransaction, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error)
		Log(source string, editor Identity, payload string)
		UpdateActivityPubLink(id int64, editor Identity) (IActivityPubLink, error)
		UpdateActivityPubLinkObject(id int64, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error)
		UpdateActivityPubLinkAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubLink, error)
		UpdateActivityPubLinkObjectAtomic(transaction ITransaction, id int64, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error)
		AddOrUpdateActivityPubLinkObject(id int64, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error)
		AddOrUpdateActivityPubLinkObjectAtomic(transaction ITransaction, id int64, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error)
		RemoveActivityPubLink(id int64, editor Identity) (IActivityPubLink, error)
		RemoveActivityPubLinkAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubLink, error)
		Find(id int64) IActivityPubLink
		ForEach(iterator ActivityPubLinkIterator)
		Filter(predicate ActivityPubLinkFilterPredicate) IActivityPubLinkCollection
		Map(predicate ActivityPubLinkMapPredicate) IActivityPubLinkCollection
	}
)
