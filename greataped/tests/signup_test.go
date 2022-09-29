package tests

import (
	"encoding/json"
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

	result := &struct{ Code string }{}
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		t.Fatal(err)
	}

	resp, err = Post("/api/v1/verify", Payload{
		"email": fmt.Sprintf("%s@%s", id, DOMAIN),
		"code":  result.Code,
	})
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.FailNow()
	}
}
