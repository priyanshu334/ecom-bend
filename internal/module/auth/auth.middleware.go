package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/priyanshu334/tw-bend/internal/config"
)

func RequireAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenStr := c.Cookies("access_token")
		if tokenStr == "" {
			return fiber.ErrUnauthorized
		}

		token, err := jwt.ParseWithClaims(
			tokenStr,
			&Claims{},
			func(t *jwt.Token) (interface{}, error) {
				return []byte(config.Cfg.JWTSecret), nil
			},
		)
		if err != nil || !token.Valid {
			return fiber.ErrUnauthorized
		}

		claims := token.Claims.(*Claims)
		c.Locals("user_id", claims.UserID)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}
