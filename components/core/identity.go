package core

import (
	"fmt"

	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/reiver/greatape/components/contracts/model"
	"github.com/reiver/greatape/components/model/repository"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/system"
)

type identity struct {
	object
	username             string
	phoneNumber          string
	phoneNumberConfirmed bool
	firstName            string
	lastName             string
	displayName          string
	email                string
	emailConfirmed       bool
	avatar               string
	banner               string
	summary              string
	token                string
	multiFactor          bool
	hash                 string
	salt                 string
	publicKey            string
	privateKey           string
	permission           uint64
	restriction          uint32
	lastLogin            int64
	loginCount           uint32
	remoteAddress        string
	userAgent            string
	systemCallHandler    func(Identity, []string) error
}

// noinspection GoUnusedExportedFunction
func InitializeIdentity() {
	_ = ENABLE_SECURITY
	_ = ENABLE_CUSTOM_ERRORS
	_ = validators.Initialize
	_ = repository.Initialize
}

func NewIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) (IIdentity, error) {
	instance := &identity{
		object: object{
			id: id,
		},
		username:             username,
		phoneNumber:          phoneNumber,
		phoneNumberConfirmed: phoneNumberConfirmed,
		firstName:            firstName,
		lastName:             lastName,
		displayName:          displayName,
		email:                email,
		emailConfirmed:       emailConfirmed,
		avatar:               avatar,
		banner:               banner,
		summary:              summary,
		token:                token,
		multiFactor:          multiFactor,
		hash:                 hash,
		salt:                 salt,
		publicKey:            publicKey,
		privateKey:           privateKey,
		permission:           permission,
		restriction:          restriction,
		lastLogin:            lastLogin,
		loginCount:           loginCount,
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewIdentityFromEntity(entity IIdentityEntity) (IIdentity, error) {
	instance := &identity{
		object: object{
			id: entity.Id(),
		},
		username:             entity.Username(),
		phoneNumber:          entity.PhoneNumber(),
		phoneNumberConfirmed: entity.PhoneNumberConfirmed(),
		firstName:            entity.FirstName(),
		lastName:             entity.LastName(),
		displayName:          entity.DisplayName(),
		email:                entity.Email(),
		emailConfirmed:       entity.EmailConfirmed(),
		avatar:               entity.Avatar(),
		banner:               entity.Banner(),
		summary:              entity.Summary(),
		token:                entity.Token(),
		multiFactor:          entity.MultiFactor(),
		hash:                 entity.Hash(),
		salt:                 entity.Salt(),
		publicKey:            entity.PublicKey(),
		privateKey:           entity.PrivateKey(),
		permission:           entity.Permission(),
		restriction:          entity.Restriction(),
		lastLogin:            entity.LastLogin(),
		loginCount:           entity.LoginCount(),
	}

	if err := instance.Validate(); err != nil {
		return nil, err
	}

	return instance, nil
}

func NewAnonymousIdentity(token, remoteAddress string, userAgent string) IIdentity {
	return &identity{
		object:        object{id: 0},
		username:      "anonymous",
		phoneNumber:   "",
		token:         token,
		remoteAddress: remoteAddress,
		userAgent:     userAgent,
		permission:    ANONYMOUS,
		restriction:   0,
	}
}

func NewSystemIdentity() IIdentity {
	return &identity{
		object:      object{id: 1},
		username:    "system",
		permission:  ADMINISTRATOR,
		restriction: 0,
	}
}

func (identity *identity) SetSystemCallHandler(handler func(Identity, []string) error) {
	identity.systemCallHandler = handler
}

func (identity *identity) SystemCall(editor Identity, args []string) error {
	if identity.systemCallHandler == nil {
		return ERROR_INITIALIZE
	}

	return identity.systemCallHandler(editor, args)
}

func (identity *identity) Username() string {
	return identity.username
}

func (identity *identity) UpdateUsername(username string, editor Identity) {
	if err := repository.Identities.UpdateUsername(identity.id, username, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.username = username
}

func (identity *identity) UpdateUsernameAtomic(transaction ITransaction, username string, editor Identity) {
	transaction.OnCommit(func() {
		identity.username = username
	})

	if err := repository.Identities.UpdateUsernameAtomic(transaction, identity.id, username, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) PhoneNumber() string {
	return identity.phoneNumber
}

func (identity *identity) UpdatePhoneNumber(phoneNumber string, editor Identity) {
	if err := repository.Identities.UpdatePhoneNumber(identity.id, phoneNumber, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.phoneNumber = phoneNumber
}

func (identity *identity) UpdatePhoneNumberAtomic(transaction ITransaction, phoneNumber string, editor Identity) {
	transaction.OnCommit(func() {
		identity.phoneNumber = phoneNumber
	})

	if err := repository.Identities.UpdatePhoneNumberAtomic(transaction, identity.id, phoneNumber, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) PhoneNumberConfirmed() bool {
	return identity.phoneNumberConfirmed
}

func (identity *identity) UpdatePhoneNumberConfirmed(phoneNumberConfirmed bool, editor Identity) {
	if err := repository.Identities.UpdatePhoneNumberConfirmed(identity.id, phoneNumberConfirmed, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.phoneNumberConfirmed = phoneNumberConfirmed
}

func (identity *identity) UpdatePhoneNumberConfirmedAtomic(transaction ITransaction, phoneNumberConfirmed bool, editor Identity) {
	transaction.OnCommit(func() {
		identity.phoneNumberConfirmed = phoneNumberConfirmed
	})

	if err := repository.Identities.UpdatePhoneNumberConfirmedAtomic(transaction, identity.id, phoneNumberConfirmed, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) FirstName() string {
	return identity.firstName
}

func (identity *identity) UpdateFirstName(firstName string, editor Identity) {
	if err := repository.Identities.UpdateFirstName(identity.id, firstName, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.firstName = firstName
}

func (identity *identity) UpdateFirstNameAtomic(transaction ITransaction, firstName string, editor Identity) {
	transaction.OnCommit(func() {
		identity.firstName = firstName
	})

	if err := repository.Identities.UpdateFirstNameAtomic(transaction, identity.id, firstName, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) LastName() string {
	return identity.lastName
}

func (identity *identity) UpdateLastName(lastName string, editor Identity) {
	if err := repository.Identities.UpdateLastName(identity.id, lastName, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.lastName = lastName
}

func (identity *identity) UpdateLastNameAtomic(transaction ITransaction, lastName string, editor Identity) {
	transaction.OnCommit(func() {
		identity.lastName = lastName
	})

	if err := repository.Identities.UpdateLastNameAtomic(transaction, identity.id, lastName, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) DisplayName() string {
	return identity.displayName
}

func (identity *identity) UpdateDisplayName(displayName string, editor Identity) {
	if err := repository.Identities.UpdateDisplayName(identity.id, displayName, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.displayName = displayName
}

func (identity *identity) UpdateDisplayNameAtomic(transaction ITransaction, displayName string, editor Identity) {
	transaction.OnCommit(func() {
		identity.displayName = displayName
	})

	if err := repository.Identities.UpdateDisplayNameAtomic(transaction, identity.id, displayName, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) Email() string {
	return identity.email
}

func (identity *identity) UpdateEmail(email string, editor Identity) {
	if err := repository.Identities.UpdateEmail(identity.id, email, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.email = email
}

func (identity *identity) UpdateEmailAtomic(transaction ITransaction, email string, editor Identity) {
	transaction.OnCommit(func() {
		identity.email = email
	})

	if err := repository.Identities.UpdateEmailAtomic(transaction, identity.id, email, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) EmailConfirmed() bool {
	return identity.emailConfirmed
}

func (identity *identity) UpdateEmailConfirmed(emailConfirmed bool, editor Identity) {
	if err := repository.Identities.UpdateEmailConfirmed(identity.id, emailConfirmed, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.emailConfirmed = emailConfirmed
}

func (identity *identity) UpdateEmailConfirmedAtomic(transaction ITransaction, emailConfirmed bool, editor Identity) {
	transaction.OnCommit(func() {
		identity.emailConfirmed = emailConfirmed
	})

	if err := repository.Identities.UpdateEmailConfirmedAtomic(transaction, identity.id, emailConfirmed, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) Avatar() string {
	return identity.avatar
}

func (identity *identity) UpdateAvatar(avatar string, editor Identity) {
	if err := repository.Identities.UpdateAvatar(identity.id, avatar, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.avatar = avatar
}

func (identity *identity) UpdateAvatarAtomic(transaction ITransaction, avatar string, editor Identity) {
	transaction.OnCommit(func() {
		identity.avatar = avatar
	})

	if err := repository.Identities.UpdateAvatarAtomic(transaction, identity.id, avatar, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) Banner() string {
	return identity.banner
}

func (identity *identity) UpdateBanner(banner string, editor Identity) {
	if err := repository.Identities.UpdateBanner(identity.id, banner, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.banner = banner
}

func (identity *identity) UpdateBannerAtomic(transaction ITransaction, banner string, editor Identity) {
	transaction.OnCommit(func() {
		identity.banner = banner
	})

	if err := repository.Identities.UpdateBannerAtomic(transaction, identity.id, banner, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) Summary() string {
	return identity.summary
}

func (identity *identity) UpdateSummary(summary string, editor Identity) {
	if err := repository.Identities.UpdateSummary(identity.id, summary, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.summary = summary
}

func (identity *identity) UpdateSummaryAtomic(transaction ITransaction, summary string, editor Identity) {
	transaction.OnCommit(func() {
		identity.summary = summary
	})

	if err := repository.Identities.UpdateSummaryAtomic(transaction, identity.id, summary, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) Token() string {
	return identity.token
}

func (identity *identity) UpdateToken(token string, editor Identity) {
	if err := repository.Identities.UpdateToken(identity.id, token, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.token = token
}

func (identity *identity) UpdateTokenAtomic(transaction ITransaction, token string, editor Identity) {
	transaction.OnCommit(func() {
		identity.token = token
	})

	if err := repository.Identities.UpdateTokenAtomic(transaction, identity.id, token, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) MultiFactor() bool {
	return identity.multiFactor
}

func (identity *identity) UpdateMultiFactor(multiFactor bool, editor Identity) {
	if err := repository.Identities.UpdateMultiFactor(identity.id, multiFactor, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.multiFactor = multiFactor
}

func (identity *identity) UpdateMultiFactorAtomic(transaction ITransaction, multiFactor bool, editor Identity) {
	transaction.OnCommit(func() {
		identity.multiFactor = multiFactor
	})

	if err := repository.Identities.UpdateMultiFactorAtomic(transaction, identity.id, multiFactor, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) Hash() string {
	return identity.hash
}

func (identity *identity) UpdateHash(hash string, editor Identity) {
	if err := repository.Identities.UpdateHash(identity.id, hash, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.hash = hash
}

func (identity *identity) UpdateHashAtomic(transaction ITransaction, hash string, editor Identity) {
	transaction.OnCommit(func() {
		identity.hash = hash
	})

	if err := repository.Identities.UpdateHashAtomic(transaction, identity.id, hash, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) Salt() string {
	return identity.salt
}

func (identity *identity) UpdateSalt(salt string, editor Identity) {
	if err := repository.Identities.UpdateSalt(identity.id, salt, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.salt = salt
}

func (identity *identity) UpdateSaltAtomic(transaction ITransaction, salt string, editor Identity) {
	transaction.OnCommit(func() {
		identity.salt = salt
	})

	if err := repository.Identities.UpdateSaltAtomic(transaction, identity.id, salt, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) PublicKey() string {
	return identity.publicKey
}

func (identity *identity) UpdatePublicKey(publicKey string, editor Identity) {
	if err := repository.Identities.UpdatePublicKey(identity.id, publicKey, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.publicKey = publicKey
}

func (identity *identity) UpdatePublicKeyAtomic(transaction ITransaction, publicKey string, editor Identity) {
	transaction.OnCommit(func() {
		identity.publicKey = publicKey
	})

	if err := repository.Identities.UpdatePublicKeyAtomic(transaction, identity.id, publicKey, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) PrivateKey() string {
	return identity.privateKey
}

func (identity *identity) UpdatePrivateKey(privateKey string, editor Identity) {
	if err := repository.Identities.UpdatePrivateKey(identity.id, privateKey, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.privateKey = privateKey
}

func (identity *identity) UpdatePrivateKeyAtomic(transaction ITransaction, privateKey string, editor Identity) {
	transaction.OnCommit(func() {
		identity.privateKey = privateKey
	})

	if err := repository.Identities.UpdatePrivateKeyAtomic(transaction, identity.id, privateKey, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) Permission() uint64 {
	return identity.permission
}

func (identity *identity) UpdatePermission(permission uint64, editor Identity) {
	if err := repository.Identities.UpdatePermission(identity.id, permission, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.permission = permission
}

func (identity *identity) UpdatePermissionAtomic(transaction ITransaction, permission uint64, editor Identity) {
	transaction.OnCommit(func() {
		identity.permission = permission
	})

	if err := repository.Identities.UpdatePermissionAtomic(transaction, identity.id, permission, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) Restriction() uint32 {
	return identity.restriction
}

func (identity *identity) UpdateRestriction(restriction uint32, editor Identity) {
	if err := repository.Identities.UpdateRestriction(identity.id, restriction, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.restriction = restriction
}

func (identity *identity) UpdateRestrictionAtomic(transaction ITransaction, restriction uint32, editor Identity) {
	transaction.OnCommit(func() {
		identity.restriction = restriction
	})

	if err := repository.Identities.UpdateRestrictionAtomic(transaction, identity.id, restriction, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) LastLogin() int64 {
	return identity.lastLogin
}

func (identity *identity) UpdateLastLogin(lastLogin int64, editor Identity) {
	if err := repository.Identities.UpdateLastLogin(identity.id, lastLogin, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.lastLogin = lastLogin
}

func (identity *identity) UpdateLastLoginAtomic(transaction ITransaction, lastLogin int64, editor Identity) {
	transaction.OnCommit(func() {
		identity.lastLogin = lastLogin
	})

	if err := repository.Identities.UpdateLastLoginAtomic(transaction, identity.id, lastLogin, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) LoginCount() uint32 {
	return identity.loginCount
}

func (identity *identity) UpdateLoginCount(loginCount uint32, editor Identity) {
	if err := repository.Identities.UpdateLoginCount(identity.id, loginCount, editor.Id()); err != nil {
		panic(err.Error())
	}

	identity.loginCount = loginCount
}

func (identity *identity) UpdateLoginCountAtomic(transaction ITransaction, loginCount uint32, editor Identity) {
	transaction.OnCommit(func() {
		identity.loginCount = loginCount
	})

	if err := repository.Identities.UpdateLoginCountAtomic(transaction, identity.id, loginCount, editor.Id()); err != nil {
		panic(err.Error())
	}
}

func (identity *identity) RemoteAddress() string {
	return identity.remoteAddress
}

func (identity *identity) SetRemoteAddress(value string) {
	identity.remoteAddress = value
}

func (identity *identity) UserAgent() string {
	return identity.userAgent
}

func (identity *identity) SetUserAgent(value string) {
	identity.userAgent = value
}

func (identity *identity) Role() Role {
	return identity.permission
}

func (identity *identity) IsInRole(role Role) bool {
	identityRole := identity.permission & role
	return identityRole == role
}

func (identity *identity) IsRestricted() bool {
	return identity.restriction != 0
}

func (identity *identity) IsNotRestricted() bool {
	return identity.restriction == 0
}

func (identity *identity) Payload() Pointer {
	return identity
}

func (identity *identity) SetToken(token string) {
	identity.token = token
}

func (identity *identity) Validate() error {
	return nil
}

func (identity *identity) String() string {
	return fmt.Sprintf("Identity (Id: %d, Username: %v, PhoneNumber: %v, PhoneNumberConfirmed: %v, FirstName: %v, LastName: %v, DisplayName: %v, Email: %v, EmailConfirmed: %v, Avatar: %v, Banner: %v, Summary: %v, Token: %v, MultiFactor: %v, Hash: %v, Salt: %v, PublicKey: %v, PrivateKey: %v, Permission: %v, Restriction: %v, LastLogin: %v, LoginCount: %v)", identity.Id(), identity.Username(), identity.PhoneNumber(), identity.PhoneNumberConfirmed(), identity.FirstName(), identity.LastName(), identity.DisplayName(), identity.Email(), identity.EmailConfirmed(), identity.Avatar(), identity.Banner(), identity.Summary(), identity.Token(), identity.MultiFactor(), identity.Hash(), identity.Salt(), identity.PublicKey(), identity.PrivateKey(), identity.Permission(), identity.Restriction(), identity.LastLogin(), identity.LoginCount())
}

//------------------------------------------------------------------------------

type identities struct {
	collection Identities
}

// NewIdentities creates an empty collection of 'Identity' which is not thread-safe.
func NewIdentities() IIdentityCollection {
	return &identities{
		collection: make(Identities, 0),
	}
}

func (identities *identities) Count() int {
	return len(identities.collection)
}

func (identities *identities) IsEmpty() bool {
	return len(identities.collection) == 0
}

func (identities *identities) IsNotEmpty() bool {
	return len(identities.collection) > 0
}

func (identities *identities) HasExactlyOneItem() bool {
	return len(identities.collection) == 1
}

func (identities *identities) HasAtLeastOneItem() bool {
	return len(identities.collection) >= 1
}

func (identities *identities) First() IIdentity {
	return identities.collection[0]
}

func (identities *identities) Append(identity IIdentity) {
	identities.collection = append(identities.collection, identity)
}

func (identities *identities) ForEach(iterator IdentityIterator) {
	if iterator == nil {
		return
	}

	for _, value := range identities.collection {
		iterator(value)
	}
}

func (identities *identities) Array() Identities {
	return identities.collection
}

//------------------------------------------------------------------------------

func (dispatcher *dispatcher) IdentityExists(id int64) bool {
	return dispatcher.conductor.IdentityManager().Exists(id)
}

func (dispatcher *dispatcher) IdentityExistsWhich(condition IdentityCondition) bool {
	return dispatcher.conductor.IdentityManager().ExistsWhich(condition)
}

func (dispatcher *dispatcher) ListIdentities() IIdentityCollection {
	return dispatcher.conductor.IdentityManager().ListIdentities(0, 0, "", dispatcher.identity)
}

func (dispatcher *dispatcher) ForEachIdentity(iterator IdentityIterator) {
	dispatcher.conductor.IdentityManager().ForEach(iterator)
}

func (dispatcher *dispatcher) FilterIdentities(predicate IdentityFilterPredicate) IIdentityCollection {
	return dispatcher.conductor.IdentityManager().Filter(predicate)
}

func (dispatcher *dispatcher) MapIdentities(predicate IdentityMapPredicate) IIdentityCollection {
	return dispatcher.conductor.IdentityManager().Map(predicate)
}

func (dispatcher *dispatcher) GetIdentity(id int64) IIdentity {
	if identity, err := dispatcher.conductor.IdentityManager().GetIdentity(id, dispatcher.identity); err != nil {
		panic(err.Error())
	} else {
		return identity
	}
}

func (dispatcher *dispatcher) AddIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) IIdentity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if identity, err := dispatcher.conductor.IdentityManager().AddIdentityAtomic(transaction, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	} else {
		if identity, err := dispatcher.conductor.IdentityManager().AddIdentity(username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	}
}

func (dispatcher *dispatcher) AddIdentityWithCustomId(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) IIdentity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if identity, err := dispatcher.conductor.IdentityManager().AddIdentityWithCustomIdAtomic(id, transaction, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	} else {
		if identity, err := dispatcher.conductor.IdentityManager().AddIdentityWithCustomId(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	}
}

func (dispatcher *dispatcher) LogIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, source string, payload string) {
	dispatcher.conductor.IdentityManager().Log(username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, source, dispatcher.identity, payload)
}

func (dispatcher *dispatcher) UpdateIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) IIdentity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if identity, err := dispatcher.conductor.IdentityManager().UpdateIdentityAtomic(transaction, id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	} else {
		if identity, err := dispatcher.conductor.IdentityManager().UpdateIdentity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	}
}

// noinspection GoUnusedParameter
func (dispatcher *dispatcher) UpdateIdentityObject(object IObject, identity IIdentity) IIdentity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if identity, err := dispatcher.conductor.IdentityManager().UpdateIdentityAtomic(transaction, object.Id(), identity.Username(), identity.PhoneNumber(), identity.PhoneNumberConfirmed(), identity.FirstName(), identity.LastName(), identity.DisplayName(), identity.Email(), identity.EmailConfirmed(), identity.Avatar(), identity.Banner(), identity.Summary(), identity.Token(), identity.MultiFactor(), identity.Hash(), identity.Salt(), identity.PublicKey(), identity.PrivateKey(), identity.Permission(), identity.Restriction(), identity.LastLogin(), identity.LoginCount(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	} else {
		if identity, err := dispatcher.conductor.IdentityManager().UpdateIdentity(object.Id(), identity.Username(), identity.PhoneNumber(), identity.PhoneNumberConfirmed(), identity.FirstName(), identity.LastName(), identity.DisplayName(), identity.Email(), identity.EmailConfirmed(), identity.Avatar(), identity.Banner(), identity.Summary(), identity.Token(), identity.MultiFactor(), identity.Hash(), identity.Salt(), identity.PublicKey(), identity.PrivateKey(), identity.Permission(), identity.Restriction(), identity.LastLogin(), identity.LoginCount(), dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	}
}

func (dispatcher *dispatcher) AddOrUpdateIdentityObject(object IObject, identity IIdentity) IIdentity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if identity, err := dispatcher.conductor.IdentityManager().AddOrUpdateIdentityObjectAtomic(transaction, object.Id(), identity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	} else {
		if identity, err := dispatcher.conductor.IdentityManager().AddOrUpdateIdentityObject(object.Id(), identity, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	}
}

func (dispatcher *dispatcher) RemoveIdentity(id int64) IIdentity {
	transaction := dispatcher.transaction
	if transaction != nil {
		if identity, err := dispatcher.conductor.IdentityManager().RemoveIdentityAtomic(transaction, id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	} else {
		if identity, err := dispatcher.conductor.IdentityManager().RemoveIdentity(id, dispatcher.identity); err != nil {
			panic(err.Error())
		} else {
			return identity
		}
	}
}
