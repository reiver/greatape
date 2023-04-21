package model

type (
	IdentityEntities []IIdentityEntity

	IIdentityEntity interface {
		IEntity
		Username() string
		PhoneNumber() string
		PhoneNumberConfirmed() bool
		FirstName() string
		LastName() string
		DisplayName() string
		Email() string
		EmailConfirmed() bool
		Avatar() string
		Banner() string
		Summary() string
		Token() string
		MultiFactor() bool
		Hash() string
		Salt() string
		PublicKey() string
		PrivateKey() string
		Permission() uint64
		Restriction() uint32
		LastLogin() int64
		LoginCount() uint32
	}

	IIdentityPipeEntity interface {
		IIdentityEntity
		IPipeEntity
	}

	IIdentitiesRepository interface {
		IRepository
		Add(entity IIdentityEntity, editor int64) error
		AddAtomic(transaction IRepositoryTransaction, entity IIdentityEntity, editor int64) error
		FetchById(editor int64) (IIdentityEntity, error)
		Update(entity IIdentityEntity, editor int64) error
		UpdateAtomic(transaction IRepositoryTransaction, entity IIdentityEntity, editor int64) error
		Remove(entity IIdentityEntity, editor int64) error
		RemoveAtomic(transaction IRepositoryTransaction, entity IIdentityEntity, editor int64) error
		FetchAll() (IdentityEntities, error)
		UpdateUsername(id int64, value string, editor int64) error
		UpdateUsernameAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdatePhoneNumber(id int64, value string, editor int64) error
		UpdatePhoneNumberAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdatePhoneNumberConfirmed(id int64, value bool, editor int64) error
		UpdatePhoneNumberConfirmedAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error
		UpdateFirstName(id int64, value string, editor int64) error
		UpdateFirstNameAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateLastName(id int64, value string, editor int64) error
		UpdateLastNameAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateDisplayName(id int64, value string, editor int64) error
		UpdateDisplayNameAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateEmail(id int64, value string, editor int64) error
		UpdateEmailAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateEmailConfirmed(id int64, value bool, editor int64) error
		UpdateEmailConfirmedAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error
		UpdateAvatar(id int64, value string, editor int64) error
		UpdateAvatarAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateBanner(id int64, value string, editor int64) error
		UpdateBannerAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateSummary(id int64, value string, editor int64) error
		UpdateSummaryAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateToken(id int64, value string, editor int64) error
		UpdateTokenAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateMultiFactor(id int64, value bool, editor int64) error
		UpdateMultiFactorAtomic(transaction IRepositoryTransaction, id int64, value bool, editor int64) error
		UpdateHash(id int64, value string, editor int64) error
		UpdateHashAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdateSalt(id int64, value string, editor int64) error
		UpdateSaltAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdatePublicKey(id int64, value string, editor int64) error
		UpdatePublicKeyAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdatePrivateKey(id int64, value string, editor int64) error
		UpdatePrivateKeyAtomic(transaction IRepositoryTransaction, id int64, value string, editor int64) error
		UpdatePermission(id int64, value uint64, editor int64) error
		UpdatePermissionAtomic(transaction IRepositoryTransaction, id int64, value uint64, editor int64) error
		UpdateRestriction(id int64, value uint32, editor int64) error
		UpdateRestrictionAtomic(transaction IRepositoryTransaction, id int64, value uint32, editor int64) error
		UpdateLastLogin(id int64, value int64, editor int64) error
		UpdateLastLoginAtomic(transaction IRepositoryTransaction, id int64, value int64, editor int64) error
		UpdateLoginCount(id int64, value uint32, editor int64) error
		UpdateLoginCountAtomic(transaction IRepositoryTransaction, id int64, value uint32, editor int64) error
	}
)
