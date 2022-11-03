package core

import (
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/system"
	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
)

// noinspection GoSnakeCaseUsage
const ACTIVITY_PUB_PUBLIC_KEY_MANAGER = "ActivityPubPublicKeyManager"

type activityPubPublicKeyManager struct {
	systemComponent
	cache ICache
}

func newActivityPubPublicKeyManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IActivityPubPublicKeyManager {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize

	manager := &activityPubPublicKeyManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *activityPubPublicKeyManager) Name() string {
	return ACTIVITY_PUB_PUBLIC_KEY_MANAGER
}

func (manager *activityPubPublicKeyManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *activityPubPublicKeyManager) Load() error {
	return nil
}

func (manager *activityPubPublicKeyManager) Reload() error {
	return manager.Load()
}

func (manager *activityPubPublicKeyManager) OnCacheChanged(callback ActivityPubPublicKeyCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *activityPubPublicKeyManager) Count() int {
	return manager.cache.Size()
}

func (manager *activityPubPublicKeyManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *activityPubPublicKeyManager) ExistsWhich(condition ActivityPubPublicKeyCondition) bool {
	var activityPubPublicKeys ActivityPubPublicKeys
	manager.ForEach(func(activityPubPublicKey IActivityPubPublicKey) {
		if condition(activityPubPublicKey) {
			activityPubPublicKeys = append(activityPubPublicKeys, activityPubPublicKey)
		}
	})

	return len(activityPubPublicKeys) > 0
}

func (manager *activityPubPublicKeyManager) ListActivityPubPublicKeys(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubPublicKeyCollection {
	return manager.Filter(ActivityPubPublicKeyPassThroughFilter)
}

func (manager *activityPubPublicKeyManager) GetActivityPubPublicKey(id int64, _ Identity) (IActivityPubPublicKey, error) {
	if activityPubPublicKey := manager.Find(id); activityPubPublicKey == nil {
		return nil, ERROR_ACTIVITY_PUB_PUBLIC_KEY_NOT_FOUND
	} else {
		return activityPubPublicKey, nil
	}
}

func (manager *activityPubPublicKeyManager) AddActivityPubPublicKey(editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) AddActivityPubPublicKeyWithCustomId(id int64, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) AddActivityPubPublicKeyObject(activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) AddActivityPubPublicKeyAtomic(transaction ITransaction, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) AddActivityPubPublicKeyWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) AddActivityPubPublicKeyObjectAtomic(transaction ITransaction, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) Log(source string, editor Identity, payload string) {
}

func (manager *activityPubPublicKeyManager) UpdateActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) UpdateActivityPubPublicKeyObject(id int64, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) UpdateActivityPubPublicKeyAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) UpdateActivityPubPublicKeyObjectAtomic(transaction ITransaction, id int64, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) AddOrUpdateActivityPubPublicKeyObject(id int64, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubPublicKeyObject(id, activityPubPublicKey, editor)
	} else {
		return manager.AddActivityPubPublicKeyObject(activityPubPublicKey, editor)
	}
}

func (manager *activityPubPublicKeyManager) AddOrUpdateActivityPubPublicKeyObjectAtomic(transaction ITransaction, id int64, activityPubPublicKey IActivityPubPublicKey, editor Identity) (IActivityPubPublicKey, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubPublicKeyObjectAtomic(transaction, id, activityPubPublicKey, editor)
	} else {
		return manager.AddActivityPubPublicKeyObjectAtomic(transaction, activityPubPublicKey, editor)
	}
}

func (manager *activityPubPublicKeyManager) RemoveActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) RemoveActivityPubPublicKeyAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubPublicKey, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *activityPubPublicKeyManager) Find(id int64) IActivityPubPublicKey {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IActivityPubPublicKey)
	}
}

func (manager *activityPubPublicKeyManager) ForEach(iterator ActivityPubPublicKeyIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IActivityPubPublicKey))
	})
}

func (manager *activityPubPublicKeyManager) Filter(predicate ActivityPubPublicKeyFilterPredicate) IActivityPubPublicKeyCollection {
	activityPubPublicKeys := NewActivityPubPublicKeys()
	if predicate == nil {
		return activityPubPublicKeys
	}

	manager.ForEach(func(activityPubPublicKey IActivityPubPublicKey) {
		if predicate(activityPubPublicKey) {
			activityPubPublicKeys.Append(activityPubPublicKey)
		}
	})

	return activityPubPublicKeys
}

func (manager *activityPubPublicKeyManager) Map(predicate ActivityPubPublicKeyMapPredicate) IActivityPubPublicKeyCollection {
	activityPubPublicKeys := NewActivityPubPublicKeys()
	if predicate == nil {
		return activityPubPublicKeys
	}

	manager.ForEach(func(activityPubPublicKey IActivityPubPublicKey) {
		activityPubPublicKeys.Append(predicate(activityPubPublicKey))
	})

	return activityPubPublicKeys
}
