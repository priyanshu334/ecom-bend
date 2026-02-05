package auth

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service Service
}

func NewHanler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(c fiber.Ctx) error {
	var req RegisterRequest

	if err := c.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}
	if req.Email == "" || len(req.Password) < 8 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid email or password",
		)
	}
	user, err := h.service.Register(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, ErrorUserAlreadyExists) {
			return fiber.NewError(
				fiber.StatusConflict,
				"user already exists",
			)
		}
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusCreated).JSON(UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	})

}

func (h *Handler) Login(c fiber.Ctx) error {
	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return fiber.ErrBadRequest
	}
	user, err := h.service.Login(req.Email, req.Password)
	if err != nil {

		return fiber.NewError(fiber.StatusBadRequest, "invalid credintials")
	}

	accessToken, _ := generateToken(user, 15*time.Minute)
	refreshToken, _ := generateToken(user, 7*24*time.Minute)

	c.Cookie(&fiber.Cookie{
		Name:  "access_token",
		Value: accessToken,

		HTTPOnly: true,
		Secure:   false, // true in prod
		SameSite: "Lax",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	return c.JSON(UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	})

}

func (h *Handler) Me(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"user_id":   c.Locals("user_id"),
		"user_role": c.Locals("user_role"),
	})
}
