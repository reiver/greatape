package contracts

import (
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/service"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
)

const (
	INITIALIZE = 0
	FINALIZE   = 100
)

type (
	SystemComponentType       int
	SystemAction              func() error
	SystemComponentsContainer map[string]ISystemComponent
	SystemObjectCache         map[int64]ISystemObject
	TransactionHandler        func(transaction ITransaction) error

	IConductor interface {
		Logger() ILogger
		Configuration() IConfiguration
		Atomic(handler TransactionHandler) error
		Schedule(spec string, callback func()) error
		GetSystemComponent(name string) ISystemComponent
		RequestActivityStream(method, url, keyId, privateKey string, data []byte, output interface{}) error
		LogRemoteCall(context IContext, eventType uint32, source string, input, result interface{}, err error)

		// Document
		DocumentManager() IDocumentManager
		DocumentExists(id int64) bool
		ListDocuments(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IDocumentCollection
		GetDocument(id int64, editor Identity) (IDocument, error)
		AddDocument(content string, editor Identity) (IDocument, error)
		AddDocumentAtomic(transaction ITransaction, content string, editor Identity) (IDocument, error)
		LogDocument(content string, source string, editor Identity, payload string)
		UpdateDocument(id int64, content string, editor Identity) (IDocument, error)
		UpdateDocumentAtomic(transaction ITransaction, id int64, content string, editor Identity) (IDocument, error)
		RemoveDocument(id int64, editor Identity) (IDocument, error)
		RemoveDocumentAtomic(transaction ITransaction, id int64, editor Identity) (IDocument, error)

		// SystemSchedule
		SystemScheduleManager() ISystemScheduleManager
		SystemScheduleExists(id int64) bool
		ListSystemSchedules(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ISystemScheduleCollection
		GetSystemSchedule(id int64, editor Identity) (ISystemSchedule, error)
		AddSystemSchedule(enabled bool, config string, editor Identity) (ISystemSchedule, error)
		AddSystemScheduleAtomic(transaction ITransaction, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		LogSystemSchedule(enabled bool, config string, source string, editor Identity, payload string)
		UpdateSystemSchedule(id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		UpdateSystemScheduleAtomic(transaction ITransaction, id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error)
		RemoveSystemSchedule(id int64, editor Identity) (ISystemSchedule, error)
		RemoveSystemScheduleAtomic(transaction ITransaction, id int64, editor Identity) (ISystemSchedule, error)

		// Identity
		IdentityManager() IIdentityManager
		IdentityExists(id int64) bool
		ListIdentities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IIdentityCollection
		GetIdentity(id int64, editor Identity) (IIdentity, error)
		AddIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		AddIdentityAtomic(transaction ITransaction, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		LogIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, source string, editor Identity, payload string)
		UpdateIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		UpdateIdentityAtomic(transaction ITransaction, id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		RemoveIdentity(id int64, editor Identity) (IIdentity, error)
		RemoveIdentityAtomic(transaction ITransaction, id int64, editor Identity) (IIdentity, error)

		// AccessControl
		AccessControlManager() IAccessControlManager
		AccessControlExists(id int64) bool
		ListAccessControls(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IAccessControlCollection
		GetAccessControl(id int64, editor Identity) (IAccessControl, error)
		AddAccessControl(key uint64, value uint64, editor Identity) (IAccessControl, error)
		AddAccessControlAtomic(transaction ITransaction, key uint64, value uint64, editor Identity) (IAccessControl, error)
		LogAccessControl(key uint64, value uint64, source string, editor Identity, payload string)
		UpdateAccessControl(id int64, key uint64, value uint64, editor Identity) (IAccessControl, error)
		UpdateAccessControlAtomic(transaction ITransaction, id int64, key uint64, value uint64, editor Identity) (IAccessControl, error)
		RemoveAccessControl(id int64, editor Identity) (IAccessControl, error)
		RemoveAccessControlAtomic(transaction ITransaction, id int64, editor Identity) (IAccessControl, error)

		// RemoteActivity
		RemoteActivityManager() IRemoteActivityManager
		RemoteActivityExists(id int64) bool
		ListRemoteActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IRemoteActivityCollection
		GetRemoteActivity(id int64, editor Identity) (IRemoteActivity, error)
		AddRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		AddRemoteActivityAtomic(transaction ITransaction, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		LogRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, source string, editor Identity, payload string)
		UpdateRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		UpdateRemoteActivityAtomic(transaction ITransaction, id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error)
		RemoveRemoteActivity(id int64, editor Identity) (IRemoteActivity, error)
		RemoveRemoteActivityAtomic(transaction ITransaction, id int64, editor Identity) (IRemoteActivity, error)

		// CategoryType
		CategoryTypeManager() ICategoryTypeManager
		CategoryTypeExists(id int64) bool
		ListCategoryTypes(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryTypeCollection
		GetCategoryType(id int64, editor Identity) (ICategoryType, error)
		AddCategoryType(description string, editor Identity) (ICategoryType, error)
		AddCategoryTypeAtomic(transaction ITransaction, description string, editor Identity) (ICategoryType, error)
		LogCategoryType(description string, source string, editor Identity, payload string)
		UpdateCategoryType(id int64, description string, editor Identity) (ICategoryType, error)
		UpdateCategoryTypeAtomic(transaction ITransaction, id int64, description string, editor Identity) (ICategoryType, error)
		RemoveCategoryType(id int64, editor Identity) (ICategoryType, error)
		RemoveCategoryTypeAtomic(transaction ITransaction, id int64, editor Identity) (ICategoryType, error)

		// Category
		CategoryManager() ICategoryManager
		CategoryExists(id int64) bool
		ListCategories(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection
		GetCategory(id int64, editor Identity) (ICategory, error)
		AddCategory(categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		AddCategoryAtomic(transaction ITransaction, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		LogCategory(categoryTypeId int64, categoryId int64, title string, description string, source string, editor Identity, payload string)
		UpdateCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		UpdateCategoryAtomic(transaction ITransaction, id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error)
		RemoveCategory(id int64, editor Identity) (ICategory, error)
		RemoveCategoryAtomic(transaction ITransaction, id int64, editor Identity) (ICategory, error)
		ListCategoriesByCategoryType(categoryTypeId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection
		ForEachCategoryByCategoryType(categoryTypeId int64, iterator CategoryIterator)
		ListCategoriesByCategory(categoryId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection
		ForEachCategoryByCategory(categoryId int64, iterator CategoryIterator)

		// User
		UserManager() IUserManager
		UserExists(id int64) bool
		ListUsers(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IUserCollection
		GetUser(id int64, editor Identity) (IUser, error)
		AddUser(identityId int64, github string, editor Identity) (IUser, error)
		AddUserAtomic(transaction ITransaction, identityId int64, github string, editor Identity) (IUser, error)
		LogUser(identityId int64, github string, source string, editor Identity, payload string)
		UpdateUser(id int64, github string, editor Identity) (IUser, error)
		UpdateUserAtomic(transaction ITransaction, id int64, github string, editor Identity) (IUser, error)
		RemoveUser(id int64, editor Identity) (IUser, error)
		RemoveUserAtomic(transaction ITransaction, id int64, editor Identity) (IUser, error)

		// ActivityPubObject
		ActivityPubObjectManager() IActivityPubObjectManager
		ActivityPubObjectExists(id int64) bool
		ListActivityPubObjects(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubObjectCollection
		GetActivityPubObject(id int64, editor Identity) (IActivityPubObject, error)
		AddActivityPubObject(editor Identity) (IActivityPubObject, error)
		AddActivityPubObjectAtomic(transaction ITransaction, editor Identity) (IActivityPubObject, error)
		LogActivityPubObject(source string, editor Identity, payload string)
		UpdateActivityPubObject(id int64, editor Identity) (IActivityPubObject, error)
		UpdateActivityPubObjectAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubObject, error)
		RemoveActivityPubObject(id int64, editor Identity) (IActivityPubObject, error)
		RemoveActivityPubObjectAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubObject, error)

		// ActivityPubActivity
		ActivityPubActivityManager() IActivityPubActivityManager
		ActivityPubActivityExists(id int64) bool
		ListActivityPubActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubActivityCollection
		GetActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error)
		AddActivityPubActivity(editor Identity) (IActivityPubActivity, error)
		AddActivityPubActivityAtomic(transaction ITransaction, editor Identity) (IActivityPubActivity, error)
		LogActivityPubActivity(source string, editor Identity, payload string)
		UpdateActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error)
		UpdateActivityPubActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubActivity, error)
		RemoveActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error)
		RemoveActivityPubActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubActivity, error)

		// ActivityPubPublicKey
		ActivityPubPublicKeyManager() IActivityPubPublicKeyManager
		ActivityPubPublicKeyExists(id int64) bool
		ListActivityPubPublicKeys(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubPublicKeyCollection
		GetActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error)
		AddActivityPubPublicKey(editor Identity) (IActivityPubPublicKey, error)
		AddActivityPubPublicKeyAtomic(transaction ITransaction, editor Identity) (IActivityPubPublicKey, error)
		LogActivityPubPublicKey(source string, editor Identity, payload string)
		UpdateActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error)
		UpdateActivityPubPublicKeyAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubPublicKey, error)
		RemoveActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error)
		RemoveActivityPubPublicKeyAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubPublicKey, error)

		// ActivityPubLink
		ActivityPubLinkManager() IActivityPubLinkManager
		ActivityPubLinkExists(id int64) bool
		ListActivityPubLinks(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubLinkCollection
		GetActivityPubLink(id int64, editor Identity) (IActivityPubLink, error)
		AddActivityPubLink(editor Identity) (IActivityPubLink, error)
		AddActivityPubLinkAtomic(transaction ITransaction, editor Identity) (IActivityPubLink, error)
		LogActivityPubLink(source string, editor Identity, payload string)
		UpdateActivityPubLink(id int64, editor Identity) (IActivityPubLink, error)
		UpdateActivityPubLinkAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubLink, error)
		RemoveActivityPubLink(id int64, editor Identity) (IActivityPubLink, error)
		RemoveActivityPubLinkAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubLink, error)

		// ActivityPubMedia
		ActivityPubMediaManager() IActivityPubMediaManager
		ActivityPubMediaExists(id int64) bool
		ListActivityPubMedias(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubMediaCollection
		GetActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error)
		AddActivityPubMedia(editor Identity) (IActivityPubMedia, error)
		AddActivityPubMediaAtomic(transaction ITransaction, editor Identity) (IActivityPubMedia, error)
		LogActivityPubMedia(source string, editor Identity, payload string)
		UpdateActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error)
		UpdateActivityPubMediaAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubMedia, error)
		RemoveActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error)
		RemoveActivityPubMediaAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubMedia, error)

		// ActivityPubIncomingActivity
		ActivityPubIncomingActivityManager() IActivityPubIncomingActivityManager
		ActivityPubIncomingActivityExists(id int64) bool
		ListActivityPubIncomingActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubIncomingActivityCollection
		GetActivityPubIncomingActivity(id int64, editor Identity) (IActivityPubIncomingActivity, error)
		AddActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		AddActivityPubIncomingActivityAtomic(transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		LogActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor Identity, payload string)
		UpdateActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		UpdateActivityPubIncomingActivityAtomic(transaction ITransaction, id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error)
		RemoveActivityPubIncomingActivity(id int64, editor Identity) (IActivityPubIncomingActivity, error)
		RemoveActivityPubIncomingActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubIncomingActivity, error)
		ListActivityPubIncomingActivitiesByIdentity(identityId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubIncomingActivityCollection
		ForEachActivityPubIncomingActivityByIdentity(identityId int64, iterator ActivityPubIncomingActivityIterator)

		// ActivityPubOutgoingActivity
		ActivityPubOutgoingActivityManager() IActivityPubOutgoingActivityManager
		ActivityPubOutgoingActivityExists(id int64) bool
		ListActivityPubOutgoingActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubOutgoingActivityCollection
		GetActivityPubOutgoingActivity(id int64, editor Identity) (IActivityPubOutgoingActivity, error)
		AddActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		AddActivityPubOutgoingActivityAtomic(transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		LogActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor Identity, payload string)
		UpdateActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		UpdateActivityPubOutgoingActivityAtomic(transaction ITransaction, id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error)
		RemoveActivityPubOutgoingActivity(id int64, editor Identity) (IActivityPubOutgoingActivity, error)
		RemoveActivityPubOutgoingActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubOutgoingActivity, error)
		ListActivityPubOutgoingActivitiesByIdentity(identityId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubOutgoingActivityCollection
		ForEachActivityPubOutgoingActivityByIdentity(identityId int64, iterator ActivityPubOutgoingActivityIterator)

		// ActivityPubFollower
		ActivityPubFollowerManager() IActivityPubFollowerManager
		ActivityPubFollowerExists(id int64) bool
		ListActivityPubFollowers(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubFollowerCollection
		GetActivityPubFollower(id int64, editor Identity) (IActivityPubFollower, error)
		AddActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		AddActivityPubFollowerAtomic(transaction ITransaction, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		LogActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool, source string, editor Identity, payload string)
		UpdateActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		UpdateActivityPubFollowerAtomic(transaction ITransaction, id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error)
		RemoveActivityPubFollower(id int64, editor Identity) (IActivityPubFollower, error)
		RemoveActivityPubFollowerAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubFollower, error)

		// Spi
		SpiManager() ISpiManager
		SpiExists(id int64) bool
		ListSpis(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ISpiCollection
		GetSpi(id int64, editor Identity) (ISpi, error)
		AddSpi(editor Identity) (ISpi, error)
		AddSpiAtomic(transaction ITransaction, editor Identity) (ISpi, error)
		LogSpi(source string, editor Identity, payload string)
		UpdateSpi(id int64, editor Identity) (ISpi, error)
		UpdateSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error)
		RemoveSpi(id int64, editor Identity) (ISpi, error)
		RemoveSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error)
		Echo(document IDocument, editor Identity) (IEchoResult, error)
		CheckUsernameAvailability(username string, editor Identity) (ICheckUsernameAvailabilityResult, error)
		Signup(username string, email string, password string, editor Identity) (ISignupResult, error)
		Verify(email string, token string, code string, editor Identity) (IVerifyResult, error)
		Login(email string, password string, editor Identity) (ILoginResult, error)
		GetProfileByUser(editor Identity) (IGetProfileByUserResult, error)
		UpdateProfileByUser(displayName string, avatar string, banner string, summary string, github string, editor Identity) (IUpdateProfileByUserResult, error)
		Logout(editor Identity) (ILogoutResult, error)
		Webfinger(resource string, editor Identity) (IWebfingerResult, error)
		GetPackages(editor Identity) (IGetPackagesResult, error)
		GetActor(username string, editor Identity) (IGetActorResult, error)
		FollowActor(username string, acct string, editor Identity) (IFollowActorResult, error)
		AuthorizeInteraction(uri string, editor Identity) (IAuthorizeInteractionResult, error)
		GetFollowers(username string, editor Identity) (IGetFollowersResult, error)
		GetFollowing(username string, editor Identity) (IGetFollowingResult, error)
		PostToOutbox(username string, body []byte, editor Identity) (IPostToOutboxResult, error)
		GetOutbox(username string, editor Identity) (IGetOutboxResult, error)
		PostToInbox(username string, body []byte, editor Identity) (IPostToInboxResult, error)
		GetInbox(username string, editor Identity) (IGetInboxResult, error)

		NewDocument(id int64, content string) (IDocument, error)
		NewSystemSchedule(id int64, enabled bool, config string) (ISystemSchedule, error)
		NewIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) (IIdentity, error)
		NewAccessControl(id int64, key uint64, value uint64) (IAccessControl, error)
		NewRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) (IRemoteActivity, error)
		NewCategoryType(id int64, description string) (ICategoryType, error)
		NewCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string) (ICategory, error)
		NewUser(id int64, github string) (IUser, error)
		NewActivityPubObject() (IActivityPubObject, error)
		NewActivityPubActivity() (IActivityPubActivity, error)
		NewActivityPubPublicKey() (IActivityPubPublicKey, error)
		NewActivityPubLink() (IActivityPubLink, error)
		NewActivityPubMedia() (IActivityPubMedia, error)
		NewActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) (IActivityPubIncomingActivity, error)
		NewActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) (IActivityPubOutgoingActivity, error)
		NewActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool) (IActivityPubFollower, error)
		NewSpi() (ISpi, error)
		NewEchoResult(document IDocument, ignored interface{}) IEchoResult
		NewCheckUsernameAvailabilityResult(isAvailable bool, ignored interface{}) ICheckUsernameAvailabilityResult
		NewSignupResult(token string, code string, ignored interface{}) ISignupResult
		NewVerifyResult(token string, ignored interface{}) IVerifyResult
		NewLoginResult(username string, token string, ignored interface{}) ILoginResult
		NewGetProfileByUserResult(username string, displayName string, avatar string, banner string, summary string, github string, ignored interface{}) IGetProfileByUserResult
		NewUpdateProfileByUserResult(displayName string, avatar string, banner string, summary string, github string, ignored interface{}) IUpdateProfileByUserResult
		NewLogoutResult(ignored interface{}) ILogoutResult
		NewWebfingerResult(aliases []string, links []IActivityPubLink, subject string, ignored interface{}) IWebfingerResult
		NewGetPackagesResult(body []byte, ignored interface{}) IGetPackagesResult
		NewGetActorResult(context []string, id string, followers string, following string, inbox string, outbox string, name string, preferredUsername string, type_ string, url string, icon IActivityPubMedia, image IActivityPubMedia, publicKey IActivityPubPublicKey, summary string, published string, ignored interface{}) IGetActorResult
		NewFollowActorResult(url string, ignored interface{}) IFollowActorResult
		NewAuthorizeInteractionResult(uri string, success bool, ignored interface{}) IAuthorizeInteractionResult
		NewGetFollowersResult(context string, id string, type_ string, totalItems int32, orderedItems []string, first string, ignored interface{}) IGetFollowersResult
		NewGetFollowingResult(context string, id string, type_ string, totalItems int32, orderedItems []string, first string, ignored interface{}) IGetFollowingResult
		NewPostToOutboxResult(body []byte, ignored interface{}) IPostToOutboxResult
		NewGetOutboxResult(context string, id string, type_ string, totalItems int32, orderedItems []IActivityPubActivity, first string, ignored interface{}) IGetOutboxResult
		NewPostToInboxResult(body []byte, ignored interface{}) IPostToInboxResult
		NewGetInboxResult(context string, id string, type_ string, totalItems int32, orderedItems []IActivityPubActivity, first string, ignored interface{}) IGetInboxResult
	}

	ISystemComponent interface {
		Name() string
		ResolveDependencies(dependencies ...ISystemComponent) error
		Load() error
		Reload() error
		IsTestEnvironment() bool
		IsDevelopmentEnvironment() bool
		IsStagingEnvironment() bool
		IsProductionEnvironment() bool
		UniqueId() int64
		Logger() ILogger
		Async(task func())
		GenerateUUID() string
		GenerateSalt() string
		GenerateHash(value string, salt string) string
		GenerateJwtToken() string
		GenerateRSAKeyPair() (string, string, error)
		VerifyJwtToken(token string) error
		GenerateCode() string
		Email(destination string, format string, args ...interface{})
		SMS(destination string, format string, args ...interface{})
		Format(format string, args ...interface{}) string
		Match(pattern string, input string) (bool, error)
		Error(interface{}) error
	}

	ISystemComponentFactory interface {
		Create(SystemComponentType, IConfiguration, ILogger, ...ISystemComponent) ISystemComponent
		Components() []ISystemComponent
	}

	IAssertionResult interface {
		Or(error)
	}

	ITransaction interface {
		OnCommit(func())
	}
)

// noinspection GoSnakeCaseUsage
const (
	SYSTEM_COMPONENT_DOCUMENT_MANAGER                       SystemComponentType = 0x00000001
	SYSTEM_COMPONENT_SYSTEM_SCHEDULE_MANAGER                SystemComponentType = 0x00000002
	SYSTEM_COMPONENT_IDENTITY_MANAGER                       SystemComponentType = 0x00000003
	SYSTEM_COMPONENT_ACCESS_CONTROL_MANAGER                 SystemComponentType = 0x00000004
	SYSTEM_COMPONENT_REMOTE_ACTIVITY_MANAGER                SystemComponentType = 0x00000005
	SYSTEM_COMPONENT_CATEGORY_TYPE_MANAGER                  SystemComponentType = 0x00000006
	SYSTEM_COMPONENT_CATEGORY_MANAGER                       SystemComponentType = 0x00000007
	SYSTEM_COMPONENT_USER_MANAGER                           SystemComponentType = 0x00000008
	SYSTEM_COMPONENT_ACTIVITY_PUB_OBJECT_MANAGER            SystemComponentType = 0x00000009
	SYSTEM_COMPONENT_ACTIVITY_PUB_ACTIVITY_MANAGER          SystemComponentType = 0x0000000A
	SYSTEM_COMPONENT_ACTIVITY_PUB_PUBLIC_KEY_MANAGER        SystemComponentType = 0x0000000B
	SYSTEM_COMPONENT_ACTIVITY_PUB_LINK_MANAGER              SystemComponentType = 0x0000000C
	SYSTEM_COMPONENT_ACTIVITY_PUB_MEDIA_MANAGER             SystemComponentType = 0x0000000D
	SYSTEM_COMPONENT_ACTIVITY_PUB_INCOMING_ACTIVITY_MANAGER SystemComponentType = 0x0000000E
	SYSTEM_COMPONENT_ACTIVITY_PUB_OUTGOING_ACTIVITY_MANAGER SystemComponentType = 0x0000000F
	SYSTEM_COMPONENT_ACTIVITY_PUB_FOLLOWER_MANAGER          SystemComponentType = 0x00000010
	SYSTEM_COMPONENT_SPI_MANAGER                            SystemComponentType = 0x00000011
)
