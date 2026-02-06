package product

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/tw-bend/internal/middleware"
	"github.com/priyanshu334/tw-bend/internal/module/auth"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	public := router.Group("/product")
	public.Get("/", handler.ListProducts)
	public.Get("/:slug", handler.GetProduct)

	admin := router.Group("/admin/product", auth.RequireAuth(), middleware.AdminOnly())
	admin.Post("/", handler.CreateProduct)
	admin.Put("/:id", handler.UpdateProduct)

	adminCategories := router.Group("/admin/categories", auth.RequireAuth(), middleware.AdminOnly())
	adminCategories.Post("/", handler.CreateCategory)
}
