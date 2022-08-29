package tests

import (
	"app/models/types"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestLogin(t *testing.T) {
	id := uuid.New().String()
	payload := types.LoginDTO{
		Email:    fmt.Sprintf("%s@%s", id, DOMAIN),
		Password: "123456",
	}

	resp, err := Post("/api/v1/login", payload)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusUnauthorized {
		t.FailNow()
	}
}
