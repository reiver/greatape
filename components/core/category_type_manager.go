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
const CATEGORY_TYPE_MANAGER = "CategoryTypeManager"

type categoryTypeManager struct {
	systemComponent
	cache ICache
}

func newCategoryTypeManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) ICategoryTypeManager {
	manager := &categoryTypeManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *categoryTypeManager) Name() string {
	return CATEGORY_TYPE_MANAGER
}

func (manager *categoryTypeManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *categoryTypeManager) Load() error {
	categoryTypeEntities, err := repository.CategoryTypes.FetchAll()
	if err != nil {
		return err
	}

	categoryTypes := make(SystemObjectCache)
	for _, categoryTypeEntity := range categoryTypeEntities {
		if categoryType, err := NewCategoryTypeFromEntity(categoryTypeEntity); err == nil {
			categoryTypes[categoryType.Id()] = categoryType
		} else {
			return err
		}
	}

	manager.cache.Load(categoryTypes)
	return nil
}

func (manager *categoryTypeManager) Reload() error {
	return manager.Load()
}

func (manager *categoryTypeManager) OnCacheChanged(callback CategoryTypeCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *categoryTypeManager) Count() int {
	return manager.cache.Size()
}

func (manager *categoryTypeManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *categoryTypeManager) ExistsWhich(condition CategoryTypeCondition) bool {
	var categoryTypes CategoryTypes
	manager.ForEach(func(categoryType ICategoryType) {
		if condition(categoryType) {
			categoryTypes = append(categoryTypes, categoryType)
		}
	})

	return len(categoryTypes) > 0
}

func (manager *categoryTypeManager) ListCategoryTypes(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) ICategoryTypeCollection {
	return manager.Filter(CategoryTypePassThroughFilter)
}

func (manager *categoryTypeManager) GetCategoryType(id int64, _ Identity) (ICategoryType, error) {
	if categoryType := manager.Find(id); categoryType == nil {
		return nil, ERROR_CATEGORY_TYPE_NOT_FOUND
	} else {
		return categoryType, nil
	}
}

func (manager *categoryTypeManager) AddCategoryType(description string, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(manager.UniqueId(), description)
	return manager.Apply(categoryTypeEntity, repository.CategoryTypes.Add, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) AddCategoryTypeWithCustomId(id int64, description string, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(id, description)
	return manager.Apply(categoryTypeEntity, repository.CategoryTypes.Add, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) AddCategoryTypeObject(categoryType ICategoryType, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(manager.UniqueId(), categoryType.Description())
	return manager.Apply(categoryTypeEntity, repository.CategoryTypes.Add, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) AddCategoryTypeAtomic(transaction ITransaction, description string, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(manager.UniqueId(), description)
	return manager.ApplyAtomic(transaction, categoryTypeEntity, repository.CategoryTypes.AddAtomic, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) AddCategoryTypeWithCustomIdAtomic(id int64, transaction ITransaction, description string, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(id, description)
	return manager.ApplyAtomic(transaction, categoryTypeEntity, repository.CategoryTypes.AddAtomic, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) AddCategoryTypeObjectAtomic(transaction ITransaction, categoryType ICategoryType, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(manager.UniqueId(), categoryType.Description())
	return manager.ApplyAtomic(transaction, categoryTypeEntity, repository.CategoryTypes.AddAtomic, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) Log(description string, source string, editor Identity, payload string) {
	categoryTypePipeEntity := NewCategoryTypePipeEntity(manager.UniqueId(), description, source, editor.Id(), payload)
	repository.Pipe.Insert(categoryTypePipeEntity)

	categoryType, err := NewCategoryTypeFromEntity(categoryTypePipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(categoryType.Id(), categoryType)
	}
}

func (manager *categoryTypeManager) UpdateCategoryType(id int64, description string, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(id, description)
	return manager.Apply(categoryTypeEntity, repository.CategoryTypes.Update, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) UpdateCategoryTypeObject(id int64, categoryType ICategoryType, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(id, categoryType.Description())
	return manager.Apply(categoryTypeEntity, repository.CategoryTypes.Update, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) UpdateCategoryTypeAtomic(transaction ITransaction, id int64, description string, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(id, description)
	return manager.ApplyAtomic(transaction, categoryTypeEntity, repository.CategoryTypes.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) UpdateCategoryTypeObjectAtomic(transaction ITransaction, id int64, categoryType ICategoryType, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(id, categoryType.Description())
	return manager.ApplyAtomic(transaction, categoryTypeEntity, repository.CategoryTypes.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *categoryTypeManager) AddOrUpdateCategoryTypeObject(id int64, categoryType ICategoryType, editor Identity) (ICategoryType, error) {
	if manager.Exists(id) {
		return manager.UpdateCategoryTypeObject(id, categoryType, editor)
	} else {
		return manager.AddCategoryTypeObject(categoryType, editor)
	}
}

func (manager *categoryTypeManager) AddOrUpdateCategoryTypeObjectAtomic(transaction ITransaction, id int64, categoryType ICategoryType, editor Identity) (ICategoryType, error) {
	if manager.Exists(id) {
		return manager.UpdateCategoryTypeObjectAtomic(transaction, id, categoryType, editor)
	} else {
		return manager.AddCategoryTypeObjectAtomic(transaction, categoryType, editor)
	}
}

func (manager *categoryTypeManager) RemoveCategoryType(id int64, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(id, "")
	return manager.Apply(categoryTypeEntity, repository.CategoryTypes.Remove, manager.cache.Remove, editor)
}

func (manager *categoryTypeManager) RemoveCategoryTypeAtomic(transaction ITransaction, id int64, editor Identity) (ICategoryType, error) {
	categoryTypeEntity := NewCategoryTypeEntity(id, "")
	return manager.ApplyAtomic(transaction, categoryTypeEntity, repository.CategoryTypes.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *categoryTypeManager) Apply(categoryTypeEntity ICategoryTypeEntity, repositoryHandler func(ICategoryTypeEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (ICategoryType, error) {
	result, err := NewCategoryTypeFromEntity(categoryTypeEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(categoryTypeEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *categoryTypeManager) ApplyAtomic(transaction ITransaction, categoryTypeEntity ICategoryTypeEntity, repositoryHandler func(IRepositoryTransaction, ICategoryTypeEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (ICategoryType, error) {
	result, err := NewCategoryTypeFromEntity(categoryTypeEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, categoryTypeEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *categoryTypeManager) Find(id int64) ICategoryType {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(ICategoryType)
	}
}

func (manager *categoryTypeManager) ForEach(iterator CategoryTypeIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(ICategoryType))
	})
}

func (manager *categoryTypeManager) Filter(predicate CategoryTypeFilterPredicate) ICategoryTypeCollection {
	categoryTypes := NewCategoryTypes()
	if predicate == nil {
		return categoryTypes
	}

	manager.ForEach(func(categoryType ICategoryType) {
		if predicate(categoryType) {
			categoryTypes.Append(categoryType)
		}
	})

	return categoryTypes
}

func (manager *categoryTypeManager) Map(predicate CategoryTypeMapPredicate) ICategoryTypeCollection {
	categoryTypes := NewCategoryTypes()
	if predicate == nil {
		return categoryTypes
	}

	manager.ForEach(func(categoryType ICategoryType) {
		categoryTypes.Append(predicate(categoryType))
	})

	return categoryTypes
}
