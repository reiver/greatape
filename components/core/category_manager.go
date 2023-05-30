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
const CATEGORY_MANAGER = "CategoryManager"

//lint:ignore U1000 GoUnused
type categoryManager struct {
	systemComponent
	cache ICache

	//Dependencies
	categoryTypeManager ICategoryTypeManager
	categoryManager     ICategoryManager
}

func newCategoryManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) ICategoryManager {
	manager := &categoryManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *categoryManager) Name() string {
	return CATEGORY_MANAGER
}

func (manager *categoryManager) ResolveDependencies(dependencies ...ISystemComponent) error {
	if len(dependencies) == 0 {
		return nil
	}

	var (
		categoryTypeManager ICategoryTypeManager
		categoryManager     ICategoryManager
	)

	for _, _dependency := range dependencies {
		if false {
		} else if dependency, ok := _dependency.(ICategoryTypeManager); ok {
			categoryTypeManager = dependency
		} else if dependency, ok := _dependency.(ICategoryManager); ok {
			categoryManager = dependency
		}
	}

	if // noinspection GoBoolExpressions
	false || categoryTypeManager == nil || categoryManager == nil {
		return ERROR_UNRESOLVED_DEPENDENCIES
	}

	return nil
}

func (manager *categoryManager) Load() error {
	categoryEntities, err := repository.Categories.FetchAll()
	if err != nil {
		return err
	}

	categories := make(SystemObjectCache)
	for _, categoryEntity := range categoryEntities {
		if category, err := NewCategoryFromEntity(categoryEntity); err == nil {
			categories[category.Id()] = category
		} else {
			return err
		}
	}

	manager.cache.Load(categories)
	return nil
}

func (manager *categoryManager) Reload() error {
	return manager.Load()
}

func (manager *categoryManager) OnCacheChanged(callback CategoryCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *categoryManager) Count() int {
	return manager.cache.Size()
}

func (manager *categoryManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *categoryManager) ExistsWhich(condition CategoryCondition) bool {
	var categories Categories
	manager.ForEach(func(category ICategory) {
		if condition(category) {
			categories = append(categories, category)
		}
	})

	return len(categories) > 0
}

func (manager *categoryManager) ListCategories(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) ICategoryCollection {
	return manager.Filter(CategoryPassThroughFilter)
}

func (manager *categoryManager) GetCategory(id int64, _ Identity) (ICategory, error) {
	if category := manager.Find(id); category == nil {
		return nil, ERROR_CATEGORY_NOT_FOUND
	} else {
		return category, nil
	}
}

func (manager *categoryManager) AddCategory(categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(manager.UniqueId(), categoryTypeId, categoryId, title, description)
	return manager.Apply(categoryEntity, repository.Categories.Add, manager.cache.Put, editor)
}

func (manager *categoryManager) AddCategoryWithCustomId(id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(id, categoryTypeId, categoryId, title, description)
	return manager.Apply(categoryEntity, repository.Categories.Add, manager.cache.Put, editor)
}

func (manager *categoryManager) AddCategoryObject(category ICategory, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(manager.UniqueId(), category.CategoryTypeId(), category.CategoryId(), category.Title(), category.Description())
	return manager.Apply(categoryEntity, repository.Categories.Add, manager.cache.Put, editor)
}

func (manager *categoryManager) AddCategoryAtomic(transaction ITransaction, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(manager.UniqueId(), categoryTypeId, categoryId, title, description)
	return manager.ApplyAtomic(transaction, categoryEntity, repository.Categories.AddAtomic, manager.cache.Put, editor)
}

func (manager *categoryManager) AddCategoryWithCustomIdAtomic(id int64, transaction ITransaction, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(id, categoryTypeId, categoryId, title, description)
	return manager.ApplyAtomic(transaction, categoryEntity, repository.Categories.AddAtomic, manager.cache.Put, editor)
}

func (manager *categoryManager) AddCategoryObjectAtomic(transaction ITransaction, category ICategory, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(manager.UniqueId(), category.CategoryTypeId(), category.CategoryId(), category.Title(), category.Description())
	return manager.ApplyAtomic(transaction, categoryEntity, repository.Categories.AddAtomic, manager.cache.Put, editor)
}

func (manager *categoryManager) Log(categoryTypeId int64, categoryId int64, title string, description string, source string, editor Identity, payload string) {
	categoryPipeEntity := NewCategoryPipeEntity(manager.UniqueId(), categoryTypeId, categoryId, title, description, source, editor.Id(), payload)
	repository.Pipe.Insert(categoryPipeEntity)

	category, err := NewCategoryFromEntity(categoryPipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(category.Id(), category)
	}
}

func (manager *categoryManager) UpdateCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(id, categoryTypeId, categoryId, title, description)
	return manager.Apply(categoryEntity, repository.Categories.Update, manager.cache.Put, editor)
}

func (manager *categoryManager) UpdateCategoryObject(id int64, category ICategory, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(id, category.Id(), category.Id(), category.Title(), category.Description())
	return manager.Apply(categoryEntity, repository.Categories.Update, manager.cache.Put, editor)
}

func (manager *categoryManager) UpdateCategoryAtomic(transaction ITransaction, id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(id, categoryTypeId, categoryId, title, description)
	return manager.ApplyAtomic(transaction, categoryEntity, repository.Categories.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *categoryManager) UpdateCategoryObjectAtomic(transaction ITransaction, id int64, category ICategory, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(id, category.Id(), category.Id(), category.Title(), category.Description())
	return manager.ApplyAtomic(transaction, categoryEntity, repository.Categories.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *categoryManager) AddOrUpdateCategoryObject(id int64, category ICategory, editor Identity) (ICategory, error) {
	if manager.Exists(id) {
		return manager.UpdateCategoryObject(id, category, editor)
	} else {
		return manager.AddCategoryObject(category, editor)
	}
}

func (manager *categoryManager) AddOrUpdateCategoryObjectAtomic(transaction ITransaction, id int64, category ICategory, editor Identity) (ICategory, error) {
	if manager.Exists(id) {
		return manager.UpdateCategoryObjectAtomic(transaction, id, category, editor)
	} else {
		return manager.AddCategoryObjectAtomic(transaction, category, editor)
	}
}

func (manager *categoryManager) RemoveCategory(id int64, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(id, 0, 0, "", "")
	return manager.Apply(categoryEntity, repository.Categories.Remove, manager.cache.Remove, editor)
}

func (manager *categoryManager) RemoveCategoryAtomic(transaction ITransaction, id int64, editor Identity) (ICategory, error) {
	categoryEntity := NewCategoryEntity(id, 0, 0, "", "")
	return manager.ApplyAtomic(transaction, categoryEntity, repository.Categories.RemoveAtomic, manager.cache.Remove, editor)
}

func (manager *categoryManager) Apply(categoryEntity ICategoryEntity, repositoryHandler func(ICategoryEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (ICategory, error) {
	result, err := NewCategoryFromEntity(categoryEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(categoryEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *categoryManager) ApplyAtomic(transaction ITransaction, categoryEntity ICategoryEntity, repositoryHandler func(IRepositoryTransaction, ICategoryEntity, int64) error, cacheHandler func(int64, ISystemObject), editor Identity) (ICategory, error) {
	result, err := NewCategoryFromEntity(categoryEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, categoryEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *categoryManager) Find(id int64) ICategory {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(ICategory)
	}
}

func (manager *categoryManager) ForEach(iterator CategoryIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(ICategory))
	})
}

func (manager *categoryManager) Filter(predicate CategoryFilterPredicate) ICategoryCollection {
	categories := NewCategories()
	if predicate == nil {
		return categories
	}

	manager.ForEach(func(category ICategory) {
		if predicate(category) {
			categories.Append(category)
		}
	})

	return categories
}

func (manager *categoryManager) Map(predicate CategoryMapPredicate) ICategoryCollection {
	categories := NewCategories()
	if predicate == nil {
		return categories
	}

	manager.ForEach(func(category ICategory) {
		categories.Append(predicate(category))
	})

	return categories
}

func (manager *categoryManager) ListCategoriesByCategoryType(categoryTypeId int64, _ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) ICategoryCollection {
	return manager.Filter(func(category ICategory) bool {
		return category.CategoryTypeId() == categoryTypeId
	})
}

func (manager *categoryManager) ForEachByCategoryType(categoryTypeId int64, iterator CategoryIterator) {
	manager.ForEach(func(category ICategory) {
		if category.CategoryTypeId() == categoryTypeId {
			iterator(category)
		}
	})
}

func (manager *categoryManager) ListCategoriesByCategory(categoryId int64, _ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) ICategoryCollection {
	return manager.Filter(func(category ICategory) bool {
		return category.CategoryId() == categoryId
	})
}

func (manager *categoryManager) ForEachByCategory(categoryId int64, iterator CategoryIterator) {
	manager.ForEach(func(category ICategory) {
		if category.CategoryId() == categoryId {
			iterator(category)
		}
	})
}
