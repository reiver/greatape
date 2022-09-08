package server

import (
	"github.com/gofiber/fiber/v2"
)

func newError(code int, message string) *fiber.Error {
	return &fiber.Error{
		Code:    code,
		Message: message,
	}
}
