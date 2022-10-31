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
	. "rail.town/infrastructure/components/contracts/model"
	. "rail.town/infrastructure/components/model/entity"
	"rail.town/infrastructure/components/model/repository"
)

// noinspection GoSnakeCaseUsage
const ACCESS_CONTROL_MANAGER = "AccessControlManager"

type accessControlManager struct {
	systemComponent
	cache ICache
}

func newAccessControlManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IAccessControlManager {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize

	manager := &accessControlManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *accessControlManager) Name() string {
	return ACCESS_CONTROL_MANAGER
}

func (manager *accessControlManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *accessControlManager) Load() error {
	accessControlEntities, err := repository.AccessControls.FetchAll()
	if err != nil {
		return err
	}

	accessControls := make(SystemObjectCache)
	for _, accessControlEntity := range accessControlEntities {
		if accessControl, err := NewAccessControlFromEntity(accessControlEntity); err == nil {
			accessControls[accessControl.Id()] = accessControl
		} else {
			return err
		}
	}

	manager.cache.Load(accessControls)
	return nil
}

func (manager *accessControlManager) Reload() error {
	return manager.Load()
}

func (manager *accessControlManager) OnCacheChanged(callback AccessControlCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *accessControlManager) Count() int {
	return manager.cache.Size()
}

func (manager *accessControlManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *accessControlManager) ExistsWhich(condition AccessControlCondition) bool {
	var accessControls AccessControls
	manager.ForEach(func(accessControl IAccessControl) {
		if condition(accessControl) {
			accessControls = append(accessControls, accessControl)
		}
	})

	return len(accessControls) > 0
}

func (manager *accessControlManager) ListAccessControls(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IAccessControlCollection {
	return manager.Filter(AccessControlPassThroughFilter)
}

func (manager *accessControlManager) GetAccessControl(id int64, _ Identity) (IAccessControl, error) {
	if accessControl := manager.Find(id); accessControl == nil {
		return nil, ERROR_ACCESS_CONTROL_NOT_FOUND
	} else {
		return accessControl, nil
	}
}

func (manager *accessControlManager) AccessControls() map[uint64]uint64 {
	result := make(map[uint64]uint64)
	manager.ForEach(func(accessControl IAccessControl) {
		result[accessControl.Key()] = accessControl.Value()
	})

	return result
}

func (manager *accessControlManager) AddOrUpdateAccessControl(key uint64, value uint64, editor Identity) error {
	var accessControl IAccessControl
	for _, _accessControl := range manager.Filter(AccessControlPassThroughFilter).Array() {
		if _accessControl.Key() == key {
			accessControl = _accessControl
			break
		}
	}

	if accessControl != nil {
		_, err := manager.UpdateAccessControl(accessControl.Id(), accessControl.Key(), value, editor)
		return err
	} else {
		_, err := manager.AddAccessControl(key, value, editor)
		return err
	}
}

func (manager *accessControlManager) AddAccessControl(key uint64, value uint64, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(manager.UniqueId(), key, value)
	return manager.Apply(accessControlEntity, repository.AccessControls.Add, manager.cache.Put, editor)
}

func (manager *accessControlManager) AddAccessControlWithCustomId(id int64, key uint64, value uint64, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(id, key, value)
	return manager.Apply(accessControlEntity, repository.AccessControls.Add, manager.cache.Put, editor)
}

func (manager *accessControlManager) AddAccessControlObject(accessControl IAccessControl, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(manager.UniqueId(), accessControl.Key(), accessControl.Value())
	return manager.Apply(accessControlEntity, repository.AccessControls.Add, manager.cache.Put, editor)
}

func (manager *accessControlManager) AddAccessControlAtomic(transaction ITransaction, key uint64, value uint64, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(manager.UniqueId(), key, value)
	return manager.ApplyAtomic(transaction, accessControlEntity, repository.AccessControls.AddAtomic, manager.cache.Put, editor)
}

func (manager *accessControlManager) AddAccessControlWithCustomIdAtomic(id int64, transaction ITransaction, key uint64, value uint64, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(id, key, value)
	return manager.ApplyAtomic(transaction, accessControlEntity, repository.AccessControls.AddAtomic, manager.cache.Put, editor)
}

func (manager *accessControlManager) AddAccessControlObjectAtomic(transaction ITransaction, accessControl IAccessControl, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(manager.UniqueId(), accessControl.Key(), accessControl.Value())
	return manager.ApplyAtomic(transaction, accessControlEntity, repository.AccessControls.AddAtomic, manager.cache.Put, editor)
}

func (manager *accessControlManager) Log(key uint64, value uint64, source string, editor Identity, payload string) {
	accessControlPipeEntity := NewAccessControlPipeEntity(manager.UniqueId(), key, value, source, editor.Id(), payload)
	repository.Pipe.Insert(accessControlPipeEntity)

	accessControl, err := NewAccessControlFromEntity(accessControlPipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(accessControl.Id(), accessControl)
	}
}

func (manager *accessControlManager) UpdateAccessControl(id int64, key uint64, value uint64, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(id, key, value)
	return manager.Apply(accessControlEntity, repository.AccessControls.Update, manager.cache.Put, editor)
}

func (manager *accessControlManager) UpdateAccessControlObject(id int64, accessControl IAccessControl, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(id, accessControl.Key(), accessControl.Value())
	return manager.Apply(accessControlEntity, repository.AccessControls.Update, manager.cache.Put, editor)
}

func (manager *accessControlManager) UpdateAccessControlAtomic(transaction ITransaction, id int64, key uint64, value uint64, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(id, key, value)
	return manager.ApplyAtomic(transaction, accessControlEntity, repository.AccessControls.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *accessControlManager) UpdateAccessControlObjectAtomic(transaction ITransaction, id int64, accessControl IAccessControl, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(id, accessControl.Key(), accessControl.Value())
	return manager.ApplyAtomic(transaction, accessControlEntity, repository.AccessControls.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *accessControlManager) AddOrUpdateAccessControlObject(id int64, accessControl IAccessControl, editor Identity) (IAccessControl, error) {
	if manager.Exists(id) {
		return manager.UpdateAccessControlObject(id, accessControl, editor)
	} else {
		return manager.AddAccessControlObject(accessControl, editor)
	}
}

func (manager *accessControlManager) AddOrUpdateAccessControlObjectAtomic(transaction ITransaction, id int64, accessControl IAccessControl, editor Identity) (IAccessControl, error) {
	if manager.Exists(id) {
		return manager.UpdateAccessControlObjectAtomic(transaction, id, accessControl, editor)
	} else {
		return manager.AddAccessControlObjectAtomic(transaction, accessControl, editor)
	}
}

func (manager *accessControlManager) RemoveAccessControl(id int64, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(id, 0, 0)
	return manager.Apply(accessControlEntity, repository.AccessControls.Remove, manager.cache.Remove, editor)
}

func (manager *accessControlManager) RemoveAccessControlAtomic(transaction ITransaction, id int64, editor Identity) (IAccessControl, error) {
	accessControlEntity := NewAccessControlEntity(id, 0, 0)
	return manager.ApplyAtomic(transaction, accessControlEntity, repository.AccessControls.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *accessControlManager) Apply(accessControlEntity IAccessControlEntity, repositoryHandler func(IAccessControlEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IAccessControl, error) {
	result, err := NewAccessControlFromEntity(accessControlEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(accessControlEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *accessControlManager) ApplyAtomic(transaction ITransaction, accessControlEntity IAccessControlEntity, repositoryHandler func(IRepositoryTransaction, IAccessControlEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IAccessControl, error) {
	result, err := NewAccessControlFromEntity(accessControlEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, accessControlEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *accessControlManager) Find(id int64) IAccessControl {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IAccessControl)
	}
}

func (manager *accessControlManager) ForEach(iterator AccessControlIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IAccessControl))
	})
}

func (manager *accessControlManager) Filter(predicate AccessControlFilterPredicate) IAccessControlCollection {
	accessControls := NewAccessControls()
	if predicate == nil {
		return accessControls
	}

	manager.ForEach(func(accessControl IAccessControl) {
		if predicate(accessControl) {
			accessControls.Append(accessControl)
		}
	})

	return accessControls
}

func (manager *accessControlManager) Map(predicate AccessControlMapPredicate) IAccessControlCollection {
	accessControls := NewAccessControls()
	if predicate == nil {
		return accessControls
	}

	manager.ForEach(func(accessControl IAccessControl) {
		accessControls.Append(predicate(accessControl))
	})

	return accessControls
}
