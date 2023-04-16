package repository

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string
	Category string
	Price    int
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo ProductRepository) ListProducts() []Product {
	var products []Product
	repo.db.Find(&products)
	return products
}

func (repo ProductRepository) CreateProduct(product *Product) (*Product, error) {
	tx := repo.db.Create(product)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return product, nil
}
