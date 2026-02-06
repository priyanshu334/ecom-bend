package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/tw-bend/internal/db"
	"github.com/priyanshu334/tw-bend/internal/module/auth"
	"github.com/priyanshu334/tw-bend/internal/module/product"
	"github.com/priyanshu334/tw-bend/internal/module/user"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	authRepo := auth.NewRepository(db.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHanler(authService)

	userRepo := user.NewRepository(db.DB)

	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	ProductRepo := product.NewRepository(db.DB)
	ProductService := product.NewService(ProductRepo)
	ProductHandler := product.NewHandler(ProductService)

	auth.RegisterRoutes(api, authHandler)
	user.RegisterRoutes(api, userHandler)
	product.RegisterRoutes(api, ProductHandler)
}
