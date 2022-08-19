package tests

import (
	"net/http"
	"testing"
)

func TestHealth(t *testing.T) {
	resp, err := http.DefaultClient.Get(Root + "/health")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.FailNow()
	}
}
