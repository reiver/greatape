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

	// ActivityPubActivity
	// ------------------------------------------------------------

	// ActivityPubActivityExists checks whether a specific 'Activity Pub Activity' with the provided
	// unique identifier or 'Id' exists in the system.
	ActivityPubActivityExists(id int64) bool
	// ActivityPubActivityExistsWhich checks whether a specific 'Activity Pub Activity' exists in the system
	// which satisfies the provided condition.
	ActivityPubActivityExistsWhich(condition ActivityPubActivityCondition) bool
	// ListActivityPubActivities returns a list of all 'Activity Pub Activity' instances in the system.
	ListActivityPubActivities() IActivityPubActivityCollection
	// ForEachActivityPubActivity loops over all 'Activity Pub Activity' instances in the system running
	// the provided iterator for each of them.
	ForEachActivityPubActivity(iterator ActivityPubActivityIterator)
	// FilterActivityPubActivities returns a filtered list of 'Activity Pub Activity' instances based
	// on the provided predicate.
	FilterActivityPubActivities(predicate ActivityPubActivityFilterPredicate) IActivityPubActivityCollection
	// MapActivityPubActivities loops over all 'Activity Pub Activity' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapActivityPubActivities(predicate ActivityPubActivityMapPredicate) IActivityPubActivityCollection
	// GetActivityPubActivity finds a specific 'Activity Pub Activity' instance using
	// the provided unique identifier or 'Id'.
	GetActivityPubActivity(id int64) IActivityPubActivity
	// AddActivityPubActivity creates a new 'Activity Pub Activity' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubActivity() IActivityPubActivity
	// AddActivityPubActivityWithCustomId creates a new 'Activity Pub Activity' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubActivityWithCustomId(id int64) IActivityPubActivity
	// LogActivityPubActivity creates a new 'Activity Pub Activity' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogActivityPubActivity(source string, payload string)
	// UpdateActivityPubActivity finds the 'Activity Pub Activity' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateActivityPubActivity(id int64) IActivityPubActivity
	// UpdateActivityPubActivityObject finds and updates the 'Activity Pub Activity' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateActivityPubActivityObject(object IObject, activityPubActivity IActivityPubActivity) IActivityPubActivity
	// AddOrUpdateActivityPubActivityObject tries to find the 'Activity Pub Activity' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateActivityPubActivityObject(object IObject, activityPubActivity IActivityPubActivity) IActivityPubActivity
	// RemoveActivityPubActivity finds the 'Activity Pub Activity' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveActivityPubActivity(id int64) IActivityPubActivity

	// ActivityPubPublicKey
	// ------------------------------------------------------------

	// ActivityPubPublicKeyExists checks whether a specific 'Activity Pub Public Key' with the provided
	// unique identifier or 'Id' exists in the system.
	ActivityPubPublicKeyExists(id int64) bool
	// ActivityPubPublicKeyExistsWhich checks whether a specific 'Activity Pub Public Key' exists in the system
	// which satisfies the provided condition.
	ActivityPubPublicKeyExistsWhich(condition ActivityPubPublicKeyCondition) bool
	// ListActivityPubPublicKeys returns a list of all 'Activity Pub Public Key' instances in the system.
	ListActivityPubPublicKeys() IActivityPubPublicKeyCollection
	// ForEachActivityPubPublicKey loops over all 'Activity Pub Public Key' instances in the system running
	// the provided iterator for each of them.
	ForEachActivityPubPublicKey(iterator ActivityPubPublicKeyIterator)
	// FilterActivityPubPublicKeys returns a filtered list of 'Activity Pub Public Key' instances based
	// on the provided predicate.
	FilterActivityPubPublicKeys(predicate ActivityPubPublicKeyFilterPredicate) IActivityPubPublicKeyCollection
	// MapActivityPubPublicKeys loops over all 'Activity Pub Public Key' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapActivityPubPublicKeys(predicate ActivityPubPublicKeyMapPredicate) IActivityPubPublicKeyCollection
	// GetActivityPubPublicKey finds a specific 'Activity Pub Public Key' instance using
	// the provided unique identifier or 'Id'.
	GetActivityPubPublicKey(id int64) IActivityPubPublicKey
	// AddActivityPubPublicKey creates a new 'Activity Pub Public Key' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubPublicKey() IActivityPubPublicKey
	// AddActivityPubPublicKeyWithCustomId creates a new 'Activity Pub Public Key' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubPublicKeyWithCustomId(id int64) IActivityPubPublicKey
	// LogActivityPubPublicKey creates a new 'Activity Pub Public Key' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogActivityPubPublicKey(source string, payload string)
	// UpdateActivityPubPublicKey finds the 'Activity Pub Public Key' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateActivityPubPublicKey(id int64) IActivityPubPublicKey
	// UpdateActivityPubPublicKeyObject finds and updates the 'Activity Pub Public Key' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateActivityPubPublicKeyObject(object IObject, activityPubPublicKey IActivityPubPublicKey) IActivityPubPublicKey
	// AddOrUpdateActivityPubPublicKeyObject tries to find the 'Activity Pub Public Key' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateActivityPubPublicKeyObject(object IObject, activityPubPublicKey IActivityPubPublicKey) IActivityPubPublicKey
	// RemoveActivityPubPublicKey finds the 'Activity Pub Public Key' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveActivityPubPublicKey(id int64) IActivityPubPublicKey

	// ActivityPubLink
	// ------------------------------------------------------------

	// ActivityPubLinkExists checks whether a specific 'Activity Pub Link' with the provided
	// unique identifier or 'Id' exists in the system.
	ActivityPubLinkExists(id int64) bool
	// ActivityPubLinkExistsWhich checks whether a specific 'Activity Pub Link' exists in the system
	// which satisfies the provided condition.
	ActivityPubLinkExistsWhich(condition ActivityPubLinkCondition) bool
	// ListActivityPubLinks returns a list of all 'Activity Pub Link' instances in the system.
	ListActivityPubLinks() IActivityPubLinkCollection
	// ForEachActivityPubLink loops over all 'Activity Pub Link' instances in the system running
	// the provided iterator for each of them.
	ForEachActivityPubLink(iterator ActivityPubLinkIterator)
	// FilterActivityPubLinks returns a filtered list of 'Activity Pub Link' instances based
	// on the provided predicate.
	FilterActivityPubLinks(predicate ActivityPubLinkFilterPredicate) IActivityPubLinkCollection
	// MapActivityPubLinks loops over all 'Activity Pub Link' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapActivityPubLinks(predicate ActivityPubLinkMapPredicate) IActivityPubLinkCollection
	// GetActivityPubLink finds a specific 'Activity Pub Link' instance using
	// the provided unique identifier or 'Id'.
	GetActivityPubLink(id int64) IActivityPubLink
	// AddActivityPubLink creates a new 'Activity Pub Link' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubLink() IActivityPubLink
	// AddActivityPubLinkWithCustomId creates a new 'Activity Pub Link' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubLinkWithCustomId(id int64) IActivityPubLink
	// LogActivityPubLink creates a new 'Activity Pub Link' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogActivityPubLink(source string, payload string)
	// UpdateActivityPubLink finds the 'Activity Pub Link' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateActivityPubLink(id int64) IActivityPubLink
	// UpdateActivityPubLinkObject finds and updates the 'Activity Pub Link' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateActivityPubLinkObject(object IObject, activityPubLink IActivityPubLink) IActivityPubLink
	// AddOrUpdateActivityPubLinkObject tries to find the 'Activity Pub Link' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateActivityPubLinkObject(object IObject, activityPubLink IActivityPubLink) IActivityPubLink
	// RemoveActivityPubLink finds the 'Activity Pub Link' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveActivityPubLink(id int64) IActivityPubLink

	// ActivityPubMedia
	// ------------------------------------------------------------

	// ActivityPubMediaExists checks whether a specific 'Activity Pub Media' with the provided
	// unique identifier or 'Id' exists in the system.
	ActivityPubMediaExists(id int64) bool
	// ActivityPubMediaExistsWhich checks whether a specific 'Activity Pub Media' exists in the system
	// which satisfies the provided condition.
	ActivityPubMediaExistsWhich(condition ActivityPubMediaCondition) bool
	// ListActivityPubMedias returns a list of all 'Activity Pub Media' instances in the system.
	ListActivityPubMedias() IActivityPubMediaCollection
	// ForEachActivityPubMedia loops over all 'Activity Pub Media' instances in the system running
	// the provided iterator for each of them.
	ForEachActivityPubMedia(iterator ActivityPubMediaIterator)
	// FilterActivityPubMedias returns a filtered list of 'Activity Pub Media' instances based
	// on the provided predicate.
	FilterActivityPubMedias(predicate ActivityPubMediaFilterPredicate) IActivityPubMediaCollection
	// MapActivityPubMedias loops over all 'Activity Pub Media' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapActivityPubMedias(predicate ActivityPubMediaMapPredicate) IActivityPubMediaCollection
	// GetActivityPubMedia finds a specific 'Activity Pub Media' instance using
	// the provided unique identifier or 'Id'.
	GetActivityPubMedia(id int64) IActivityPubMedia
	// AddActivityPubMedia creates a new 'Activity Pub Media' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubMedia() IActivityPubMedia
	// AddActivityPubMediaWithCustomId creates a new 'Activity Pub Media' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubMediaWithCustomId(id int64) IActivityPubMedia
	// LogActivityPubMedia creates a new 'Activity Pub Media' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogActivityPubMedia(source string, payload string)
	// UpdateActivityPubMedia finds the 'Activity Pub Media' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateActivityPubMedia(id int64) IActivityPubMedia
	// UpdateActivityPubMediaObject finds and updates the 'Activity Pub Media' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateActivityPubMediaObject(object IObject, activityPubMedia IActivityPubMedia) IActivityPubMedia
	// AddOrUpdateActivityPubMediaObject tries to find the 'Activity Pub Media' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateActivityPubMediaObject(object IObject, activityPubMedia IActivityPubMedia) IActivityPubMedia
	// RemoveActivityPubMedia finds the 'Activity Pub Media' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveActivityPubMedia(id int64) IActivityPubMedia

	// ActivityPubIncomingActivity
	// ------------------------------------------------------------

	// ActivityPubIncomingActivityExists checks whether a specific 'Activity Pub Incoming Activity' with the provided
	// unique identifier or 'Id' exists in the system.
	ActivityPubIncomingActivityExists(id int64) bool
	// ActivityPubIncomingActivityExistsWhich checks whether a specific 'Activity Pub Incoming Activity' exists in the system
	// which satisfies the provided condition.
	ActivityPubIncomingActivityExistsWhich(condition ActivityPubIncomingActivityCondition) bool
	// ListActivityPubIncomingActivities returns a list of all 'Activity Pub Incoming Activity' instances in the system.
	ListActivityPubIncomingActivities() IActivityPubIncomingActivityCollection
	// ForEachActivityPubIncomingActivity loops over all 'Activity Pub Incoming Activity' instances in the system running
	// the provided iterator for each of them.
	ForEachActivityPubIncomingActivity(iterator ActivityPubIncomingActivityIterator)
	// FilterActivityPubIncomingActivities returns a filtered list of 'Activity Pub Incoming Activity' instances based
	// on the provided predicate.
	FilterActivityPubIncomingActivities(predicate ActivityPubIncomingActivityFilterPredicate) IActivityPubIncomingActivityCollection
	// MapActivityPubIncomingActivities loops over all 'Activity Pub Incoming Activity' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapActivityPubIncomingActivities(predicate ActivityPubIncomingActivityMapPredicate) IActivityPubIncomingActivityCollection
	// GetActivityPubIncomingActivity finds a specific 'Activity Pub Incoming Activity' instance using
	// the provided unique identifier or 'Id'.
	GetActivityPubIncomingActivity(id int64) IActivityPubIncomingActivity
	// AddActivityPubIncomingActivity creates a new 'Activity Pub Incoming Activity' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubIncomingActivity
	// AddActivityPubIncomingActivityWithCustomId creates a new 'Activity Pub Incoming Activity' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubIncomingActivityWithCustomId(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubIncomingActivity
	// LogActivityPubIncomingActivity creates a new 'Activity Pub Incoming Activity' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, payload string)
	// UpdateActivityPubIncomingActivity finds the 'Activity Pub Incoming Activity' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubIncomingActivity
	// UpdateActivityPubIncomingActivityObject finds and updates the 'Activity Pub Incoming Activity' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateActivityPubIncomingActivityObject(object IObject, activityPubIncomingActivity IActivityPubIncomingActivity) IActivityPubIncomingActivity
	// AddOrUpdateActivityPubIncomingActivityObject tries to find the 'Activity Pub Incoming Activity' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateActivityPubIncomingActivityObject(object IObject, activityPubIncomingActivity IActivityPubIncomingActivity) IActivityPubIncomingActivity
	// RemoveActivityPubIncomingActivity finds the 'Activity Pub Incoming Activity' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveActivityPubIncomingActivity(id int64) IActivityPubIncomingActivity
	// ListActivityPubIncomingActivitiesByIdentity returns a list of all 'Activity Pub Incoming Activity' instances in the system
	// that are children of the provided 'Identity' instance.
	ListActivityPubIncomingActivitiesByIdentity(identity IIdentity) IActivityPubIncomingActivityCollection
	// ListActivityPubIncomingActivitiesByIdentityId returns a list of all 'Activity Pub Incoming Activity' instances in the system that are
	// children of the 'Identity' instance with the provided unique identifier.
	ListActivityPubIncomingActivitiesByIdentityId(identityId int64) IActivityPubIncomingActivityCollection
	// ForEachActivityPubIncomingActivityByIdentity loops over all 'Activity Pub Incoming Activity' instances in the system that are children
	// of the provided 'Identity' instance, running the provided iterator for each of them.
	ForEachActivityPubIncomingActivityByIdentity(identity IIdentity, iterator ActivityPubIncomingActivityIterator)
	// ForEachActivityPubIncomingActivityByIdentityId loops over all 'Activity Pub Incoming Activity' instances in the system that are children
	// of the 'Identity' instance with the provided unique identifier,
	// running the provided iterator for each of them.
	ForEachActivityPubIncomingActivityByIdentityId(identityId int64, iterator ActivityPubIncomingActivityIterator)

	// ActivityPubOutgoingActivity
	// ------------------------------------------------------------

	// ActivityPubOutgoingActivityExists checks whether a specific 'Activity Pub Outgoing Activity' with the provided
	// unique identifier or 'Id' exists in the system.
	ActivityPubOutgoingActivityExists(id int64) bool
	// ActivityPubOutgoingActivityExistsWhich checks whether a specific 'Activity Pub Outgoing Activity' exists in the system
	// which satisfies the provided condition.
	ActivityPubOutgoingActivityExistsWhich(condition ActivityPubOutgoingActivityCondition) bool
	// ListActivityPubOutgoingActivities returns a list of all 'Activity Pub Outgoing Activity' instances in the system.
	ListActivityPubOutgoingActivities() IActivityPubOutgoingActivityCollection
	// ForEachActivityPubOutgoingActivity loops over all 'Activity Pub Outgoing Activity' instances in the system running
	// the provided iterator for each of them.
	ForEachActivityPubOutgoingActivity(iterator ActivityPubOutgoingActivityIterator)
	// FilterActivityPubOutgoingActivities returns a filtered list of 'Activity Pub Outgoing Activity' instances based
	// on the provided predicate.
	FilterActivityPubOutgoingActivities(predicate ActivityPubOutgoingActivityFilterPredicate) IActivityPubOutgoingActivityCollection
	// MapActivityPubOutgoingActivities loops over all 'Activity Pub Outgoing Activity' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapActivityPubOutgoingActivities(predicate ActivityPubOutgoingActivityMapPredicate) IActivityPubOutgoingActivityCollection
	// GetActivityPubOutgoingActivity finds a specific 'Activity Pub Outgoing Activity' instance using
	// the provided unique identifier or 'Id'.
	GetActivityPubOutgoingActivity(id int64) IActivityPubOutgoingActivity
	// AddActivityPubOutgoingActivity creates a new 'Activity Pub Outgoing Activity' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubOutgoingActivity
	// AddActivityPubOutgoingActivityWithCustomId creates a new 'Activity Pub Outgoing Activity' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubOutgoingActivityWithCustomId(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubOutgoingActivity
	// LogActivityPubOutgoingActivity creates a new 'Activity Pub Outgoing Activity' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, payload string)
	// UpdateActivityPubOutgoingActivity finds the 'Activity Pub Outgoing Activity' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) IActivityPubOutgoingActivity
	// UpdateActivityPubOutgoingActivityObject finds and updates the 'Activity Pub Outgoing Activity' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateActivityPubOutgoingActivityObject(object IObject, activityPubOutgoingActivity IActivityPubOutgoingActivity) IActivityPubOutgoingActivity
	// AddOrUpdateActivityPubOutgoingActivityObject tries to find the 'Activity Pub Outgoing Activity' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateActivityPubOutgoingActivityObject(object IObject, activityPubOutgoingActivity IActivityPubOutgoingActivity) IActivityPubOutgoingActivity
	// RemoveActivityPubOutgoingActivity finds the 'Activity Pub Outgoing Activity' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveActivityPubOutgoingActivity(id int64) IActivityPubOutgoingActivity
	// ListActivityPubOutgoingActivitiesByIdentity returns a list of all 'Activity Pub Outgoing Activity' instances in the system
	// that are children of the provided 'Identity' instance.
	ListActivityPubOutgoingActivitiesByIdentity(identity IIdentity) IActivityPubOutgoingActivityCollection
	// ListActivityPubOutgoingActivitiesByIdentityId returns a list of all 'Activity Pub Outgoing Activity' instances in the system that are
	// children of the 'Identity' instance with the provided unique identifier.
	ListActivityPubOutgoingActivitiesByIdentityId(identityId int64) IActivityPubOutgoingActivityCollection
	// ForEachActivityPubOutgoingActivityByIdentity loops over all 'Activity Pub Outgoing Activity' instances in the system that are children
	// of the provided 'Identity' instance, running the provided iterator for each of them.
	ForEachActivityPubOutgoingActivityByIdentity(identity IIdentity, iterator ActivityPubOutgoingActivityIterator)
	// ForEachActivityPubOutgoingActivityByIdentityId loops over all 'Activity Pub Outgoing Activity' instances in the system that are children
	// of the 'Identity' instance with the provided unique identifier,
	// running the provided iterator for each of them.
	ForEachActivityPubOutgoingActivityByIdentityId(identityId int64, iterator ActivityPubOutgoingActivityIterator)

	// ActivityPubFollower
	// ------------------------------------------------------------

	// ActivityPubFollowerExists checks whether a specific 'Activity Pub Follower' with the provided
	// unique identifier or 'Id' exists in the system.
	ActivityPubFollowerExists(id int64) bool
	// ActivityPubFollowerExistsWhich checks whether a specific 'Activity Pub Follower' exists in the system
	// which satisfies the provided condition.
	ActivityPubFollowerExistsWhich(condition ActivityPubFollowerCondition) bool
	// ListActivityPubFollowers returns a list of all 'Activity Pub Follower' instances in the system.
	ListActivityPubFollowers() IActivityPubFollowerCollection
	// ForEachActivityPubFollower loops over all 'Activity Pub Follower' instances in the system running
	// the provided iterator for each of them.
	ForEachActivityPubFollower(iterator ActivityPubFollowerIterator)
	// FilterActivityPubFollowers returns a filtered list of 'Activity Pub Follower' instances based
	// on the provided predicate.
	FilterActivityPubFollowers(predicate ActivityPubFollowerFilterPredicate) IActivityPubFollowerCollection
	// MapActivityPubFollowers loops over all 'Activity Pub Follower' instances in the system and
	// returns a transformed list based on the provided predicate.
	MapActivityPubFollowers(predicate ActivityPubFollowerMapPredicate) IActivityPubFollowerCollection
	// GetActivityPubFollower finds a specific 'Activity Pub Follower' instance using
	// the provided unique identifier or 'Id'.
	GetActivityPubFollower(id int64) IActivityPubFollower
	// AddActivityPubFollower creates a new 'Activity Pub Follower' instance with an auto-generated unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool) IActivityPubFollower
	// AddActivityPubFollowerWithCustomId creates a new 'Activity Pub Follower' instance with a custom unique identifier using the
	// provided property values and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is synchronous.
	AddActivityPubFollowerWithCustomId(id int64, handle string, inbox string, subject string, activity string, accepted bool) IActivityPubFollower
	// LogActivityPubFollower creates a new 'Activity Pub Follower' instance using the provided property values
	// and adds it to persistent data store and system cache.
	// The method is smart enough to respect the transaction if used in an
	// x.Atomic context. This method is asynchronous.
	LogActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool, source string, payload string)
	// UpdateActivityPubFollower finds the 'Activity Pub Follower' instance using the provided unique identifier and updates it using
	// the provided property values and reflects the changes to persistent data store and system
	// cache. The method is smart enough to respect the transaction if used in an x.Atomic context.
	// This method is synchronous.
	UpdateActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool) IActivityPubFollower
	// UpdateActivityPubFollowerObject finds and updates the 'Activity Pub Follower' using the provided instance and reflects the
	// changes to persistent data store and system cache. The method is smart enough to
	// respect the transaction if used in an x.Atomic context. This method is synchronous.
	UpdateActivityPubFollowerObject(object IObject, activityPubFollower IActivityPubFollower) IActivityPubFollower
	// AddOrUpdateActivityPubFollowerObject tries to find the 'Activity Pub Follower' using the provided instance, then updates it in persistent
	// data store and system cache or creates it if doesn't already exist. The method is smart enough
	// to respect the transaction if used in an x.Atomic context. This method is synchronous.
	AddOrUpdateActivityPubFollowerObject(object IObject, activityPubFollower IActivityPubFollower) IActivityPubFollower
	// RemoveActivityPubFollower finds the 'Activity Pub Follower' instance using the provided unique identifier and
	// removes it from the system cache. The method is smart enough to respect
	// the transaction if used in an x.Atomic context. This method is synchronous.
	RemoveActivityPubFollower(id int64) IActivityPubFollower

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
	Signup(username string, email string, password string) (ISignupResult, error)
	Verify(email string, token string, code string) (IVerifyResult, error)
	Login(email string, password string) (ILoginResult, error)
	GetProfileByUser() (IGetProfileByUserResult, error)
	UpdateProfileByUser(displayName string, avatar string, banner string, summary string, github string) (IUpdateProfileByUserResult, error)
	Logout() (ILogoutResult, error)
	Webfinger(resource string) (IWebfingerResult, error)
	GetActor(username string) (IGetActorResult, error)
	FollowActor(username string, acct string) (IFollowActorResult, error)
	AuthorizeInteraction(uri string) (IAuthorizeInteractionResult, error)

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
	// NewActivityPubActivity creates a new 'Activity Pub Activity' instance using the provided property values.
	NewActivityPubActivity() (IActivityPubActivity, error)
	// NewActivityPubActivities creates an empty in-memory 'Activity Pub Activity' collection which is not thread-safe.
	NewActivityPubActivities() IActivityPubActivityCollection
	// NewActivityPubPublicKey creates a new 'Activity Pub Public Key' instance using the provided property values.
	NewActivityPubPublicKey() (IActivityPubPublicKey, error)
	// NewActivityPubPublicKeys creates an empty in-memory 'Activity Pub Public Key' collection which is not thread-safe.
	NewActivityPubPublicKeys() IActivityPubPublicKeyCollection
	// NewActivityPubLink creates a new 'Activity Pub Link' instance using the provided property values.
	NewActivityPubLink() (IActivityPubLink, error)
	// NewActivityPubLinks creates an empty in-memory 'Activity Pub Link' collection which is not thread-safe.
	NewActivityPubLinks() IActivityPubLinkCollection
	// NewActivityPubMedia creates a new 'Activity Pub Media' instance using the provided property values.
	NewActivityPubMedia() (IActivityPubMedia, error)
	// NewActivityPubMedias creates an empty in-memory 'Activity Pub Media' collection which is not thread-safe.
	NewActivityPubMedias() IActivityPubMediaCollection
	// NewActivityPubIncomingActivity creates a new 'Activity Pub Incoming Activity' instance using the provided property values.
	NewActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) (IActivityPubIncomingActivity, error)
	// NewActivityPubIncomingActivities creates an empty in-memory 'Activity Pub Incoming Activity' collection which is not thread-safe.
	NewActivityPubIncomingActivities() IActivityPubIncomingActivityCollection
	// NewActivityPubOutgoingActivity creates a new 'Activity Pub Outgoing Activity' instance using the provided property values.
	NewActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) (IActivityPubOutgoingActivity, error)
	// NewActivityPubOutgoingActivities creates an empty in-memory 'Activity Pub Outgoing Activity' collection which is not thread-safe.
	NewActivityPubOutgoingActivities() IActivityPubOutgoingActivityCollection
	// NewActivityPubFollower creates a new 'Activity Pub Follower' instance using the provided property values.
	NewActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool) (IActivityPubFollower, error)
	// NewActivityPubFollowers creates an empty in-memory 'Activity Pub Follower' collection which is not thread-safe.
	NewActivityPubFollowers() IActivityPubFollowerCollection
	// NewSpi creates a new 'Spi' instance using the provided property values.
	NewSpi() (ISpi, error)
	// NewSpis creates an empty in-memory 'Spi' collection which is not thread-safe.
	NewSpis() ISpiCollection
	// NewEchoResult creates a new result container for 'Echo' system action.
	NewEchoResult(document IDocument) IEchoResult
	// NewSignupResult creates a new result container for 'Signup' system action.
	NewSignupResult(token string, code string) ISignupResult
	// NewVerifyResult creates a new result container for 'Verify' system action.
	NewVerifyResult(token string) IVerifyResult
	// NewLoginResult creates a new result container for 'Login' system action.
	NewLoginResult(username string, token string) ILoginResult
	// NewGetProfileByUserResult creates a new result container for 'Get Profile By User' system action.
	NewGetProfileByUserResult(username string, displayName string, avatar string, banner string, summary string, github string) IGetProfileByUserResult
	// NewUpdateProfileByUserResult creates a new result container for 'Update Profile By User' system action.
	NewUpdateProfileByUserResult(displayName string, avatar string, banner string, summary string, github string) IUpdateProfileByUserResult
	// NewLogoutResult creates a new result container for 'Logout' system action.
	NewLogoutResult() ILogoutResult
	// NewWebfingerResult creates a new result container for 'Webfinger' system action.
	NewWebfingerResult(aliases []string, links []IActivityPubLink, subject string) IWebfingerResult
	// NewGetActorResult creates a new result container for 'Get Actor' system action.
	NewGetActorResult(context []string, id string, followers string, following string, inbox string, outbox string, name string, preferredUsername string, type_ string, url string, icon IActivityPubMedia, image IActivityPubMedia, publicKey IActivityPubPublicKey, summary string, published string) IGetActorResult
	// NewFollowActorResult creates a new result container for 'Follow Actor' system action.
	NewFollowActorResult(url string) IFollowActorResult
	// NewAuthorizeInteractionResult creates a new result container for 'Authorize Interaction' system action.
	NewAuthorizeInteractionResult(uri string, success bool) IAuthorizeInteractionResult
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
