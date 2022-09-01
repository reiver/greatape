package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	for _, statusCode := range []int{
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusNotFound,
		http.StatusInternalServerError,
	} {
		url := fmt.Sprintf("/error/v1/%d/error_message", statusCode)
		resp, err := Get(url)
		if err != nil {
			t.Fatal(err)
		}

		if resp.StatusCode != statusCode {
			t.FailNow()
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		data := &struct {
			Type    string
			Version int
			Payload any
		}{}

		if err := json.Unmarshal(body, data); err != nil {
			t.Fatal(err)
		}

		if data.Type != "server_error" ||
			data.Version != 1 ||
			data.Payload != "error_message" {
			t.FailNow()
		}
	}
}
