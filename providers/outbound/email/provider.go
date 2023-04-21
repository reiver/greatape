package email

import (
	"github.com/reiver/greatape/providers/outbound/common/messaging"
	"github.com/xeronith/diamante/contracts/logging"

	"github.com/xeronith/diamante/contracts/email"
	. "github.com/xeronith/diamante/contracts/messaging"
)

const EmailProvider = "EMAIL_PROVIDER"

type provider struct {
	messaging IMessagingProvider
}

func NewProvider(logger logging.ILogger) email.IEmailProvider {
	return &provider{
		messaging: messaging.NewProvider(EmailProvider, logger, handler),
	}
}

func (provider *provider) Send(receiver, message string) error {
	return provider.messaging.Send(receiver, message)
}
