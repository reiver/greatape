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
const ACTIVITY_PUB_ACTIVITY_MANAGER = "ActivityPubActivityManager"

type activityPubActivityManager struct {
	systemComponent
	cache ICache
}

func newActivityPubActivityManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IActivityPubActivityManager {
	manager := &activityPubActivityManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *activityPubActivityManager) Name() string {
	return ACTIVITY_PUB_ACTIVITY_MANAGER
}

func (manager *activityPubActivityManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *activityPubActivityManager) Load() error {
	return nil
}

func (manager *activityPubActivityManager) Reload() error {
	return manager.Load()
}

func (manager *activityPubActivityManager) OnCacheChanged(callback ActivityPubActivityCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *activityPubActivityManager) Count() int {
	return manager.cache.Size()
}

func (manager *activityPubActivityManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *activityPubActivityManager) ExistsWhich(condition ActivityPubActivityCondition) bool {
	var activityPubActivities ActivityPubActivities
	manager.ForEach(func(activityPubActivity IActivityPubActivity) {
		if condition(activityPubActivity) {
			activityPubActivities = append(activityPubActivities, activityPubActivity)
		}
	})

	return len(activityPubActivities) > 0
}

func (manager *activityPubActivityManager) ListActivityPubActivities(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubActivityCollection {
	return manager.Filter(ActivityPubActivityPassThroughFilter)
}

func (manager *activityPubActivityManager) GetActivityPubActivity(id int64, _ Identity) (IActivityPubActivity, error) {
	if activityPubActivity := manager.Find(id); activityPubActivity == nil {
		return nil, ERROR_ACTIVITY_PUB_ACTIVITY_NOT_FOUND
	} else {
		return activityPubActivity, nil
	}
}

func (manager *activityPubActivityManager) AddActivityPubActivity(editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) AddActivityPubActivityWithCustomId(id int64, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) AddActivityPubActivityObject(activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) AddActivityPubActivityAtomic(transaction ITransaction, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) AddActivityPubActivityWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) AddActivityPubActivityObjectAtomic(transaction ITransaction, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) Log(source string, editor Identity, payload string) {
}

func (manager *activityPubActivityManager) UpdateActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) UpdateActivityPubActivityObject(id int64, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) UpdateActivityPubActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) UpdateActivityPubActivityObjectAtomic(transaction ITransaction, id int64, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) AddOrUpdateActivityPubActivityObject(id int64, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubActivityObject(id, activityPubActivity, editor)
	} else {
		return manager.AddActivityPubActivityObject(activityPubActivity, editor)
	}
}

func (manager *activityPubActivityManager) AddOrUpdateActivityPubActivityObjectAtomic(transaction ITransaction, id int64, activityPubActivity IActivityPubActivity, editor Identity) (IActivityPubActivity, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubActivityObjectAtomic(transaction, id, activityPubActivity, editor)
	} else {
		return manager.AddActivityPubActivityObjectAtomic(transaction, activityPubActivity, editor)
	}
}

func (manager *activityPubActivityManager) RemoveActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) RemoveActivityPubActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubActivity, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubActivityManager) Find(id int64) IActivityPubActivity {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IActivityPubActivity)
	}
}

func (manager *activityPubActivityManager) ForEach(iterator ActivityPubActivityIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IActivityPubActivity))
	})
}

func (manager *activityPubActivityManager) Filter(predicate ActivityPubActivityFilterPredicate) IActivityPubActivityCollection {
	activityPubActivities := NewActivityPubActivities()
	if predicate == nil {
		return activityPubActivities
	}

	manager.ForEach(func(activityPubActivity IActivityPubActivity) {
		if predicate(activityPubActivity) {
			activityPubActivities.Append(activityPubActivity)
		}
	})

	return activityPubActivities
}

func (manager *activityPubActivityManager) Map(predicate ActivityPubActivityMapPredicate) IActivityPubActivityCollection {
	activityPubActivities := NewActivityPubActivities()
	if predicate == nil {
		return activityPubActivities
	}

	manager.ForEach(func(activityPubActivity IActivityPubActivity) {
		activityPubActivities.Append(predicate(activityPubActivity))
	})

	return activityPubActivities
}
