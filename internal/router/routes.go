package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/tw-bend/internal/db"
	"github.com/priyanshu334/tw-bend/internal/module/auth"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	userRepo := auth.NewRepository(db.DB)
	userService := auth.NewService(userRepo)
	userHandler := auth.NewHanler(userService)

	auth.RegisterRoutes(api, userHandler)
}
