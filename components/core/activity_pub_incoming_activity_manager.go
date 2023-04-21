package core

import (
	"github.com/reiver/greatape/app/validators"
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
const ACTIVITY_PUB_INCOMING_ACTIVITY_MANAGER = "ActivityPubIncomingActivityManager"

//lint:ignore U1000 GoUnused
type activityPubIncomingActivityManager struct {
	systemComponent
	cache ICache

	//Dependencies
	identityManager IIdentityManager
}

func newActivityPubIncomingActivityManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IActivityPubIncomingActivityManager {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize

	manager := &activityPubIncomingActivityManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *activityPubIncomingActivityManager) Name() string {
	return ACTIVITY_PUB_INCOMING_ACTIVITY_MANAGER
}

func (manager *activityPubIncomingActivityManager) ResolveDependencies(dependencies ...ISystemComponent) error {
	if len(dependencies) == 0 {
		return nil
	}

	var (
		identityManager IIdentityManager
	)

	for _, _dependency := range dependencies {
		if false {
		} else if dependency, ok := _dependency.(IIdentityManager); ok {
			identityManager = dependency
		}
	}

	if // noinspection GoBoolExpressions
	false || identityManager == nil {
		return ERROR_UNRESOLVED_DEPENDENCIES
	}

	return nil
}

func (manager *activityPubIncomingActivityManager) Load() error {
	activityPubIncomingActivityEntities, err := repository.ActivityPubIncomingActivities.FetchAll()
	if err != nil {
		return err
	}

	activityPubIncomingActivities := make(SystemObjectCache)
	for _, activityPubIncomingActivityEntity := range activityPubIncomingActivityEntities {
		if activityPubIncomingActivity, err := NewActivityPubIncomingActivityFromEntity(activityPubIncomingActivityEntity); err == nil {
			activityPubIncomingActivities[activityPubIncomingActivity.Id()] = activityPubIncomingActivity
		} else {
			return err
		}
	}

	manager.cache.Load(activityPubIncomingActivities)
	return nil
}

func (manager *activityPubIncomingActivityManager) Reload() error {
	return manager.Load()
}

func (manager *activityPubIncomingActivityManager) OnCacheChanged(callback ActivityPubIncomingActivityCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *activityPubIncomingActivityManager) Count() int {
	return manager.cache.Size()
}

func (manager *activityPubIncomingActivityManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *activityPubIncomingActivityManager) ExistsWhich(condition ActivityPubIncomingActivityCondition) bool {
	var activityPubIncomingActivities ActivityPubIncomingActivities
	manager.ForEach(func(activityPubIncomingActivity IActivityPubIncomingActivity) {
		if condition(activityPubIncomingActivity) {
			activityPubIncomingActivities = append(activityPubIncomingActivities, activityPubIncomingActivity)
		}
	})

	return len(activityPubIncomingActivities) > 0
}

func (manager *activityPubIncomingActivityManager) ListActivityPubIncomingActivities(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubIncomingActivityCollection {
	return manager.Filter(ActivityPubIncomingActivityPassThroughFilter)
}

func (manager *activityPubIncomingActivityManager) GetActivityPubIncomingActivity(id int64, _ Identity) (IActivityPubIncomingActivity, error) {
	if activityPubIncomingActivity := manager.Find(id); activityPubIncomingActivity == nil {
		return nil, ERROR_ACTIVITY_PUB_INCOMING_ACTIVITY_NOT_FOUND
	} else {
		return activityPubIncomingActivity, nil
	}
}

func (manager *activityPubIncomingActivityManager) AddActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(manager.UniqueId(), identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.Apply(activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.Add, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) AddActivityPubIncomingActivityWithCustomId(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.Apply(activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.Add, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) AddActivityPubIncomingActivityObject(activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(manager.UniqueId(), activityPubIncomingActivity.IdentityId(), activityPubIncomingActivity.UniqueIdentifier(), activityPubIncomingActivity.Timestamp(), activityPubIncomingActivity.From(), activityPubIncomingActivity.To(), activityPubIncomingActivity.Content(), activityPubIncomingActivity.Raw())
	return manager.Apply(activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.Add, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) AddActivityPubIncomingActivityAtomic(transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(manager.UniqueId(), identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.ApplyAtomic(transaction, activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.AddAtomic, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) AddActivityPubIncomingActivityWithCustomIdAtomic(id int64, transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.ApplyAtomic(transaction, activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.AddAtomic, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) AddActivityPubIncomingActivityObjectAtomic(transaction ITransaction, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(manager.UniqueId(), activityPubIncomingActivity.IdentityId(), activityPubIncomingActivity.UniqueIdentifier(), activityPubIncomingActivity.Timestamp(), activityPubIncomingActivity.From(), activityPubIncomingActivity.To(), activityPubIncomingActivity.Content(), activityPubIncomingActivity.Raw())
	return manager.ApplyAtomic(transaction, activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.AddAtomic, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) Log(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor Identity, payload string) {
	activityPubIncomingActivityPipeEntity := NewActivityPubIncomingActivityPipeEntity(manager.UniqueId(), identityId, uniqueIdentifier, timestamp, from, to, content, raw, source, editor.Id(), payload)
	repository.Pipe.Insert(activityPubIncomingActivityPipeEntity)

	activityPubIncomingActivity, err := NewActivityPubIncomingActivityFromEntity(activityPubIncomingActivityPipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(activityPubIncomingActivity.Id(), activityPubIncomingActivity)
	}
}

func (manager *activityPubIncomingActivityManager) UpdateActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.Apply(activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.Update, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) UpdateActivityPubIncomingActivityObject(id int64, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(id, activityPubIncomingActivity.Id(), activityPubIncomingActivity.UniqueIdentifier(), activityPubIncomingActivity.Timestamp(), activityPubIncomingActivity.From(), activityPubIncomingActivity.To(), activityPubIncomingActivity.Content(), activityPubIncomingActivity.Raw())
	return manager.Apply(activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.Update, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) UpdateActivityPubIncomingActivityAtomic(transaction ITransaction, id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.ApplyAtomic(transaction, activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) UpdateActivityPubIncomingActivityObjectAtomic(transaction ITransaction, id int64, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(id, activityPubIncomingActivity.Id(), activityPubIncomingActivity.UniqueIdentifier(), activityPubIncomingActivity.Timestamp(), activityPubIncomingActivity.From(), activityPubIncomingActivity.To(), activityPubIncomingActivity.Content(), activityPubIncomingActivity.Raw())
	return manager.ApplyAtomic(transaction, activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *activityPubIncomingActivityManager) AddOrUpdateActivityPubIncomingActivityObject(id int64, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubIncomingActivityObject(id, activityPubIncomingActivity, editor)
	} else {
		return manager.AddActivityPubIncomingActivityObject(activityPubIncomingActivity, editor)
	}
}

func (manager *activityPubIncomingActivityManager) AddOrUpdateActivityPubIncomingActivityObjectAtomic(transaction ITransaction, id int64, activityPubIncomingActivity IActivityPubIncomingActivity, editor Identity) (IActivityPubIncomingActivity, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubIncomingActivityObjectAtomic(transaction, id, activityPubIncomingActivity, editor)
	} else {
		return manager.AddActivityPubIncomingActivityObjectAtomic(transaction, activityPubIncomingActivity, editor)
	}
}

func (manager *activityPubIncomingActivityManager) RemoveActivityPubIncomingActivity(id int64, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(id, 0, "", 0, "", "", "", "")
	return manager.Apply(activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.Remove, manager.cache.Remove, editor)
}

func (manager *activityPubIncomingActivityManager) RemoveActivityPubIncomingActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubIncomingActivity, error) {
	activityPubIncomingActivityEntity := NewActivityPubIncomingActivityEntity(id, 0, "", 0, "", "", "", "")
	return manager.ApplyAtomic(transaction, activityPubIncomingActivityEntity, repository.ActivityPubIncomingActivities.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *activityPubIncomingActivityManager) Apply(activityPubIncomingActivityEntity IActivityPubIncomingActivityEntity, repositoryHandler func(IActivityPubIncomingActivityEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IActivityPubIncomingActivity, error) {
	result, err := NewActivityPubIncomingActivityFromEntity(activityPubIncomingActivityEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(activityPubIncomingActivityEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *activityPubIncomingActivityManager) ApplyAtomic(transaction ITransaction, activityPubIncomingActivityEntity IActivityPubIncomingActivityEntity, repositoryHandler func(IRepositoryTransaction, IActivityPubIncomingActivityEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IActivityPubIncomingActivity, error) {
	result, err := NewActivityPubIncomingActivityFromEntity(activityPubIncomingActivityEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, activityPubIncomingActivityEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *activityPubIncomingActivityManager) Find(id int64) IActivityPubIncomingActivity {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IActivityPubIncomingActivity)
	}
}

func (manager *activityPubIncomingActivityManager) ForEach(iterator ActivityPubIncomingActivityIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IActivityPubIncomingActivity))
	})
}

func (manager *activityPubIncomingActivityManager) Filter(predicate ActivityPubIncomingActivityFilterPredicate) IActivityPubIncomingActivityCollection {
	activityPubIncomingActivities := NewActivityPubIncomingActivities()
	if predicate == nil {
		return activityPubIncomingActivities
	}

	manager.ForEach(func(activityPubIncomingActivity IActivityPubIncomingActivity) {
		if predicate(activityPubIncomingActivity) {
			activityPubIncomingActivities.Append(activityPubIncomingActivity)
		}
	})

	return activityPubIncomingActivities
}

func (manager *activityPubIncomingActivityManager) Map(predicate ActivityPubIncomingActivityMapPredicate) IActivityPubIncomingActivityCollection {
	activityPubIncomingActivities := NewActivityPubIncomingActivities()
	if predicate == nil {
		return activityPubIncomingActivities
	}

	manager.ForEach(func(activityPubIncomingActivity IActivityPubIncomingActivity) {
		activityPubIncomingActivities.Append(predicate(activityPubIncomingActivity))
	})

	return activityPubIncomingActivities
}

func (manager *activityPubIncomingActivityManager) ListActivityPubIncomingActivitiesByIdentity(identityId int64, _ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubIncomingActivityCollection {
	return manager.Filter(func(activityPubIncomingActivity IActivityPubIncomingActivity) bool {
		return activityPubIncomingActivity.IdentityId() == identityId
	})
}

func (manager *activityPubIncomingActivityManager) ForEachByIdentity(identityId int64, iterator ActivityPubIncomingActivityIterator) {
	manager.ForEach(func(activityPubIncomingActivity IActivityPubIncomingActivity) {
		if activityPubIncomingActivity.IdentityId() == identityId {
			iterator(activityPubIncomingActivity)
		}
	})
}
