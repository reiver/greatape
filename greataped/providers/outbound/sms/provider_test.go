package sms_test

import (
	"testing"

	"github.com/xeronith/diamante/logging"
	"rail.town/infrastructure/providers/outbound/sms"
)

func Test_Send(t *testing.T) {
	logger := logging.NewLogger(false)
	provider := sms.NewProvider(logger)

	if err := provider.Send("09123456789", "Message content"); err != nil {
		t.Fatal(err)
	}
}
