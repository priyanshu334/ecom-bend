package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/tw-bend/internal/module/auth"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	group := router.Group("/users", auth.RequireAuth())

	group.Get("/me", handler.GetProfile)
	group.Put("/me", handler.UpdateProfile)
	group.Get("/me/address", handler.ListAddress)
	group.Post("/me/address", handler.CreateAddress)
	group.Delete("/me/address/:id", handler.DeleteAddress)

}
