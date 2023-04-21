package contracts

import (
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/system"
)

var IdentityPassThroughFilter = func(IIdentity) bool { return true }

type (
	Identities              []IIdentity
	IdentityIterator        func(IIdentity)
	IdentityCondition       func(IIdentity) bool
	IdentityFilterPredicate func(IIdentity) bool
	IdentityMapPredicate    func(IIdentity) IIdentity
	IdentityCacheCallback   func()

	IIdentity interface {
		IObject
		// Username returns 'Username' of this 'Identity' instance.
		Username() string
		// UpdateUsername directly updates 'Username' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateUsername(username string, editor Identity)
		// UpdateUsernameAtomic updates 'Username' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateUsernameAtomic(transaction ITransaction, username string, editor Identity)
		// PhoneNumber returns 'PhoneNumber' of this 'Identity' instance.
		PhoneNumber() string
		// UpdatePhoneNumber directly updates 'PhoneNumber' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdatePhoneNumber(phoneNumber string, editor Identity)
		// UpdatePhoneNumberAtomic updates 'PhoneNumber' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdatePhoneNumberAtomic(transaction ITransaction, phoneNumber string, editor Identity)
		// PhoneNumberConfirmed returns 'PhoneNumberConfirmed' of this 'Identity' instance.
		PhoneNumberConfirmed() bool
		// UpdatePhoneNumberConfirmed directly updates 'PhoneNumberConfirmed' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdatePhoneNumberConfirmed(phoneNumberConfirmed bool, editor Identity)
		// UpdatePhoneNumberConfirmedAtomic updates 'PhoneNumberConfirmed' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdatePhoneNumberConfirmedAtomic(transaction ITransaction, phoneNumberConfirmed bool, editor Identity)
		// FirstName returns 'FirstName' of this 'Identity' instance.
		FirstName() string
		// UpdateFirstName directly updates 'FirstName' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateFirstName(firstName string, editor Identity)
		// UpdateFirstNameAtomic updates 'FirstName' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateFirstNameAtomic(transaction ITransaction, firstName string, editor Identity)
		// LastName returns 'LastName' of this 'Identity' instance.
		LastName() string
		// UpdateLastName directly updates 'LastName' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateLastName(lastName string, editor Identity)
		// UpdateLastNameAtomic updates 'LastName' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateLastNameAtomic(transaction ITransaction, lastName string, editor Identity)
		// DisplayName returns 'DisplayName' of this 'Identity' instance.
		DisplayName() string
		// UpdateDisplayName directly updates 'DisplayName' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateDisplayName(displayName string, editor Identity)
		// UpdateDisplayNameAtomic updates 'DisplayName' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateDisplayNameAtomic(transaction ITransaction, displayName string, editor Identity)
		// Email returns 'Email' of this 'Identity' instance.
		Email() string
		// UpdateEmail directly updates 'Email' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateEmail(email string, editor Identity)
		// UpdateEmailAtomic updates 'Email' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateEmailAtomic(transaction ITransaction, email string, editor Identity)
		// EmailConfirmed returns 'EmailConfirmed' of this 'Identity' instance.
		EmailConfirmed() bool
		// UpdateEmailConfirmed directly updates 'EmailConfirmed' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateEmailConfirmed(emailConfirmed bool, editor Identity)
		// UpdateEmailConfirmedAtomic updates 'EmailConfirmed' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateEmailConfirmedAtomic(transaction ITransaction, emailConfirmed bool, editor Identity)
		// Avatar returns 'Avatar' of this 'Identity' instance.
		Avatar() string
		// UpdateAvatar directly updates 'Avatar' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateAvatar(avatar string, editor Identity)
		// UpdateAvatarAtomic updates 'Avatar' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateAvatarAtomic(transaction ITransaction, avatar string, editor Identity)
		// Banner returns 'Banner' of this 'Identity' instance.
		Banner() string
		// UpdateBanner directly updates 'Banner' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateBanner(banner string, editor Identity)
		// UpdateBannerAtomic updates 'Banner' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateBannerAtomic(transaction ITransaction, banner string, editor Identity)
		// Summary returns 'Summary' of this 'Identity' instance.
		Summary() string
		// UpdateSummary directly updates 'Summary' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateSummary(summary string, editor Identity)
		// UpdateSummaryAtomic updates 'Summary' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateSummaryAtomic(transaction ITransaction, summary string, editor Identity)
		// Token returns 'Token' of this 'Identity' instance.
		Token() string
		// UpdateToken directly updates 'Token' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateToken(token string, editor Identity)
		// UpdateTokenAtomic updates 'Token' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateTokenAtomic(transaction ITransaction, token string, editor Identity)
		// MultiFactor returns 'MultiFactor' of this 'Identity' instance.
		MultiFactor() bool
		// UpdateMultiFactor directly updates 'MultiFactor' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateMultiFactor(multiFactor bool, editor Identity)
		// UpdateMultiFactorAtomic updates 'MultiFactor' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateMultiFactorAtomic(transaction ITransaction, multiFactor bool, editor Identity)
		// Hash returns 'Hash' of this 'Identity' instance.
		Hash() string
		// UpdateHash directly updates 'Hash' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateHash(hash string, editor Identity)
		// UpdateHashAtomic updates 'Hash' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateHashAtomic(transaction ITransaction, hash string, editor Identity)
		// Salt returns 'Salt' of this 'Identity' instance.
		Salt() string
		// UpdateSalt directly updates 'Salt' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateSalt(salt string, editor Identity)
		// UpdateSaltAtomic updates 'Salt' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateSaltAtomic(transaction ITransaction, salt string, editor Identity)
		// PublicKey returns 'PublicKey' of this 'Identity' instance.
		PublicKey() string
		// UpdatePublicKey directly updates 'PublicKey' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdatePublicKey(publicKey string, editor Identity)
		// UpdatePublicKeyAtomic updates 'PublicKey' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdatePublicKeyAtomic(transaction ITransaction, publicKey string, editor Identity)
		// PrivateKey returns 'PrivateKey' of this 'Identity' instance.
		PrivateKey() string
		// UpdatePrivateKey directly updates 'PrivateKey' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdatePrivateKey(privateKey string, editor Identity)
		// UpdatePrivateKeyAtomic updates 'PrivateKey' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdatePrivateKeyAtomic(transaction ITransaction, privateKey string, editor Identity)
		// Permission returns 'Permission' of this 'Identity' instance.
		Permission() uint64
		// UpdatePermission directly updates 'Permission' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdatePermission(permission uint64, editor Identity)
		// UpdatePermissionAtomic updates 'Permission' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdatePermissionAtomic(transaction ITransaction, permission uint64, editor Identity)
		// Restriction returns 'Restriction' of this 'Identity' instance.
		Restriction() uint32
		// UpdateRestriction directly updates 'Restriction' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateRestriction(restriction uint32, editor Identity)
		// UpdateRestrictionAtomic updates 'Restriction' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateRestrictionAtomic(transaction ITransaction, restriction uint32, editor Identity)
		// LastLogin returns 'LastLogin' of this 'Identity' instance.
		LastLogin() int64
		// UpdateLastLogin directly updates 'LastLogin' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateLastLogin(lastLogin int64, editor Identity)
		// UpdateLastLoginAtomic updates 'LastLogin' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateLastLoginAtomic(transaction ITransaction, lastLogin int64, editor Identity)
		// LoginCount returns 'LoginCount' of this 'Identity' instance.
		LoginCount() uint32
		// UpdateLoginCount directly updates 'LoginCount' into persistent data store and
		// refreshes the in-memory cache after successful update.
		UpdateLoginCount(loginCount uint32, editor Identity)
		// UpdateLoginCountAtomic updates 'LoginCount' into persistent data store through a transaction and
		// refreshes the in-memory cache after successful commit.
		UpdateLoginCountAtomic(transaction ITransaction, loginCount uint32, editor Identity)
		// RemoteAddress returns 'RemoteAddress' of this 'Identity' instance.
		RemoteAddress() string
		// SetRemoteAddress sets 'RemoteAddress' in-memory value of this 'Identity' instance.
		// This doesn't affect the persistent data store.
		SetRemoteAddress(remoteAddress string)
		// UserAgent returns 'UserAgent' of this 'Identity' instance.
		UserAgent() string
		// SetUserAgent sets 'UserAgent' in-memory value of this 'Identity' instance.
		// This doesn't affect the persistent data store.
		SetUserAgent(userAgent string)
		Role() Role
		IsInRole(role Role) bool
		IsRestricted() bool
		IsNotRestricted() bool
		Payload() Pointer
		SetToken(token string)
		SetSystemCallHandler(func(Identity, []string) error)
		SystemCall(Identity, []string) error
	}

	IIdentityCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() IIdentity
		Append(identity IIdentity)
		ForEach(IdentityIterator)
		Array() Identities
	}

	IIdentityManager interface {
		ISystemComponent
		ISecurityHandler
		OnCacheChanged(IdentityCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition IdentityCondition) bool
		ListIdentities(pageIndex uint32, pageSize uint32, criteria string, editor Identity) IIdentityCollection
		GetIdentity(id int64, editor Identity) (IIdentity, error)
		AddIdentity(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		AddIdentityWithCustomId(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		AddIdentityObject(identity IIdentity, editor Identity) (IIdentity, error)
		AddIdentityAtomic(transaction ITransaction, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		AddIdentityWithCustomIdAtomic(id int64, transaction ITransaction, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		AddIdentityObjectAtomic(transaction ITransaction, identity IIdentity, editor Identity) (IIdentity, error)
		Log(username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, source string, editor Identity, payload string)
		UpdateIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		UpdateIdentityObject(id int64, identity IIdentity, editor Identity) (IIdentity, error)
		UpdateIdentityAtomic(transaction ITransaction, id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32, editor Identity) (IIdentity, error)
		UpdateIdentityObjectAtomic(transaction ITransaction, id int64, identity IIdentity, editor Identity) (IIdentity, error)
		AddOrUpdateIdentityObject(id int64, identity IIdentity, editor Identity) (IIdentity, error)
		AddOrUpdateIdentityObjectAtomic(transaction ITransaction, id int64, identity IIdentity, editor Identity) (IIdentity, error)
		RemoveIdentity(id int64, editor Identity) (IIdentity, error)
		RemoveIdentityAtomic(transaction ITransaction, id int64, editor Identity) (IIdentity, error)
		Find(id int64) IIdentity
		ForEach(iterator IdentityIterator)
		Filter(predicate IdentityFilterPredicate) IIdentityCollection
		Map(predicate IdentityMapPredicate) IIdentityCollection
	}
)
