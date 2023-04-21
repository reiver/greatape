package core

import (
	"sync"
	"time"

	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/security"
)

//region IDispatcherCache Implementation

var EPOCH = time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)

type dispatcherCache struct {
	conductor IConductor
	identity  Identity
	timeout   time.Duration

	categoriesByCategoryTypeCacheInstance                *categoriesByCategoryTypeCache
	categoriesByCategoryCacheInstance                    *categoriesByCategoryCache
	activityPubIncomingActivitiesByIdentityCacheInstance *activityPubIncomingActivitiesByIdentityCache
	activityPubOutgoingActivitiesByIdentityCacheInstance *activityPubOutgoingActivitiesByIdentityCache
}

func newDispatcherCache(conductor IConductor, identity Identity) IDispatcherCache {
	instance := &dispatcherCache{
		conductor: conductor,
		identity:  identity,
		timeout:   30 * time.Minute,

		categoriesByCategoryTypeCacheInstance:                newCategoriesByCategoryTypeCache(),
		categoriesByCategoryCacheInstance:                    newCategoriesByCategoryCache(),
		activityPubIncomingActivitiesByIdentityCacheInstance: newActivityPubIncomingActivitiesByIdentityCache(),
		activityPubOutgoingActivitiesByIdentityCacheInstance: newActivityPubOutgoingActivitiesByIdentityCache(),
	}

	// Categories
	instance.conductor.CategoryManager().OnCacheChanged(func() {
		instance.categoriesByCategoryTypeCacheInstance.Invalidate()
		instance.categoriesByCategoryCacheInstance.Invalidate()
	})

	// ActivityPubIncomingActivities
	instance.conductor.ActivityPubIncomingActivityManager().OnCacheChanged(func() {
		instance.activityPubIncomingActivitiesByIdentityCacheInstance.Invalidate()
	})

	// ActivityPubOutgoingActivities
	instance.conductor.ActivityPubOutgoingActivityManager().OnCacheChanged(func() {
		instance.activityPubOutgoingActivitiesByIdentityCacheInstance.Invalidate()
	})

	return instance
}

// CategoriesByCategoryType
// ------------------------------------------------------------

type categoriesByCategoryTypeCache struct {
	sync.RWMutex
	items      map[int64]ICategoryCollection
	lastUpdate time.Time
}

func newCategoriesByCategoryTypeCache() *categoriesByCategoryTypeCache {
	return &categoriesByCategoryTypeCache{
		items:      make(map[int64]ICategoryCollection),
		lastUpdate: EPOCH,
	}
}

func (cache *categoriesByCategoryTypeCache) Invalidate() {
	cache.Lock()
	defer cache.Unlock()

	cache.lastUpdate = EPOCH
}

func (cache *dispatcherCache) ListCategoriesByCategoryType(categoryType ICategoryType) ICategoryCollection {
	return cache.ListCategoriesByCategoryTypeId(categoryType.Id())
}

func (cache *dispatcherCache) ListCategoriesByCategoryTypeId(categoryTypeId int64) ICategoryCollection {
	instance := cache.categoriesByCategoryTypeCacheInstance

	func() {
		instance.Lock()
		defer instance.Unlock()

		if time.Since(instance.lastUpdate) > cache.timeout {
			cache.conductor.CategoryManager().ForEach(func(category ICategory) {
				instance.items[category.CategoryTypeId()].Append(category)
			})

			instance.lastUpdate = time.Now()
		}
	}()

	instance.RLock()
	defer instance.RUnlock()

	if item, exists := instance.items[categoryTypeId]; exists {
		return item
	}

	return NewCategories()
}

func (cache *dispatcherCache) ForEachCategoryByCategoryType(categoryType ICategoryType, iterator CategoryIterator) {
	cache.ForEachCategoryByCategoryTypeId(categoryType.Id(), iterator)
}

func (cache *dispatcherCache) ForEachCategoryByCategoryTypeId(categoryTypeId int64, iterator CategoryIterator) {
	cache.ListCategoriesByCategoryTypeId(categoryTypeId).ForEach(iterator)
}

// CategoriesByCategory
// ------------------------------------------------------------

type categoriesByCategoryCache struct {
	sync.RWMutex
	items      map[int64]ICategoryCollection
	lastUpdate time.Time
}

func newCategoriesByCategoryCache() *categoriesByCategoryCache {
	return &categoriesByCategoryCache{
		items:      make(map[int64]ICategoryCollection),
		lastUpdate: EPOCH,
	}
}

func (cache *categoriesByCategoryCache) Invalidate() {
	cache.Lock()
	defer cache.Unlock()

	cache.lastUpdate = EPOCH
}

func (cache *dispatcherCache) ListCategoriesByCategory(category ICategory) ICategoryCollection {
	return cache.ListCategoriesByCategoryId(category.Id())
}

func (cache *dispatcherCache) ListCategoriesByCategoryId(categoryId int64) ICategoryCollection {
	instance := cache.categoriesByCategoryCacheInstance

	func() {
		instance.Lock()
		defer instance.Unlock()

		if time.Since(instance.lastUpdate) > cache.timeout {
			cache.conductor.CategoryManager().ForEach(func(category ICategory) {
				instance.items[category.CategoryId()].Append(category)
			})

			instance.lastUpdate = time.Now()
		}
	}()

	instance.RLock()
	defer instance.RUnlock()

	if item, exists := instance.items[categoryId]; exists {
		return item
	}

	return NewCategories()
}

func (cache *dispatcherCache) ForEachCategoryByCategory(category ICategory, iterator CategoryIterator) {
	cache.ForEachCategoryByCategoryId(category.Id(), iterator)
}

func (cache *dispatcherCache) ForEachCategoryByCategoryId(categoryId int64, iterator CategoryIterator) {
	cache.ListCategoriesByCategoryId(categoryId).ForEach(iterator)
}

// ActivityPubIncomingActivitiesByIdentity
// ------------------------------------------------------------

type activityPubIncomingActivitiesByIdentityCache struct {
	sync.RWMutex
	items      map[int64]IActivityPubIncomingActivityCollection
	lastUpdate time.Time
}

func newActivityPubIncomingActivitiesByIdentityCache() *activityPubIncomingActivitiesByIdentityCache {
	return &activityPubIncomingActivitiesByIdentityCache{
		items:      make(map[int64]IActivityPubIncomingActivityCollection),
		lastUpdate: EPOCH,
	}
}

func (cache *activityPubIncomingActivitiesByIdentityCache) Invalidate() {
	cache.Lock()
	defer cache.Unlock()

	cache.lastUpdate = EPOCH
}

func (cache *dispatcherCache) ListActivityPubIncomingActivitiesByIdentity(identity IIdentity) IActivityPubIncomingActivityCollection {
	return cache.ListActivityPubIncomingActivitiesByIdentityId(identity.Id())
}

func (cache *dispatcherCache) ListActivityPubIncomingActivitiesByIdentityId(identityId int64) IActivityPubIncomingActivityCollection {
	instance := cache.activityPubIncomingActivitiesByIdentityCacheInstance

	func() {
		instance.Lock()
		defer instance.Unlock()

		if time.Since(instance.lastUpdate) > cache.timeout {
			cache.conductor.ActivityPubIncomingActivityManager().ForEach(func(activityPubIncomingActivity IActivityPubIncomingActivity) {
				instance.items[activityPubIncomingActivity.IdentityId()].Append(activityPubIncomingActivity)
			})

			instance.lastUpdate = time.Now()
		}
	}()

	instance.RLock()
	defer instance.RUnlock()

	if item, exists := instance.items[identityId]; exists {
		return item
	}

	return NewActivityPubIncomingActivities()
}

func (cache *dispatcherCache) ForEachActivityPubIncomingActivityByIdentity(identity IIdentity, iterator ActivityPubIncomingActivityIterator) {
	cache.ForEachActivityPubIncomingActivityByIdentityId(identity.Id(), iterator)
}

func (cache *dispatcherCache) ForEachActivityPubIncomingActivityByIdentityId(identityId int64, iterator ActivityPubIncomingActivityIterator) {
	cache.ListActivityPubIncomingActivitiesByIdentityId(identityId).ForEach(iterator)
}

// ActivityPubOutgoingActivitiesByIdentity
// ------------------------------------------------------------

type activityPubOutgoingActivitiesByIdentityCache struct {
	sync.RWMutex
	items      map[int64]IActivityPubOutgoingActivityCollection
	lastUpdate time.Time
}

func newActivityPubOutgoingActivitiesByIdentityCache() *activityPubOutgoingActivitiesByIdentityCache {
	return &activityPubOutgoingActivitiesByIdentityCache{
		items:      make(map[int64]IActivityPubOutgoingActivityCollection),
		lastUpdate: EPOCH,
	}
}

func (cache *activityPubOutgoingActivitiesByIdentityCache) Invalidate() {
	cache.Lock()
	defer cache.Unlock()

	cache.lastUpdate = EPOCH
}

func (cache *dispatcherCache) ListActivityPubOutgoingActivitiesByIdentity(identity IIdentity) IActivityPubOutgoingActivityCollection {
	return cache.ListActivityPubOutgoingActivitiesByIdentityId(identity.Id())
}

func (cache *dispatcherCache) ListActivityPubOutgoingActivitiesByIdentityId(identityId int64) IActivityPubOutgoingActivityCollection {
	instance := cache.activityPubOutgoingActivitiesByIdentityCacheInstance

	func() {
		instance.Lock()
		defer instance.Unlock()

		if time.Since(instance.lastUpdate) > cache.timeout {
			cache.conductor.ActivityPubOutgoingActivityManager().ForEach(func(activityPubOutgoingActivity IActivityPubOutgoingActivity) {
				instance.items[activityPubOutgoingActivity.IdentityId()].Append(activityPubOutgoingActivity)
			})

			instance.lastUpdate = time.Now()
		}
	}()

	instance.RLock()
	defer instance.RUnlock()

	if item, exists := instance.items[identityId]; exists {
		return item
	}

	return NewActivityPubOutgoingActivities()
}

func (cache *dispatcherCache) ForEachActivityPubOutgoingActivityByIdentity(identity IIdentity, iterator ActivityPubOutgoingActivityIterator) {
	cache.ForEachActivityPubOutgoingActivityByIdentityId(identity.Id(), iterator)
}

func (cache *dispatcherCache) ForEachActivityPubOutgoingActivityByIdentityId(identityId int64, iterator ActivityPubOutgoingActivityIterator) {
	cache.ListActivityPubOutgoingActivitiesByIdentityId(identityId).ForEach(iterator)
}

//endregion
