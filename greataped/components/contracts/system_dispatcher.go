package contracts

import (
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
)

type IDispatcher interface {
	Logger() ILogger
	Config() IConfiguration
	FQDN() string
	PublicUrl() string
	Accelerator() IDispatcherCache
	// Atomic initializes an atomic context using a new transaction.
	Atomic(SystemAction)
	// Schedule creates a new recurring event. First parameter is the cron style
	// schedule spec with a callback function as second parameter.
	Schedule(id int64, spec string, callback func(IDispatcher, string)) error
	// Transaction returns the current running transaction.
	Transaction() ITransaction
	// IdentityManager returns system identity manager.
	IdentityManager() IIdentityManager
	// Identity returns the current identity.
	Identity() Identity
	// CurrentUser returns the current user.
	CurrentUser() IUser
	// SignOut logs the current user out.
	SignOut() error

	// Document
	// ------------------------------------------------------------

	// DocumentExists checks whether a specific 'Document' with the provided
	// unique identifier or 'Id' exists in the system.
	DocumentExists(id int64) bool
	// DocumentExistsWhich checks whether a specific 'Document' exists in the system
	// which satisfies the provided condition.
	DocumentExistsWhich(condition DocumentCondition) bool
	// ListDocuments returns a list of all 'Document' instances in the system.
	ListDocuments() IDocumentCollection
	// ForEachDocument loops over all 'Document' instances in the system running
	// the provided iterator for each of them.
	ForEachDocument(iterator DocumentIterator)
	// FilterDocuments returns a filtered list of 'Document' instances based
	// on the provided predicate.
	FilterDocuments(predicate DocumentFilterPredicate) IDocumentCollection
	// MapDocuments loops over all 'Document' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapDocuments(predicate DocumentMapPredicate) IDocumentCollection
	// GetDocument finds a specific 'Document' instance using
	// the provided unique identifier or 'Id'.
	GetDocument(id int64) IDocument
	// AddDocument creates a new 'Document' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddDocument(content string) IDocument
	// AddDocumentWithCustomId creates a new 'Document' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddDocumentWithCustomId(id int64, content string) IDocument
	// LogDocument creates a new 'Document' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogDocument(content string, source string, payload string)
	// UpdateDocument finds the 'Document' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateDocument(id int64, content string) IDocument
	// UpdateDocumentObject finds and updates the 'Document' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateDocumentObject(object IObject, document IDocument) IDocument
	// AddOrUpdateDocumentObject tries to find the 'Document' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateDocumentObject(object IObject, document IDocument) IDocument
	// RemoveDocument finds the 'Document' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveDocument(id int64) IDocument

	// SystemSchedule
	// ------------------------------------------------------------

	// SystemScheduleExists checks whether a specific 'System Schedule' with the provided
	// unique identifier or 'Id' exists in the system.
	SystemScheduleExists(id int64) bool
	// SystemScheduleExistsWhich checks whether a specific 'System Schedule' exists in the system
	// which satisfies the provided condition.
	SystemScheduleExistsWhich(condition SystemScheduleCondition) bool
	// ListSystemSchedules returns a list of all 'System Schedule' instances in the system.
	ListSystemSchedules() ISystemScheduleCollection
	// ForEachSystemSchedule loops over all 'System Schedule' instances in the system running
	// the provided iterator for each of them.
	ForEachSystemSchedule(iterator SystemScheduleIterator)
	// FilterSystemSchedules returns a filtered list of 'System Schedule' instances based
	// on the provided predicate.
	FilterSystemSchedules(predicate SystemScheduleFilterPredicate) ISystemScheduleCollection
	// MapSystemSchedules loops over all 'System Schedule' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapSystemSchedules(predicate SystemScheduleMapPredicate) ISystemScheduleCollection
	// GetSystemSchedule finds a specific 'System Schedule' instance using
	// the provided unique identifier or 'Id'.
	GetSystemSchedule(id int64) ISystemSchedule
	// AddSystemSchedule creates a new 'System Schedule' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddSystemSchedule(enabled bool, config string) ISystemSchedule
	// AddSystemScheduleWithCustomId creates a new 'System Schedule' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddSystemScheduleWithCustomId(id int64, enabled bool, config string) ISystemSchedule
	// LogSystemSchedule creates a new 'System Schedule' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogSystemSchedule(enabled bool, config string, source string, payload string)
	// UpdateSystemSchedule finds the 'System Schedule' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateSystemSchedule(id int64, enabled bool, config string) ISystemSchedule
	// UpdateSystemScheduleObject finds and updates the 'System Schedule' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateSystemScheduleObject(object IObject, systemSchedule ISystemSchedule) ISystemSchedule
	// AddOrUpdateSystemScheduleObject tries to find the 'System Schedule' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateSystemScheduleObject(object IObject, systemSchedule ISystemSchedule) ISystemSchedule
	// RemoveSystemSchedule finds the 'System Schedule' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveSystemSchedule(id int64) ISystemSchedule

	// Identity
	// ------------------------------------------------------------

	// IdentityExists checks whether a specific 'Identity' with the provided
	// unique identifier or 'Id' exists in the system.
	IdentityExists(id int64) bool
	// IdentityExistsWhich checks whether a specific 'Identity' exists in the system
	// which satisfies the provided condition.
	IdentityExistsWhich(condition IdentityCondition) bool
	// ListIdentities returns a list of all 'Identity' instances in the system.
	ListIdentities() IIdentityCollection
	// ForEachIdentity loops over all 'Identity' instances in the system running
	// the provided iterator for each of them.
	ForEachIdentity(iterator IdentityIterator)
	// FilterIdentities returns a filtered list of 'Identity' instances based
	// on the provided predicate.
	FilterIdentities(predicate IdentityFilterPredicate) IIdentityCollection
	// MapIdentities loops over all 'Identity' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapIdentities(predicate IdentityMapPredicate) IIdentityCollection
	// GetIdentity finds a specific 'Identity' instance using
	// the provided unique identifier or 'Id'.
	GetIdentity(id int64) IIdentity
	// AddIdentity creates a new 'Identity' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) IIdentity
	// AddIdentityWithCustomId creates a new 'Identity' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddIdentityWithCustomId(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) IIdentity
	// LogIdentity creates a new 'Identity' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, source string, payload string)
	// UpdateIdentity finds the 'Identity' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) IIdentity
	// UpdateIdentityObject finds and updates the 'Identity' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateIdentityObject(object IObject, identity IIdentity) IIdentity
	// AddOrUpdateIdentityObject tries to find the 'Identity' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateIdentityObject(object IObject, identity IIdentity) IIdentity
	// RemoveIdentity finds the 'Identity' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveIdentity(id int64) IIdentity

	// AccessControl
	// ------------------------------------------------------------

	// AccessControlExists checks whether a specific 'Access Control' with the provided
	// unique identifier or 'Id' exists in the system.
	AccessControlExists(id int64) bool
	// AccessControlExistsWhich checks whether a specific 'Access Control' exists in the system
	// which satisfies the provided condition.
	AccessControlExistsWhich(condition AccessControlCondition) bool
	// ListAccessControls returns a list of all 'Access Control' instances in the system.
	ListAccessControls() IAccessControlCollection
	// ForEachAccessControl loops over all 'Access Control' instances in the system running
	// the provided iterator for each of them.
	ForEachAccessControl(iterator AccessControlIterator)
	// FilterAccessControls returns a filtered list of 'Access Control' instances based
	// on the provided predicate.
	FilterAccessControls(predicate AccessControlFilterPredicate) IAccessControlCollection
	// MapAccessControls loops over all 'Access Control' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapAccessControls(predicate AccessControlMapPredicate) IAccessControlCollection
	// GetAccessControl finds a specific 'Access Control' instance using
	// the provided unique identifier or 'Id'.
	GetAccessControl(id int64) IAccessControl
	// AddAccessControl creates a new 'Access Control' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddAccessControl(key uint64, value uint64) IAccessControl
	// AddAccessControlWithCustomId creates a new 'Access Control' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddAccessControlWithCustomId(id int64, key uint64, value uint64) IAccessControl
	// LogAccessControl creates a new 'Access Control' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogAccessControl(key uint64, value uint64, source string, payload string)
	// UpdateAccessControl finds the 'Access Control' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateAccessControl(id int64, key uint64, value uint64) IAccessControl
	// UpdateAccessControlObject finds and updates the 'Access Control' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateAccessControlObject(object IObject, accessControl IAccessControl) IAccessControl
	// AddOrUpdateAccessControlObject tries to find the 'Access Control' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateAccessControlObject(object IObject, accessControl IAccessControl) IAccessControl
	// RemoveAccessControl finds the 'Access Control' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveAccessControl(id int64) IAccessControl

	// RemoteActivity
	// ------------------------------------------------------------

	// RemoteActivityExists checks whether a specific 'Remote Activity' with the provided
	// unique identifier or 'Id' exists in the system.
	RemoteActivityExists(id int64) bool
	// RemoteActivityExistsWhich checks whether a specific 'Remote Activity' exists in the system
	// which satisfies the provided condition.
	RemoteActivityExistsWhich(condition RemoteActivityCondition) bool
	// ListRemoteActivities returns a list of all 'Remote Activity' instances in the system.
	ListRemoteActivities() IRemoteActivityCollection
	// ForEachRemoteActivity loops over all 'Remote Activity' instances in the system running
	// the provided iterator for each of them.
	ForEachRemoteActivity(iterator RemoteActivityIterator)
	// FilterRemoteActivities returns a filtered list of 'Remote Activity' instances based
	// on the provided predicate.
	FilterRemoteActivities(predicate RemoteActivityFilterPredicate) IRemoteActivityCollection
	// MapRemoteActivities loops over all 'Remote Activity' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapRemoteActivities(predicate RemoteActivityMapPredicate) IRemoteActivityCollection
	// GetRemoteActivity finds a specific 'Remote Activity' instance using
	// the provided unique identifier or 'Id'.
	GetRemoteActivity(id int64) IRemoteActivity
	// AddRemoteActivity creates a new 'Remote Activity' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) IRemoteActivity
	// AddRemoteActivityWithCustomId creates a new 'Remote Activity' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddRemoteActivityWithCustomId(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) IRemoteActivity
	// LogRemoteActivity creates a new 'Remote Activity' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, source string, payload string)
	// UpdateRemoteActivity finds the 'Remote Activity' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) IRemoteActivity
	// UpdateRemoteActivityObject finds and updates the 'Remote Activity' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateRemoteActivityObject(object IObject, remoteActivity IRemoteActivity) IRemoteActivity
	// AddOrUpdateRemoteActivityObject tries to find the 'Remote Activity' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateRemoteActivityObject(object IObject, remoteActivity IRemoteActivity) IRemoteActivity
	// RemoveRemoteActivity finds the 'Remote Activity' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveRemoteActivity(id int64) IRemoteActivity

	// CategoryType
	// ------------------------------------------------------------

	// CategoryTypeExists checks whether a specific 'Category Type' with the provided
	// unique identifier or 'Id' exists in the system.
	CategoryTypeExists(id int64) bool
	// CategoryTypeExistsWhich checks whether a specific 'Category Type' exists in the system
	// which satisfies the provided condition.
	CategoryTypeExistsWhich(condition CategoryTypeCondition) bool
	// ListCategoryTypes returns a list of all 'Category Type' instances in the system.
	ListCategoryTypes() ICategoryTypeCollection
	// ForEachCategoryType loops over all 'Category Type' instances in the system running
	// the provided iterator for each of them.
	ForEachCategoryType(iterator CategoryTypeIterator)
	// FilterCategoryTypes returns a filtered list of 'Category Type' instances based
	// on the provided predicate.
	FilterCategoryTypes(predicate CategoryTypeFilterPredicate) ICategoryTypeCollection
	// MapCategoryTypes loops over all 'Category Type' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapCategoryTypes(predicate CategoryTypeMapPredicate) ICategoryTypeCollection
	// GetCategoryType finds a specific 'Category Type' instance using
	// the provided unique identifier or 'Id'.
	GetCategoryType(id int64) ICategoryType
	// AddCategoryType creates a new 'Category Type' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddCategoryType(description string) ICategoryType
	// AddCategoryTypeWithCustomId creates a new 'Category Type' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddCategoryTypeWithCustomId(id int64, description string) ICategoryType
	// LogCategoryType creates a new 'Category Type' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogCategoryType(description string, source string, payload string)
	// UpdateCategoryType finds the 'Category Type' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateCategoryType(id int64, description string) ICategoryType
	// UpdateCategoryTypeObject finds and updates the 'Category Type' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateCategoryTypeObject(object IObject, categoryType ICategoryType) ICategoryType
	// AddOrUpdateCategoryTypeObject tries to find the 'Category Type' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateCategoryTypeObject(object IObject, categoryType ICategoryType) ICategoryType
	// RemoveCategoryType finds the 'Category Type' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveCategoryType(id int64) ICategoryType

	// Category
	// ------------------------------------------------------------

	// CategoryExists checks whether a specific 'Category' with the provided
	// unique identifier or 'Id' exists in the system.
	CategoryExists(id int64) bool
	// CategoryExistsWhich checks whether a specific 'Category' exists in the system
	// which satisfies the provided condition.
	CategoryExistsWhich(condition CategoryCondition) bool
	// ListCategories returns a list of all 'Category' instances in the system.
	ListCategories() ICategoryCollection
	// ForEachCategory loops over all 'Category' instances in the system running
	// the provided iterator for each of them.
	ForEachCategory(iterator CategoryIterator)
	// FilterCategories returns a filtered list of 'Category' instances based
	// on the provided predicate.
	FilterCategories(predicate CategoryFilterPredicate) ICategoryCollection
	// MapCategories loops over all 'Category' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapCategories(predicate CategoryMapPredicate) ICategoryCollection
	// GetCategory finds a specific 'Category' instance using
	// the provided unique identifier or 'Id'.
	GetCategory(id int64) ICategory
	// AddCategory creates a new 'Category' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddCategory(categoryTypeId int64, categoryId int64, title string, description string) ICategory
	// AddCategoryWithCustomId creates a new 'Category' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddCategoryWithCustomId(id int64, categoryTypeId int64, categoryId int64, title string, description string) ICategory
	// LogCategory creates a new 'Category' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogCategory(categoryTypeId int64, categoryId int64, title string, description string, source string, payload string)
	// UpdateCategory finds the 'Category' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string) ICategory
	// UpdateCategoryObject finds and updates the 'Category' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateCategoryObject(object IObject, category ICategory) ICategory
	// AddOrUpdateCategoryObject tries to find the 'Category' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateCategoryObject(object IObject, category ICategory) ICategory
	// RemoveCategory finds the 'Category' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveCategory(id int64) ICategory
	// ListCategoriesByCategoryType returns a list of all 'Category' instances in the system
	// that are children of the provided 'Category Type' instance.
	ListCategoriesByCategoryType(categoryType ICategoryType) ICategoryCollection
	// ListCategoriesByCategoryTypeId returns a list of all 'Category' instances in the system that are
	// children of the 'Category Type' instance with the provided unique identifier.
	ListCategoriesByCategoryTypeId(categoryTypeId int64) ICategoryCollection
	// ForEachCategoryByCategoryType loops over all 'Category' instances in the system that are children
	// of the provided 'Category Type' instance, running the provided iterator for each of them.
	ForEachCategoryByCategoryType(categoryType ICategoryType, iterator CategoryIterator)
	// ForEachCategoryByCategoryTypeId loops over all 'Category' instances in the system that are children
	// of the 'Category Type' instance with the provided unique identifier,
	// running the provided iterator for each of them.
	ForEachCategoryByCategoryTypeId(categoryTypeId int64, iterator CategoryIterator)
	// ListCategoriesByCategory returns a list of all 'Category' instances in the system
	// that are children of the provided 'Category' instance.
	ListCategoriesByCategory(category ICategory) ICategoryCollection
	// ListCategoriesByCategoryId returns a list of all 'Category' instances in the system that are
	// children of the 'Category' instance with the provided unique identifier.
	ListCategoriesByCategoryId(categoryId int64) ICategoryCollection
	// ForEachCategoryByCategory loops over all 'Category' instances in the system that are children
	// of the provided 'Category' instance, running the provided iterator for each of them.
	ForEachCategoryByCategory(category ICategory, iterator CategoryIterator)
	// ForEachCategoryByCategoryId loops over all 'Category' instances in the system that are children
	// of the 'Category' instance with the provided unique identifier,
	// running the provided iterator for each of them.
	ForEachCategoryByCategoryId(categoryId int64, iterator CategoryIterator)

	// User
	// ------------------------------------------------------------

	// UserExists checks whether a specific 'User' with the provided
	// unique identifier or 'Id' exists in the system.
	UserExists(id int64) bool
	// UserExistsWhich checks whether a specific 'User' exists in the system
	// which satisfies the provided condition.
	UserExistsWhich(condition UserCondition) bool
	// ListUsers returns a list of all 'User' instances in the system.
	ListUsers() IUserCollection
	// ForEachUser loops over all 'User' instances in the system running
	// the provided iterator for each of them.
	ForEachUser(iterator UserIterator)
	// FilterUsers returns a filtered list of 'User' instances based
	// on the provided predicate.
	FilterUsers(predicate UserFilterPredicate) IUserCollection
	// MapUsers loops over all 'User' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapUsers(predicate UserMapPredicate) IUserCollection
	// GetUser finds a specific 'User' instance using
	// the provided unique identifier or 'Id'.
	GetUser(id int64) IUser
	// AddUser creates a new 'User' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddUser(identityId int64, github string) IUser
	// AddUserObject creates a new 'User' instance with an auto-generated unique identifier using
	// the property values in provided instance and adds it to persistent data store and
	// system cache. The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddUserObject(identity IIdentity, user IUser) IUser
	// LogUser creates a new 'User' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogUser(identityId int64, github string, source string, payload string)
	// UpdateUser finds the 'User' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateUser(id int64, github string) IUser
	// UpdateUserObject finds and updates the 'User' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateUserObject(object IObject, user IUser) IUser
	// AddOrUpdateUserObject tries to find the 'User' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateUserObject(object IObject, user IUser) IUser
	// RemoveUser finds the 'User' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveUser(id int64) IUser

	// ActivityPubObject
	// ------------------------------------------------------------

	// ActivityPubObjectExists checks whether a specific 'Activity Pub Object' with the provided
	// unique identifier or 'Id' exists in the system.
	ActivityPubObjectExists(id int64) bool
	// ActivityPubObjectExistsWhich checks whether a specific 'Activity Pub Object' exists in the system
	// which satisfies the provided condition.
	ActivityPubObjectExistsWhich(condition ActivityPubObjectCondition) bool
	// ListActivityPubObjects returns a list of all 'Activity Pub Object' instances in the system.
	ListActivityPubObjects() IActivityPubObjectCollection
	// ForEachActivityPubObject loops over all 'Activity Pub Object' instances in the system running
	// the provided iterator for each of them.
	ForEachActivityPubObject(iterator ActivityPubObjectIterator)
	// FilterActivityPubObjects returns a filtered list of 'Activity Pub Object' instances based
	// on the provided predicate.
	FilterActivityPubObjects(predicate ActivityPubObjectFilterPredicate) IActivityPubObjectCollection
	// MapActivityPubObjects loops over all 'Activity Pub Object' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapActivityPubObjects(predicate ActivityPubObjectMapPredicate) IActivityPubObjectCollection
	// GetActivityPubObject finds a specific 'Activity Pub Object' instance using
	// the provided unique identifier or 'Id'.
	GetActivityPubObject(id int64) IActivityPubObject
	// AddActivityPubObject creates a new 'Activity Pub Object' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubObject() IActivityPubObject
	// AddActivityPubObjectWithCustomId creates a new 'Activity Pub Object' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubObjectWithCustomId(id int64) IActivityPubObject
	// LogActivityPubObject creates a new 'Activity Pub Object' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogActivityPubObject(source string, payload string)
	// UpdateActivityPubObject finds the 'Activity Pub Object' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateActivityPubObject(id int64) IActivityPubObject
	// UpdateActivityPubObjectObject finds and updates the 'Activity Pub Object' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateActivityPubObjectObject(object IObject, activityPubObject IActivityPubObject) IActivityPubObject
	// AddOrUpdateActivityPubObjectObject tries to find the 'Activity Pub Object' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateActivityPubObjectObject(object IObject, activityPubObject IActivityPubObject) IActivityPubObject
	// RemoveActivityPubObject finds the 'Activity Pub Object' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveActivityPubObject(id int64) IActivityPubObject

	// Spi
	// ------------------------------------------------------------

	// SpiExists checks whether a specific 'Spi' with the provided
	// unique identifier or 'Id' exists in the system.
	SpiExists(id int64) bool
	// SpiExistsWhich checks whether a specific 'Spi' exists in the system
	// which satisfies the provided condition.
	SpiExistsWhich(condition SpiCondition) bool
	// ListSpis returns a list of all 'Spi' instances in the system.
	ListSpis() ISpiCollection
	// ForEachSpi loops over all 'Spi' instances in the system running
	// the provided iterator for each of them.
	ForEachSpi(iterator SpiIterator)
	// FilterSpis returns a filtered list of 'Spi' instances based
	// on the provided predicate.
	FilterSpis(predicate SpiFilterPredicate) ISpiCollection
	// MapSpis loops over all 'Spi' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapSpis(predicate SpiMapPredicate) ISpiCollection
	// GetSpi finds a specific 'Spi' instance using
	// the provided unique identifier or 'Id'.
	GetSpi(id int64) ISpi
	// AddSpi creates a new 'Spi' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddSpi() ISpi
	// AddSpiWithCustomId creates a new 'Spi' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddSpiWithCustomId(id int64) ISpi
	// LogSpi creates a new 'Spi' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogSpi(source string, payload string)
	// UpdateSpi finds the 'Spi' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateSpi(id int64) ISpi
	// UpdateSpiObject finds and updates the 'Spi' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateSpiObject(object IObject, spi ISpi) ISpi
	// AddOrUpdateSpiObject tries to find the 'Spi' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateSpiObject(object IObject, spi ISpi) ISpi
	// RemoveSpi finds the 'Spi' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveSpi(id int64) ISpi
	Echo(document IDocument) (IEchoResult, error)

	// CustomError
	// ------------------------------------------------------------

	// CustomErrorExists checks whether a specific 'Custom Error' with the provided
	// unique identifier or 'Id' exists in the system.
	CustomErrorExists(id int64) bool
	// CustomErrorExistsWhich checks whether a specific 'Custom Error' exists in the system
	// which satisfies the provided condition.
	CustomErrorExistsWhich(condition CustomErrorCondition) bool
	// ListCustomErrors returns a list of all 'Custom Error' instances in the system.
	ListCustomErrors() ICustomErrorCollection
	// ForEachCustomError loops over all 'Custom Error' instances in the system running
	// the provided iterator for each of them.
	ForEachCustomError(iterator CustomErrorIterator)
	// FilterCustomErrors returns a filtered list of 'Custom Error' instances based
	// on the provided predicate.
	FilterCustomErrors(predicate CustomErrorFilterPredicate) ICustomErrorCollection
	// MapCustomErrors loops over all 'Custom Error' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapCustomErrors(predicate CustomErrorMapPredicate) ICustomErrorCollection
	// GetCustomError finds a specific 'Custom Error' instance using
	// the provided unique identifier or 'Id'.
	GetCustomError(id int64) ICustomError
	// AddCustomError creates a new 'Custom Error' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddCustomError() ICustomError
	// AddCustomErrorWithCustomId creates a new 'Custom Error' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddCustomErrorWithCustomId(id int64) ICustomError
	// LogCustomError creates a new 'Custom Error' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogCustomError(source string, payload string)
	// UpdateCustomError finds the 'Custom Error' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateCustomError(id int64) ICustomError
	// UpdateCustomErrorObject finds and updates the 'Custom Error' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateCustomErrorObject(object IObject, customError ICustomError) ICustomError
	// AddOrUpdateCustomErrorObject tries to find the 'Custom Error' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateCustomErrorObject(object IObject, customError ICustomError) ICustomError
	// RemoveCustomError finds the 'Custom Error' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveCustomError(id int64) ICustomError
	ResolveError(document IDocument) (IResolveErrorResult, error)

	// NewDocument creates a new 'Document' instance using the provided property values.
	NewDocument(id int64, content string) (IDocument, error)
	// NewDocuments creates an empty in-memory 'Document' collection which is not thread-safe.
	NewDocuments() IDocumentCollection
	// NewSystemSchedule creates a new 'System Schedule' instance using the provided property values.
	NewSystemSchedule(id int64, enabled bool, config string) (ISystemSchedule, error)
	// NewSystemSchedules creates an empty in-memory 'System Schedule' collection which is not thread-safe.
	NewSystemSchedules() ISystemScheduleCollection
	// NewIdentity creates a new 'Identity' instance using the provided property values.
	NewIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) (IIdentity, error)
	// NewIdentities creates an empty in-memory 'Identity' collection which is not thread-safe.
	NewIdentities() IIdentityCollection
	// NewAccessControl creates a new 'Access Control' instance using the provided property values.
	NewAccessControl(id int64, key uint64, value uint64) (IAccessControl, error)
	// NewAccessControls creates an empty in-memory 'Access Control' collection which is not thread-safe.
	NewAccessControls() IAccessControlCollection
	// NewRemoteActivity creates a new 'Remote Activity' instance using the provided property values.
	NewRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) (IRemoteActivity, error)
	// NewRemoteActivities creates an empty in-memory 'Remote Activity' collection which is not thread-safe.
	NewRemoteActivities() IRemoteActivityCollection
	// NewCategoryType creates a new 'Category Type' instance using the provided property values.
	NewCategoryType(id int64, description string) (ICategoryType, error)
	// NewCategoryTypes creates an empty in-memory 'Category Type' collection which is not thread-safe.
	NewCategoryTypes() ICategoryTypeCollection
	// NewCategory creates a new 'Category' instance using the provided property values.
	NewCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string) (ICategory, error)
	// NewCategories creates an empty in-memory 'Category' collection which is not thread-safe.
	NewCategories() ICategoryCollection
	// NewUser creates a new 'User' instance using the provided property values.
	NewUser(id int64, github string) (IUser, error)
	// NewUsers creates an empty in-memory 'User' collection which is not thread-safe.
	NewUsers() IUserCollection
	// NewActivityPubObject creates a new 'Activity Pub Object' instance using the provided property values.
	NewActivityPubObject() (IActivityPubObject, error)
	// NewActivityPubObjects creates an empty in-memory 'Activity Pub Object' collection which is not thread-safe.
	NewActivityPubObjects() IActivityPubObjectCollection
	// NewSpi creates a new 'Spi' instance using the provided property values.
	NewSpi() (ISpi, error)
	// NewSpis creates an empty in-memory 'Spi' collection which is not thread-safe.
	NewSpis() ISpiCollection
	// NewCustomError creates a new 'Custom Error' instance using the provided property values.
	NewCustomError() (ICustomError, error)
	// NewCustomErrors creates an empty in-memory 'Custom Error' collection which is not thread-safe.
	NewCustomErrors() ICustomErrorCollection
	// NewEchoResult creates a new result container for 'Echo' system action.
	NewEchoResult(document IDocument) IEchoResult
	// NewResolveErrorResult creates a new result container for 'Resolve Error' system action.
	NewResolveErrorResult() IResolveErrorResult
	// Assert asserts the provided condition and panics if the assertion is not valid.
	Assert(condition bool) IAssertionResult
	// AssertNoError panics if the provided error is not nil.
	AssertNoError(error)
	// Ensure panics if any of the provided errors is not nil.
	Ensure(...error)
	// AssertNull panics if the provided interfaces is not nil.
	AssertNull(interface{}) IAssertionResult
	// AssertNotNull panics if the provided interfaces is nil.
	AssertNotNull(interface{}) IAssertionResult
	// AssertEmpty panic if the provided string is not empty. Trims the spaces in the string first.
	AssertEmpty(input string) IAssertionResult
	// AssertNotEmpty panic if the provided string is empty. Trims the spaces in the string first.
	AssertNotEmpty(input string) IAssertionResult
	// Format provides a wrapper around fmt.Sprintf
	Format(format string, args ...interface{}) string
	// Sort sorts the provided slice using the provided comparator function.
	Sort(slice interface{}, less func(a, b int) bool)
	// Search searches the input for any or all of the words in criteria.
	Search(input, criteria string) bool
	// Email sends an email message asynchronously.
	Email(destination string, format string, args ...interface{})
	// SMS sends an sms message asynchronously.
	SMS(destination string, format string, args ...interface{})
	// GenerateTrackingNumber returns a new random tracking number between 100000 and 999999.
	GenerateTrackingNumber() uint32
	// GenerateUUID returns a new universal unique identifier.
	GenerateUUID() string
	// GenerateSalt returns a random salt string.
	GenerateSalt() string
	// GenerateHash returns a new hash from a string value and a salt.
	GenerateHash(value string, salt string) string
	// GenerateJwtToken returns a new jwt token.
	GenerateJwtToken() string
	// VerifyJwtToken verifies jwt token.
	VerifyJwtToken(token string) error
	// GenerateCode returns a random string code.
	GenerateCode() string
	// GenerateRSAKeyPair returns a new pair of public and private keys.
	GenerateRSAKeyPair() (string, string, error)
	// UnixNano returns the number of nanoseconds elapsed
	// since January 1, 1970 UTC. The result is undefined if the Unix time
	// in nanoseconds cannot be represented by an int64 (a date before the year
	// 1678 or after 2262). The result does not depend on the location
	UnixNano() int64
	// Trim trims the spaces in the input string
	Trim(input string) string
	// Contains reports whether substr is within s.
	Contains(input, substr string) bool
	// IsEmpty checks whether the provided string is empty. Trims the spaces in the string first.
	IsEmpty(input string) bool
	// IsNotEmpty checks whether the provided string is not empty. Trims the spaces in the string first.
	IsNotEmpty(input string) bool
	// IsSet returns true if the provided int64 is not zero.
	IsSet(id int64) bool
	// Join concatenates the elements of its first argument to create a single string. The separator
	// string is placed between elements in the resulting string.
	Join(elements []string, separator string) string

	IsTestEnvironment() bool
	IsDevelopmentEnvironment() bool
	IsStagingEnvironment() bool
	IsProductionEnvironment() bool

	GetActivityStream(url string, data []byte, output interface{}) error
	PostActivityStream(url string, data []byte, output interface{}) error
	GetActivityStreamSigned(url, keyId, privateKey string, data []byte, output interface{}) error
	PostActivityStreamSigned(url, keyId, privateKey string, data []byte, output interface{}) error
}
