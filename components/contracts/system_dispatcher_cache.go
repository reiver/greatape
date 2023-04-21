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

	// ActivityPubIncomingActivity
	// ------------------------------------------------------------

	// Returns a list of all 'Activity Pub Incoming Activity' instances in the system
	// that are children of the provided 'Identity' instance.
	ListActivityPubIncomingActivitiesByIdentity(identity IIdentity) IActivityPubIncomingActivityCollection
	// Returns a list of all 'Activity Pub Incoming Activity' instances in the system that are
	// children of the 'Identity' instance with the provided unique identifier.
	ListActivityPubIncomingActivitiesByIdentityId(identityId int64) IActivityPubIncomingActivityCollection
	// Loops over all 'Activity Pub Incoming Activity' instances in the system that are children
	// of the provided 'Identity' instance, running the provided iterator for each of them.
	ForEachActivityPubIncomingActivityByIdentity(identity IIdentity, iterator ActivityPubIncomingActivityIterator)
	// Loops over all 'Activity Pub Incoming Activity' instances in the system that are children
	// of the 'Identity' instance with the provided unique identifier,
	// running the provided iterator for each of them.
	ForEachActivityPubIncomingActivityByIdentityId(identityId int64, iterator ActivityPubIncomingActivityIterator)

	// ActivityPubOutgoingActivity
	// ------------------------------------------------------------

	// Returns a list of all 'Activity Pub Outgoing Activity' instances in the system
	// that are children of the provided 'Identity' instance.
	ListActivityPubOutgoingActivitiesByIdentity(identity IIdentity) IActivityPubOutgoingActivityCollection
	// Returns a list of all 'Activity Pub Outgoing Activity' instances in the system that are
	// children of the 'Identity' instance with the provided unique identifier.
	ListActivityPubOutgoingActivitiesByIdentityId(identityId int64) IActivityPubOutgoingActivityCollection
	// Loops over all 'Activity Pub Outgoing Activity' instances in the system that are children
	// of the provided 'Identity' instance, running the provided iterator for each of them.
	ForEachActivityPubOutgoingActivityByIdentity(identity IIdentity, iterator ActivityPubOutgoingActivityIterator)
	// Loops over all 'Activity Pub Outgoing Activity' instances in the system that are children
	// of the 'Identity' instance with the provided unique identifier,
	// running the provided iterator for each of them.
	ForEachActivityPubOutgoingActivityByIdentityId(identityId int64, iterator ActivityPubOutgoingActivityIterator)
}
