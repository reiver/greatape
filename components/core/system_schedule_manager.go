package core

import (
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	. "github.com/reiver/greatape/components/model/entity"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/system"
)

// noinspection GoSnakeCaseUsage
const SYSTEM_SCHEDULE_MANAGER = "SystemScheduleManager"

type systemScheduleManager struct {
	systemComponent
	cache ICache
}

func newSystemScheduleManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) ISystemScheduleManager {
	manager := &systemScheduleManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *systemScheduleManager) Name() string {
	return SYSTEM_SCHEDULE_MANAGER
}

func (manager *systemScheduleManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *systemScheduleManager) Load() error {
	systemScheduleEntities, err := repository.SystemSchedules.FetchAll()
	if err != nil {
		return err
	}

	systemSchedules := make(SystemObjectCache)
	for _, systemScheduleEntity := range systemScheduleEntities {
		if systemSchedule, err := NewSystemScheduleFromEntity(systemScheduleEntity); err == nil {
			systemSchedules[systemSchedule.Id()] = systemSchedule
		} else {
			return err
		}
	}

	manager.cache.Load(systemSchedules)
	return nil
}

func (manager *systemScheduleManager) Reload() error {
	return manager.Load()
}

func (manager *systemScheduleManager) OnCacheChanged(callback SystemScheduleCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *systemScheduleManager) Count() int {
	return manager.cache.Size()
}

func (manager *systemScheduleManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *systemScheduleManager) ExistsWhich(condition SystemScheduleCondition) bool {
	var systemSchedules SystemSchedules
	manager.ForEach(func(systemSchedule ISystemSchedule) {
		if condition(systemSchedule) {
			systemSchedules = append(systemSchedules, systemSchedule)
		}
	})

	return len(systemSchedules) > 0
}

func (manager *systemScheduleManager) ListSystemSchedules(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) ISystemScheduleCollection {
	return manager.Filter(SystemSchedulePassThroughFilter)
}

func (manager *systemScheduleManager) GetSystemSchedule(id int64, _ Identity) (ISystemSchedule, error) {
	if systemSchedule := manager.Find(id); systemSchedule == nil {
		return nil, ERROR_SYSTEM_SCHEDULE_NOT_FOUND
	} else {
		return systemSchedule, nil
	}
}

func (manager *systemScheduleManager) AddSystemSchedule(enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(manager.UniqueId(), enabled, config)
	return manager.Apply(systemScheduleEntity, repository.SystemSchedules.Add, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) AddSystemScheduleWithCustomId(id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(id, enabled, config)
	return manager.Apply(systemScheduleEntity, repository.SystemSchedules.Add, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) AddSystemScheduleObject(systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(manager.UniqueId(), systemSchedule.Enabled(), systemSchedule.Config())
	return manager.Apply(systemScheduleEntity, repository.SystemSchedules.Add, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) AddSystemScheduleAtomic(transaction ITransaction, enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(manager.UniqueId(), enabled, config)
	return manager.ApplyAtomic(transaction, systemScheduleEntity, repository.SystemSchedules.AddAtomic, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) AddSystemScheduleWithCustomIdAtomic(id int64, transaction ITransaction, enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(id, enabled, config)
	return manager.ApplyAtomic(transaction, systemScheduleEntity, repository.SystemSchedules.AddAtomic, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) AddSystemScheduleObjectAtomic(transaction ITransaction, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(manager.UniqueId(), systemSchedule.Enabled(), systemSchedule.Config())
	return manager.ApplyAtomic(transaction, systemScheduleEntity, repository.SystemSchedules.AddAtomic, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) Log(enabled bool, config string, source string, editor Identity, payload string) {
	systemSchedulePipeEntity := NewSystemSchedulePipeEntity(manager.UniqueId(), enabled, config, source, editor.Id(), payload)
	repository.Pipe.Insert(systemSchedulePipeEntity)

	systemSchedule, err := NewSystemScheduleFromEntity(systemSchedulePipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(systemSchedule.Id(), systemSchedule)
	}
}

func (manager *systemScheduleManager) UpdateSystemSchedule(id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(id, enabled, config)
	return manager.Apply(systemScheduleEntity, repository.SystemSchedules.Update, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) UpdateSystemScheduleObject(id int64, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(id, systemSchedule.Enabled(), systemSchedule.Config())
	return manager.Apply(systemScheduleEntity, repository.SystemSchedules.Update, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) UpdateSystemScheduleAtomic(transaction ITransaction, id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(id, enabled, config)
	return manager.ApplyAtomic(transaction, systemScheduleEntity, repository.SystemSchedules.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) UpdateSystemScheduleObjectAtomic(transaction ITransaction, id int64, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(id, systemSchedule.Enabled(), systemSchedule.Config())
	return manager.ApplyAtomic(transaction, systemScheduleEntity, repository.SystemSchedules.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *systemScheduleManager) AddOrUpdateSystemScheduleObject(id int64, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error) {
	if manager.Exists(id) {
		return manager.UpdateSystemScheduleObject(id, systemSchedule, editor)
	} else {
		return manager.AddSystemScheduleObject(systemSchedule, editor)
	}
}

func (manager *systemScheduleManager) AddOrUpdateSystemScheduleObjectAtomic(transaction ITransaction, id int64, systemSchedule ISystemSchedule, editor Identity) (ISystemSchedule, error) {
	if manager.Exists(id) {
		return manager.UpdateSystemScheduleObjectAtomic(transaction, id, systemSchedule, editor)
	} else {
		return manager.AddSystemScheduleObjectAtomic(transaction, systemSchedule, editor)
	}
}

func (manager *systemScheduleManager) RemoveSystemSchedule(id int64, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(id, false, "")
	return manager.Apply(systemScheduleEntity, repository.SystemSchedules.Remove, manager.cache.Remove, editor)
}

func (manager *systemScheduleManager) RemoveSystemScheduleAtomic(transaction ITransaction, id int64, editor Identity) (ISystemSchedule, error) {
	systemScheduleEntity := NewSystemScheduleEntity(id, false, "")
	return manager.ApplyAtomic(transaction, systemScheduleEntity, repository.SystemSchedules.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *systemScheduleManager) Apply(systemScheduleEntity ISystemScheduleEntity, repositoryHandler func(ISystemScheduleEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (ISystemSchedule, error) {
	result, err := NewSystemScheduleFromEntity(systemScheduleEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(systemScheduleEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *systemScheduleManager) ApplyAtomic(transaction ITransaction, systemScheduleEntity ISystemScheduleEntity, repositoryHandler func(IRepositoryTransaction, ISystemScheduleEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (ISystemSchedule, error) {
	result, err := NewSystemScheduleFromEntity(systemScheduleEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, systemScheduleEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *systemScheduleManager) Find(id int64) ISystemSchedule {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(ISystemSchedule)
	}
}

func (manager *systemScheduleManager) ForEach(iterator SystemScheduleIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(ISystemSchedule))
	})
}

func (manager *systemScheduleManager) Filter(predicate SystemScheduleFilterPredicate) ISystemScheduleCollection {
	systemSchedules := NewSystemSchedules()
	if predicate == nil {
		return systemSchedules
	}

	manager.ForEach(func(systemSchedule ISystemSchedule) {
		if predicate(systemSchedule) {
			systemSchedules.Append(systemSchedule)
		}
	})

	return systemSchedules
}

func (manager *systemScheduleManager) Map(predicate SystemScheduleMapPredicate) ISystemScheduleCollection {
	systemSchedules := NewSystemSchedules()
	if predicate == nil {
		return systemSchedules
	}

	manager.ForEach(func(systemSchedule ISystemSchedule) {
		systemSchedules.Append(predicate(systemSchedule))
	})

	return systemSchedules
}
