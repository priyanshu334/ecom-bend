package middleware

import "github.com/gofiber/fiber/v3"

func AdminOnly() fiber.Handler {
	return func(c fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok || role != "admin" {
			return fiber.NewError(fiber.StatusForbidden, "admin access required")
		}
		return c.Next()

	}
}
