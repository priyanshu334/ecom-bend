package product

type ProductResponse struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type CategoryResponse struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
	Slug string `json:"slug" validate:"required"`
}

type CreateProductRequest struct {
	CategoryID  uint   `json:"category_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description"`
	Price       int64  `json:"price" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	IsActive    *bool  `json:"is_active"`
}
