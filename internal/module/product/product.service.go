package product

import (
	"gorm.io/gorm"
)

type Service interface {
	ListProducts() ([]Product, error)
	GetProductBySlug(slug string) (*Product, error)
	GetCategories() ([]Category, error)

	CreateCategory(name, slug string) error
	CreateProduct(req CreateProductRequest) error
	UpdateProduct(id uint, req UpdateProductRequest) error
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) ListProducts() ([]Product, error) {
	return s.repo.ListProducts()
}

func (s *service) GetProductBySlug(slug string) (*Product, error) {
	product, err := s.repo.GetProductBySlug(slug)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	return product, nil
}

func (s *service) GetCategories() ([]Category, error) {
	return s.repo.ListCategories()
}

func (s *service) CreateCategory(name, slug string) error {
	return s.repo.CreateCategory(&Category{
		Name: name,
		Slug: slug,
	})
}

func (s *service) CreateProduct(req CreateProductRequest) error {
	return s.repo.CreateProduct(&Product{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Price:       req.Price,
		IsActive:    true,
	})

}

func (s *service) UpdateProduct(id uint, req UpdateProductRequest) error {
	product, err := s.repo.GetProductByID(id)

	if err != nil {
		return ErrProductNotFound
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.IsActive != nil {
		product.IsActive = *req.IsActive
	}
	return s.repo.UpdateProduct(product)
}
