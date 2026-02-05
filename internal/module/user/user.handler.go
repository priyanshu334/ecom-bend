package user

import (
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func GetUserId(c fiber.Ctx) (uint, error) {
	UserID, ok := c.Locals("user_id").(uint)
	if !ok {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}
	return UserID, nil

}

func (h *Handler) GetProfile(c fiber.Ctx) error {
	userID, err := GetUserId(c)
	if err != nil {
		return err

	}
	profile, err := h.service.GetOrCreateProfile(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get and create profile")
	}
	return c.Status(fiber.StatusOK).JSON(ProfileResponse{
		FullName: profile.FullName,
		Phone:    profile.Phone,
	})
}

func (h *Handler) UpdateProfile(c fiber.Ctx) error {
	userID, err := GetUserId(c)
	if err != nil {
		return err
	}
	var req UpdateProfileRequest
	if err := c.Bind().Body(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	profile, err := h.service.UpdateProfile(userID, req.FullName, req.Phone)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to UpdateProfile")
	}
	return c.Status(fiber.StatusOK).JSON(
		ProfileResponse{
			FullName: profile.FullName,
			Phone:    profile.Phone,
		},
	)

}

func (h *Handler) CreateAddress(c fiber.Ctx) error {
	userID, err := GetUserId(c)
	if err != nil {
		return err
	}
	var req CreateAddressRequest
	if err := c.Bind().Body(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid input")
	}
	if err := h.service.AddAddress(userID, Address{
		Label:      req.Label,
		Line1:      req.Line1,
		Line2:      req.Line2,
		City:       req.City,
		State:      req.State,
		PostalCode: req.PostalCode,
		Country:    req.Country,
	}); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create adddresss")
	}
	return c.SendStatus(fiber.StatusOK)

}

func (h *Handler) DeleteAddress(c fiber.Ctx) error {
	userID, err := GetUserId(c)
	if err != nil {
		return err
	}
	addressID := fiber.Params[int](c, "id")

	if err := h.service.RemoveAddress(userID, uint(addressID)); err != nil {
		return fiber.NewError(fiber.StatusForbidden, "not allowed")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) ListAddress(c fiber.Ctx) error {
	userID, err := GetUserId(c)
	if err != nil {
		return err
	}
	address, err := h.service.ListAddresses(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "not found")
	}
	resp := make([]AddressResponse, 0)
	for _, a := range address {
		resp = append(resp, AddressResponse{
			ID:         a.ID,
			Label:      a.Label,
			Line1:      a.Line1,
			Line2:      a.Line2,
			City:       a.City,
			State:      a.State,
			PostalCode: a.PostalCode,
			Country:    a.Country,
			IsDefault:  a.IsDefault,
		})
	}

	return c.JSON(resp)

}
