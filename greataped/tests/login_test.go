package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestLogin(t *testing.T) {
	id := uuid.New().String()

	resp, err := Post("/api/v1/login", Payload{
		"email":    fmt.Sprintf("%s@%s", id, DOMAIN),
		"password": "123456",
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.FailNow()
	}
}
