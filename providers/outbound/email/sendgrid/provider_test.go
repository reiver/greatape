package sendgrid_test

import (
	"testing"

	"github.com/reiver/greatape/providers/outbound/email/sendgrid"
	"github.com/xeronith/diamante/logging"
)

func Test_Send(t *testing.T) {
	logger := logging.NewLogger(false)
	provider := sendgrid.NewProvider(logger)

	if err := provider.Send("somebody@somewhere.com", "Message content"); err != nil {
		t.Fatal(err)
	}
}
