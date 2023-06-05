package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	app "github.com/reiver/greatape/app"
	. "github.com/reiver/greatape/components/contracts"
	"github.com/reiver/greatape/components/contracts/model"
	"github.com/reiver/greatape/components/model/repository"
	schedule "github.com/robfig/cron"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/service"
	. "github.com/xeronith/diamante/contracts/settings"
	"github.com/xeronith/diamante/utility/httpsig"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	// @formatter:off
	Conductor           IConductor
	componentsContainer SystemComponentsContainer
	buildNumber         = "0"
	runningInContainer  string
	Dockerized          = runningInContainer == "true"
	// @formatter:on
)

func Initialize(configuration IConfiguration, logger ILogger) error {
	logger.SysComp("┄ Booting up ...")

	environment := configuration.GetEnvironment()

	if Dockerized {
		environment += " (Docker)"
	} else {
		environment += " (Metal)"
	}

	logger.SysComp(fmt.Sprintf("┄ Environment: %s", environment))
	logger.SysComp("┄ Initializing system components ...")

	factory := newSystemComponentFactory()

	// Initializing System Components
	documentManager := factory.Create(SYSTEM_COMPONENT_DOCUMENT_MANAGER, configuration, logger).(IDocumentManager)
	systemScheduleManager := factory.Create(SYSTEM_COMPONENT_SYSTEM_SCHEDULE_MANAGER, configuration, logger).(ISystemScheduleManager)
	identityManager := factory.Create(SYSTEM_COMPONENT_IDENTITY_MANAGER, configuration, logger).(IIdentityManager)
	accessControlManager := factory.Create(SYSTEM_COMPONENT_ACCESS_CONTROL_MANAGER, configuration, logger).(IAccessControlManager)
	remoteActivityManager := factory.Create(SYSTEM_COMPONENT_REMOTE_ACTIVITY_MANAGER, configuration, logger).(IRemoteActivityManager)
	categoryTypeManager := factory.Create(SYSTEM_COMPONENT_CATEGORY_TYPE_MANAGER, configuration, logger).(ICategoryTypeManager)
	categoryManager := factory.Create(SYSTEM_COMPONENT_CATEGORY_MANAGER, configuration, logger).(ICategoryManager)
	userManager := factory.Create(SYSTEM_COMPONENT_USER_MANAGER, configuration, logger).(IUserManager)
	activityPubObjectManager := factory.Create(SYSTEM_COMPONENT_ACTIVITY_PUB_OBJECT_MANAGER, configuration, logger).(IActivityPubObjectManager)
	activityPubActivityManager := factory.Create(SYSTEM_COMPONENT_ACTIVITY_PUB_ACTIVITY_MANAGER, configuration, logger).(IActivityPubActivityManager)
	activityPubPublicKeyManager := factory.Create(SYSTEM_COMPONENT_ACTIVITY_PUB_PUBLIC_KEY_MANAGER, configuration, logger).(IActivityPubPublicKeyManager)
	activityPubLinkManager := factory.Create(SYSTEM_COMPONENT_ACTIVITY_PUB_LINK_MANAGER, configuration, logger).(IActivityPubLinkManager)
	activityPubMediaManager := factory.Create(SYSTEM_COMPONENT_ACTIVITY_PUB_MEDIA_MANAGER, configuration, logger).(IActivityPubMediaManager)
	activityPubIncomingActivityManager := factory.Create(SYSTEM_COMPONENT_ACTIVITY_PUB_INCOMING_ACTIVITY_MANAGER, configuration, logger).(IActivityPubIncomingActivityManager)
	activityPubOutgoingActivityManager := factory.Create(SYSTEM_COMPONENT_ACTIVITY_PUB_OUTGOING_ACTIVITY_MANAGER, configuration, logger).(IActivityPubOutgoingActivityManager)
	activityPubFollowerManager := factory.Create(SYSTEM_COMPONENT_ACTIVITY_PUB_FOLLOWER_MANAGER, configuration, logger).(IActivityPubFollowerManager)
	spiManager := factory.Create(SYSTEM_COMPONENT_SPI_MANAGER, configuration, logger).(ISpiManager)

	// Resolving Dependencies
	// @formatter:off
	if err := categoryManager.ResolveDependencies(nil, categoryTypeManager, categoryManager); err != nil {
		return err
	}
	if err := activityPubIncomingActivityManager.ResolveDependencies(nil, identityManager); err != nil {
		return err
	}
	if err := activityPubOutgoingActivityManager.ResolveDependencies(nil, identityManager); err != nil {
		return err
	}
	// @formatter:on

	identityManager.SetAccessControlHandler(accessControlManager)
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	scheduler := schedule.NewWithLocation(location)
	scheduler.Start()

	// Aggregating System Components
	Conductor = &conductor{
		// @formatter:off
		documentManager:                    documentManager,
		systemScheduleManager:              systemScheduleManager,
		identityManager:                    identityManager,
		accessControlManager:               accessControlManager,
		remoteActivityManager:              remoteActivityManager,
		categoryTypeManager:                categoryTypeManager,
		categoryManager:                    categoryManager,
		userManager:                        userManager,
		activityPubObjectManager:           activityPubObjectManager,
		activityPubActivityManager:         activityPubActivityManager,
		activityPubPublicKeyManager:        activityPubPublicKeyManager,
		activityPubLinkManager:             activityPubLinkManager,
		activityPubMediaManager:            activityPubMediaManager,
		activityPubIncomingActivityManager: activityPubIncomingActivityManager,
		activityPubOutgoingActivityManager: activityPubOutgoingActivityManager,
		activityPubFollowerManager:         activityPubFollowerManager,
		spiManager:                         spiManager,
		logger:                             logger,
		configuration:                      configuration,
		scheduler:                          scheduler,
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
		// @formatter:on
	}

	logger.SysComp("┄ Loading system components ...")

	var totalDuration float64 = 0
	componentsContainer = make(SystemComponentsContainer)
	for _, component := range factory.Components() {
		start := time.Now()
		componentName := component.Name()
		if _, exists := componentsContainer[componentName]; exists {
			return fmt.Errorf("%s already registered", componentName)
		}

		if err := component.Load(); err != nil {
			return err
		}

		componentsContainer[componentName] = component
		duration := time.Since(start).Seconds()
		totalDuration += duration
		logger.SysComp(fmt.Sprintf("✓ %s: %.2fs", componentName, duration))
	}

	serverBuildNumber, err := strconv.ParseInt(buildNumber, 10, 32)
	if err != nil {
		return err
	}

	configuration.GetServerConfiguration().SetBuildNumber(int32(serverBuildNumber))
	logger.SysComp(fmt.Sprintf("┄ All system components loaded in %.2fs", totalDuration))
	logger.SysComp(fmt.Sprintf("┄ Runtime: %s/%s %s build %s", runtime.GOOS, runtime.GOARCH, runtime.Version(), buildNumber))
	if err := app.Initialize(NewDispatcher(Conductor, NewSystemIdentity())); err != nil {
		return err
	}

	logger.SysComp("┄ System operational")

	return nil
}

//region IConductor Implementation

type conductor struct {
	// @formatter:off
	documentManager                    IDocumentManager
	systemScheduleManager              ISystemScheduleManager
	identityManager                    IIdentityManager
	accessControlManager               IAccessControlManager
	remoteActivityManager              IRemoteActivityManager
	categoryTypeManager                ICategoryTypeManager
	categoryManager                    ICategoryManager
	userManager                        IUserManager
	activityPubObjectManager           IActivityPubObjectManager
	activityPubActivityManager         IActivityPubActivityManager
	activityPubPublicKeyManager        IActivityPubPublicKeyManager
	activityPubLinkManager             IActivityPubLinkManager
	activityPubMediaManager            IActivityPubMediaManager
	activityPubIncomingActivityManager IActivityPubIncomingActivityManager
	activityPubOutgoingActivityManager IActivityPubOutgoingActivityManager
	activityPubFollowerManager         IActivityPubFollowerManager
	spiManager                         ISpiManager
	logger                             ILogger
	configuration                      IConfiguration
	scheduler                          *schedule.Cron
	httpClient                         *http.Client
	// @formatter:on
}

func (conductor *conductor) Logger() ILogger {
	return conductor.logger
}

func (conductor *conductor) Configuration() IConfiguration {
	return conductor.configuration
}

func (conductor *conductor) Atomic(handler TransactionHandler) error {
	return repository.WithTransaction(func(transaction model.IRepositoryTransaction) error {
		return handler(transaction)
	})
}

func (conductor *conductor) Schedule(spec string, callback func()) error {
	return conductor.scheduler.AddFunc(spec, callback)
}

func (conductor *conductor) GetSystemComponent(name string) ISystemComponent {
	if component, exists := componentsContainer[name]; exists {
		return component
	} else {
		return nil
	}
}

func (conductor *conductor) SignRequest(keyId, privateKey string, data []byte, req *http.Request) error {
	privKey, err := httpsig.ParseRsaPrivateKeyFromPemStr(privateKey)
	if err != nil {
		return err
	}

	signer := httpsig.NewRSASHA256Signer(keyId, privKey, []string{"Date", "Digest"})
	if data != nil {
		hasher := sha256.New()
		hasher.Write(data)
		sum := hasher.Sum(nil)
		encodedHash := base64.StdEncoding.EncodeToString(sum)
		digest := fmt.Sprintf("sha-256=%s", encodedHash)
		req.Header.Set("Content-Type", "application/activity+json; charset=utf-8")
		req.Header.Set("Digest", digest)
	}

	if err := signer.Sign(req); err != nil {
		return err
	}

	return nil
}

func (conductor *conductor) RequestActivityStream(method, url, keyId, privateKey string, data []byte, output interface{}) error {
	var reader io.Reader
	if data != nil {
		reader = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/activity+json")

	if privateKey != "" {
		if err := conductor.SignRequest(keyId, privateKey, data, req); err != nil {
			return err
		}
	}

	res, err := conductor.httpClient.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	if res.StatusCode != http.StatusOK &&
		res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("%s", res.Status)
	}

	if output != nil {
		if err := json.NewDecoder(res.Body).Decode(output); err != nil {
			return err
		}
	}

	return nil
}

// Document

func (conductor *conductor) DocumentManager() IDocumentManager {
	return conductor.documentManager
}

func (conductor *conductor) DocumentExists(id int64) bool {
	return conductor.documentManager.Exists(id)
}

func (conductor *conductor) ListDocuments(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IDocumentCollection {
	return conductor.documentManager.ListDocuments(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetDocument(id int64, editor Identity) (IDocument, error) {
	return conductor.documentManager.GetDocument(id, editor)
}

func (conductor *conductor) AddDocument(content string, editor Identity) (IDocument, error) {
	return conductor.documentManager.AddDocument(content, editor)
}

func (conductor *conductor) AddDocumentAtomic(transaction ITransaction, content string, editor Identity) (IDocument, error) {
	return conductor.documentManager.AddDocumentAtomic(transaction, content, editor)
}

func (conductor *conductor) LogDocument(content string, source string, editor Identity, payload string) {
	conductor.documentManager.Log(content, source, editor, payload)
}

func (conductor *conductor) UpdateDocument(id int64, content string, editor Identity) (IDocument, error) {
	return conductor.documentManager.UpdateDocument(id, content, editor)
}

func (conductor *conductor) UpdateDocumentAtomic(transaction ITransaction, id int64, content string, editor Identity) (IDocument, error) {
	return conductor.documentManager.UpdateDocumentAtomic(transaction, id, content, editor)
}

func (conductor *conductor) RemoveDocument(id int64, editor Identity) (IDocument, error) {
	return conductor.documentManager.RemoveDocument(id, editor)
}

func (conductor *conductor) RemoveDocumentAtomic(transaction ITransaction, id int64, editor Identity) (IDocument, error) {
	return conductor.documentManager.RemoveDocumentAtomic(transaction, id, editor)
}

// SystemSchedule

func (conductor *conductor) SystemScheduleManager() ISystemScheduleManager {
	return conductor.systemScheduleManager
}

func (conductor *conductor) SystemScheduleExists(id int64) bool {
	return conductor.systemScheduleManager.Exists(id)
}

func (conductor *conductor) ListSystemSchedules(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ISystemScheduleCollection {
	return conductor.systemScheduleManager.ListSystemSchedules(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetSystemSchedule(id int64, editor Identity) (ISystemSchedule, error) {
	return conductor.systemScheduleManager.GetSystemSchedule(id, editor)
}

func (conductor *conductor) AddSystemSchedule(enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	return conductor.systemScheduleManager.AddSystemSchedule(enabled, config, editor)
}

func (conductor *conductor) AddSystemScheduleAtomic(transaction ITransaction, enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	return conductor.systemScheduleManager.AddSystemScheduleAtomic(transaction, enabled, config, editor)
}

func (conductor *conductor) LogSystemSchedule(enabled bool, config string, source string, editor Identity, payload string) {
	conductor.systemScheduleManager.Log(enabled, config, source, editor, payload)
}

func (conductor *conductor) UpdateSystemSchedule(id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	return conductor.systemScheduleManager.UpdateSystemSchedule(id, enabled, config, editor)
}

func (conductor *conductor) UpdateSystemScheduleAtomic(transaction ITransaction, id int64, enabled bool, config string, editor Identity) (ISystemSchedule, error) {
	return conductor.systemScheduleManager.UpdateSystemScheduleAtomic(transaction, id, enabled, config, editor)
}

func (conductor *conductor) RemoveSystemSchedule(id int64, editor Identity) (ISystemSchedule, error) {
	return conductor.systemScheduleManager.RemoveSystemSchedule(id, editor)
}

func (conductor *conductor) RemoveSystemScheduleAtomic(transaction ITransaction, id int64, editor Identity) (ISystemSchedule, error) {
	return conductor.systemScheduleManager.RemoveSystemScheduleAtomic(transaction, id, editor)
}

// Identity

func (conductor *conductor) IdentityManager() IIdentityManager {
	return conductor.identityManager
}

func (conductor *conductor) IdentityExists(id int64) bool {
	return conductor.identityManager.Exists(id)
}

func (conductor *conductor) ListIdentities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IIdentityCollection {
	return conductor.identityManager.ListIdentities(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetIdentity(id int64, editor Identity) (IIdentity, error) {
	return conductor.identityManager.GetIdentity(id, editor)
}

func (conductor *conductor) AddIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	return conductor.identityManager.AddIdentity(username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, editor)
}

func (conductor *conductor) AddIdentityAtomic(transaction ITransaction, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	return conductor.identityManager.AddIdentityAtomic(transaction, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, editor)
}

func (conductor *conductor) LogIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, source string, editor Identity, payload string) {
	conductor.identityManager.Log(username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, source, editor, payload)
}

func (conductor *conductor) UpdateIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	return conductor.identityManager.UpdateIdentity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, editor)
}

func (conductor *conductor) UpdateIdentityAtomic(transaction ITransaction, id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	return conductor.identityManager.UpdateIdentityAtomic(transaction, id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, editor)
}

func (conductor *conductor) RemoveIdentity(id int64, editor Identity) (IIdentity, error) {
	return conductor.identityManager.RemoveIdentity(id, editor)
}

func (conductor *conductor) RemoveIdentityAtomic(transaction ITransaction, id int64, editor Identity) (IIdentity, error) {
	return conductor.identityManager.RemoveIdentityAtomic(transaction, id, editor)
}

// AccessControl

func (conductor *conductor) AccessControlManager() IAccessControlManager {
	return conductor.accessControlManager
}

func (conductor *conductor) AccessControlExists(id int64) bool {
	return conductor.accessControlManager.Exists(id)
}

func (conductor *conductor) ListAccessControls(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IAccessControlCollection {
	return conductor.accessControlManager.ListAccessControls(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetAccessControl(id int64, editor Identity) (IAccessControl, error) {
	return conductor.accessControlManager.GetAccessControl(id, editor)
}

func (conductor *conductor) AddAccessControl(key uint64, value uint64, editor Identity) (IAccessControl, error) {
	return conductor.accessControlManager.AddAccessControl(key, value, editor)
}

func (conductor *conductor) AddAccessControlAtomic(transaction ITransaction, key uint64, value uint64, editor Identity) (IAccessControl, error) {
	return conductor.accessControlManager.AddAccessControlAtomic(transaction, key, value, editor)
}

func (conductor *conductor) LogAccessControl(key uint64, value uint64, source string, editor Identity, payload string) {
	conductor.accessControlManager.Log(key, value, source, editor, payload)
}

func (conductor *conductor) UpdateAccessControl(id int64, key uint64, value uint64, editor Identity) (IAccessControl, error) {
	return conductor.accessControlManager.UpdateAccessControl(id, key, value, editor)
}

func (conductor *conductor) UpdateAccessControlAtomic(transaction ITransaction, id int64, key uint64, value uint64, editor Identity) (IAccessControl, error) {
	return conductor.accessControlManager.UpdateAccessControlAtomic(transaction, id, key, value, editor)
}

func (conductor *conductor) RemoveAccessControl(id int64, editor Identity) (IAccessControl, error) {
	return conductor.accessControlManager.RemoveAccessControl(id, editor)
}

func (conductor *conductor) RemoveAccessControlAtomic(transaction ITransaction, id int64, editor Identity) (IAccessControl, error) {
	return conductor.accessControlManager.RemoveAccessControlAtomic(transaction, id, editor)
}

// RemoteActivity

func (conductor *conductor) RemoteActivityManager() IRemoteActivityManager {
	return conductor.remoteActivityManager
}

func (conductor *conductor) RemoteActivityExists(id int64) bool {
	return conductor.remoteActivityManager.Exists(id)
}

func (conductor *conductor) ListRemoteActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IRemoteActivityCollection {
	return conductor.remoteActivityManager.ListRemoteActivities(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetRemoteActivity(id int64, editor Identity) (IRemoteActivity, error) {
	return conductor.remoteActivityManager.GetRemoteActivity(id, editor)
}

func (conductor *conductor) AddRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	return conductor.remoteActivityManager.AddRemoteActivity(entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, editor)
}

func (conductor *conductor) AddRemoteActivityAtomic(transaction ITransaction, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	return conductor.remoteActivityManager.AddRemoteActivityAtomic(transaction, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, editor)
}

func (conductor *conductor) LogRemoteActivity(entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, source string, editor Identity, payload string) {
	conductor.remoteActivityManager.Log(entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, source, editor, payload)
}

func (conductor *conductor) UpdateRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	return conductor.remoteActivityManager.UpdateRemoteActivity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, editor)
}

func (conductor *conductor) UpdateRemoteActivityAtomic(transaction ITransaction, id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64, editor Identity) (IRemoteActivity, error) {
	return conductor.remoteActivityManager.UpdateRemoteActivityAtomic(transaction, id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp, editor)
}

func (conductor *conductor) RemoveRemoteActivity(id int64, editor Identity) (IRemoteActivity, error) {
	return conductor.remoteActivityManager.RemoveRemoteActivity(id, editor)
}

func (conductor *conductor) RemoveRemoteActivityAtomic(transaction ITransaction, id int64, editor Identity) (IRemoteActivity, error) {
	return conductor.remoteActivityManager.RemoveRemoteActivityAtomic(transaction, id, editor)
}

// CategoryType

func (conductor *conductor) CategoryTypeManager() ICategoryTypeManager {
	return conductor.categoryTypeManager
}

func (conductor *conductor) CategoryTypeExists(id int64) bool {
	return conductor.categoryTypeManager.Exists(id)
}

func (conductor *conductor) ListCategoryTypes(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryTypeCollection {
	return conductor.categoryTypeManager.ListCategoryTypes(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetCategoryType(id int64, editor Identity) (ICategoryType, error) {
	return conductor.categoryTypeManager.GetCategoryType(id, editor)
}

func (conductor *conductor) AddCategoryType(description string, editor Identity) (ICategoryType, error) {
	return conductor.categoryTypeManager.AddCategoryType(description, editor)
}

func (conductor *conductor) AddCategoryTypeAtomic(transaction ITransaction, description string, editor Identity) (ICategoryType, error) {
	return conductor.categoryTypeManager.AddCategoryTypeAtomic(transaction, description, editor)
}

func (conductor *conductor) LogCategoryType(description string, source string, editor Identity, payload string) {
	conductor.categoryTypeManager.Log(description, source, editor, payload)
}

func (conductor *conductor) UpdateCategoryType(id int64, description string, editor Identity) (ICategoryType, error) {
	return conductor.categoryTypeManager.UpdateCategoryType(id, description, editor)
}

func (conductor *conductor) UpdateCategoryTypeAtomic(transaction ITransaction, id int64, description string, editor Identity) (ICategoryType, error) {
	return conductor.categoryTypeManager.UpdateCategoryTypeAtomic(transaction, id, description, editor)
}

func (conductor *conductor) RemoveCategoryType(id int64, editor Identity) (ICategoryType, error) {
	return conductor.categoryTypeManager.RemoveCategoryType(id, editor)
}

func (conductor *conductor) RemoveCategoryTypeAtomic(transaction ITransaction, id int64, editor Identity) (ICategoryType, error) {
	return conductor.categoryTypeManager.RemoveCategoryTypeAtomic(transaction, id, editor)
}

// Category

func (conductor *conductor) CategoryManager() ICategoryManager {
	return conductor.categoryManager
}

func (conductor *conductor) CategoryExists(id int64) bool {
	return conductor.categoryManager.Exists(id)
}

func (conductor *conductor) ListCategories(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection {
	return conductor.categoryManager.ListCategories(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetCategory(id int64, editor Identity) (ICategory, error) {
	return conductor.categoryManager.GetCategory(id, editor)
}

func (conductor *conductor) AddCategory(categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	return conductor.categoryManager.AddCategory(categoryTypeId, categoryId, title, description, editor)
}

func (conductor *conductor) AddCategoryAtomic(transaction ITransaction, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	return conductor.categoryManager.AddCategoryAtomic(transaction, categoryTypeId, categoryId, title, description, editor)
}

func (conductor *conductor) LogCategory(categoryTypeId int64, categoryId int64, title string, description string, source string, editor Identity, payload string) {
	conductor.categoryManager.Log(categoryTypeId, categoryId, title, description, source, editor, payload)
}

func (conductor *conductor) UpdateCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	return conductor.categoryManager.UpdateCategory(id, categoryTypeId, categoryId, title, description, editor)
}

func (conductor *conductor) UpdateCategoryAtomic(transaction ITransaction, id int64, categoryTypeId int64, categoryId int64, title string, description string, editor Identity) (ICategory, error) {
	return conductor.categoryManager.UpdateCategoryAtomic(transaction, id, categoryTypeId, categoryId, title, description, editor)
}

func (conductor *conductor) RemoveCategory(id int64, editor Identity) (ICategory, error) {
	return conductor.categoryManager.RemoveCategory(id, editor)
}

func (conductor *conductor) RemoveCategoryAtomic(transaction ITransaction, id int64, editor Identity) (ICategory, error) {
	return conductor.categoryManager.RemoveCategoryAtomic(transaction, id, editor)
}

func (conductor *conductor) ListCategoriesByCategoryType(categoryTypeId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection {
	return conductor.categoryManager.ListCategoriesByCategoryType(categoryTypeId, pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) ForEachCategoryByCategoryType(categoryTypeId int64, iterator CategoryIterator) {
	conductor.categoryManager.ForEachByCategoryType(categoryTypeId, iterator)
}

func (conductor *conductor) ListCategoriesByCategory(categoryId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) ICategoryCollection {
	return conductor.categoryManager.ListCategoriesByCategory(categoryId, pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) ForEachCategoryByCategory(categoryId int64, iterator CategoryIterator) {
	conductor.categoryManager.ForEachByCategory(categoryId, iterator)
}

// User

func (conductor *conductor) UserManager() IUserManager {
	return conductor.userManager
}

func (conductor *conductor) UserExists(id int64) bool {
	return conductor.userManager.Exists(id)
}

func (conductor *conductor) ListUsers(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IUserCollection {
	return conductor.userManager.ListUsers(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetUser(id int64, editor Identity) (IUser, error) {
	return conductor.userManager.GetUser(id, editor)
}

func (conductor *conductor) AddUser(identityId int64, github string, editor Identity) (IUser, error) {
	return conductor.userManager.AddUser(identityId, github, editor)
}

func (conductor *conductor) AddUserAtomic(transaction ITransaction, identityId int64, github string, editor Identity) (IUser, error) {
	return conductor.userManager.AddUserAtomic(transaction, identityId, github, editor)
}

func (conductor *conductor) LogUser(identityId int64, github string, source string, editor Identity, payload string) {
	conductor.userManager.Log(identityId, github, source, editor, payload)
}

func (conductor *conductor) UpdateUser(id int64, github string, editor Identity) (IUser, error) {
	return conductor.userManager.UpdateUser(id, github, editor)
}

func (conductor *conductor) UpdateUserAtomic(transaction ITransaction, id int64, github string, editor Identity) (IUser, error) {
	return conductor.userManager.UpdateUserAtomic(transaction, id, github, editor)
}

func (conductor *conductor) RemoveUser(id int64, editor Identity) (IUser, error) {
	return conductor.userManager.RemoveUser(id, editor)
}

func (conductor *conductor) RemoveUserAtomic(transaction ITransaction, id int64, editor Identity) (IUser, error) {
	return conductor.userManager.RemoveUserAtomic(transaction, id, editor)
}

// ActivityPubObject

func (conductor *conductor) ActivityPubObjectManager() IActivityPubObjectManager {
	return conductor.activityPubObjectManager
}

func (conductor *conductor) ActivityPubObjectExists(id int64) bool {
	return conductor.activityPubObjectManager.Exists(id)
}

func (conductor *conductor) ListActivityPubObjects(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubObjectCollection {
	return conductor.activityPubObjectManager.ListActivityPubObjects(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetActivityPubObject(id int64, editor Identity) (IActivityPubObject, error) {
	return conductor.activityPubObjectManager.GetActivityPubObject(id, editor)
}

func (conductor *conductor) AddActivityPubObject(editor Identity) (IActivityPubObject, error) {
	return conductor.activityPubObjectManager.AddActivityPubObject(editor)
}

func (conductor *conductor) AddActivityPubObjectAtomic(transaction ITransaction, editor Identity) (IActivityPubObject, error) {
	return conductor.activityPubObjectManager.AddActivityPubObjectAtomic(transaction, editor)
}

func (conductor *conductor) LogActivityPubObject(source string, editor Identity, payload string) {
	conductor.activityPubObjectManager.Log(source, editor, payload)
}

func (conductor *conductor) UpdateActivityPubObject(id int64, editor Identity) (IActivityPubObject, error) {
	return conductor.activityPubObjectManager.UpdateActivityPubObject(id, editor)
}

func (conductor *conductor) UpdateActivityPubObjectAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubObject, error) {
	return conductor.activityPubObjectManager.UpdateActivityPubObjectAtomic(transaction, id, editor)
}

func (conductor *conductor) RemoveActivityPubObject(id int64, editor Identity) (IActivityPubObject, error) {
	return conductor.activityPubObjectManager.RemoveActivityPubObject(id, editor)
}

func (conductor *conductor) RemoveActivityPubObjectAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubObject, error) {
	return conductor.activityPubObjectManager.RemoveActivityPubObjectAtomic(transaction, id, editor)
}

// ActivityPubActivity

func (conductor *conductor) ActivityPubActivityManager() IActivityPubActivityManager {
	return conductor.activityPubActivityManager
}

func (conductor *conductor) ActivityPubActivityExists(id int64) bool {
	return conductor.activityPubActivityManager.Exists(id)
}

func (conductor *conductor) ListActivityPubActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubActivityCollection {
	return conductor.activityPubActivityManager.ListActivityPubActivities(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error) {
	return conductor.activityPubActivityManager.GetActivityPubActivity(id, editor)
}

func (conductor *conductor) AddActivityPubActivity(editor Identity) (IActivityPubActivity, error) {
	return conductor.activityPubActivityManager.AddActivityPubActivity(editor)
}

func (conductor *conductor) AddActivityPubActivityAtomic(transaction ITransaction, editor Identity) (IActivityPubActivity, error) {
	return conductor.activityPubActivityManager.AddActivityPubActivityAtomic(transaction, editor)
}

func (conductor *conductor) LogActivityPubActivity(source string, editor Identity, payload string) {
	conductor.activityPubActivityManager.Log(source, editor, payload)
}

func (conductor *conductor) UpdateActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error) {
	return conductor.activityPubActivityManager.UpdateActivityPubActivity(id, editor)
}

func (conductor *conductor) UpdateActivityPubActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubActivity, error) {
	return conductor.activityPubActivityManager.UpdateActivityPubActivityAtomic(transaction, id, editor)
}

func (conductor *conductor) RemoveActivityPubActivity(id int64, editor Identity) (IActivityPubActivity, error) {
	return conductor.activityPubActivityManager.RemoveActivityPubActivity(id, editor)
}

func (conductor *conductor) RemoveActivityPubActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubActivity, error) {
	return conductor.activityPubActivityManager.RemoveActivityPubActivityAtomic(transaction, id, editor)
}

// ActivityPubPublicKey

func (conductor *conductor) ActivityPubPublicKeyManager() IActivityPubPublicKeyManager {
	return conductor.activityPubPublicKeyManager
}

func (conductor *conductor) ActivityPubPublicKeyExists(id int64) bool {
	return conductor.activityPubPublicKeyManager.Exists(id)
}

func (conductor *conductor) ListActivityPubPublicKeys(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubPublicKeyCollection {
	return conductor.activityPubPublicKeyManager.ListActivityPubPublicKeys(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error) {
	return conductor.activityPubPublicKeyManager.GetActivityPubPublicKey(id, editor)
}

func (conductor *conductor) AddActivityPubPublicKey(editor Identity) (IActivityPubPublicKey, error) {
	return conductor.activityPubPublicKeyManager.AddActivityPubPublicKey(editor)
}

func (conductor *conductor) AddActivityPubPublicKeyAtomic(transaction ITransaction, editor Identity) (IActivityPubPublicKey, error) {
	return conductor.activityPubPublicKeyManager.AddActivityPubPublicKeyAtomic(transaction, editor)
}

func (conductor *conductor) LogActivityPubPublicKey(source string, editor Identity, payload string) {
	conductor.activityPubPublicKeyManager.Log(source, editor, payload)
}

func (conductor *conductor) UpdateActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error) {
	return conductor.activityPubPublicKeyManager.UpdateActivityPubPublicKey(id, editor)
}

func (conductor *conductor) UpdateActivityPubPublicKeyAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubPublicKey, error) {
	return conductor.activityPubPublicKeyManager.UpdateActivityPubPublicKeyAtomic(transaction, id, editor)
}

func (conductor *conductor) RemoveActivityPubPublicKey(id int64, editor Identity) (IActivityPubPublicKey, error) {
	return conductor.activityPubPublicKeyManager.RemoveActivityPubPublicKey(id, editor)
}

func (conductor *conductor) RemoveActivityPubPublicKeyAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubPublicKey, error) {
	return conductor.activityPubPublicKeyManager.RemoveActivityPubPublicKeyAtomic(transaction, id, editor)
}

// ActivityPubLink

func (conductor *conductor) ActivityPubLinkManager() IActivityPubLinkManager {
	return conductor.activityPubLinkManager
}

func (conductor *conductor) ActivityPubLinkExists(id int64) bool {
	return conductor.activityPubLinkManager.Exists(id)
}

func (conductor *conductor) ListActivityPubLinks(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubLinkCollection {
	return conductor.activityPubLinkManager.ListActivityPubLinks(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetActivityPubLink(id int64, editor Identity) (IActivityPubLink, error) {
	return conductor.activityPubLinkManager.GetActivityPubLink(id, editor)
}

func (conductor *conductor) AddActivityPubLink(editor Identity) (IActivityPubLink, error) {
	return conductor.activityPubLinkManager.AddActivityPubLink(editor)
}

func (conductor *conductor) AddActivityPubLinkAtomic(transaction ITransaction, editor Identity) (IActivityPubLink, error) {
	return conductor.activityPubLinkManager.AddActivityPubLinkAtomic(transaction, editor)
}

func (conductor *conductor) LogActivityPubLink(source string, editor Identity, payload string) {
	conductor.activityPubLinkManager.Log(source, editor, payload)
}

func (conductor *conductor) UpdateActivityPubLink(id int64, editor Identity) (IActivityPubLink, error) {
	return conductor.activityPubLinkManager.UpdateActivityPubLink(id, editor)
}

func (conductor *conductor) UpdateActivityPubLinkAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubLink, error) {
	return conductor.activityPubLinkManager.UpdateActivityPubLinkAtomic(transaction, id, editor)
}

func (conductor *conductor) RemoveActivityPubLink(id int64, editor Identity) (IActivityPubLink, error) {
	return conductor.activityPubLinkManager.RemoveActivityPubLink(id, editor)
}

func (conductor *conductor) RemoveActivityPubLinkAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubLink, error) {
	return conductor.activityPubLinkManager.RemoveActivityPubLinkAtomic(transaction, id, editor)
}

// ActivityPubMedia

func (conductor *conductor) ActivityPubMediaManager() IActivityPubMediaManager {
	return conductor.activityPubMediaManager
}

func (conductor *conductor) ActivityPubMediaExists(id int64) bool {
	return conductor.activityPubMediaManager.Exists(id)
}

func (conductor *conductor) ListActivityPubMedias(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubMediaCollection {
	return conductor.activityPubMediaManager.ListActivityPubMedias(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error) {
	return conductor.activityPubMediaManager.GetActivityPubMedia(id, editor)
}

func (conductor *conductor) AddActivityPubMedia(editor Identity) (IActivityPubMedia, error) {
	return conductor.activityPubMediaManager.AddActivityPubMedia(editor)
}

func (conductor *conductor) AddActivityPubMediaAtomic(transaction ITransaction, editor Identity) (IActivityPubMedia, error) {
	return conductor.activityPubMediaManager.AddActivityPubMediaAtomic(transaction, editor)
}

func (conductor *conductor) LogActivityPubMedia(source string, editor Identity, payload string) {
	conductor.activityPubMediaManager.Log(source, editor, payload)
}

func (conductor *conductor) UpdateActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error) {
	return conductor.activityPubMediaManager.UpdateActivityPubMedia(id, editor)
}

func (conductor *conductor) UpdateActivityPubMediaAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubMedia, error) {
	return conductor.activityPubMediaManager.UpdateActivityPubMediaAtomic(transaction, id, editor)
}

func (conductor *conductor) RemoveActivityPubMedia(id int64, editor Identity) (IActivityPubMedia, error) {
	return conductor.activityPubMediaManager.RemoveActivityPubMedia(id, editor)
}

func (conductor *conductor) RemoveActivityPubMediaAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubMedia, error) {
	return conductor.activityPubMediaManager.RemoveActivityPubMediaAtomic(transaction, id, editor)
}

// ActivityPubIncomingActivity

func (conductor *conductor) ActivityPubIncomingActivityManager() IActivityPubIncomingActivityManager {
	return conductor.activityPubIncomingActivityManager
}

func (conductor *conductor) ActivityPubIncomingActivityExists(id int64) bool {
	return conductor.activityPubIncomingActivityManager.Exists(id)
}

func (conductor *conductor) ListActivityPubIncomingActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubIncomingActivityCollection {
	return conductor.activityPubIncomingActivityManager.ListActivityPubIncomingActivities(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetActivityPubIncomingActivity(id int64, editor Identity) (IActivityPubIncomingActivity, error) {
	return conductor.activityPubIncomingActivityManager.GetActivityPubIncomingActivity(id, editor)
}

func (conductor *conductor) AddActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	return conductor.activityPubIncomingActivityManager.AddActivityPubIncomingActivity(identityId, uniqueIdentifier, timestamp, from, to, content, raw, editor)
}

func (conductor *conductor) AddActivityPubIncomingActivityAtomic(transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	return conductor.activityPubIncomingActivityManager.AddActivityPubIncomingActivityAtomic(transaction, identityId, uniqueIdentifier, timestamp, from, to, content, raw, editor)
}

func (conductor *conductor) LogActivityPubIncomingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor Identity, payload string) {
	conductor.activityPubIncomingActivityManager.Log(identityId, uniqueIdentifier, timestamp, from, to, content, raw, source, editor, payload)
}

func (conductor *conductor) UpdateActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	return conductor.activityPubIncomingActivityManager.UpdateActivityPubIncomingActivity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, editor)
}

func (conductor *conductor) UpdateActivityPubIncomingActivityAtomic(transaction ITransaction, id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubIncomingActivity, error) {
	return conductor.activityPubIncomingActivityManager.UpdateActivityPubIncomingActivityAtomic(transaction, id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, editor)
}

func (conductor *conductor) RemoveActivityPubIncomingActivity(id int64, editor Identity) (IActivityPubIncomingActivity, error) {
	return conductor.activityPubIncomingActivityManager.RemoveActivityPubIncomingActivity(id, editor)
}

func (conductor *conductor) RemoveActivityPubIncomingActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubIncomingActivity, error) {
	return conductor.activityPubIncomingActivityManager.RemoveActivityPubIncomingActivityAtomic(transaction, id, editor)
}

func (conductor *conductor) ListActivityPubIncomingActivitiesByIdentity(identityId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubIncomingActivityCollection {
	return conductor.activityPubIncomingActivityManager.ListActivityPubIncomingActivitiesByIdentity(identityId, pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) ForEachActivityPubIncomingActivityByIdentity(identityId int64, iterator ActivityPubIncomingActivityIterator) {
	conductor.activityPubIncomingActivityManager.ForEachByIdentity(identityId, iterator)
}

// ActivityPubOutgoingActivity

func (conductor *conductor) ActivityPubOutgoingActivityManager() IActivityPubOutgoingActivityManager {
	return conductor.activityPubOutgoingActivityManager
}

func (conductor *conductor) ActivityPubOutgoingActivityExists(id int64) bool {
	return conductor.activityPubOutgoingActivityManager.Exists(id)
}

func (conductor *conductor) ListActivityPubOutgoingActivities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubOutgoingActivityCollection {
	return conductor.activityPubOutgoingActivityManager.ListActivityPubOutgoingActivities(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetActivityPubOutgoingActivity(id int64, editor Identity) (IActivityPubOutgoingActivity, error) {
	return conductor.activityPubOutgoingActivityManager.GetActivityPubOutgoingActivity(id, editor)
}

func (conductor *conductor) AddActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	return conductor.activityPubOutgoingActivityManager.AddActivityPubOutgoingActivity(identityId, uniqueIdentifier, timestamp, from, to, content, raw, editor)
}

func (conductor *conductor) AddActivityPubOutgoingActivityAtomic(transaction ITransaction, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	return conductor.activityPubOutgoingActivityManager.AddActivityPubOutgoingActivityAtomic(transaction, identityId, uniqueIdentifier, timestamp, from, to, content, raw, editor)
}

func (conductor *conductor) LogActivityPubOutgoingActivity(identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, source string, editor Identity, payload string) {
	conductor.activityPubOutgoingActivityManager.Log(identityId, uniqueIdentifier, timestamp, from, to, content, raw, source, editor, payload)
}

func (conductor *conductor) UpdateActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	return conductor.activityPubOutgoingActivityManager.UpdateActivityPubOutgoingActivity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, editor)
}

func (conductor *conductor) UpdateActivityPubOutgoingActivityAtomic(transaction ITransaction, id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string, editor Identity) (IActivityPubOutgoingActivity, error) {
	return conductor.activityPubOutgoingActivityManager.UpdateActivityPubOutgoingActivityAtomic(transaction, id, identityId, uniqueIdentifier, timestamp, from, to, content, raw, editor)
}

func (conductor *conductor) RemoveActivityPubOutgoingActivity(id int64, editor Identity) (IActivityPubOutgoingActivity, error) {
	return conductor.activityPubOutgoingActivityManager.RemoveActivityPubOutgoingActivity(id, editor)
}

func (conductor *conductor) RemoveActivityPubOutgoingActivityAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubOutgoingActivity, error) {
	return conductor.activityPubOutgoingActivityManager.RemoveActivityPubOutgoingActivityAtomic(transaction, id, editor)
}

func (conductor *conductor) ListActivityPubOutgoingActivitiesByIdentity(identityId int64, pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubOutgoingActivityCollection {
	return conductor.activityPubOutgoingActivityManager.ListActivityPubOutgoingActivitiesByIdentity(identityId, pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) ForEachActivityPubOutgoingActivityByIdentity(identityId int64, iterator ActivityPubOutgoingActivityIterator) {
	conductor.activityPubOutgoingActivityManager.ForEachByIdentity(identityId, iterator)
}

// ActivityPubFollower

func (conductor *conductor) ActivityPubFollowerManager() IActivityPubFollowerManager {
	return conductor.activityPubFollowerManager
}

func (conductor *conductor) ActivityPubFollowerExists(id int64) bool {
	return conductor.activityPubFollowerManager.Exists(id)
}

func (conductor *conductor) ListActivityPubFollowers(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IActivityPubFollowerCollection {
	return conductor.activityPubFollowerManager.ListActivityPubFollowers(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetActivityPubFollower(id int64, editor Identity) (IActivityPubFollower, error) {
	return conductor.activityPubFollowerManager.GetActivityPubFollower(id, editor)
}

func (conductor *conductor) AddActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	return conductor.activityPubFollowerManager.AddActivityPubFollower(handle, inbox, subject, activity, accepted, editor)
}

func (conductor *conductor) AddActivityPubFollowerAtomic(transaction ITransaction, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	return conductor.activityPubFollowerManager.AddActivityPubFollowerAtomic(transaction, handle, inbox, subject, activity, accepted, editor)
}

func (conductor *conductor) LogActivityPubFollower(handle string, inbox string, subject string, activity string, accepted bool, source string, editor Identity, payload string) {
	conductor.activityPubFollowerManager.Log(handle, inbox, subject, activity, accepted, source, editor, payload)
}

func (conductor *conductor) UpdateActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	return conductor.activityPubFollowerManager.UpdateActivityPubFollower(id, handle, inbox, subject, activity, accepted, editor)
}

func (conductor *conductor) UpdateActivityPubFollowerAtomic(transaction ITransaction, id int64, handle string, inbox string, subject string, activity string, accepted bool, editor Identity) (IActivityPubFollower, error) {
	return conductor.activityPubFollowerManager.UpdateActivityPubFollowerAtomic(transaction, id, handle, inbox, subject, activity, accepted, editor)
}

func (conductor *conductor) RemoveActivityPubFollower(id int64, editor Identity) (IActivityPubFollower, error) {
	return conductor.activityPubFollowerManager.RemoveActivityPubFollower(id, editor)
}

func (conductor *conductor) RemoveActivityPubFollowerAtomic(transaction ITransaction, id int64, editor Identity) (IActivityPubFollower, error) {
	return conductor.activityPubFollowerManager.RemoveActivityPubFollowerAtomic(transaction, id, editor)
}

// Spi

func (conductor *conductor) SpiManager() ISpiManager {
	return conductor.spiManager
}

func (conductor *conductor) SpiExists(id int64) bool {
	return conductor.spiManager.Exists(id)
}

func (conductor *conductor) ListSpis(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ISpiCollection {
	return conductor.spiManager.ListSpis(pageIndex, pageSize, criteria, editor)
}

func (conductor *conductor) GetSpi(id int64, editor Identity) (ISpi, error) {
	return conductor.spiManager.GetSpi(id, editor)
}

func (conductor *conductor) AddSpi(editor Identity) (ISpi, error) {
	return conductor.spiManager.AddSpi(editor)
}

func (conductor *conductor) AddSpiAtomic(transaction ITransaction, editor Identity) (ISpi, error) {
	return conductor.spiManager.AddSpiAtomic(transaction, editor)
}

func (conductor *conductor) LogSpi(source string, editor Identity, payload string) {
	conductor.spiManager.Log(source, editor, payload)
}

func (conductor *conductor) UpdateSpi(id int64, editor Identity) (ISpi, error) {
	return conductor.spiManager.UpdateSpi(id, editor)
}

func (conductor *conductor) UpdateSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error) {
	return conductor.spiManager.UpdateSpiAtomic(transaction, id, editor)
}

func (conductor *conductor) RemoveSpi(id int64, editor Identity) (ISpi, error) {
	return conductor.spiManager.RemoveSpi(id, editor)
}

func (conductor *conductor) RemoveSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error) {
	return conductor.spiManager.RemoveSpiAtomic(transaction, id, editor)
}

func (conductor *conductor) Echo(document IDocument, editor Identity) (IEchoResult, error) {
	return conductor.spiManager.Echo(document, editor)
}

func (conductor *conductor) CheckUsernameAvailability(username string, editor Identity) (ICheckUsernameAvailabilityResult, error) {
	return conductor.spiManager.CheckUsernameAvailability(username, editor)
}

func (conductor *conductor) Signup(username string, email string, password string, editor Identity) (ISignupResult, error) {
	return conductor.spiManager.Signup(username, email, password, editor)
}

func (conductor *conductor) Verify(email string, token string, code string, editor Identity) (IVerifyResult, error) {
	return conductor.spiManager.Verify(email, token, code, editor)
}

func (conductor *conductor) Login(email string, password string, editor Identity) (ILoginResult, error) {
	return conductor.spiManager.Login(email, password, editor)
}

func (conductor *conductor) GetProfileByUser(editor Identity) (IGetProfileByUserResult, error) {
	return conductor.spiManager.GetProfileByUser(editor)
}

func (conductor *conductor) UpdateProfileByUser(displayName string, avatar string, banner string, summary string, github string, editor Identity) (IUpdateProfileByUserResult, error) {
	return conductor.spiManager.UpdateProfileByUser(displayName, avatar, banner, summary, github, editor)
}

func (conductor *conductor) ChangePassword(currentPassword string, newPassword string, editor Identity) (IChangePasswordResult, error) {
	return conductor.spiManager.ChangePassword(currentPassword, newPassword, editor)
}

func (conductor *conductor) Logout(editor Identity) (ILogoutResult, error) {
	return conductor.spiManager.Logout(editor)
}

func (conductor *conductor) Webfinger(resource string, editor Identity) (IWebfingerResult, error) {
	return conductor.spiManager.Webfinger(resource, editor)
}

func (conductor *conductor) GetPackages(editor Identity) (IGetPackagesResult, error) {
	return conductor.spiManager.GetPackages(editor)
}

func (conductor *conductor) GetActor(username string, editor Identity) (IGetActorResult, error) {
	return conductor.spiManager.GetActor(username, editor)
}

func (conductor *conductor) FollowActor(username string, acct string, editor Identity) (IFollowActorResult, error) {
	return conductor.spiManager.FollowActor(username, acct, editor)
}

func (conductor *conductor) AuthorizeInteraction(uri string, editor Identity) (IAuthorizeInteractionResult, error) {
	return conductor.spiManager.AuthorizeInteraction(uri, editor)
}

func (conductor *conductor) GetFollowers(username string, editor Identity) (IGetFollowersResult, error) {
	return conductor.spiManager.GetFollowers(username, editor)
}

func (conductor *conductor) GetFollowing(username string, editor Identity) (IGetFollowingResult, error) {
	return conductor.spiManager.GetFollowing(username, editor)
}

func (conductor *conductor) PostToOutbox(username string, body []byte, editor Identity) (IPostToOutboxResult, error) {
	return conductor.spiManager.PostToOutbox(username, body, editor)
}

func (conductor *conductor) GetOutbox(username string, editor Identity) (IGetOutboxResult, error) {
	return conductor.spiManager.GetOutbox(username, editor)
}

func (conductor *conductor) PostToInbox(username string, body []byte, editor Identity) (IPostToInboxResult, error) {
	return conductor.spiManager.PostToInbox(username, body, editor)
}

func (conductor *conductor) GetInbox(username string, editor Identity) (IGetInboxResult, error) {
	return conductor.spiManager.GetInbox(username, editor)
}

func (conductor *conductor) NewDocument(id int64, content string) (IDocument, error) {
	return NewDocument(id, content)
}

func (conductor *conductor) NewSystemSchedule(id int64, enabled bool, config string) (ISystemSchedule, error) {
	return NewSystemSchedule(id, enabled, config)
}

func (conductor *conductor) NewIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) (IIdentity, error) {
	return NewIdentity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount)
}

func (conductor *conductor) NewAccessControl(id int64, key uint64, value uint64) (IAccessControl, error) {
	return NewAccessControl(id, key, value)
}

func (conductor *conductor) NewRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) (IRemoteActivity, error) {
	return NewRemoteActivity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp)
}

func (conductor *conductor) NewCategoryType(id int64, description string) (ICategoryType, error) {
	return NewCategoryType(id, description)
}

func (conductor *conductor) NewCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string) (ICategory, error) {
	return NewCategory(id, categoryTypeId, categoryId, title, description)
}

func (conductor *conductor) NewUser(id int64, github string) (IUser, error) {
	return NewUser(id, github)
}

func (conductor *conductor) NewActivityPubObject() (IActivityPubObject, error) {
	return NewActivityPubObject()
}

func (conductor *conductor) NewActivityPubActivity() (IActivityPubActivity, error) {
	return NewActivityPubActivity()
}

func (conductor *conductor) NewActivityPubPublicKey() (IActivityPubPublicKey, error) {
	return NewActivityPubPublicKey()
}

func (conductor *conductor) NewActivityPubLink() (IActivityPubLink, error) {
	return NewActivityPubLink()
}

func (conductor *conductor) NewActivityPubMedia() (IActivityPubMedia, error) {
	return NewActivityPubMedia()
}

func (conductor *conductor) NewActivityPubIncomingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) (IActivityPubIncomingActivity, error) {
	return NewActivityPubIncomingActivity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
}

func (conductor *conductor) NewActivityPubOutgoingActivity(id int64, identityId int64, uniqueIdentifier string, timestamp int64, from string, to string, content string, raw string) (IActivityPubOutgoingActivity, error) {
	return NewActivityPubOutgoingActivity(id, identityId, uniqueIdentifier, timestamp, from, to, content, raw)
}

func (conductor *conductor) NewActivityPubFollower(id int64, handle string, inbox string, subject string, activity string, accepted bool) (IActivityPubFollower, error) {
	return NewActivityPubFollower(id, handle, inbox, subject, activity, accepted)
}

func (conductor *conductor) NewSpi() (ISpi, error) {
	return NewSpi()
}

func (conductor *conductor) NewEchoResult(document IDocument, _ interface{}) IEchoResult {
	return NewEchoResult(document, nil)
}

func (conductor *conductor) NewCheckUsernameAvailabilityResult(isAvailable bool, _ interface{}) ICheckUsernameAvailabilityResult {
	return NewCheckUsernameAvailabilityResult(isAvailable, nil)
}

func (conductor *conductor) NewSignupResult(token string, code string, _ interface{}) ISignupResult {
	return NewSignupResult(token, code, nil)
}

func (conductor *conductor) NewVerifyResult(token string, _ interface{}) IVerifyResult {
	return NewVerifyResult(token, nil)
}

func (conductor *conductor) NewLoginResult(username string, token string, _ interface{}) ILoginResult {
	return NewLoginResult(username, token, nil)
}

func (conductor *conductor) NewGetProfileByUserResult(username string, displayName string, avatar string, banner string, summary string, github string, _ interface{}) IGetProfileByUserResult {
	return NewGetProfileByUserResult(username, displayName, avatar, banner, summary, github, nil)
}

func (conductor *conductor) NewUpdateProfileByUserResult(displayName string, avatar string, banner string, summary string, github string, _ interface{}) IUpdateProfileByUserResult {
	return NewUpdateProfileByUserResult(displayName, avatar, banner, summary, github, nil)
}

func (conductor *conductor) NewChangePasswordResult(_ interface{}) IChangePasswordResult {
	return NewChangePasswordResult(nil)
}

func (conductor *conductor) NewLogoutResult(_ interface{}) ILogoutResult {
	return NewLogoutResult(nil)
}

func (conductor *conductor) NewWebfingerResult(aliases []string, links []IActivityPubLink, subject string, _ interface{}) IWebfingerResult {
	return NewWebfingerResult(aliases, links, subject, nil)
}

func (conductor *conductor) NewGetPackagesResult(body []byte, _ interface{}) IGetPackagesResult {
	return NewGetPackagesResult(body, nil)
}

func (conductor *conductor) NewGetActorResult(context []string, id string, followers string, following string, inbox string, outbox string, name string, preferredUsername string, type_ string, url string, icon IActivityPubMedia, image IActivityPubMedia, publicKey IActivityPubPublicKey, summary string, published string, _ interface{}) IGetActorResult {
	return NewGetActorResult(context, id, followers, following, inbox, outbox, name, preferredUsername, type_, url, icon, image, publicKey, summary, published, nil)
}

func (conductor *conductor) NewFollowActorResult(url string, _ interface{}) IFollowActorResult {
	return NewFollowActorResult(url, nil)
}

func (conductor *conductor) NewAuthorizeInteractionResult(uri string, success bool, _ interface{}) IAuthorizeInteractionResult {
	return NewAuthorizeInteractionResult(uri, success, nil)
}

func (conductor *conductor) NewGetFollowersResult(context string, id string, type_ string, totalItems int32, orderedItems []string, first string, _ interface{}) IGetFollowersResult {
	return NewGetFollowersResult(context, id, type_, totalItems, orderedItems, first, nil)
}

func (conductor *conductor) NewGetFollowingResult(context string, id string, type_ string, totalItems int32, orderedItems []string, first string, _ interface{}) IGetFollowingResult {
	return NewGetFollowingResult(context, id, type_, totalItems, orderedItems, first, nil)
}

func (conductor *conductor) NewPostToOutboxResult(body []byte, _ interface{}) IPostToOutboxResult {
	return NewPostToOutboxResult(body, nil)
}

func (conductor *conductor) NewGetOutboxResult(context string, id string, type_ string, totalItems int32, orderedItems []IActivityPubActivity, first string, _ interface{}) IGetOutboxResult {
	return NewGetOutboxResult(context, id, type_, totalItems, orderedItems, first, nil)
}

func (conductor *conductor) NewPostToInboxResult(body []byte, _ interface{}) IPostToInboxResult {
	return NewPostToInboxResult(body, nil)
}

func (conductor *conductor) NewGetInboxResult(context string, id string, type_ string, totalItems int32, orderedItems []IActivityPubActivity, first string, _ interface{}) IGetInboxResult {
	return NewGetInboxResult(context, id, type_, totalItems, orderedItems, first, nil)
}

func (conductor *conductor) LogRemoteCall(context IContext, eventType uint32, source string, input, result interface{}, err error) {
	if !context.Configuration().IsTrafficRecordEnabled() {
		return
	}

	errorMessage := ""
	if err != nil {
		errorMessage = strings.TrimPrefix(err.Error(), "ERROR_MESSAGE_")
	}

	if _, marshalError := json.Marshal(input); marshalError != nil {
		input = fmt.Sprintf("%s", input)
		context.Logger().Error(fmt.Sprintf("LRC_JSON_INPUT: %s %s", marshalError, input))
	}

	if _, marshalError := json.Marshal(result); marshalError != nil {
		result = fmt.Sprintf("%s", result)
		context.Logger().Error(fmt.Sprintf("LRC_JSON_RESULT: %s %s", marshalError, result))
	}

	data, marshalError := json.Marshal(&struct {
		Operation           string      `json:"operation"`
		Identity            int64       `json:"identity"`
		Token               string      `json:"token"`
		RequestId           uint64      `json:"request_id"`
		ClientName          string      `json:"client_name"`
		ClientVersion       int32       `json:"client_version"`
		ClientLatestVersion int32       `json:"client_latest_version"`
		ServerVersion       int32       `json:"server_version"`
		ApiVersion          int32       `json:"api_version"`
		Input               interface{} `json:"input"`
		Result              interface{} `json:"result"`
		Error               string      `json:"error"`
		Timestamp           int64       `json:"timestamp"`
	}{
		Operation:           source,
		Identity:            context.Identity().Id(),
		Token:               context.Token(),
		RequestId:           context.RequestId(),
		ClientName:          context.ClientName(),
		ClientVersion:       context.ClientVersion(),
		ClientLatestVersion: context.ClientLatestVersion(),
		ServerVersion:       context.ServerVersion(),
		ApiVersion:          context.ApiVersion(),
		Input:               input,
		Result:              result,
		Error:               errorMessage,
		Timestamp:           context.Timestamp().UnixNano(),
	})

	if marshalError != nil {
		data = []byte("{}")
		context.Logger().Error(fmt.Sprintf("LRC_JSON: %s %s %s", marshalError, input, result))
	}

	identity := context.Identity()
	conductor.RemoteActivityManager().Log(source, time.Since(context.Timestamp()).Nanoseconds(), err == nil, errorMessage, identity.RemoteAddress(), identity.UserAgent(), eventType, context.Timestamp().UnixNano(), source, identity, string(data))
}

//endregion

//region IAssertionResult Implementation

type assertionResult struct {
	condition bool
}

func (result *assertionResult) Or(err error) {
	if !result.condition {
		panic(err.Error())
	}
}

//endregion
