package core

import (
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"

	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/security"
	. "github.com/xeronith/diamante/contracts/settings"
	"github.com/xeronith/diamante/utility/search"
	. "rail.town/infrastructure/components/contracts"
)

//region IDispatcher Implementation

type dispatcher struct {
	cache       IDispatcherCache
	conductor   IConductor
	identity    Identity
	transaction ITransaction
}

func NewDispatcher(conductor IConductor, identity Identity) IDispatcher {
	return &dispatcher{
		cache:     newDispatcherCache(conductor, identity),
		conductor: conductor,
		identity:  identity,
	}
}

func (dispatcher *dispatcher) Logger() ILogger {
	return dispatcher.conductor.Logger()
}

func (dispatcher *dispatcher) Config() IConfiguration {
	return dispatcher.conductor.Configuration()
}

func (dispatcher *dispatcher) FQDN() string {
	config := dispatcher.conductor.Configuration().GetServerConfiguration()
	return config.GetFQDN()
}

func (dispatcher *dispatcher) PublicUrl() string {
	config := dispatcher.conductor.Configuration().GetServerConfiguration()
	return fmt.Sprintf("%s://%s", config.GetProtocol(), config.GetFQDN())
}

func (dispatcher *dispatcher) Accelerator() IDispatcherCache {
	return dispatcher.cache
}

func (dispatcher *dispatcher) Atomic(action SystemAction) {
	if err := dispatcher.conductor.Atomic(func(transaction ITransaction) error {
		defer func() {
			dispatcher.transaction = nil
		}()

		dispatcher.transaction = transaction
		return action()
	}); err != nil {
		panic(err.Error())
	}
}

func (dispatcher *dispatcher) Schedule(id int64, spec string, callback func(IDispatcher, string)) error {
	return dispatcher.conductor.Schedule(spec, func() {
		if !dispatcher.SystemScheduleExists(id) {
			dispatcher.AddSystemScheduleWithCustomId(id, true, "")
		}

		systemSchedule := dispatcher.GetSystemSchedule(id)
		if !systemSchedule.Enabled() {
			return
		}

		defer func() {
			if reason := recover(); reason != nil {
				dispatcher.conductor.Logger().Panic(fmt.Sprintf("JOB_SCHEDULER: %s", reason))
			}
		}()

		callback(dispatcher, systemSchedule.Config())
	})
}

func (dispatcher *dispatcher) Transaction() ITransaction {
	return dispatcher.transaction
}

func (dispatcher *dispatcher) IdentityManager() IIdentityManager {
	return dispatcher.conductor.IdentityManager()
}

func (dispatcher *dispatcher) Identity() Identity {
	return dispatcher.identity
}

func (dispatcher *dispatcher) CurrentUser() IUser {
	return dispatcher.GetUser(dispatcher.identity.Id())
}

func (dispatcher *dispatcher) SignOut() error {
	return dispatcher.conductor.IdentityManager().SignOut(dispatcher.identity)
}

func (dispatcher *dispatcher) NewDocument(id int64, content string) (IDocument, error) {
	return NewDocument(id, content)
}

func (dispatcher *dispatcher) NewSystemSchedule(id int64, enabled bool, config string) (ISystemSchedule, error) {
	return NewSystemSchedule(id, enabled, config)
}

func (dispatcher *dispatcher) NewIdentity(id int64, username string, phoneNumber string, phoneNumberConfirmed bool, firstName string, lastName string, displayName string, email string, emailConfirmed bool, avatar string, banner string, summary string, token string, multiFactor bool, hash string, salt string, publicKey string, privateKey string, permission uint64, restriction uint32, lastLogin int64, loginCount uint32) (IIdentity, error) {
	return NewIdentity(id, username, phoneNumber, phoneNumberConfirmed, firstName, lastName, displayName, email, emailConfirmed, avatar, banner, summary, token, multiFactor, hash, salt, publicKey, privateKey, permission, restriction, lastLogin, loginCount)
}

func (dispatcher *dispatcher) NewAccessControl(id int64, key uint64, value uint64) (IAccessControl, error) {
	return NewAccessControl(id, key, value)
}

func (dispatcher *dispatcher) NewRemoteActivity(id int64, entryPoint string, duration int64, successful bool, errorMessage string, remoteAddress string, userAgent string, eventType uint32, timestamp int64) (IRemoteActivity, error) {
	return NewRemoteActivity(id, entryPoint, duration, successful, errorMessage, remoteAddress, userAgent, eventType, timestamp)
}

func (dispatcher *dispatcher) NewCategoryType(id int64, description string) (ICategoryType, error) {
	return NewCategoryType(id, description)
}

func (dispatcher *dispatcher) NewCategory(id int64, categoryTypeId int64, categoryId int64, title string, description string) (ICategory, error) {
	return NewCategory(id, categoryTypeId, categoryId, title, description)
}

func (dispatcher *dispatcher) NewUser(id int64, github string) (IUser, error) {
	return NewUser(id, github)
}

func (dispatcher *dispatcher) NewActivityPubObject() (IActivityPubObject, error) {
	return NewActivityPubObject()
}

func (dispatcher *dispatcher) NewActivityPubActivity() (IActivityPubActivity, error) {
	return NewActivityPubActivity()
}

func (dispatcher *dispatcher) NewSpi() (ISpi, error) {
	return NewSpi()
}

func (dispatcher *dispatcher) Assert(condition bool) IAssertionResult {
	return &assertionResult{
		condition: condition,
	}
}

func (dispatcher *dispatcher) AssertNoError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (dispatcher *dispatcher) Ensure(errors ...error) {
	for _, err := range errors {
		if err != nil {
			panic(err.Error())
		}
	}
}

func (dispatcher *dispatcher) AssertNull(x interface{}) IAssertionResult {
	return &assertionResult{
		condition: x == nil,
	}
}

func (dispatcher *dispatcher) AssertNotNull(x interface{}) IAssertionResult {
	return &assertionResult{
		condition: x != nil,
	}
}

func (dispatcher *dispatcher) AssertEmpty(x string) IAssertionResult {
	return &assertionResult{
		condition: strings.TrimSpace(x) == "",
	}
}

func (dispatcher *dispatcher) AssertNotEmpty(x string) IAssertionResult {
	return &assertionResult{
		condition: strings.TrimSpace(x) != "",
	}
}

func (dispatcher *dispatcher) Format(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func (dispatcher *dispatcher) Sort(slice interface{}, less func(a, b int) bool) {
	sort.Slice(slice, less)
}

func (dispatcher *dispatcher) Search(input, criteria string) bool {
	return search.MatchAny(input, criteria)
}

func (dispatcher *dispatcher) Email(destination string, format string, args ...interface{}) {
	dispatcher.conductor.IdentityManager().Email(destination, format, args...)
}

func (dispatcher *dispatcher) SMS(destination string, format string, args ...interface{}) {
	dispatcher.conductor.IdentityManager().SMS(destination, format, args...)
}

func (dispatcher *dispatcher) GenerateTrackingNumber() uint32 {
	rand.Seed(time.Now().UnixNano())
	return 100000 + uint32(rand.Intn(899999))
}

func (dispatcher *dispatcher) GenerateUUID() string {
	return dispatcher.conductor.IdentityManager().GenerateUUID()
}

func (dispatcher *dispatcher) GenerateSalt() string {
	return dispatcher.conductor.IdentityManager().GenerateSalt()
}

func (dispatcher *dispatcher) GenerateHash(value string, salt string) string {
	return dispatcher.conductor.IdentityManager().GenerateHash(value, salt)
}

func (dispatcher *dispatcher) GenerateJwtToken() string {
	return dispatcher.conductor.IdentityManager().GenerateJwtToken()
}

func (dispatcher *dispatcher) VerifyJwtToken(token string) error {
	return dispatcher.conductor.IdentityManager().VerifyJwtToken(token)
}

func (dispatcher *dispatcher) GenerateCode() string {
	return dispatcher.conductor.IdentityManager().GenerateCode()
}

func (dispatcher *dispatcher) GenerateRSAKeyPair() (string, string, error) {
	return dispatcher.conductor.IdentityManager().GenerateRSAKeyPair()
}

func (dispatcher *dispatcher) UnixNano() int64 {
	return time.Now().UnixNano()
}

func (dispatcher *dispatcher) Trim(input string) string {
	return strings.TrimSpace(input)
}

func (dispatcher *dispatcher) Contains(input, substr string) bool {
	return strings.Contains(input, substr)
}

func (dispatcher *dispatcher) IsEmpty(input string) bool {
	return strings.TrimSpace(input) == ""
}

func (dispatcher *dispatcher) IsNotEmpty(input string) bool {
	return strings.TrimSpace(input) != ""
}

func (dispatcher *dispatcher) IsSet(id int64) bool {
	return id != 0
}

func (dispatcher *dispatcher) Join(elements []string, separator string) string {
	return strings.Join(elements, separator)
}

func (dispatcher *dispatcher) IsTestEnvironment() bool {
	return dispatcher.conductor.IdentityManager().IsTestEnvironment()
}

func (dispatcher *dispatcher) IsDevelopmentEnvironment() bool {
	return dispatcher.conductor.IdentityManager().IsDevelopmentEnvironment()
}

func (dispatcher *dispatcher) IsStagingEnvironment() bool {
	return dispatcher.conductor.IdentityManager().IsStagingEnvironment()
}

func (dispatcher *dispatcher) IsProductionEnvironment() bool {
	return dispatcher.conductor.IdentityManager().IsProductionEnvironment()
}

func (dispatcher *dispatcher) GetActivityStream(url string, data []byte, output interface{}) error {
	return dispatcher.conductor.RequestActivityStream(http.MethodGet, url, "", "", data, output)
}

func (dispatcher *dispatcher) PostActivityStream(url string, data []byte, output interface{}) error {
	return dispatcher.conductor.RequestActivityStream(http.MethodPost, url, "", "", data, output)
}

func (dispatcher *dispatcher) GetActivityStreamSigned(url, keyId, privateKey string, data []byte, output interface{}) error {
	return dispatcher.conductor.RequestActivityStream(http.MethodGet, url, keyId, privateKey, data, output)
}

func (dispatcher *dispatcher) PostActivityStreamSigned(url, keyId, privateKey string, data []byte, output interface{}) error {
	return dispatcher.conductor.RequestActivityStream(http.MethodPost, url, keyId, privateKey, data, output)
}

//endregion
