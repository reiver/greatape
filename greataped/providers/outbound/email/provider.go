package email

import (
	"github.com/xeronith/diamante/contracts/logging"
	"rail.town/infrastructure/providers/outbound/common/messaging"

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
