package postmark_test

import (
	"testing"

	"github.com/reiver/greatape/providers/outbound/email/postmark"
	"github.com/xeronith/diamante/logging"
)

func Test_Send(t *testing.T) {
	logger := logging.NewLogger(false)
	provider := postmark.NewProvider(logger)

	if err := provider.Send("somebody@somewhere.com", "template-alias", map[string]interface{}{}); err != nil {
		t.Fatal(err)
	}
}
