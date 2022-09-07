package authorize

import (
	"strings"
	"utility/jwt"

	"github.com/gofiber/fiber/v2"
)

func New() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		header := ctx.Get(fiber.HeaderAuthorization)
		if header == "" {
			return fiber.ErrUnauthorized
		}

		chunks := strings.Split(header, " ")
		if len(chunks) < 2 || chunks[0] != "Bearer" {
			return fiber.ErrUnauthorized
		}

		user, err := jwt.Verify(chunks[1])
		if err != nil {
			return fiber.ErrUnauthorized
		}

		ctx.Locals("USER", user.ID)
		return ctx.Next()
	}
}
