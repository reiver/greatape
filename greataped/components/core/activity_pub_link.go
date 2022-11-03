package core

import (
	"fmt"

	. "github.com/xeronith/diamante/contracts/security"
	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	"rail.town/infrastructure/components/model/repository"
)

type activityPubLink struct {
	href     string
	rel      string
	type_    string
	template string
}

// noinspection GoUnusedExportedFunction
func InitializeActivityPubLink() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewActivityPubLink() (IActivityPubLink, error) {
	instance := &activityPubLink{}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (activityPubLink *activityPubLink) Href() string {
	return activityPubLink.href
}

func (activityPubLink *activityPubLink) SetHref(value string) {
	activityPubLink.href = value
}

func (activityPubLink *activityPubLink) Rel() string {
	return activityPubLink.rel
}

func (activityPubLink *activityPubLink) SetRel(value string) {
	activityPubLink.rel = value
}

func (activityPubLink *activityPubLink) Type() string {
	return activityPubLink.type_
}

func (activityPubLink *activityPubLink) SetType(value string) {
	activityPubLink.type_ = value
}

func (activityPubLink *activityPubLink) Template() string {
	return activityPubLink.template
}

func (activityPubLink *activityPubLink) SetTemplate(value string) {
	activityPubLink.template = value
}

func (activityPubLink *activityPubLink) Validate() error {
	return nil
}

func (activityPubLink *activityPubLink) String() string {
	return fmt.Sprintf("ActivityPubLink (Id: %d)", 0)
}

//------------------------------------------------------------------------------

type activityPubLinks struct {
	collection ActivityPubLinks
}

// NewActivityPubLinks creates an empty collection of 'Activity Pub Link' which is not thread-safe.
func NewActivityPubLinks() IActivityPubLinkCollection {
	return &activityPubLinks{
		collection: make(ActivityPubLinks, 0),
	}
}

func (activityPubLinks *activityPubLinks) Count() int {
	return len(activityPubLinks.collection)
}

func (activityPubLinks *activityPubLinks) IsEmpty() bool {
	return len(activityPubLinks.collection) == 0
}

func (activityPubLinks *activityPubLinks) IsNotEmpty() bool {
	return len(activityPubLinks.collection) > 0
}

func (activityPubLinks *activityPubLinks) HasExactlyOneItem() bool {
	return len(activityPubLinks.collection) == 1
}

func (activityPubLinks *activityPubLinks) HasAtLeastOneItem() bool {
	return len(activityPubLinks.collection) >= 1
}

func (activityPubLinks *activityPubLinks) First() IActivityPubLink {
	return activityPubLinks.collection[0]
}

func (activityPubLinks *activityPubLinks) Append(activityPubLink IActivityPubLink) {
	activityPubLinks.collection = append(activityPubLinks.collection, activityPubLink)
}

func (activityPubLinks *activityPubLinks) ForEach(iterator ActivityPubLinkIterator) {
	if iterator == nil {
		return
	}

	for _, value := range activityPubLinks.collection {
		iterator(value)
	}
}

func (activityPubLinks *activityPubLinks) Array() ActivityPubLinks {
	return activityPubLinks.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) ActivityPubLinkExists(id int64) bool {
	return dispatcher.conductor.ActivityPubLinkManager().Exists(id)
}

func (dispatcher *dispatcher) ActivityPubLinkExistsWhich(condition ActivityPubLinkCondition) bool {
	return dispatcher.conductor.ActivityPubLinkManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListActivityPubLinks() IActivityPubLinkCollection {
	return dispatcher.conductor.ActivityPubLinkManager().ListActivityPubLinks(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachActivityPubLink(iterator ActivityPubLinkIterator) {
	dispatcher.conductor.ActivityPubLinkManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterActivityPubLinks(predicate ActivityPubLinkFilterPredicate) IActivityPubLinkCollection {
	return dispatcher.conductor.ActivityPubLinkManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapActivityPubLinks(predicate ActivityPubLinkMapPredicate) IActivityPubLinkCollection {
	return dispatcher.conductor.ActivityPubLinkManager().Map(predicate)
}

func (dispatcher *dispatcher) GetActivityPubLink(id int64) IActivityPubLink {
	if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().GetActivityPubLink(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return activityPubLink
	}
}

func (dispatcher *dispatcher) AddActivityPubLink() IActivityPubLink {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().AddActivityPubLinkAtomic(transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	} else {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().AddActivityPubLink(dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	}
}

func (dispatcher *dispatcher) AddActivityPubLinkWithCustomId(id int64) IActivityPubLink {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().AddActivityPubLinkWithCustomIdAtomic(id, transaction, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	} else {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().AddActivityPubLinkWithCustomId(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	}
}

func (dispatcher *dispatcher) LogActivityPubLink(source string, payload string) {
	dispatcher.conductor.ActivityPubLinkManager().Log(source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateActivityPubLink(id int64) IActivityPubLink {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().UpdateActivityPubLinkAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	} else {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().UpdateActivityPubLink(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateActivityPubLinkObject(object IObject, activityPubLink IActivityPubLink) IActivityPubLink {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().UpdateActivityPubLinkAtomic(transaction, object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	} else {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().UpdateActivityPubLink(object.Id(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateActivityPubLinkObject(object IObject, activityPubLink IActivityPubLink) IActivityPubLink {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().AddOrUpdateActivityPubLinkObjectAtomic(transaction, object.Id(), activityPubLink, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	} else {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().AddOrUpdateActivityPubLinkObject(object.Id(), activityPubLink, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	}
}

func (dispatcher *dispatcher) RemoveActivityPubLink(id int64) IActivityPubLink {
	transaction := dispatcher.transaction
	if transaction != nil {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().RemoveActivityPubLinkAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	} else {
		if activityPubLink, err := dispatcher.conductor.ActivityPubLinkManager().RemoveActivityPubLink(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return activityPubLink
		}
	}
}
