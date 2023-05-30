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
const ACTIVITY_PUB_OBJECT_MANAGER = "ActivityPubObjectManager"

type activityPubObjectManager struct {
	systemComponent
	cache ICache
}

func newActivityPubObjectManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IActivityPubObjectManager {
	manager := &activityPubObjectManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *activityPubObjectManager) Name() string {
	return ACTIVITY_PUB_OBJECT_MANAGER
}

func (manager *activityPubObjectManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *activityPubObjectManager) Load() error {
	return nil
}

func (manager *activityPubObjectManager) Reload() error {
	return manager.Load()
}

func (manager *activityPubObjectManager) OnCacheChanged(callback ActivityPubObjectCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *activityPubObjectManager) Count() int {
	return manager.cache.Size()
}

func (manager *activityPubObjectManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *activityPubObjectManager) ExistsWhich(condition ActivityPubObjectCondition) bool {
	var activityPubObjects ActivityPubObjects
	manager.ForEach(func(activityPubObject IActivityPubObject) {
		if condition(activityPubObject) {
			activityPubObjects = append(activityPubObjects, activityPubObject)
		}
	})

	return len(activityPubObjects) > 0
}

func (manager *activityPubObjectManager) ListActivityPubObjects(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubObjectCollection {
	return manager.Filter(ActivityPubObjectPassThroughFilter)
}

func (manager *activityPubObjectManager) GetActivityPubObject(id int64, _ Identity) (IActivityPubObject, error) {
	if activityPubObject := manager.Find(id); activityPubObject == nil {
		return nil, ERROR_ACTIVITY_PUB_OBJECT_NOT_FOUND
	} else {
		return activityPubObject, nil
	}
}

func (manager *activityPubObjectManager) AddActivityPubObject(editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) AddActivityPubObjectWithCustomId(id int64, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) AddActivityPubObjectObject(activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) AddActivityPubObjectAtomic(transaction ITransaction, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) AddActivityPubObjectWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) AddActivityPubObjectObjectAtomic(transaction ITransaction, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) Log(source string, editor Identity, payload string) {
}

func (manager *activityPubObjectManager) UpdateActivityPubObject(id int64, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) UpdateActivityPubObjectObject(id int64, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) UpdateActivityPubObjectAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) UpdateActivityPubObjectObjectAtomic(transaction ITransaction, id int64, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) AddOrUpdateActivityPubObjectObject(id int64, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubObjectObject(id, activityPubObject, editor)
	} else {
		return manager.AddActivityPubObjectObject(activityPubObject, editor)
	}
}

func (manager *activityPubObjectManager) AddOrUpdateActivityPubObjectObjectAtomic(transaction ITransaction, id int64, activityPubObject IActivityPubObject, editor Identity) (IActivityPubObject, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubObjectObjectAtomic(transaction, id, activityPubObject, editor)
	} else {
		return manager.AddActivityPubObjectObjectAtomic(transaction, activityPubObject, editor)
	}
}

func (manager *activityPubObjectManager) RemoveActivityPubObject(id int64, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) RemoveActivityPubObjectAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubObject, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubObjectManager) Find(id int64) IActivityPubObject {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IActivityPubObject)
	}
}

func (manager *activityPubObjectManager) ForEach(iterator ActivityPubObjectIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IActivityPubObject))
	})
}

func (manager *activityPubObjectManager) Filter(predicate ActivityPubObjectFilterPredicate) IActivityPubObjectCollection {
	activityPubObjects := NewActivityPubObjects()
	if predicate == nil {
		return activityPubObjects
	}

	manager.ForEach(func(activityPubObject IActivityPubObject) {
		if predicate(activityPubObject) {
			activityPubObjects.Append(activityPubObject)
		}
	})

	return activityPubObjects
}

func (manager *activityPubObjectManager) Map(predicate ActivityPubObjectMapPredicate) IActivityPubObjectCollection {
	activityPubObjects := NewActivityPubObjects()
	if predicate == nil {
		return activityPubObjects
	}

	manager.ForEach(func(activityPubObject IActivityPubObject) {
		activityPubObjects.Append(predicate(activityPubObject))
	})

	return activityPubObjects
}
