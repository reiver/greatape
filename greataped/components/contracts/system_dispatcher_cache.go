package contracts

type IDispatcherCache interface {

	// Category
	// ------------------------------------------------------------

	// Returns a list of all 'Category' instances in the system
	// that are children of the provided 'Category Type' instance.
	ListCategoriesByCategoryType(categoryType ICategoryType) ICategoryCollection
	// Returns a list of all 'Category' instances in the system that are
	// children of the 'Category Type' instance with the provided unique identifier.
	ListCategoriesByCategoryTypeId(categoryTypeId int64) ICategoryCollection
	// Loops over all 'Category' instances in the system that are children
	// of the provided 'Category Type' instance, running the provided iterator for each of them.
	ForEachCategoryByCategoryType(categoryType ICategoryType, iterator CategoryIterator)
	// Loops over all 'Category' instances in the system that are children
	// of the 'Category Type' instance with the provided unique identifier,
	// running the provided iterator for each of them.
	ForEachCategoryByCategoryTypeId(categoryTypeId int64, iterator CategoryIterator)
	// Returns a list of all 'Category' instances in the system
	// that are children of the provided 'Category' instance.
	ListCategoriesByCategory(category ICategory) ICategoryCollection
	// Returns a list of all 'Category' instances in the system that are
	// children of the 'Category' instance with the provided unique identifier.
	ListCategoriesByCategoryId(categoryId int64) ICategoryCollection
	// Loops over all 'Category' instances in the system that are children
	// of the provided 'Category' instance, running the provided iterator for each of them.
	ForEachCategoryByCategory(category ICategory, iterator CategoryIterator)
	// Loops over all 'Category' instances in the system that are children
	// of the 'Category' instance with the provided unique identifier,
	// running the provided iterator for each of them.
	ForEachCategoryByCategoryId(categoryId int64, iterator CategoryIterator)
}
