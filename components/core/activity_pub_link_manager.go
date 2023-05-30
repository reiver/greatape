package core

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/system"
)

// noinspection GoSnakeCaseUsage
const ACTIVITY_PUB_LINK_MANAGER = "ActivityPubLinkManager"

type activityPubLinkManager struct {
	systemComponent
	cache ICache
}

func newActivityPubLinkManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IActivityPubLinkManager {
	manager := &activityPubLinkManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *activityPubLinkManager) Name() string {
	return ACTIVITY_PUB_LINK_MANAGER
}

func (manager *activityPubLinkManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *activityPubLinkManager) Load() error {
	return nil
}

func (manager *activityPubLinkManager) Reload() error {
	return manager.Load()
}

func (manager *activityPubLinkManager) OnCacheChanged(callback ActivityPubLinkCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *activityPubLinkManager) Count() int {
	return manager.cache.Size()
}

func (manager *activityPubLinkManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *activityPubLinkManager) ExistsWhich(condition ActivityPubLinkCondition) bool {
	var activityPubLinks ActivityPubLinks
	manager.ForEach(func(activityPubLink IActivityPubLink) {
		if condition(activityPubLink) {
			activityPubLinks = append(activityPubLinks, activityPubLink)
		}
	})

	return len(activityPubLinks) > 0
}

func (manager *activityPubLinkManager) ListActivityPubLinks(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubLinkCollection {
	return manager.Filter(ActivityPubLinkPassThroughFilter)
}

func (manager *activityPubLinkManager) GetActivityPubLink(id int64, _ Identity) (IActivityPubLink, error) {
	if activityPubLink := manager.Find(id); activityPubLink == nil {
		return nil, ERROR_ACTIVITY_PUB_LINK_NOT_FOUND
	} else {
		return activityPubLink, nil
	}
}

func (manager *activityPubLinkManager) AddActivityPubLink(editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) AddActivityPubLinkWithCustomId(id int64, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) AddActivityPubLinkObject(activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) AddActivityPubLinkAtomic(transaction ITransaction, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) AddActivityPubLinkWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) AddActivityPubLinkObjectAtomic(transaction ITransaction, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) Log(source string, editor Identity, payload string) {
}

func (manager *activityPubLinkManager) UpdateActivityPubLink(id int64, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) UpdateActivityPubLinkObject(id int64, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) UpdateActivityPubLinkAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) UpdateActivityPubLinkObjectAtomic(transaction ITransaction, id int64, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) AddOrUpdateActivityPubLinkObject(id int64, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubLinkObject(id, activityPubLink, editor)
	} else {
		return manager.AddActivityPubLinkObject(activityPubLink, editor)
	}
}

func (manager *activityPubLinkManager) AddOrUpdateActivityPubLinkObjectAtomic(transaction ITransaction, id int64, activityPubLink IActivityPubLink, editor Identity) (IActivityPubLink, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubLinkObjectAtomic(transaction, id, activityPubLink, editor)
	} else {
		return manager.AddActivityPubLinkObjectAtomic(transaction, activityPubLink, editor)
	}
}

func (manager *activityPubLinkManager) RemoveActivityPubLink(id int64, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) RemoveActivityPubLinkAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubLink, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubLinkManager) Find(id int64) IActivityPubLink {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IActivityPubLink)
	}
}

func (manager *activityPubLinkManager) ForEach(iterator ActivityPubLinkIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IActivityPubLink))
	})
}

func (manager *activityPubLinkManager) Filter(predicate ActivityPubLinkFilterPredicate) IActivityPubLinkCollection {
	activityPubLinks := NewActivityPubLinks()
	if predicate == nil {
		return activityPubLinks
	}

	manager.ForEach(func(activityPubLink IActivityPubLink) {
		if predicate(activityPubLink) {
			activityPubLinks.Append(activityPubLink)
		}
	})

	return activityPubLinks
}

func (manager *activityPubLinkManager) Map(predicate ActivityPubLinkMapPredicate) IActivityPubLinkCollection {
	activityPubLinks := NewActivityPubLinks()
	if predicate == nil {
		return activityPubLinks
	}

	manager.ForEach(func(activityPubLink IActivityPubLink) {
		activityPubLinks.Append(predicate(activityPubLink))
	})

	return activityPubLinks
}
