package core

import (
	"github.com/reiver/greatape/app/commands"
	"github.com/reiver/greatape/app/validators"
	. "github.com/reiver/greatape/components/constants"
	. "github.com/reiver/greatape/components/contracts"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/system"
	. "github.com/xeronith/diamante/system"
)

// noinspection GoSnakeCaseUsage
const SPI_MANAGER = "SpiManager"

type spiManager struct {
	systemComponent
	cache ICache
}

func newSpiManager(configuration IConfiguration, logger ILogger, dependencies ...ISystemComponent) ISpiManager {
	manager := &spiManager{
		systemComponent: newSystemComponent(configuration, logger),
		cache:           NewCache(),
	}

	if err := manager.ResolveDependencies(dependencies...); err != nil {
		return nil
	}

	return manager
}

func (manager *spiManager) Name() string {
	return SPI_MANAGER
}

func (manager *spiManager) ResolveDependencies(_ ...ISystemComponent) error {
	return nil
}

func (manager *spiManager) Load() error {
	return nil
}

func (manager *spiManager) Reload() error {
	return manager.Load()
}

func (manager *spiManager) OnCacheChanged(callback SpiCacheCallback) {
	manager.cache.OnChanged(callback)
}

func (manager *spiManager) Count() int {
	return manager.cache.Size()
}

func (manager *spiManager) Exists(id int64) bool {
	return manager.Find(id) != nil
}

func (manager *spiManager) ExistsWhich(condition SpiCondition) bool {
	var spis Spis
	manager.ForEach(func(spi ISpi) {
		if condition(spi) {
			spis = append(spis, spi)
		}
	})

	return len(spis) > 0
}

func (manager *spiManager) ListSpis(_ /* pageIndex */ uint32, _ /* pageSize */ uint32, _ /* criteria */ string, _ Identity) ISpiCollection {
	return manager.Filter(SpiPassThroughFilter)
}

func (manager *spiManager) GetSpi(id int64, _ Identity) (ISpi, error) {
	if spi := manager.Find(id); spi == nil {
		return nil, ERROR_SPI_NOT_FOUND
	} else {
		return spi, nil
	}
}

func (manager *spiManager) AddSpi(editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiWithCustomId(id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiObject(spi ISpi, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiAtomic(transaction ITransaction, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiWithCustomIdAtomic(id int64, transaction ITransaction, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddSpiObjectAtomic(transaction ITransaction, spi ISpi, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) Log(source string, editor Identity, payload string) {
}

func (manager *spiManager) UpdateSpi(id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) UpdateSpiObject(id int64, spi ISpi, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) UpdateSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) UpdateSpiObjectAtomic(transaction ITransaction, id int64, spi ISpi, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) AddOrUpdateSpiObject(id int64, spi ISpi, editor Identity) (ISpi, error) {
	if manager.Exists(id) {
		return manager.UpdateSpiObject(id, spi, editor)
	} else {
		return manager.AddSpiObject(spi, editor)
	}
}

func (manager *spiManager) AddOrUpdateSpiObjectAtomic(transaction ITransaction, id int64, spi ISpi, editor Identity) (ISpi, error) {
	if manager.Exists(id) {
		return manager.UpdateSpiObjectAtomic(transaction, id, spi, editor)
	} else {
		return manager.AddSpiObjectAtomic(transaction, spi, editor)
	}
}

func (manager *spiManager) RemoveSpi(id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) RemoveSpiAtomic(transaction ITransaction, id int64, editor Identity) (ISpi, error) {
	return nil, ERROR_NOT_IMPLEMENTED
}

func (manager *spiManager) Find(id int64) ISpi {
	if object, exists := manager.cache.Get(id); !exists {
		return nil
	} else {
		return object.(ISpi)
	}
}

func (manager *spiManager) ForEach(iterator SpiIterator) {
	manager.cache.ForEachValue(func(object ISystemObject) {
		iterator(object.(ISpi))
	})
}

func (manager *spiManager) Filter(predicate SpiFilterPredicate) ISpiCollection {
	spis := NewSpis()
	if predicate == nil {
		return spis
	}

	manager.ForEach(func(spi ISpi) {
		if predicate(spi) {
			spis.Append(spi)
		}
	})

	return spis
}

func (manager *spiManager) Map(predicate SpiMapPredicate) ISpiCollection {
	spis := NewSpis()
	if predicate == nil {
		return spis
	}

	manager.ForEach(func(spi ISpi) {
		spis.Append(predicate(spi))
	})

	return spis
}

//region IEchoResult Implementation

type echoResult struct {
	document IDocument
}

func NewEchoResult(document IDocument, _ interface{}) IEchoResult {
	return &echoResult{
		document: document,
	}
}

func (result echoResult) Document() IDocument {
	return result.document
}

//endregion

func (manager *spiManager) Echo(document IDocument, editor Identity) (result IEchoResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(ECHO_REQUEST)
	defer editor.Unlock(ECHO_REQUEST)

	if result, err = commands.Echo(NewDispatcher(Conductor, editor), document); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IGetServerConfigurationResult Implementation

type getServerConfigurationResult struct {
	product     string
	environment string
	fqdn        string
}

func NewGetServerConfigurationResult(product string, environment string, fqdn string, _ interface{}) IGetServerConfigurationResult {
	return &getServerConfigurationResult{
		product:     product,
		environment: environment,
		fqdn:        fqdn,
	}
}

func (result getServerConfigurationResult) Product() string {
	return result.product
}

func (result getServerConfigurationResult) Environment() string {
	return result.environment
}

func (result getServerConfigurationResult) Fqdn() string {
	return result.fqdn
}

//endregion

func (manager *spiManager) GetServerConfiguration(editor Identity) (result IGetServerConfigurationResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(GET_SERVER_CONFIGURATION_REQUEST)
	defer editor.Unlock(GET_SERVER_CONFIGURATION_REQUEST)

	if result, err = commands.GetServerConfiguration(NewDispatcher(Conductor, editor)); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region ICheckUsernameAvailabilityResult Implementation

type checkUsernameAvailabilityResult struct {
	isAvailable bool
}

func NewCheckUsernameAvailabilityResult(isAvailable bool, _ interface{}) ICheckUsernameAvailabilityResult {
	return &checkUsernameAvailabilityResult{
		isAvailable: isAvailable,
	}
}

func (result checkUsernameAvailabilityResult) IsAvailable() bool {
	return result.isAvailable
}

//endregion

func (manager *spiManager) CheckUsernameAvailability(username string, editor Identity) (result ICheckUsernameAvailabilityResult, err error) {
	if !validators.UsernameIsValid(username) {
		return nil, ERROR_INVALID_USERNAME_FOR_CHECK_USERNAME_AVAILABILITY
	}

	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(CHECK_USERNAME_AVAILABILITY_REQUEST)
	defer editor.Unlock(CHECK_USERNAME_AVAILABILITY_REQUEST)

	if result, err = commands.CheckUsernameAvailability(NewDispatcher(Conductor, editor), username); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region ISignupResult Implementation

type signupResult struct {
	token string
	code  string
}

func NewSignupResult(token string, code string, _ interface{}) ISignupResult {
	return &signupResult{
		token: token,
		code:  code,
	}
}

func (result signupResult) Token() string {
	return result.token
}

func (result signupResult) Code() string {
	return result.code
}

//endregion

func (manager *spiManager) Signup(username string, email string, password string, editor Identity) (result ISignupResult, err error) {
	if email != "" {
		if match, err := manager.Match(EMAIL, email); err != nil {
			return nil, err
		} else if !match {
			return nil, ERROR_INVALID_EMAIL_FOR_SIGNUP
		}
	}

	if !validators.UsernameIsValid(username) {
		return nil, ERROR_INVALID_USERNAME_FOR_SIGNUP
	}

	if !validators.PasswordIsValid(password) {
		return nil, ERROR_INVALID_PASSWORD_FOR_SIGNUP
	}

	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(SIGNUP_REQUEST)
	defer editor.Unlock(SIGNUP_REQUEST)

	if result, err = commands.Signup(NewDispatcher(Conductor, editor), username, email, password); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IResendVerificationCodeResult Implementation

type resendVerificationCodeResult struct {
	code string
}

func NewResendVerificationCodeResult(code string, _ interface{}) IResendVerificationCodeResult {
	return &resendVerificationCodeResult{
		code: code,
	}
}

func (result resendVerificationCodeResult) Code() string {
	return result.code
}

//endregion

func (manager *spiManager) ResendVerificationCode(email string, editor Identity) (result IResendVerificationCodeResult, err error) {
	if email != "" {
		if match, err := manager.Match(EMAIL, email); err != nil {
			return nil, err
		} else if !match {
			return nil, ERROR_INVALID_EMAIL_FOR_RESEND_VERIFICATION_CODE
		}
	}

	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(RESEND_VERIFICATION_CODE_REQUEST)
	defer editor.Unlock(RESEND_VERIFICATION_CODE_REQUEST)

	if result, err = commands.ResendVerificationCode(NewDispatcher(Conductor, editor), email); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IVerifyResult Implementation

type verifyResult struct {
	token string
}

func NewVerifyResult(token string, _ interface{}) IVerifyResult {
	return &verifyResult{
		token: token,
	}
}

func (result verifyResult) Token() string {
	return result.token
}

//endregion

func (manager *spiManager) Verify(email string, token string, code string, editor Identity) (result IVerifyResult, err error) {
	if email != "" {
		if match, err := manager.Match(EMAIL, email); err != nil {
			return nil, err
		} else if !match {
			return nil, ERROR_INVALID_EMAIL_FOR_VERIFY
		}
	}

	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(VERIFY_REQUEST)
	defer editor.Unlock(VERIFY_REQUEST)

	if result, err = commands.Verify(NewDispatcher(Conductor, editor), email, token, code); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region ILoginResult Implementation

type loginResult struct {
	username string
	token    string
}

func NewLoginResult(username string, token string, _ interface{}) ILoginResult {
	return &loginResult{
		username: username,
		token:    token,
	}
}

func (result loginResult) Username() string {
	return result.username
}

func (result loginResult) Token() string {
	return result.token
}

//endregion

func (manager *spiManager) Login(email string, password string, editor Identity) (result ILoginResult, err error) {
	if email != "" {
		if match, err := manager.Match(EMAIL, email); err != nil {
			return nil, err
		} else if !match {
			return nil, ERROR_INVALID_EMAIL_FOR_LOGIN
		}
	}

	if !validators.PasswordIsValid(password) {
		return nil, ERROR_INVALID_PASSWORD_FOR_LOGIN
	}

	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(LOGIN_REQUEST)
	defer editor.Unlock(LOGIN_REQUEST)

	if result, err = commands.Login(NewDispatcher(Conductor, editor), email, password); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IGetProfileByUserResult Implementation

type getProfileByUserResult struct {
	username    string
	displayName string
	avatar      string
	banner      string
	summary     string
	github      string
}

func NewGetProfileByUserResult(username string, displayName string, avatar string, banner string, summary string, github string, _ interface{}) IGetProfileByUserResult {
	return &getProfileByUserResult{
		username:    username,
		displayName: displayName,
		avatar:      avatar,
		banner:      banner,
		summary:     summary,
		github:      github,
	}
}

func (result getProfileByUserResult) Username() string {
	return result.username
}

func (result getProfileByUserResult) DisplayName() string {
	return result.displayName
}

func (result getProfileByUserResult) Avatar() string {
	return result.avatar
}

func (result getProfileByUserResult) Banner() string {
	return result.banner
}

func (result getProfileByUserResult) Summary() string {
	return result.summary
}

func (result getProfileByUserResult) Github() string {
	return result.github
}

//endregion

func (manager *spiManager) GetProfileByUser(editor Identity) (result IGetProfileByUserResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(GET_PROFILE_BY_USER_REQUEST)
	defer editor.Unlock(GET_PROFILE_BY_USER_REQUEST)

	if result, err = commands.GetProfileByUser(NewDispatcher(Conductor, editor)); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IUpdateProfileByUserResult Implementation

type updateProfileByUserResult struct {
	displayName string
	avatar      string
	banner      string
	summary     string
	github      string
}

func NewUpdateProfileByUserResult(displayName string, avatar string, banner string, summary string, github string, _ interface{}) IUpdateProfileByUserResult {
	return &updateProfileByUserResult{
		displayName: displayName,
		avatar:      avatar,
		banner:      banner,
		summary:     summary,
		github:      github,
	}
}

func (result updateProfileByUserResult) DisplayName() string {
	return result.displayName
}

func (result updateProfileByUserResult) Avatar() string {
	return result.avatar
}

func (result updateProfileByUserResult) Banner() string {
	return result.banner
}

func (result updateProfileByUserResult) Summary() string {
	return result.summary
}

func (result updateProfileByUserResult) Github() string {
	return result.github
}

//endregion

func (manager *spiManager) UpdateProfileByUser(displayName string, avatar string, banner string, summary string, github string, editor Identity) (result IUpdateProfileByUserResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(UPDATE_PROFILE_BY_USER_REQUEST)
	defer editor.Unlock(UPDATE_PROFILE_BY_USER_REQUEST)

	if result, err = commands.UpdateProfileByUser(NewDispatcher(Conductor, editor), displayName, avatar, banner, summary, github); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IChangePasswordResult Implementation

type changePasswordResult struct {
}

func NewChangePasswordResult(_ interface{}) IChangePasswordResult {
	return &changePasswordResult{}
}

//endregion

func (manager *spiManager) ChangePassword(currentPassword string, newPassword string, editor Identity) (result IChangePasswordResult, err error) {
	if !validators.PasswordIsValid(currentPassword) {
		return nil, ERROR_INVALID_CURRENT_PASSWORD_FOR_CHANGE_PASSWORD
	}

	if !validators.PasswordIsValid(newPassword) {
		return nil, ERROR_INVALID_NEW_PASSWORD_FOR_CHANGE_PASSWORD
	}

	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(CHANGE_PASSWORD_REQUEST)
	defer editor.Unlock(CHANGE_PASSWORD_REQUEST)

	if result, err = commands.ChangePassword(NewDispatcher(Conductor, editor), currentPassword, newPassword); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IResetPasswordResult Implementation

type resetPasswordResult struct {
}

func NewResetPasswordResult(_ interface{}) IResetPasswordResult {
	return &resetPasswordResult{}
}

//endregion

func (manager *spiManager) ResetPassword(usernameOrEmail string, editor Identity) (result IResetPasswordResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(RESET_PASSWORD_REQUEST)
	defer editor.Unlock(RESET_PASSWORD_REQUEST)

	if result, err = commands.ResetPassword(NewDispatcher(Conductor, editor), usernameOrEmail); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region ILogoutResult Implementation

type logoutResult struct {
}

func NewLogoutResult(_ interface{}) ILogoutResult {
	return &logoutResult{}
}

//endregion

func (manager *spiManager) Logout(editor Identity) (result ILogoutResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(LOGOUT_REQUEST)
	defer editor.Unlock(LOGOUT_REQUEST)

	if result, err = commands.Logout(NewDispatcher(Conductor, editor)); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IWebfingerResult Implementation

type webfingerResult struct {
	aliases []string
	links   []IActivityPubLink
	subject string
}

func NewWebfingerResult(aliases []string, links []IActivityPubLink, subject string, _ interface{}) IWebfingerResult {
	return &webfingerResult{
		aliases: aliases,
		links:   links,
		subject: subject,
	}
}

func (result webfingerResult) Aliases() []string {
	return result.aliases
}

func (result webfingerResult) Links() []IActivityPubLink {
	return result.links
}

func (result webfingerResult) Subject() string {
	return result.subject
}

//endregion

func (manager *spiManager) Webfinger(resource string, editor Identity) (result IWebfingerResult, err error) {
	if match, err := manager.Match(WEBFINGER, resource); err != nil {
		return nil, err
	} else if !match {
		return nil, ERROR_INVALID_RESOURCE_FOR_WEBFINGER
	}

	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(WEBFINGER_REQUEST)
	defer editor.Unlock(WEBFINGER_REQUEST)

	if result, err = commands.Webfinger(NewDispatcher(Conductor, editor), resource); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IGetPackagesResult Implementation

type getPackagesResult struct {
	body []byte
}

func NewGetPackagesResult(body []byte, _ interface{}) IGetPackagesResult {
	return &getPackagesResult{
		body: body,
	}
}

func (result getPackagesResult) Body() []byte {
	return result.body
}

//endregion

func (manager *spiManager) GetPackages(editor Identity) (result IGetPackagesResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(GET_PACKAGES_REQUEST)
	defer editor.Unlock(GET_PACKAGES_REQUEST)

	if result, err = commands.GetPackages(NewDispatcher(Conductor, editor)); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IGetActorResult Implementation

type getActorResult struct {
	context           []string
	id                string
	followers         string
	following         string
	inbox             string
	outbox            string
	name              string
	preferredUsername string
	type_             string
	url               string
	icon              IActivityPubMedia
	image             IActivityPubMedia
	publicKey         IActivityPubPublicKey
	summary           string
	published         string
}

func NewGetActorResult(context []string, id string, followers string, following string, inbox string, outbox string, name string, preferredUsername string, type_ string, url string, icon IActivityPubMedia, image IActivityPubMedia, publicKey IActivityPubPublicKey, summary string, published string, _ interface{}) IGetActorResult {
	return &getActorResult{
		context:           context,
		id:                id,
		followers:         followers,
		following:         following,
		inbox:             inbox,
		outbox:            outbox,
		name:              name,
		preferredUsername: preferredUsername,
		type_:             type_,
		url:               url,
		icon:              icon,
		image:             image,
		publicKey:         publicKey,
		summary:           summary,
		published:         published,
	}
}

func (result getActorResult) Context() []string {
	return result.context
}

func (result getActorResult) Id() string {
	return result.id
}

func (result getActorResult) Followers() string {
	return result.followers
}

func (result getActorResult) Following() string {
	return result.following
}

func (result getActorResult) Inbox() string {
	return result.inbox
}

func (result getActorResult) Outbox() string {
	return result.outbox
}

func (result getActorResult) Name() string {
	return result.name
}

func (result getActorResult) PreferredUsername() string {
	return result.preferredUsername
}

func (result getActorResult) Type() string {
	return result.type_
}

func (result getActorResult) Url() string {
	return result.url
}

func (result getActorResult) Icon() IActivityPubMedia {
	return result.icon
}

func (result getActorResult) Image() IActivityPubMedia {
	return result.image
}

func (result getActorResult) PublicKey() IActivityPubPublicKey {
	return result.publicKey
}

func (result getActorResult) Summary() string {
	return result.summary
}

func (result getActorResult) Published() string {
	return result.published
}

//endregion

func (manager *spiManager) GetActor(username string, editor Identity) (result IGetActorResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(GET_ACTOR_REQUEST)
	defer editor.Unlock(GET_ACTOR_REQUEST)

	if result, err = commands.GetActor(NewDispatcher(Conductor, editor), username); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IFollowActorResult Implementation

type followActorResult struct {
	url string
}

func NewFollowActorResult(url string, _ interface{}) IFollowActorResult {
	return &followActorResult{
		url: url,
	}
}

func (result followActorResult) Url() string {
	return result.url
}

//endregion

func (manager *spiManager) FollowActor(username string, acct string, editor Identity) (result IFollowActorResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(FOLLOW_ACTOR_REQUEST)
	defer editor.Unlock(FOLLOW_ACTOR_REQUEST)

	if result, err = commands.FollowActor(NewDispatcher(Conductor, editor), username, acct); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IAuthorizeInteractionResult Implementation

type authorizeInteractionResult struct {
	uri     string
	success bool
}

func NewAuthorizeInteractionResult(uri string, success bool, _ interface{}) IAuthorizeInteractionResult {
	return &authorizeInteractionResult{
		uri:     uri,
		success: success,
	}
}

func (result authorizeInteractionResult) Uri() string {
	return result.uri
}

func (result authorizeInteractionResult) Success() bool {
	return result.success
}

//endregion

func (manager *spiManager) AuthorizeInteraction(uri string, editor Identity) (result IAuthorizeInteractionResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(AUTHORIZE_INTERACTION_REQUEST)
	defer editor.Unlock(AUTHORIZE_INTERACTION_REQUEST)

	if result, err = commands.AuthorizeInteraction(NewDispatcher(Conductor, editor), uri); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IGetFollowersResult Implementation

type getFollowersResult struct {
	context      string
	id           string
	type_        string
	totalItems   int32
	orderedItems []string
	first        string
}

func NewGetFollowersResult(context string, id string, type_ string, totalItems int32, orderedItems []string, first string, _ interface{}) IGetFollowersResult {
	return &getFollowersResult{
		context:      context,
		id:           id,
		type_:        type_,
		totalItems:   totalItems,
		orderedItems: orderedItems,
		first:        first,
	}
}

func (result getFollowersResult) Context() string {
	return result.context
}

func (result getFollowersResult) Id() string {
	return result.id
}

func (result getFollowersResult) Type() string {
	return result.type_
}

func (result getFollowersResult) TotalItems() int32 {
	return result.totalItems
}

func (result getFollowersResult) OrderedItems() []string {
	return result.orderedItems
}

func (result getFollowersResult) First() string {
	return result.first
}

//endregion

func (manager *spiManager) GetFollowers(username string, editor Identity) (result IGetFollowersResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(GET_FOLLOWERS_REQUEST)
	defer editor.Unlock(GET_FOLLOWERS_REQUEST)

	if result, err = commands.GetFollowers(NewDispatcher(Conductor, editor), username); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IGetFollowingResult Implementation

type getFollowingResult struct {
	context      string
	id           string
	type_        string
	totalItems   int32
	orderedItems []string
	first        string
}

func NewGetFollowingResult(context string, id string, type_ string, totalItems int32, orderedItems []string, first string, _ interface{}) IGetFollowingResult {
	return &getFollowingResult{
		context:      context,
		id:           id,
		type_:        type_,
		totalItems:   totalItems,
		orderedItems: orderedItems,
		first:        first,
	}
}

func (result getFollowingResult) Context() string {
	return result.context
}

func (result getFollowingResult) Id() string {
	return result.id
}

func (result getFollowingResult) Type() string {
	return result.type_
}

func (result getFollowingResult) TotalItems() int32 {
	return result.totalItems
}

func (result getFollowingResult) OrderedItems() []string {
	return result.orderedItems
}

func (result getFollowingResult) First() string {
	return result.first
}

//endregion

func (manager *spiManager) GetFollowing(username string, editor Identity) (result IGetFollowingResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(GET_FOLLOWING_REQUEST)
	defer editor.Unlock(GET_FOLLOWING_REQUEST)

	if result, err = commands.GetFollowing(NewDispatcher(Conductor, editor), username); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IPostToOutboxResult Implementation

type postToOutboxResult struct {
	body []byte
}

func NewPostToOutboxResult(body []byte, _ interface{}) IPostToOutboxResult {
	return &postToOutboxResult{
		body: body,
	}
}

func (result postToOutboxResult) Body() []byte {
	return result.body
}

//endregion

func (manager *spiManager) PostToOutbox(username string, body []byte, editor Identity) (result IPostToOutboxResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(POST_TO_OUTBOX_REQUEST)
	defer editor.Unlock(POST_TO_OUTBOX_REQUEST)

	if result, err = commands.PostToOutbox(NewDispatcher(Conductor, editor), username, body); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IGetOutboxResult Implementation

type getOutboxResult struct {
	context      string
	id           string
	type_        string
	totalItems   int32
	orderedItems []IActivityPubActivity
	first        string
}

func NewGetOutboxResult(context string, id string, type_ string, totalItems int32, orderedItems []IActivityPubActivity, first string, _ interface{}) IGetOutboxResult {
	return &getOutboxResult{
		context:      context,
		id:           id,
		type_:        type_,
		totalItems:   totalItems,
		orderedItems: orderedItems,
		first:        first,
	}
}

func (result getOutboxResult) Context() string {
	return result.context
}

func (result getOutboxResult) Id() string {
	return result.id
}

func (result getOutboxResult) Type() string {
	return result.type_
}

func (result getOutboxResult) TotalItems() int32 {
	return result.totalItems
}

func (result getOutboxResult) OrderedItems() []IActivityPubActivity {
	return result.orderedItems
}

func (result getOutboxResult) First() string {
	return result.first
}

//endregion

func (manager *spiManager) GetOutbox(username string, editor Identity) (result IGetOutboxResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(GET_OUTBOX_REQUEST)
	defer editor.Unlock(GET_OUTBOX_REQUEST)

	if result, err = commands.GetOutbox(NewDispatcher(Conductor, editor), username); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IPostToInboxResult Implementation

type postToInboxResult struct {
	body []byte
}

func NewPostToInboxResult(body []byte, _ interface{}) IPostToInboxResult {
	return &postToInboxResult{
		body: body,
	}
}

func (result postToInboxResult) Body() []byte {
	return result.body
}

//endregion

func (manager *spiManager) PostToInbox(username string, body []byte, editor Identity) (result IPostToInboxResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(POST_TO_INBOX_REQUEST)
	defer editor.Unlock(POST_TO_INBOX_REQUEST)

	if result, err = commands.PostToInbox(NewDispatcher(Conductor, editor), username, body); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

//region IGetInboxResult Implementation

type getInboxResult struct {
	context      string
	id           string
	type_        string
	totalItems   int32
	orderedItems []IActivityPubActivity
	first        string
}

func NewGetInboxResult(context string, id string, type_ string, totalItems int32, orderedItems []IActivityPubActivity, first string, _ interface{}) IGetInboxResult {
	return &getInboxResult{
		context:      context,
		id:           id,
		type_:        type_,
		totalItems:   totalItems,
		orderedItems: orderedItems,
		first:        first,
	}
}

func (result getInboxResult) Context() string {
	return result.context
}

func (result getInboxResult) Id() string {
	return result.id
}

func (result getInboxResult) Type() string {
	return result.type_
}

func (result getInboxResult) TotalItems() int32 {
	return result.totalItems
}

func (result getInboxResult) OrderedItems() []IActivityPubActivity {
	return result.orderedItems
}

func (result getInboxResult) First() string {
	return result.first
}

//endregion

func (manager *spiManager) GetInbox(username string, editor Identity) (result IGetInboxResult, err error) {
	defer func() {
		if reason := recover(); reason != nil {
			err = manager.Error(reason)
		}
	}()

	editor.Lock(GET_INBOX_REQUEST)
	defer editor.Unlock(GET_INBOX_REQUEST)

	if result, err = commands.GetInbox(NewDispatcher(Conductor, editor), username); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
