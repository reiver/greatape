package email_test

import (
	"testing"

	"github.com/xeronith/diamante/logging"
	"rail.town/infrastructure/providers/outbound/email"
)

func Test_Send(t *testing.T) {
	logger := logging.NewLogger(false)
	provider := email.NewProvider(logger)

	if err := provider.Send("somebody@somewhere.com", "Message content"); err != nil {
		t.Fatal(err)
	}
}
