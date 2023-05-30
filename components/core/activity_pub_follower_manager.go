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
const ACTIVITY_PUB_FOLLOWER_MANAGER = "ActivityPubFollowerManager"

type activityPubFollowerManager struct {
	systemComponent
	cache ICache
}

func newActivityPubFollowerManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IActivityPubFollowerManager {
	manager := &activityPubFollowerManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *activityPubFollowerManager) Name() string {
	return ACTIVITY_PUB_FOLLOWER_MANAGER
}

func (manager *activityPubFollowerManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *activityPubFollowerManager) Load() error {
	activityPubFollowerEntities, err := repository.ActivityPubFollowers.FetchAll()
	if err != nil {
		return err
	}

	activityPubFollowers := make(SystemObjectCache)
	for _, activityPubFollowerEntity := range activityPubFollowerEntities {
		if activityPubFollower, err := NewActivityPubFollowerFromEntity(activityPubFollowerEntity); err == nil {
			activityPubFollowers[activityPubFollower.Id()] = activityPubFollower
		} else {
			return err
		}
	}

	manager.cache.Load(activityPubFollowers)
	return nil
}

func (manager *activityPubFollowerManager) Reload() error {
	return manager.Load()
}

func (manager *activityPubFollowerManager) OnCacheChanged(callback ActivityPubFollowerCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *activityPubFollowerManager) Count() int {
	return manager.cache.Size()
}

func (manager *activityPubFollowerManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *activityPubFollowerManager) ExistsWhich(condition ActivityPubFollowerCondition) bool {
	var activityPubFollowers ActivityPubFollowers
	manager.ForEach(func(activityPubFollower IActivityPubFollower) {
		if condition(activityPubFollower) {
			activityPubFollowers = append(activityPubFollowers, activityPubFollower)
		}
	})

	return len(activityPubFollowers) > 0
}

func (manager *activityPubFollowerManager) ListActivityPubFollowers(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubFollowerCollection {
	return manager.Filter(ActivityPubFollowerPassThroughFilter)
}

func (manager *activityPubFollowerManager) GetActivityPubFollower(id int64, _ Identity) (IActivityPubFollower, error) {
	if activityPubFollower := manager.Find(id); activityPubFollower == nil {
		return nil, ERROR_ACTIVITY_PUB_FOLLOWER_NOT_FOUND
	} else {
		return activityPubFollower, nil
	}
}

func (manager *activityPubFollowerManager) AddActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(manager.UniqueId(), handle, inbox, subject, activity, accepted)
	return manager.Apply(activityPubFollowerEntity, repository.ActivityPubFollowers.Add, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) AddActivityPubFollowerWithCustomId(id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(id, handle, inbox, subject, activity, accepted)
	return manager.Apply(activityPubFollowerEntity, repository.ActivityPubFollowers.Add, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) AddActivityPubFollowerObject(activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(manager.UniqueId(), activityPubFollower.Handle(), activityPubFollower.Inbox(), activityPubFollower.Subject(), activityPubFollower.Activity(), activityPubFollower.Accepted())
	return manager.Apply(activityPubFollowerEntity, repository.ActivityPubFollowers.Add, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) AddActivityPubFollowerAtomic(transaction ITransaction, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(manager.UniqueId(), handle, inbox, subject, activity, accepted)
	return manager.ApplyAtomic(transaction, activityPubFollowerEntity, repository.ActivityPubFollowers.AddAtomic, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) AddActivityPubFollowerWithCustomIdAtomic(id int64, transaction ITransaction, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(id, handle, inbox, subject, activity, accepted)
	return manager.ApplyAtomic(transaction, activityPubFollowerEntity, repository.ActivityPubFollowers.AddAtomic, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) AddActivityPubFollowerObjectAtomic(transaction ITransaction, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(manager.UniqueId(), activityPubFollower.Handle(), activityPubFollower.Inbox(), activityPubFollower.Subject(), activityPubFollower.Activity(), activityPubFollower.Accepted())
	return manager.ApplyAtomic(transaction, activityPubFollowerEntity, repository.ActivityPubFollowers.AddAtomic, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) Log(handle string, inbox string, subject string, activity string, accepted bool, source string, editor Identity, payload string) {
	activityPubFollowerPipeEntity := NewActivityPubFollowerPipeEntity(manager.UniqueId(), handle, inbox, subject, activity, accepted, source, editor.Id(), payload)
	repository.Pipe.Insert(activityPubFollowerPipeEntity)

	activityPubFollower, err := NewActivityPubFollowerFromEntity(activityPubFollowerPipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(activityPubFollower.Id(), activityPubFollower)
	}
}

func (manager *activityPubFollowerManager) UpdateActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(id, handle, inbox, subject, activity, accepted)
	return manager.Apply(activityPubFollowerEntity, repository.ActivityPubFollowers.Update, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) UpdateActivityPubFollowerObject(id int64, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(id, activityPubFollower.Handle(), activityPubFollower.Inbox(), activityPubFollower.Subject(), activityPubFollower.Activity(), activityPubFollower.Accepted())
	return manager.Apply(activityPubFollowerEntity, repository.ActivityPubFollowers.Update, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) UpdateActivityPubFollowerAtomic(transaction ITransaction, id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(id, handle, inbox, subject, activity, accepted)
	return manager.ApplyAtomic(transaction, activityPubFollowerEntity, repository.ActivityPubFollowers.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) UpdateActivityPubFollowerObjectAtomic(transaction ITransaction, id int64, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(id, activityPubFollower.Handle(), activityPubFollower.Inbox(), activityPubFollower.Subject(), activityPubFollower.Activity(), activityPubFollower.Accepted())
	return manager.ApplyAtomic(transaction, activityPubFollowerEntity, repository.ActivityPubFollowers.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *activityPubFollowerManager) AddOrUpdateActivityPubFollowerObject(id int64, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubFollowerObject(id, activityPubFollower, editor)
	} else {
		return manager.AddActivityPubFollowerObject(activityPubFollower, editor)
	}
}

func (manager *activityPubFollowerManager) AddOrUpdateActivityPubFollowerObjectAtomic(transaction ITransaction, id int64, activityPubFollower IActivityPubFollower, editor Identity) (IActivityPubFollower, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubFollowerObjectAtomic(transaction, id, activityPubFollower, editor)
	} else {
		return manager.AddActivityPubFollowerObjectAtomic(transaction, activityPubFollower, editor)
	}
}

func (manager *activityPubFollowerManager) RemoveActivityPubFollower(id int64, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(id, "", "", "", "", false)
	return manager.Apply(activityPubFollowerEntity, repository.ActivityPubFollowers.Remove, manager.cache.Remove, editor)
}

func (manager *activityPubFollowerManager) RemoveActivityPubFollowerAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubFollower, error) {
	activityPubFollowerEntity := NewActivityPubFollowerEntity(id, "", "", "", "", false)
	return manager.ApplyAtomic(transaction, activityPubFollowerEntity, repository.ActivityPubFollowers.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *activityPubFollowerManager) Apply(activityPubFollowerEntity IActivityPubFollowerEntity, repositoryHandler func(IActivityPubFollowerEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IActivityPubFollower, error) {
	result, err := NewActivityPubFollowerFromEntity(activityPubFollowerEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(activityPubFollowerEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *activityPubFollowerManager) ApplyAtomic(transaction ITransaction, activityPubFollowerEntity IActivityPubFollowerEntity, repositoryHandler func(IRepositoryTransaction, IActivityPubFollowerEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IActivityPubFollower, error) {
	result, err := NewActivityPubFollowerFromEntity(activityPubFollowerEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, activityPubFollowerEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *activityPubFollowerManager) Find(id int64) IActivityPubFollower {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IActivityPubFollower)
	}
}

func (manager *activityPubFollowerManager) ForEach(iterator ActivityPubFollowerIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IActivityPubFollower))
	})
}

func (manager *activityPubFollowerManager) Filter(predicate ActivityPubFollowerFilterPredicate) IActivityPubFollowerCollection {
	activityPubFollowers := NewActivityPubFollowers()
	if predicate == nil {
		return activityPubFollowers
	}

	manager.ForEach(func(activityPubFollower IActivityPubFollower) {
		if predicate(activityPubFollower) {
			activityPubFollowers.Append(activityPubFollower)
		}
	})

	return activityPubFollowers
}

func (manager *activityPubFollowerManager) Map(predicate ActivityPubFollowerMapPredicate) IActivityPubFollowerCollection {
	activityPubFollowers := NewActivityPubFollowers()
	if predicate == nil {
		return activityPubFollowers
	}

	manager.ForEach(func(activityPubFollower IActivityPubFollower) {
		activityPubFollowers.Append(predicate(activityPubFollower))
	})

	return activityPubFollowers
}
