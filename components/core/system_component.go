package core

import (
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/reiver/greatape/providers/outbound/email"
	"github.com/reiver/greatape/providers/outbound/sms"
	. "github.com/xeronith/diamante/contracts/email"
	. "github.com/xeronith/diamante/contracts/logging"
	. "github.com/xeronith/diamante/contracts/settings"
	. "github.com/xeronith/diamante/contracts/sms"
	"github.com/xeronith/diamante/utility"
	"github.com/xeronith/diamante/utility/concurrent"
	"github.com/xeronith/diamante/utility/jwt"
)

type systemComponent struct {
	sync.RWMutex
	lastId        int64
	logger        ILogger
	configuration IConfiguration
	emailProvider IEmailProvider
	smsProvider   ISMSProvider
}

func newSystemComponent(configuration IConfiguration, logger ILogger) systemComponent {
	return systemComponent{
		logger:        logger,
		configuration: configuration,
		emailProvider: email.NewProvider(logger),
		smsProvider:   sms.NewProvider(logger),
	}
}

func (component *systemComponent) IsTestEnvironment() bool {
	return component.configuration.IsTestEnvironment()
}

func (component *systemComponent) IsDevelopmentEnvironment() bool {
	return component.configuration.IsDevelopmentEnvironment()
}

func (component *systemComponent) IsStagingEnvironment() bool {
	return component.configuration.IsStagingEnvironment()
}

func (component *systemComponent) IsProductionEnvironment() bool {
	return component.configuration.IsProductionEnvironment()
}

func (component *systemComponent) UniqueId() int64 {
	component.Lock()
	defer component.Unlock()

	var id int64
	for {
		id = time.Now().UnixNano() / 1000
		if id != component.lastId {
			break
		}

		time.Sleep(time.Microsecond)
	}

	component.lastId = id
	return id
}

func (component *systemComponent) Logger() ILogger {
	return component.logger
}

func (component *systemComponent) Async(runnable func()) {
	concurrent.NewAsyncTask(runnable).Run()
}

// Utility

func (component *systemComponent) UnixNano() int64 {
	return time.Now().UnixNano()
}

func (component *systemComponent) GenerateUUID() string {
	return utility.GenerateUUID()
}

func (component *systemComponent) GenerateSalt() string {
	return utility.GenerateUUID()
}

func (component *systemComponent) GenerateHash(value string, salt string) string {
	return utility.GenerateHash(value, salt)
}

func (component *systemComponent) GenerateJwtToken() string {
	return jwt.Generate(
		component.configuration.GetServerConfiguration().GetJwtTokenKey(),
		component.configuration.GetServerConfiguration().GetJwtTokenExpiration(),
	)
}

func (component *systemComponent) VerifyJwtToken(token string) error {
	_, err := jwt.Verify(token, component.configuration.GetServerConfiguration().GetJwtTokenKey())
	return err
}

func (component *systemComponent) GenerateCode() string {
	return utility.GenerateConfirmationCode()
}

func (component *systemComponent) GenerateRSAKeyPair() (string, string, error) {
	return utility.GenerateRSAKeyPair()
}

func (component *systemComponent) Email(destination string, format string, args ...interface{}) {
	component.Async(func() {
		message := fmt.Sprintf(format, args...)
		component.emailProvider.Send(destination, message)
	})
}

func (component *systemComponent) SMS(destination string, format string, args ...interface{}) {
	component.Async(func() {
		message := fmt.Sprintf(format, args...)
		component.smsProvider.Send(destination, message)
	})
}

func (component *systemComponent) Format(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func (component *systemComponent) Match(pattern, input string) (bool, error) {
	return regexp.MatchString(pattern, input)
}

func (component *systemComponent) Error(arg interface{}) error {
	return fmt.Errorf("%v", arg)
}
