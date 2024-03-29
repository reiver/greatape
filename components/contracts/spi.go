package contracts

import . "github.com/xeronith/diamante/contracts/security"

var SpiPassThroughFilter = func(ISpi) bool { return true }

type (
	Spis               []ISpi
	SpiIterator        func(ISpi)
	SpiCondition       func(ISpi) bool
	SpiFilterPredicate func(ISpi) bool
	SpiMapPredicate    func(ISpi) ISpi
	SpiCacheCallback   func()

	ISpi interface {
	}

	ISpiCollection interface {
		Count() int
		IsEmpty() bool
		IsNotEmpty() bool
		HasExactlyOneItem() bool
		HasAtLeastOneItem() bool
		First() ISpi
		Append(spi ISpi)
		ForEach(SpiIterator)
		Reverse() ISpiCollection
		Array() Spis
	}

	ISpiManager interface {
		ISystemComponent
		OnCacheChanged(SpiCacheCallback)
		Count() int
		Exists(id int64) bool
		ExistsWhich(condition SpiCondition) bool
		ListSpis(pageIndex uint32, pageSize uint32, criteria string, editor Identity) ISpiCollection
		GetSpi(id int64, editor Identity) (ISpi, error)
		AddSpi(editor Identity) (ISpi, error)
		AddSpiWithCustomId(id int64, editor Identity) (ISpi, error)
		AddSpiObject(spi ISpi, editor Identity) (ISpi, error)
		AddSpiAtomic(transaction ITransaction, editor Identity) (ISpi, error)
		AddSpiWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (ISpi, error)
		AddSpiObjectAtomic(transaction ITransaction, spi ISpi, editor Identity) (ISpi, error)
		Log(source string, editor Identity, payload string)
		UpdateSpi(id int64, editor Identity) (ISpi, error)
		UpdateSpiObject(id int64, spi ISpi, editor Identity) (ISpi, error)
		UpdateSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error)
		UpdateSpiObjectAtomic(transaction ITransaction, id int64, spi ISpi, editor Identity) (ISpi, error)
		AddOrUpdateSpiObject(id int64, spi ISpi, editor Identity) (ISpi, error)
		AddOrUpdateSpiObjectAtomic(transaction ITransaction, id int64, spi ISpi, editor Identity) (ISpi, error)
		RemoveSpi(id int64, editor Identity) (ISpi, error)
		RemoveSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error)
		Find(id int64) ISpi
		ForEach(iterator SpiIterator)
		Filter(predicate SpiFilterPredicate) ISpiCollection
		Map(predicate SpiMapPredicate) ISpiCollection
		Echo(document IDocument, editor Identity) (IEchoResult, error)
		GetServerConfiguration(editor Identity) (IGetServerConfigurationResult, error)
		CheckUsernameAvailability(username string, editor Identity) (ICheckUsernameAvailabilityResult, error)
		Signup(username string, email string, password string, editor Identity) (ISignupResult, error)
		ResendVerificationCode(email string, editor Identity) (IResendVerificationCodeResult, error)
		Verify(email string, token string, code string, editor Identity) (IVerifyResult, error)
		Login(email string, password string, editor Identity) (ILoginResult, error)
		GetProfileByUser(editor Identity) (IGetProfileByUserResult, error)
		UpdateProfileByUser(displayName string, avatar string, banner string, summary string, github string, editor Identity) (IUpdateProfileByUserResult, error)
		ChangePassword(currentPassword string, newPassword string, editor Identity) (IChangePasswordResult, error)
		ResetPassword(usernameOrEmail string, editor Identity) (IResetPasswordResult, error)
		Logout(editor Identity) (ILogoutResult, error)
		Webfinger(resource string, editor Identity) (IWebfingerResult, error)
		GetPackages(editor Identity) (IGetPackagesResult, error)
		GetActor(username string, editor Identity) (IGetActorResult, error)
		FollowActor(username string, account string, editor Identity) (IFollowActorResult, error)
		AuthorizeInteraction(uri string, editor Identity) (IAuthorizeInteractionResult, error)
		GetFollowers(username string, editor Identity) (IGetFollowersResult, error)
		GetFollowing(username string, editor Identity) (IGetFollowingResult, error)
		PostToOutbox(username string, body []byte, editor Identity) (IPostToOutboxResult, error)
		GetOutbox(username string, editor Identity) (IGetOutboxResult, error)
		PostToInbox(username string, body []byte, editor Identity) (IPostToInboxResult, error)
		GetInbox(username string, editor Identity) (IGetInboxResult, error)
	}

	IEchoResult interface {
		Document() IDocument
	}

	IGetServerConfigurationResult interface {
		Product() string
		Environment() string
		Fqdn() string
	}

	ICheckUsernameAvailabilityResult interface {
		IsAvailable() bool
	}

	ISignupResult interface {
		Token() string
		Code() string
	}

	IResendVerificationCodeResult interface {
		Code() string
	}

	IVerifyResult interface {
		Token() string
	}

	ILoginResult interface {
		Username() string
		Token() string
	}

	IGetProfileByUserResult interface {
		Username() string
		DisplayName() string
		Avatar() string
		Banner() string
		Summary() string
		Github() string
	}

	IUpdateProfileByUserResult interface {
		DisplayName() string
		Avatar() string
		Banner() string
		Summary() string
		Github() string
	}

	IChangePasswordResult interface {
	}

	IResetPasswordResult interface {
	}

	ILogoutResult interface {
	}

	IWebfingerResult interface {
		Aliases() []string
		Links() []IActivityPubLink
		Subject() string
	}

	IGetPackagesResult interface {
		Body() []byte
	}

	IGetActorResult interface {
		Context() []string
		Id() string
		Followers() string
		Following() string
		Inbox() string
		Outbox() string
		Name() string
		PreferredUsername() string
		Type() string
		Url() string
		Icon() IActivityPubMedia
		Image() IActivityPubMedia
		PublicKey() IActivityPubPublicKey
		Summary() string
		Published() string
	}

	IFollowActorResult interface {
	}

	IAuthorizeInteractionResult interface {
		Uri() string
		Success() bool
	}

	IGetFollowersResult interface {
		Context() string
		Id() string
		Type() string
		TotalItems() int32
		OrderedItems() []string
		First() string
	}

	IGetFollowingResult interface {
		Context() string
		Id() string
		Type() string
		TotalItems() int32
		OrderedItems() []string
		First() string
	}

	IPostToOutboxResult interface {
		Body() []byte
	}

	IGetOutboxResult interface {
		Context() string
		Id() string
		Type() string
		TotalItems() int32
		OrderedItems() []IActivityPubActivity
		First() string
	}

	IPostToInboxResult interface {
		Body() []byte
	}

	IGetInboxResult interface {
		Context() string
		Id() string
		Type() string
		TotalItems() int32
		OrderedItems() []IActivityPubActivity
		First() string
	}
)
