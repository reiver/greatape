package core

import (
	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/system"
)

// noinspection GoSnakeCaseUsage
const ACTIVITY_PUB_MEDIA_MANAGER = "ActivityPubMediaManager"

type activityPubMediaManager struct {
	systemComponent
	cache ICache
}

func newActivityPubMediaManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IActivityPubMediaManager {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize

	manager := &activityPubMediaManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *activityPubMediaManager) Name() string {
	return ACTIVITY_PUB_MEDIA_MANAGER
}

func (manager *activityPubMediaManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *activityPubMediaManager) Load() error {
	return nil
}

func (manager *activityPubMediaManager) Reload() error {
	return manager.Load()
}

func (manager *activityPubMediaManager) OnCacheChanged(callback ActivityPubMediaCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *activityPubMediaManager) Count() int {
	return manager.cache.Size()
}

func (manager *activityPubMediaManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *activityPubMediaManager) ExistsWhich(condition ActivityPubMediaCondition) bool {
	var activityPubMedias ActivityPubMedias
	manager.ForEach(func(activityPubMedia IActivityPubMedia) {
		if condition(activityPubMedia) {
			activityPubMedias = append(activityPubMedias, activityPubMedia)
		}
	})

	return len(activityPubMedias) > 0
}

func (manager *activityPubMediaManager) ListActivityPubMedias(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubMediaCollection {
	return manager.Filter(ActivityPubMediaPassThroughFilter)
}

func (manager *activityPubMediaManager) GetActivityPubMedia(id int64, _ Identity) (IActivityPubMedia, error) {
	if activityPubMedia := manager.Find(id); activityPubMedia == nil {
		return nil, ERROR_ACTIVITY_PUB_MEDIA_NOT_FOUND
	} else {
		return activityPubMedia, nil
	}
}

func (manager *activityPubMediaManager) AddActivityPubMedia(editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) AddActivityPubMediaWithCustomId(id int64, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) AddActivityPubMediaObject(activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) AddActivityPubMediaAtomic(transaction ITransaction, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) AddActivityPubMediaWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) AddActivityPubMediaObjectAtomic(transaction ITransaction, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) Log(source string, editor Identity, payload string) {
}

func (manager *activityPubMediaManager) UpdateActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) UpdateActivityPubMediaObject(id int64, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) UpdateActivityPubMediaAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) UpdateActivityPubMediaObjectAtomic(transaction ITransaction, id int64, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) AddOrUpdateActivityPubMediaObject(id int64, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubMediaObject(id, activityPubMedia, editor)
	} else {
		return manager.AddActivityPubMediaObject(activityPubMedia, editor)
	}
}

func (manager *activityPubMediaManager) AddOrUpdateActivityPubMediaObjectAtomic(transaction ITransaction, id int64, activityPubMedia IActivityPubMedia, editor Identity) (IActivityPubMedia, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubMediaObjectAtomic(transaction, id, activityPubMedia, editor)
	} else {
		return manager.AddActivityPubMediaObjectAtomic(transaction, activityPubMedia, editor)
	}
}

func (manager *activityPubMediaManager) RemoveActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) RemoveActivityPubMediaAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubMedia, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubMediaManager) Find(id int64) IActivityPubMedia {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IActivityPubMedia)
	}
}

func (manager *activityPubMediaManager) ForEach(iterator ActivityPubMediaIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IActivityPubMedia))
	})
}

func (manager *activityPubMediaManager) Filter(predicate ActivityPubMediaFilterPredicate) IActivityPubMediaCollection {
	activityPubMedias := NewActivityPubMedias()
	if predicate == nil {
		return activityPubMedias
	}

	manager.ForEach(func(activityPubMedia IActivityPubMedia) {
		if predicate(activityPubMedia) {
			activityPubMedias.Append(activityPubMedia)
		}
	})

	return activityPubMedias
}

func (manager *activityPubMediaManager) Map(predicate ActivityPubMediaMapPredicate) IActivityPubMediaCollection {
	activityPubMedias := NewActivityPubMedias()
	if predicate == nil {
		return activityPubMedias
	}

	manager.ForEach(func(activityPubMedia IActivityPubMedia) {
		activityPubMedias.Append(predicate(activityPubMedia))
	})

	return activityPubMedias
}
