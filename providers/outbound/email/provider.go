package email

import (
	"github.com/xeronith/diamante/contracts/logging"

	"github.com/reiver/greatape/providers/outbound/email/postmark"
	"github.com/xeronith/diamante/contracts/email"
)

func NewProvider(logger logging.ILogger) email.IEmailProvider {
	return postmark.NewProvider(logger)
}
