package core

import (
	"fmt"

	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
)

type category struct {
	object
	categoryTypeId int64
	categoryId     int64
	title          string
	description    string
}

func NewCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string) (ICategory, error) {
	instance := &category{
		object: object{
			id: id,
		},
		categoryTypeId: categoryTypeId,
		categoryId:     categoryId,
		title:          title,
		description:    description,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewCategoryFromEntity(entity ICategoryEntity) (ICategory, error) {
	instance := &category{
		object: object{
			id:        entity.Id(),
			sortOrder: entity.SortOrder(),
		},
		categoryTypeId: entity.CategoryTypeId(),
		categoryId:     entity.CategoryId(),
		title:          entity.Title(),
		description:    entity.Description(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (category *category) DependenciesAreUnknown() bool {
	// noinspection GoBoolExpressions
	return category.categoryTypeId == 0 || category.categoryId == 0 || false
}

func (category *category) CategoryTypeId() int64 {
	return category.categoryTypeId
}

func (category *category) AssertBelongsToCategoryType(_categoryType ICategoryType) {
	if category.categoryTypeId != _categoryType.Id() {
		panic(ERROR_MESSAGE_CATEGORY_NOT_FOUND)
	}
}

func (category *category) CategoryTypeIsUnknown() bool {
	return category.categoryTypeId == 0
}

func (category *category) AssertCategoryTypeIsProvided() {
	if category.categoryTypeId == 0 {
		panic(ERROR_MESSAGE_UNKNOWN_CATEGORY_TYPE)
	}
}

func (category *category) AssertCategoryType(categoryTypeId int64) {
	if category.categoryTypeId != categoryTypeId {
		panic(ERROR_MESSAGE_UNKNOWN_CATEGORY_TYPE)
	}
}

func (category *category) CategoryId() int64 {
	return category.categoryId
}

func (category *category) AssertBelongsToCategory(_category ICategory) {
	if category.categoryId != _category.Id() {
		panic(ERROR_MESSAGE_CATEGORY_NOT_FOUND)
	}
}

func (category *category) CategoryIsUnknown() bool {
	return category.categoryId == 0
}

func (category *category) AssertCategoryIsProvided() {
	if category.categoryId == 0 {
		panic(ERROR_MESSAGE_UNKNOWN_CATEGORY)
	}
}

func (category *category) AssertCategory(categoryId int64) {
	if category.categoryId != categoryId {
		panic(ERROR_MESSAGE_UNKNOWN_CATEGORY)
	}
}

func (category *category) Title() string {
	return category.title
}

func (category *category) UpdateTitle(title string, editor Identity) {
	if err := repository.Categories.UpdateTitle(category.id, title, editor.Id()); err != nil {
		panic(err.Error())
	}

	category.title = title
}

func (category *category) UpdateTitleAtomic(transaction ITransaction, title string, editor Identity) {
	transaction.OnCommit(func() {
		category.title = title
	})

	if err := repository.Categories.UpdateTitleAtomic(transaction, category.id, title, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (category *category) Description() string {
	return category.description
}

func (category *category) UpdateDescription(description string, editor Identity) {
	if err := repository.Categories.UpdateDescription(category.id, description, editor.Id()); err != nil {
		panic(err.Error())
	}

	category.description = description
}

func (category *category) UpdateDescriptionAtomic(transaction ITransaction, description string, editor Identity) {
	transaction.OnCommit(func() {
		category.description = description
	})

	if err := repository.Categories.UpdateDescriptionAtomic(transaction, category.id, description, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (category *category) SortOrder() float32 {
	return category.sortOrder
}

func (category *category) UpdateSortOrder(sortOrder float32, editor Identity) {
	if err := repository.Categories.UpdateSortOrder(category.id, sortOrder, editor.Id()); err != nil {
		panic(err.Error())
	}

	category.sortOrder = sortOrder
}

func (category *category) UpdateSortOrderAtomic(transaction ITransaction, sortOrder float32, editor Identity) {
	transaction.OnCommit(func() {
		category.sortOrder = sortOrder
	})

	if err := repository.Categories.UpdateSortOrderAtomic(transaction, category.id, sortOrder, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (category *category) Validate() error {
	return nil
}

func (category *category) String() string {
	return fmt.Sprintf("Category (Id: %d, CategoryTypeId: %d, CategoryId: %d, Title: %v, Description: %v)", category.Id(), category.CategoryTypeId(), category.CategoryId(), category.Title(), category.Description())
}

//------------------------------------------------------------------------------

type categories struct {
	collection Categories
}

// NewCategories creates an empty collection of 'Category' which is not thread-safe.
func NewCategories() ICategoryCollection {
	return &categories{
		collection: make(Categories, 0),
	}
}

func (categories *categories) Count() int {
	return len(categories.collection)
}

func (categories *categories) IsEmpty() bool {
	return len(categories.collection) == 0
}

func (categories *categories) IsNotEmpty() bool {
	return len(categories.collection) > 0
}

func (categories *categories) HasExactlyOneItem() bool {
	return len(categories.collection) == 1
}

func (categories *categories) HasAtLeastOneItem() bool {
	return len(categories.collection) >= 1
}

func (categories *categories) First() ICategory {
	return categories.collection[0]
}

func (categories *categories) Append(category ICategory) {
	categories.collection = append(categories.collection, category)
}

func (categories *categories) Reverse() ICategoryCollection {
	slice := categories.collection

	start := 0
	end := len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}

	categories.collection = slice

	return categories
}

func (categories *categories) ForEach(iterator CategoryIterator) {
	if iterator == nil {
		return
	}

	for _, value := range categories.collection {
		iterator(value)
	}
}

func (categories *categories) Array() Categories {
	return categories.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) CategoryExists(id int64) bool {
	return dispatcher.conductor.CategoryManager().Exists(id)
}

func (dispatcher *dispatcher) CategoryExistsWhich(condition CategoryCondition) bool {
	return dispatcher.conductor.CategoryManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListCategories() ICategoryCollection {
	return dispatcher.conductor.CategoryManager().ListCategories(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachCategory(iterator CategoryIterator) {
	dispatcher.conductor.CategoryManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterCategories(predicate CategoryFilterPredicate) ICategoryCollection {
	return dispatcher.conductor.CategoryManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapCategories(predicate CategoryMapPredicate) ICategoryCollection {
	return dispatcher.conductor.CategoryManager().Map(predicate)
}

func (dispatcher *dispatcher) GetCategory(id int64) ICategory {
	if category, err := dispatcher.conductor.CategoryManager().GetCategory(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return category
	}
}

func (dispatcher *dispatcher) AddCategory(categoryTypeId int64, categoryId int64, title string, description string) ICategory {
	transaction := dispatcher.transaction
	if transaction != nil {
		if category, err := dispatcher.conductor.CategoryManager().AddCategoryAtomic(transaction, categoryTypeId, categoryId, title, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	} else {
		if category, err := dispatcher.conductor.CategoryManager().AddCategory(categoryTypeId, categoryId, title, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	}
}

func (dispatcher *dispatcher) AddCategoryWithCustomId(id int64, categoryTypeId int64, categoryId int64, title string, description string) ICategory {
	transaction := dispatcher.transaction
	if transaction != nil {
		if category, err := dispatcher.conductor.CategoryManager().AddCategoryWithCustomIdAtomic(id, transaction, categoryTypeId, categoryId, title, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	} else {
		if category, err := dispatcher.conductor.CategoryManager().AddCategoryWithCustomId(id, categoryTypeId, categoryId, title, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	}
}

func (dispatcher *dispatcher) LogCategory(categoryTypeId int64, categoryId int64, title string, description string, source string, payload string) {
	dispatcher.conductor.CategoryManager().Log(categoryTypeId, categoryId, title, description, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string) ICategory {
	transaction := dispatcher.transaction
	if transaction != nil {
		if category, err := dispatcher.conductor.CategoryManager().UpdateCategoryAtomic(transaction, id, categoryTypeId, categoryId, title, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	} else {
		if category, err := dispatcher.conductor.CategoryManager().UpdateCategory(id, categoryTypeId, categoryId, title, description, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateCategoryObject(object IObject, category ICategory) ICategory {
	transaction := dispatcher.transaction
	if transaction != nil {
		if category, err := dispatcher.conductor.CategoryManager().UpdateCategoryAtomic(transaction, object.Id(), category.CategoryTypeId(), category.CategoryId(), category.Title(), category.Description(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	} else {
		if category, err := dispatcher.conductor.CategoryManager().UpdateCategory(object.Id(), category.CategoryTypeId(), category.CategoryId(), category.Title(), category.Description(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateCategoryObject(object IObject, category ICategory) ICategory {
	transaction := dispatcher.transaction
	if transaction != nil {
		if category, err := dispatcher.conductor.CategoryManager().AddOrUpdateCategoryObjectAtomic(transaction, object.Id(), category, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	} else {
		if category, err := dispatcher.conductor.CategoryManager().AddOrUpdateCategoryObject(object.Id(), category, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	}
}

func (dispatcher *dispatcher) RemoveCategory(id int64) ICategory {
	transaction := dispatcher.transaction
	if transaction != nil {
		if category, err := dispatcher.conductor.CategoryManager().RemoveCategoryAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	} else {
		if category, err := dispatcher.conductor.CategoryManager().RemoveCategory(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return category
		}
	}
}

func (dispatcher *dispatcher) ListCategoriesByCategoryType(categoryType ICategoryType) ICategoryCollection {
	return dispatcher.conductor.CategoryManager().ListCategoriesByCategoryType(categoryType.Id(), 0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ListCategoriesByCategoryTypeId(categoryTypeId int64) ICategoryCollection {
	return dispatcher.conductor.CategoryManager().ListCategoriesByCategoryType(categoryTypeId, 0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachCategoryByCategoryType(categoryType ICategoryType, iterator CategoryIterator) {
	dispatcher.conductor.CategoryManager().ForEachByCategoryType(categoryType.Id(), iterator)
}

func (dispatcher *dispatcher) ForEachCategoryByCategoryTypeId(categoryTypeId int64, iterator CategoryIterator) {
	dispatcher.conductor.CategoryManager().ForEachByCategoryType(categoryTypeId, iterator)
}

func (dispatcher *dispatcher) ListCategoriesByCategory(category ICategory) ICategoryCollection {
	return dispatcher.conductor.CategoryManager().ListCategoriesByCategory(category.Id(), 0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ListCategoriesByCategoryId(categoryId int64) ICategoryCollection {
	return dispatcher.conductor.CategoryManager().ListCategoriesByCategory(categoryId, 0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachCategoryByCategory(category ICategory, iterator CategoryIterator) {
	dispatcher.conductor.CategoryManager().ForEachByCategory(category.Id(), iterator)
}

func (dispatcher *dispatcher) ForEachCategoryByCategoryId(categoryId int64, iterator CategoryIterator) {
	dispatcher.conductor.CategoryManager().ForEachByCategory(categoryId, iterator)
}
