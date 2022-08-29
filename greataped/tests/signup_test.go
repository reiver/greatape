package tests

import (
	"app/models/types"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestSignup(t *testing.T) {
	id := uuid.New().String()
	payload := types.SignupDTO{
		LoginDTO: types.LoginDTO{
			Email:    fmt.Sprintf("%s@%s", id, DOMAIN),
			Password: "123456",
		},
		Username: id,
	}

	resp, err := Post("/api/v1/signup", payload)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.FailNow()
	}
}
