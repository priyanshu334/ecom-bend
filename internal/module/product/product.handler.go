package product

import "github.com/gofiber/fiber/v3"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetProduct(c fiber.Ctx) error {
	slug := c.Params("slug")
	product, err := h.service.GetProductBySlug(slug)
	if err != nil {
		if err == ErrProductNotFound {
			return fiber.NewError(fiber.StatusNotFound, "product not found")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get product")
	}
	return c.Status(fiber.StatusOK).JSON(ProductResponse{
		Name:        product.Name,
		Slug:        product.Slug,
		Description: product.Description,
		Price:       int(product.Price),
	})
}

func (h *Handler) ListProducts(c fiber.Ctx) error {
	products, err := h.service.ListProducts()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to list products")
	}

	resp := make([]ProductResponse, 0)
	for _, p := range products {
		resp = append(resp, ProductResponse{
			Name:        p.Name,
			Slug:        p.Slug,
			Description: p.Description,
			Price:       int(p.Price),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *Handler) CreateCategory(c fiber.Ctx) error {
	var req CreateCategoryRequest

	if err := c.Bind().Body(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	if err := h.service.CreateCategory(req.Name, req.Slug); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "falid to create category")
	}
	return c.SendStatus(fiber.StatusCreated)

}

func (h *Handler) CreateProduct(c fiber.Ctx) error {
	var req CreateProductRequest

	if err := c.Bind().Body(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	if err := h.service.CreateProduct(req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create prodcut")
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) UpdateProduct(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")

	var req UpdateProductRequest

	if err := c.Bind().Body(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	if err := h.service.UpdateProduct(uint(id), req); err != nil {
		return fiber.NewError(fiber.StatusNotFound, "prodcut not found")
	}
	return c.SendStatus(fiber.StatusOK)
}
