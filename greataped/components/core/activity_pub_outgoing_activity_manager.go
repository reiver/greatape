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
const ACTIVITY_PUB_OUTGOING_ACTIVITY_MANAGER = "ActivityPubOutgoingActivityManager"

//lint:ignore U1000 GoUnused
type activityPubOutgoingActivityManager struct {
	systemComponent
	cache ICache

	//Dependencies
	identityManager IIdentityManager
}

func newActivityPubOutgoingActivityManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IActivityPubOutgoingActivityManager {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize

	manager := &activityPubOutgoingActivityManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *activityPubOutgoingActivityManager) Name() string {
	return ACTIVITY_PUB_OUTGOING_ACTIVITY_MANAGER
}

func (manager *activityPubOutgoingActivityManager) ResolveDependencies(dependencies ...ISystemComponent) error {
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

func (manager *activityPubOutgoingActivityManager) Load() error {
	activityPubOutgoingActivityEntities, err := repository.ActivityPubOutgoingActivities.FetchAll()
	if err != nil {
		return err
	}

	activityPubOutgoingActivities := make(SystemObjectCache)
	for _, activityPubOutgoingActivityEntity := range activityPubOutgoingActivityEntities {
		if activityPubOutgoingActivity, err := NewActivityPubOutgoingActivityFromEntity(activityPubOutgoingActivityEntity); err == nil {
			activityPubOutgoingActivities[activityPubOutgoingActivity.Id()] = activityPubOutgoingActivity
		} else {
			return err
		}
	}

	manager.cache.Load(activityPubOutgoingActivities)
	return nil
}

func (manager *activityPubOutgoingActivityManager) Reload() error {
	return manager.Load()
}

func (manager *activityPubOutgoingActivityManager) OnCacheChanged(callback ActivityPubOutgoingActivityCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *activityPubOutgoingActivityManager) Count() int {
	return manager.cache.Size()
}

func (manager *activityPubOutgoingActivityManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *activityPubOutgoingActivityManager) ExistsWhich(condition ActivityPubOutgoingActivityCondition) bool {
	var activityPubOutgoingActivities ActivityPubOutgoingActivities
	manager.ForEach(func(activityPubOutgoingActivity IActivityPubOutgoingActivity) {
		if condition(activityPubOutgoingActivity) {
			activityPubOutgoingActivities = append(activityPubOutgoingActivities, activityPubOutgoingActivity)
		}
	})

	return len(activityPubOutgoingActivities) > 0
}

func (manager *activityPubOutgoingActivityManager) ListActivityPubOutgoingActivities(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubOutgoingActivityCollection {
	return manager.Filter(ActivityPubOutgoingActivityPassThroughFilter)
}

func (manager *activityPubOutgoingActivityManager) GetActivityPubOutgoingActivity(id int64, _ Identity) (IActivityPubOutgoingActivity, error) {
	if activityPubOutgoingActivity := manager.Find(id); activityPubOutgoingActivity == nil {
		return nil, ERROR_ACTIVITY_PUB_OUTGOING_ACTIVITY_NOT_FOUND
	} else {
		return activityPubOutgoingActivity, nil
	}
}

func (manager *activityPubOutgoingActivityManager) AddActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(manager.UniqueId(), identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.Apply(activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.Add, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) AddActivityPubOutgoingActivityWithCustomId(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.Apply(activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.Add, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) AddActivityPubOutgoingActivityObject(activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(manager.UniqueId(), activityPubOutgoingActivity.IdentityId(), activityPubOutgoingActivity.UniqueIdentifier(), activityPubOutgoingActivity.Timestamp(), activityPubOutgoingActivity.From(), activityPubOutgoingActivity.To(), activityPubOutgoingActivity.Content(), activityPubOutgoingActivity.Raw())
	return manager.Apply(activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.Add, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) AddActivityPubOutgoingActivityAtomic(transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(manager.UniqueId(), identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.ApplyAtomic(transaction, activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.AddAtomic, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) AddActivityPubOutgoingActivityWithCustomIdAtomic(id int64, transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.ApplyAtomic(transaction, activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.AddAtomic, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) AddActivityPubOutgoingActivityObjectAtomic(transaction ITransaction, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(manager.UniqueId(), activityPubOutgoingActivity.IdentityId(), activityPubOutgoingActivity.UniqueIdentifier(), activityPubOutgoingActivity.Timestamp(), activityPubOutgoingActivity.From(), activityPubOutgoingActivity.To(), activityPubOutgoingActivity.Content(), activityPubOutgoingActivity.Raw())
	return manager.ApplyAtomic(transaction, activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.AddAtomic, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) Log(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor Identity, payload string) {
	activityPubOutgoingActivityPipeEntity := NewActivityPubOutgoingActivityPipeEntity(manager.UniqueId(), identityId, uniqueIdentifier, timestamp, from, to, content, raw, source, editor.Id(), payload)
	repository.Pipe.Insert(activityPubOutgoingActivityPipeEntity)

	activityPubOutgoingActivity, err := NewActivityPubOutgoingActivityFromEntity(activityPubOutgoingActivityPipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(activityPubOutgoingActivity.Id(), activityPubOutgoingActivity)
	}
}

func (manager *activityPubOutgoingActivityManager) UpdateActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.Apply(activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.Update, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) UpdateActivityPubOutgoingActivityObject(id int64, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(id, activityPubOutgoingActivity.Id(), activityPubOutgoingActivity.UniqueIdentifier(), activityPubOutgoingActivity.Timestamp(), activityPubOutgoingActivity.From(), activityPubOutgoingActivity.To(), activityPubOutgoingActivity.Content(), activityPubOutgoingActivity.Raw())
	return manager.Apply(activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.Update, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) UpdateActivityPubOutgoingActivityAtomic(transaction ITransaction, id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
	return manager.ApplyAtomic(transaction, activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) UpdateActivityPubOutgoingActivityObjectAtomic(transaction ITransaction, id int64, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(id, activityPubOutgoingActivity.Id(), activityPubOutgoingActivity.UniqueIdentifier(), activityPubOutgoingActivity.Timestamp(), activityPubOutgoingActivity.From(), activityPubOutgoingActivity.To(), activityPubOutgoingActivity.Content(), activityPubOutgoingActivity.Raw())
	return manager.ApplyAtomic(transaction, activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *activityPubOutgoingActivityManager) AddOrUpdateActivityPubOutgoingActivityObject(id int64, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubOutgoingActivityObject(id, activityPubOutgoingActivity, editor)
	} else {
		return manager.AddActivityPubOutgoingActivityObject(activityPubOutgoingActivity, editor)
	}
}

func (manager *activityPubOutgoingActivityManager) AddOrUpdateActivityPubOutgoingActivityObjectAtomic(transaction ITransaction, id int64, activityPubOutgoingActivity IActivityPubOutgoingActivity, editor Identity) (IActivityPubOutgoingActivity, error) {
	if manager.Exists(id) {
		return manager.UpdateActivityPubOutgoingActivityObjectAtomic(transaction, id, activityPubOutgoingActivity, editor)
	} else {
		return manager.AddActivityPubOutgoingActivityObjectAtomic(transaction, activityPubOutgoingActivity, editor)
	}
}

func (manager *activityPubOutgoingActivityManager) RemoveActivityPubOutgoingActivity(id int64, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(id, 0, "", 0, "", "", "", "")
	return manager.Apply(activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.Remove, manager.cache.Remove, editor)
}

func (manager *activityPubOutgoingActivityManager) RemoveActivityPubOutgoingActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubOutgoingActivity, error) {
	activityPubOutgoingActivityEntity := NewActivityPubOutgoingActivityEntity(id, 0, "", 0, "", "", "", "")
	return manager.ApplyAtomic(transaction, activityPubOutgoingActivityEntity, repository.ActivityPubOutgoingActivities.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *activityPubOutgoingActivityManager) Apply(activityPubOutgoingActivityEntity IActivityPubOutgoingActivityEntity, repositoryHandler func(IActivityPubOutgoingActivityEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IActivityPubOutgoingActivity, error) {
	result, err := NewActivityPubOutgoingActivityFromEntity(activityPubOutgoingActivityEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(activityPubOutgoingActivityEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *activityPubOutgoingActivityManager) ApplyAtomic(transaction ITransaction, activityPubOutgoingActivityEntity IActivityPubOutgoingActivityEntity, repositoryHandler func(IRepositoryTransaction, IActivityPubOutgoingActivityEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (IActivityPubOutgoingActivity, error) {
	result, err := NewActivityPubOutgoingActivityFromEntity(activityPubOutgoingActivityEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, activityPubOutgoingActivityEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *activityPubOutgoingActivityManager) Find(id int64) IActivityPubOutgoingActivity {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IActivityPubOutgoingActivity)
	}
}

func (manager *activityPubOutgoingActivityManager) ForEach(iterator ActivityPubOutgoingActivityIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IActivityPubOutgoingActivity))
	})
}

func (manager *activityPubOutgoingActivityManager) Filter(predicate ActivityPubOutgoingActivityFilterPredicate) IActivityPubOutgoingActivityCollection {
	activityPubOutgoingActivities := NewActivityPubOutgoingActivities()
	if predicate == nil {
		return activityPubOutgoingActivities
	}

	manager.ForEach(func(activityPubOutgoingActivity IActivityPubOutgoingActivity) {
		if predicate(activityPubOutgoingActivity) {
			activityPubOutgoingActivities.Append(activityPubOutgoingActivity)
		}
	})

	return activityPubOutgoingActivities
}

func (manager *activityPubOutgoingActivityManager) Map(predicate ActivityPubOutgoingActivityMapPredicate) IActivityPubOutgoingActivityCollection {
	activityPubOutgoingActivities := NewActivityPubOutgoingActivities()
	if predicate == nil {
		return activityPubOutgoingActivities
	}

	manager.ForEach(func(activityPubOutgoingActivity IActivityPubOutgoingActivity) {
		activityPubOutgoingActivities.Append(predicate(activityPubOutgoingActivity))
	})

	return activityPubOutgoingActivities
}

func (manager *activityPubOutgoingActivityManager) ListActivityPubOutgoingActivitiesByIdentity(identityId int64, _ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IActivityPubOutgoingActivityCollection {
	return manager.Filter(func(activityPubOutgoingActivity IActivityPubOutgoingActivity) bool {
		return activityPubOutgoingActivity.IdentityId() == identityId
	})
}

func (manager *activityPubOutgoingActivityManager) ForEachByIdentity(identityId int64, iterator ActivityPubOutgoingActivityIterator) {
	manager.ForEach(func(activityPubOutgoingActivity IActivityPubOutgoingActivity) {
		if activityPubOutgoingActivity.IdentityId() == identityId {
			iterator(activityPubOutgoingActivity)
		}
	})
}
