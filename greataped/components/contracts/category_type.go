package contracts

import . "github.com/xeronith/diamante/contracts/security"

var CategoryTypePassThroughFilter = func(ICategoryType) bool { return true }

type (
	CategoryTypes               []ICategoryType
	CategoryTypeIterator        func(ICategoryType)
	CategoryTypeCondition       func(ICategoryType) bool
	CategoryTypeFilterPredicate func(ICategoryType) bool
	CategoryTypeMapPredicate    func(ICategoryType) ICategoryType
	CategoryTypeCacheCallback   func()

	ICategoryType interface {
		IObject
		// Description returns 'Description' of this 'CategoryType' instance.
		Description() string
		// UpdateDescription directly updates 'Description' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateDescription(description string, editor Identity)
		// UpdateDescriptionAtomic updates 'Description' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateDescriptionAtomic(transaction ITransaction, description string, editor Identity)
	}

	ICategoryTypeCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() ICategoryType
		Append(categoryType ICategoryType)
		ForEach(CategoryTypeIterator)
		Array() CategoryTypes
	}

	ICategoryTypeManager interface {
		ISystemComponent
		OnCacheChanged(CategoryTypeCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition CategoryTypeCondition) bool
		ListCategoryTypes(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryTypeCollection
		GetCategoryType(id int64, editor Identity) (ICategoryType, error)
		AddCategoryType(description string, editor Identity) (ICategoryType, error)
		AddCategoryTypeWithCustomId(id int64, description string, editor Identity) (ICategoryType, error)
		AddCategoryTypeObject(categoryType ICategoryType, editor Identity) (ICategoryType, error)
		AddCategoryTypeAtomic(transaction ITransaction, description string, editor Identity) (ICategoryType, error)
		AddCategoryTypeWithCustomIdAtomic(id int64, transaction ITransaction, description string, editor Identity) (ICategoryType, error)
		AddCategoryTypeObjectAtomic(transaction ITransaction, categoryType ICategoryType, editor Identity) (ICategoryType, error)
		Log(description string, source string, editor Identity, payload string)
		UpdateCategoryType(id int64, description string, editor Identity) (ICategoryType, error)
		UpdateCategoryTypeObject(id int64, categoryType ICategoryType, editor Identity) (ICategoryType, error)
		UpdateCategoryTypeAtomic(transaction ITransaction, id int64, description string, editor Identity) (ICategoryType, error)
		UpdateCategoryTypeObjectAtomic(transaction ITransaction, id int64, categoryType ICategoryType, editor Identity) (ICategoryType, error)
		AddOrUpdateCategoryTypeObject(id int64, categoryType ICategoryType, editor Identity) (ICategoryType, error)
		AddOrUpdateCategoryTypeObjectAtomic(transaction ITransaction, id int64, categoryType ICategoryType, editor Identity) (ICategoryType, error)
		RemoveCategoryType(id int64, editor Identity) (ICategoryType, error)
		RemoveCategoryTypeAtomic(transaction ITransaction, id int64, editor Identity) (ICategoryType, error)
		Find(id int64) ICategoryType
		ForEach(iterator CategoryTypeIterator)
		Filter(predicate CategoryTypeFilterPredicate) ICategoryTypeCollection
		Map(predicate CategoryTypeMapPredicate) ICategoryTypeCollection
	}
)
