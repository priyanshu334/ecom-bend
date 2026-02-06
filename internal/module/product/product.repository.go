package product

import (
	"gorm.io/gorm"
)

type Repository interface {
	ListProducts() ([]Product, error)
	GetProductBySlug(slug string) (*Product, error)
	ListCategories() ([]Category, error)

	CreateCategory(category *Category) error
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	GetProductByID(id uint) (*Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) ListProducts() ([]Product, error) {
	var products []Product
	err := r.db.Where("is_active=?", true).Find(&products).Error
	return products, err
}

func (r *repository) GetProductBySlug(slug string) (*Product, error) {
	var product Product
	err := r.db.Where("slug=? AND is_active=?", slug, true).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
func (r *repository) ListCategories() ([]Category, error) {
	var catergories []Category
	err := r.db.Find(&catergories).Error
	return catergories, err
}

func (r *repository) CreateCategory(category *Category) error {
	return r.db.Create(category).Error
}

func (r *repository) CreateProduct(product *Product) error {
	return r.db.Create(product).Error
}

func (r *repository) UpdateProduct(product *Product) error {
	return r.db.Save(product).Error
}

func (r *repository) GetProductByID(id uint) (*Product, error) {
	var product Product

	err := r.db.Where("id=?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil

}
