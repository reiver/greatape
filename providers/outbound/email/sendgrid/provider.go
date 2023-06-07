package sendgrid

import (
	"github.com/reiver/greatape/providers/outbound/common/messaging"
	"github.com/xeronith/diamante/contracts/email"
	"github.com/xeronith/diamante/contracts/logging"
)

const ProviderName = "SENDGRID_EMAIL_PROVIDER"

type provider struct {
	messaging messaging.IMessagingProvider
}

func NewProvider(logger logging.ILogger) email.IEmailProvider {
	return &provider{
		messaging: messaging.NewProvider(ProviderName, logger, handler),
	}
}

func (provider *provider) Send(receiver, template string, args ...interface{}) error {
	model := map[string]interface{}{}
	if len(args) > 0 {
		if data, ok := args[0].(map[string]interface{}); ok {
			model = data
		}
	}

	return provider.messaging.Send(receiver, template, model)
}
