package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestSignup(t *testing.T) {
	id := uuid.New().String()

	resp, err := Post("/api/v1/signup", Payload{
		"email":    fmt.Sprintf("%s@%s", id, DOMAIN),
		"password": "123456",
		"username": id,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.FailNow()
	}
}
