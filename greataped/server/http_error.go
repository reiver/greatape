package server

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func newError(code int, message string) *fiber.Error {
	data, _ := json.Marshal(struct {
		Type    string
		Version int
		Payload any
	}{
		Type:    "server_error",
		Version: 1,
		Payload: message,
	})

	return &fiber.Error{
		Code:    code,
		Message: string(data),
	}
}
