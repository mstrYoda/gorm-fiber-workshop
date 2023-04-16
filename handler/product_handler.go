package handler

import (
	"goworkshop/repository"
)

type ProductHandlerInterface interface {
	ListProducts() []repository.Product
	CreateProduct(product *repository.Product) (*repository.Product, error)
}

type ProductHandler struct {
	ProductRepository *repository.ProductRepository
}

func NewProductHandler(productRepository *repository.ProductRepository) ProductHandlerInterface {
	return &ProductHandler{ProductRepository: productRepository}
}

func (handler ProductHandler) ListProducts() []repository.Product {
	return handler.ProductRepository.ListProducts()
}

func (handler ProductHandler) CreateProduct(product *repository.Product) (*repository.Product, error) {
	return handler.ProductRepository.CreateProduct(product)
}
