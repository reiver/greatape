package core

import (
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/security"
	"rail.town/infrastructure/app/validators"
	. "rail.town/infrastructure/components/constants"
	. "rail.town/infrastructure/components/contracts"
	. "rail.town/infrastructure/components/contracts/model"
	. "rail.town/infrastructure/components/model/entity"
	"rail.town/infrastructure/components/model/repository"
)

// noinspection GoSnakeCaseUsage
const IDENTITY_MANAGER = "IdentityManager"

type identityManager struct {
	systemComponent
	cache                IdentityCache
	accessControlHandler IAccessControlHandler
}

func newIdentityManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) IIdentityManager {
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize

	manager := &identityManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewIdentityCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *identityManager) Name() string {
	return IDENTITY_MANAGER
}

func (manager *identityManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *identityManager) Load() error {
	identityEntities, err := repository.Identities.FetchAll()
	if err != nil {
		return err
	}

	identities := make(SystemObjectCache)
	for _, identityEntity := range identityEntities {
		if identity, err := NewIdentityFromEntity(identityEntity); err == nil {
			identities[identity.Id()] = identity
		} else {
			return err
		}
	}

	manager.cache.Load(identities)
	return nil
}

func (manager *identityManager) Reload() error {
	return manager.Load()
}

func (manager *identityManager) OnCacheChanged(callback IdentityCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *identityManager) Count() int {
	return manager.cache.Size()
}

func (manager *identityManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *identityManager) ExistsWhich(condition IdentityCondition) bool {
	var identities Identities
	manager.ForEach(func(identity IIdentity) {
		if condition(identity) {
			identities = append(identities, identity)
		}
	})

	return len(identities) > 0
}

func (manager *identityManager) ListIdentities(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) IIdentityCollection {
	return manager.Filter(IdentityPassThroughFilter)
}

func (manager *identityManager) GetIdentity(id int64, _ Identity) (IIdentity, error) {
	if identity := manager.Find(id); identity == nil {
		return nil, ERROR_IDENTITY_NOT_FOUND
	} else {
		return identity, nil
	}
}

func (manager *identityManager) AddIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(manager.UniqueId(), username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount)
	return manager.Apply(identityEntity, repository.Identities.Add, manager.cache.Put, editor)
}

func (manager *identityManager) AddIdentityWithCustomId(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount)
	return manager.Apply(identityEntity, repository.Identities.Add, manager.cache.Put, editor)
}

func (manager *identityManager) AddIdentityObject(identity IIdentity, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(manager.UniqueId(), identity.Username(), identity.PhoneNumber(), identity.PhoneNumberConfirmed(), identity.FirstName(), identity.LastName(), identity.DisplayName(), identity.Email(), identity.EmailConfirmed(), identity.Avatar(), identity.Banner(), identity.Summary(), identity.Token(), identity.MultiFactor(), identity.Hash(), identity.Salt(), identity.PublicKey(), identity.PrivateKey(), identity.Permission(), identity.Restriction(), identity.LastLogin(), identity.LoginCount())
	return manager.Apply(identityEntity, repository.Identities.Add, manager.cache.Put, editor)
}

func (manager *identityManager) AddIdentityAtomic(transaction ITransaction, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(manager.UniqueId(), username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount)
	return manager.ApplyAtomic(transaction, identityEntity, repository.Identities.AddAtomic, manager.cache.Put, editor)
}

func (manager *identityManager) AddIdentityWithCustomIdAtomic(id int64, transaction ITransaction, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount)
	return manager.ApplyAtomic(transaction, identityEntity, repository.Identities.AddAtomic, manager.cache.Put, editor)
}

func (manager *identityManager) AddIdentityObjectAtomic(transaction ITransaction, identity IIdentity, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(manager.UniqueId(), identity.Username(), identity.PhoneNumber(), identity.PhoneNumberConfirmed(), identity.FirstName(), identity.LastName(), identity.DisplayName(), identity.Email(), identity.EmailConfirmed(), identity.Avatar(), identity.Banner(), identity.Summary(), identity.Token(), identity.MultiFactor(), identity.Hash(), identity.Salt(), identity.PublicKey(), identity.PrivateKey(), identity.Permission(), identity.Restriction(), identity.LastLogin(), identity.LoginCount())
	return manager.ApplyAtomic(transaction, identityEntity, repository.Identities.AddAtomic, manager.cache.Put, editor)
}

func (manager *identityManager) Log(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, source string, editor Identity, payload string) {
	identityPipeEntity := NewIdentityPipeEntity(manager.UniqueId(), username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, source, editor.Id(), payload)
	repository.Pipe.Insert(identityPipeEntity)

	identity, err := NewIdentityFromEntity(identityPipeEntity)
	if err != nil {
		manager.Logger().Error(err)
	} else {
		manager.cache.Put(identity.Id(), identity)
	}
}

func (manager *identityManager) UpdateIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount)
	return manager.Apply(identityEntity, repository.Identities.Update, manager.cache.Put, editor)
}

func (manager *identityManager) UpdateIdentityObject(id int64, identity IIdentity, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(id, identity.Username(), identity.PhoneNumber(), identity.PhoneNumberConfirmed(), identity.FirstName(), identity.LastName(), identity.DisplayName(), identity.Email(), identity.EmailConfirmed(), identity.Avatar(), identity.Banner(), identity.Summary(), identity.Token(), identity.MultiFactor(), identity.Hash(), identity.Salt(), identity.PublicKey(), identity.PrivateKey(), identity.Permission(), identity.Restriction(), identity.LastLogin(), identity.LoginCount())
	return manager.Apply(identityEntity, repository.Identities.Update, manager.cache.Put, editor)
}

func (manager *identityManager) UpdateIdentityAtomic(transaction ITransaction, id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount)
	return manager.ApplyAtomic(transaction, identityEntity, repository.Identities.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *identityManager) UpdateIdentityObjectAtomic(transaction ITransaction, id int64, identity IIdentity, editor Identity) (IIdentity, error) {
	identityEntity := NewIdentityEntity(id, identity.Username(), identity.PhoneNumber(), identity.PhoneNumberConfirmed(), identity.FirstName(), identity.LastName(), identity.DisplayName(), identity.Email(), identity.EmailConfirmed(), identity.Avatar(), identity.Banner(), identity.Summary(), identity.Token(), identity.MultiFactor(), identity.Hash(), identity.Salt(), identity.PublicKey(), identity.PrivateKey(), identity.Permission(), identity.Restriction(), identity.LastLogin(), identity.LoginCount())
	return manager.ApplyAtomic(transaction, identityEntity, repository.Identities.UpdateAtomic, manager.cache.Put, editor)
}

func (manager *identityManager) AddOrUpdateIdentityObject(id int64, identity IIdentity, editor Identity) (IIdentity, error) {
	if manager.Exists(id) {
		return manager.UpdateIdentityObject(id, identity, editor)
	} else {
		return manager.AddIdentityObject(identity, editor)
	}
}

func (manager *identityManager) AddOrUpdateIdentityObjectAtomic(transaction ITransaction, id int64, identity IIdentity, editor Identity) (IIdentity, error) {
	if manager.Exists(id) {
		return manager.UpdateIdentityObjectAtomic(transaction, id, identity, editor)
	} else {
		return manager.AddIdentityObjectAtomic(transaction, identity, editor)
	}
}

func (manager *identityManager) RemoveIdentity(_ int64, _ Identity) (IIdentity, error) {
	return nil, ERROR_OPERATION_NOT_SUPPORTED
}

func (manager *identityManager) RemoveIdentityAtomic(_ ITransaction, _ int64, _ Identity) (IIdentity, error) {
	return nil, ERROR_OPERATION_NOT_SUPPORTED
}

func (manager *identityManager) Apply(identityEntity IIdentityEntity, repositoryHandler func(IIdentityEntity, int64) error, cacheHandler func(int64, Identity), editor Identity) (IIdentity, error) {
	result, err := NewIdentityFromEntity(identityEntity)
	if err != nil {
		return nil, err
	}

	if err := repositoryHandler(identityEntity, editor.Id()); err != nil {
		return nil, err
	}

	cacheHandler(result.Id(), result)
	return result, nil
}

func (manager *identityManager) ApplyAtomic(transaction ITransaction, identityEntity IIdentityEntity, repositoryHandler func(IRepositoryTransaction, IIdentityEntity, int64) error, cacheHandler func(int64, Identity), editor Identity) (IIdentity, error) {
	result, err := NewIdentityFromEntity(identityEntity)
	if err != nil {
		return nil, err
	}

	transaction.OnCommit(func() {
		cacheHandler(result.Id(), result)
	})

	if err := repositoryHandler(transaction, identityEntity, editor.Id()); err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *identityManager) Find(id int64) IIdentity {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(IIdentity)
	}
}

func (manager *identityManager) ForEach(iterator IdentityIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(IIdentity))
	})
}

func (manager *identityManager) Filter(predicate IdentityFilterPredicate) IIdentityCollection {
	identities := NewIdentities()
	if predicate == nil {
		return identities
	}

	manager.ForEach(func(identity IIdentity) {
		if predicate(identity) {
			identities.Append(identity)
		}
	})

	return identities
}

func (manager *identityManager) Map(predicate IdentityMapPredicate) IIdentityCollection {
	identities := NewIdentities()
	if predicate == nil {
		return identities
	}

	manager.ForEach(func(identity IIdentity) {
		identities.Append(predicate(identity))
	})

	return identities
}

//region ISecurityHandler Implementation

func (manager *identityManager) AccessControlHandler() IAccessControlHandler {
	return manager.accessControlHandler
}

func (manager *identityManager) SetAccessControlHandler(handler IAccessControlHandler) {
	manager.accessControlHandler = handler
}

func (manager *identityManager) Validate(phoneNumber string, password string) (string, error) {
	identity, exists := manager.cache.GetByPhoneNumber(phoneNumber)
	if !exists {
		return "", ERROR_USER_NOT_REGISTERED
	}

	if identity.IsRestricted() {
		return "", ERROR_ACCOUNT_BLOCKED
	}

	if identity.MultiFactor() && manager.GenerateHash(password, identity.Salt()) != identity.Hash() {
		return "", ERROR_INVALID_CREDENTIALS
	}

	confirmationCode := manager.GenerateCode()
	if manager.IsStagingEnvironment() || identity.MultiFactor() {
		confirmationCode = "123456"
	}

	token := manager.GenerateHash(phoneNumber, confirmationCode)
	manager.cache.StoreAuthorizationInfo(token, phoneNumber, confirmationCode)
	if manager.IsProductionEnvironment() && !identity.MultiFactor() {
		manager.SMS(phoneNumber, "Confirmation Code: %s", confirmationCode)
	}

	return token, nil
}

func (manager *identityManager) Verify(token string, confirmationCode string) (string, uint64, error) {
	storedPhoneNumber, storedConfirmationCode, err := manager.cache.RetrieveAuthorizationInfo(token)
	if err != nil {
		return "", 0, ERROR_INVALID_TOKEN
	}

	if confirmationCode != storedConfirmationCode {
		return "", 0, ERROR_INVALID_CONFIRMATION_CODE
	}

	identity, exists := manager.cache.GetByPhoneNumber(storedPhoneNumber)
	if !exists {
		return "", 0, ERROR_USER_NOT_REGISTERED
	}

	if _, knownUser := identity.(IIdentity); !knownUser {
		return "", 0, ERROR_UNKNOWN_USER
	}

	newToken := manager.GenerateHash(storedPhoneNumber, manager.GenerateCode())
	identity.(IIdentity).UpdateLastLogin(manager.UnixNano(), identity)
	identity.(IIdentity).UpdateLoginCount(identity.(IIdentity).LoginCount()+1, identity)
	if err := manager.UpdateToken(identity, newToken); err != nil {
		return "", 0, err
	}

	return newToken, identity.Permission(), nil
}

func (manager *identityManager) RefreshTokenCache(identity Identity, token string) error {
	manager.cache.RefreshToken(identity, token)
	return nil
}

func (manager *identityManager) Authenticate(token string, role Role, remoteAddress string, userAgent string) Identity {
	switch role {
	case ANONYMOUS:
		return NewAnonymousIdentity(token, remoteAddress, userAgent)
	default:
		if identity, exists := manager.cache.GetByToken(token); exists && identity.IsInRole(role) && identity.IsNotRestricted() {
			identity.SetRemoteAddress(remoteAddress)
			identity.SetUserAgent(userAgent)
			return identity
		}
	}

	return nil
}

func (manager *identityManager) SignOut(identity Identity) error {
	return manager.UpdateToken(identity, manager.GenerateUUID())
}

func (manager *identityManager) UpdateToken(identity Identity, token string) error {
	if err := repository.Identities.UpdateToken(identity.Id(), token, identity.Id()); err != nil {
		return err
	}

	manager.cache.RefreshToken(identity, token)
	return nil
}

//endregion
