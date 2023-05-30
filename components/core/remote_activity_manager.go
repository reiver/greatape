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
const REMOTE_ACTIVITY_MANAGER = "RemoteActivityManager"

type remoteActivityManager struct {
	systemComponent
	cache ICache
}

func newRemoteActivityManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IRemoteActivityManager {
	manager := &remoteActivityManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *remoteActivityManager) Name() string {
	return REMOTE_ACTIVITY_MANAGER
}

func (manager *remoteActivityManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *remoteActivityManager) Load() error {
	return nil
}

func (manager *remoteActivityManager) Reload() error {
	return manager.Load()
}

func (manager *remoteActivityManager) OnCacheChanged(callback RemoteActivityCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *remoteActivityManager) Count() int {
	return manager.cache.Size()
}

func (manager *remoteActivityManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *remoteActivityManager) ExistsWhich(condition RemoteActivityCondition) bool {
	var remoteActivities RemoteActivities
	manager.ForEach(func(remoteActivity IRemoteActivity) {
		if condition(remoteActivity) {
			remoteActivities = append(remoteActivities, remoteActivity)
		}
	})

	return len(remoteActivities) > 0
}

func (manager *remoteActivityManager) ListRemoteActivities(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IRemoteActivityCollection {
	return manager.Filter(RemoteActivityPassThroughFilter)
}

func (manager *remoteActivityManager) GetRemoteActivity(id int64, _ Identity) (IRemoteActivity, error) {
	if remoteActivity := manager.Find(id); remoteActivity == nil {
		return nil, ERROR_REMOTE_ACTIVITY_NOT_FOUND
	} else {
		return remoteActivity, nil
	}
}

func (manager *remoteActivityManager) AddRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(manager.UniqueId(), entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp)
	return manager.Apply(remoteActivityEntity, repository.RemoteActivities.Add, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) AddRemoteActivityWithCustomId(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp)
	return manager.Apply(remoteActivityEntity, repository.RemoteActivities.Add, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) AddRemoteActivityObject(remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(manager.UniqueId(), remoteActivity.EntryPoint(), remoteActivity.Duration(), remoteActivity.Successful(), remoteActivity.ErrorMessage(), remoteActivity.RemoteAddress(), remoteActivity.UserAgent(), remoteActivity.EventType(), remoteActivity.Timestamp())
	return manager.Apply(remoteActivityEntity, repository.RemoteActivities.Add, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) AddRemoteActivityAtomic(transaction ITransaction, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(manager.UniqueId(), entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp)
	return manager.ApplyAtomic(transaction, remoteActivityEntity, repository.RemoteActivities.AddAtomic, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) AddRemoteActivityWithCustomIdAtomic(id int64, transaction ITransaction, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp)
	return manager.ApplyAtomic(transaction, remoteActivityEntity, repository.RemoteActivities.AddAtomic, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) AddRemoteActivityObjectAtomic(transaction ITransaction, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(manager.UniqueId(), remoteActivity.EntryPoint(), remoteActivity.Duration(), remoteActivity.Successful(), remoteActivity.ErrorMessage(), remoteActivity.RemoteAddress(), remoteActivity.UserAgent(), remoteActivity.EventType(), remoteActivity.Timestamp())
	return manager.ApplyAtomic(transaction, remoteActivityEntity, repository.RemoteActivities.AddAtomic, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) Log(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, source string, editor Identity, payload string) {
	remoteActivityPipeEntity := NewRemoteActivityPipeEntity(manager.UniqueId(), entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, source, editor.Id(), payload)
	repository.Pipe.Insert(remoteActivityPipeEntity)

	remoteActivity, err := NewRemoteActivityFromEntity(remoteActivityPipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(remoteActivity.Id(), remoteActivity)
	}
}

func (manager *remoteActivityManager) UpdateRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp)
	return manager.Apply(remoteActivityEntity, repository.RemoteActivities.Update, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) UpdateRemoteActivityObject(id int64, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(id, remoteActivity.EntryPoint(), remoteActivity.Duration(), remoteActivity.Successful(), remoteActivity.ErrorMessage(), remoteActivity.RemoteAddress(), remoteActivity.UserAgent(), remoteActivity.EventType(), remoteActivity.Timestamp())
	return manager.Apply(remoteActivityEntity, repository.RemoteActivities.Update, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) UpdateRemoteActivityAtomic(transaction ITransaction, id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp)
	return manager.ApplyAtomic(transaction, remoteActivityEntity, repository.RemoteActivities.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) UpdateRemoteActivityObjectAtomic(transaction ITransaction, id int64, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(id, remoteActivity.EntryPoint(), remoteActivity.Duration(), remoteActivity.Successful(), remoteActivity.ErrorMessage(), remoteActivity.RemoteAddress(), remoteActivity.UserAgent(), remoteActivity.EventType(), remoteActivity.Timestamp())
	return manager.ApplyAtomic(transaction, remoteActivityEntity, repository.RemoteActivities.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *remoteActivityManager) AddOrUpdateRemoteActivityObject(id int64, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error) {
	if manager.Exists(id) {
		return manager.UpdateRemoteActivityObject(id, remoteActivity, editor)
	} else {
		return manager.AddRemoteActivityObject(remoteActivity, editor)
	}
}

func (manager *remoteActivityManager) AddOrUpdateRemoteActivityObjectAtomic(transaction ITransaction, id int64, remoteActivity IRemoteActivity, editor Identity) (IRemoteActivity, error) {
	if manager.Exists(id) {
		return manager.UpdateRemoteActivityObjectAtomic(transaction, id, remoteActivity, editor)
	} else {
		return manager.AddRemoteActivityObjectAtomic(transaction, remoteActivity, editor)
	}
}

func (manager *remoteActivityManager) RemoveRemoteActivity(id int64, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(id, "", 0, false, "", "", "", 0, 0)
	return manager.Apply(remoteActivityEntity, repository.RemoteActivities.Remove, manager.cache.Remove, editor)
}

func (manager *remoteActivityManager) RemoveRemoteActivityAtomic(transaction ITransaction, id int64, editor Identity) (IRemoteActivity, error) {
	remoteActivityEntity := NewRemoteActivityEntity(id, "", 0, false, "", "", "", 0, 0)
	return manager.ApplyAtomic(transaction, remoteActivityEntity, repository.RemoteActivities.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *remoteActivityManager) Apply(remoteActivityEntity IRemoteActivityEntity, repositoryHandler func(IRemoteActivityEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IRemoteActivity, error) {
	result, err := NewRemoteActivityFromEntity(remoteActivityEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(remoteActivityEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *remoteActivityManager) ApplyAtomic(transaction ITransaction, remoteActivityEntity IRemoteActivityEntity, repositoryHandler func(IRepositoryTransaction, IRemoteActivityEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IRemoteActivity, error) {
	result, err := NewRemoteActivityFromEntity(remoteActivityEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, remoteActivityEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *remoteActivityManager) Find(id int64) IRemoteActivity {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IRemoteActivity)
	}
}

func (manager *remoteActivityManager) ForEach(iterator RemoteActivityIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IRemoteActivity))
	})
}

func (manager *remoteActivityManager) Filter(predicate RemoteActivityFilterPredicate) IRemoteActivityCollection {
	remoteActivities := NewRemoteActivities()
	if predicate == nil {
		return remoteActivities
	}

	manager.ForEach(func(remoteActivity IRemoteActivity) {
		if predicate(remoteActivity) {
			remoteActivities.Append(remoteActivity)
		}
	})

	return remoteActivities
}

func (manager *remoteActivityManager) Map(predicate RemoteActivityMapPredicate) IRemoteActivityCollection {
	remoteActivities := NewRemoteActivities()
	if predicate == nil {
		return remoteActivities
	}

	manager.ForEach(func(remoteActivity IRemoteActivity) {
		remoteActivities.Append(predicate(remoteActivity))
	})

	return remoteActivities
}
