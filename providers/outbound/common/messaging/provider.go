package messaging

import (
	"fmt"

	. "github.com/xeronith/diamante/contracts/logging"
)

type IMessagingHandler func(receiver string, template string, model map[string]interface{}) error

type IMessagingProvider interface {
	Send(receiver string, template string, model map[string]interface{}) error
}

type provider struct {
	name    string
	logger  ILogger
	handler IMessagingHandler
}

func NewProvider(name string, logger ILogger, handler IMessagingHandler) IMessagingProvider {
	return &provider{
		name:    name,
		logger:  logger,
		handler: handler,
	}
}

func (provider *provider) Send(receiver, message string, model map[string]interface{}) error {
	if err := provider.handler(receiver, message, model); err != nil {
		provider.logger.Error(fmt.Sprintf("%s: %s", provider.name, err.Error()))
		return err
	} else {
		provider.logger.Debug(fmt.Sprintf("%s: %s %s", provider.name, receiver, message))
		return nil
	}
}
