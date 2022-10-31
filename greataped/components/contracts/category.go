package contracts

import . "github.com/xeronith/diamante/contracts/security"

var CategoryPassThroughFilter = func(ICategory) bool { return true }

type (
	Categories              []ICategory
	CategoryIterator        func(ICategory)
	CategoryCondition       func(ICategory) bool
	CategoryFilterPredicate func(ICategory) bool
	CategoryMapPredicate    func(ICategory) ICategory
	CategoryCacheCallback   func()

	ICategory interface {
		IObject
		// DependenciesAreUnknown scans all dependencies to make sure a valid parent is set for all of them.
		DependenciesAreUnknown() bool
		// CategoryTypeId returns parent 'CategoryTypeId' of this 'Category' instance.
		CategoryTypeId() int64
		// AssertBelongsToCategoryType checks whether this 'Category' instance is a child of the specified 'CategoryType'
		AssertBelongsToCategoryType(categoryType ICategoryType)
		// CategoryTypeIsUnknown checks whether a valid parent 'CategoryTypeId' is provided for this 'Category' instance or not.
		CategoryTypeIsUnknown() bool
		// AssertCategoryTypeIsProvided asserts that a valid 'CategoryTypeId' is provided for this 'Category' instance. A panic will occur if the assertion is not valid.
		AssertCategoryTypeIsProvided()
		// AssertCategoryType asserts the given 'CategoryTypeId' is in fact the parent of this 'Category' instance. A panic will occur if the assertion is not valid.
		AssertCategoryType(categoryTypeId int64)
		// CategoryId returns parent 'CategoryId' of this 'Category' instance.
		CategoryId() int64
		// AssertBelongsToCategory checks whether this 'Category' instance is a child of the specified 'Category'
		AssertBelongsToCategory(category ICategory)
		// CategoryIsUnknown checks whether a valid parent 'CategoryId' is provided for this 'Category' instance or not.
		CategoryIsUnknown() bool
		// AssertCategoryIsProvided asserts that a valid 'CategoryId' is provided for this 'Category' instance. A panic will occur if the assertion is not valid.
		AssertCategoryIsProvided()
		// AssertCategory asserts the given 'CategoryId' is in fact the parent of this 'Category' instance. A panic will occur if the assertion is not valid.
		AssertCategory(categoryId int64)
		// Title returns 'Title' of this 'Category' instance.
		Title() string
		// UpdateTitle directly updates 'Title' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateTitle(title string, editor Identity)
		// UpdateTitleAtomic updates 'Title' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateTitleAtomic(transaction ITransaction, title string, editor Identity)
		// Description returns 'Description' of this 'Category' instance.
		Description() string
		// UpdateDescription directly updates 'Description' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateDescription(description string, editor Identity)
		// UpdateDescriptionAtomic updates 'Description' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateDescriptionAtomic(transaction ITransaction, description string, editor Identity)
	}

	ICategoryCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() ICategory
		Append(category ICategory)
		ForEach(CategoryIterator)
		Array() Categories
	}

	ICategoryManager interface {
		ISystemComponent
		OnCacheChanged(CategoryCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition CategoryCondition) bool
		ListCategories(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection
		GetCategory(id int64, editor Identity) (ICategory, error)
		AddCategory(categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		AddCategoryWithCustomId(id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		AddCategoryObject(category ICategory, editor Identity) (ICategory, error)
		AddCategoryAtomic(transaction ITransaction, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		AddCategoryWithCustomIdAtomic(id int64, transaction ITransaction, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		AddCategoryObjectAtomic(transaction ITransaction, category ICategory, editor Identity) (ICategory, error)
		Log(categoryTypeId int64, categoryId int64, title string, description string, source string, editor Identity, payload string)
		UpdateCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		UpdateCategoryObject(id int64, category ICategory, editor Identity) (ICategory, error)
		UpdateCategoryAtomic(transaction ITransaction, id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		UpdateCategoryObjectAtomic(transaction ITransaction, id int64, category ICategory, editor Identity) (ICategory, error)
		AddOrUpdateCategoryObject(id int64, category ICategory, editor Identity) (ICategory, error)
		AddOrUpdateCategoryObjectAtomic(transaction ITransaction, id int64, category ICategory, editor Identity) (ICategory, error)
		RemoveCategory(id int64, editor Identity) (ICategory, error)
		RemoveCategoryAtomic(transaction ITransaction, id int64, editor Identity) (ICategory, error)
		Find(id int64) ICategory
		ForEach(iterator CategoryIterator)
		Filter(predicate CategoryFilterPredicate) ICategoryCollection
		Map(predicate CategoryMapPredicate) ICategoryCollection
		ListCategoriesByCategoryType(categoryTypeId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection
		ForEachByCategoryType(categoryTypeId int64, iterator CategoryIterator)
		ListCategoriesByCategory(categoryId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection
		ForEachByCategory(categoryId int64, iterator CategoryIterator)
	}
)
