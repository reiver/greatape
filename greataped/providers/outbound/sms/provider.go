package sms

import (
	"github.com/xeronith/diamante/contracts/logging"
	"rail.town/infrastructure/providers/outbound/common/messaging"

	. "github.com/xeronith/diamante/contracts/messaging"
	"github.com/xeronith/diamante/contracts/sms"
)

const SMSProvider = "SMS_PROVIDER"

type provider struct {
	messaging IMessagingProvider
}

func NewProvider(logger logging.ILogger) sms.ISMSProvider {
	return &provider{
		messaging: messaging.NewProvider(SMSProvider, logger, handler),
	}
}

func (provider *provider) Send(receiver, message string) error {
	return provider.messaging.Send(receiver, message)
}
