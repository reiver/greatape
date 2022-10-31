package core

import (
	"sync"
	"time"

	. "github.com/xeronith/diamante/contracts/security"
	. "rail.town/infrastructure/components/contracts"
)

//region IDispatcherCache Implementation

var EPOCH = time.Date(1970, time.January, 0, 0, 0, 0, 0, time.UTC)

type dispatcherCache struct {
	conductor IConductor
	identity  Identity
	timeout   time.Duration

	categoriesByCategoryTypeCacheInstance *categoriesByCategoryTypeCache
	categoriesByCategoryCacheInstance     *categoriesByCategoryCache
}

func newDispatcherCache(conductor IConductor, identity Identity) IDispatcherCache {
	instance := &dispatcherCache{
		conductor: conductor,
		identity:  identity,
		timeout:   30 * time.Minute,

		categoriesByCategoryTypeCacheInstance: newCategoriesByCategoryTypeCache(),
		categoriesByCategoryCacheInstance:     newCategoriesByCategoryCache(),
	}

	// Categories
	instance.conductor.CategoryManager().OnCacheChanged(func() {
		instance.categoriesByCategoryTypeCacheInstance.Invalidate()
		instance.categoriesByCategoryCacheInstance.Invalidate()
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

//endregion
