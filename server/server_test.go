package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"goworkshop/repository"
)

type ProductHandlerMock struct {
}

func (handler ProductHandlerMock) ListProducts() []repository.Product {
	return []repository.Product{
		{
			Name: "bmw",
		},
	}
}

func (handler ProductHandlerMock) CreateProduct(product *repository.Product) (*repository.Product, error) {
	return &repository.Product{
		Name:     "bmw",
		Category: "araba",
		Price:    1,
	}, nil
}

func TestServerListProduct(t *testing.T) {
	mockHandler := ProductHandlerMock{}
	server := NewServer(mockHandler)

	request, _ := http.NewRequest("GET", "http://localhost:8080/products", nil)

	resp, err := server.Test(request)
	if err != nil {
		t.Error(err)
	}

	var products []repository.Product
	json.NewDecoder(resp.Body).Decode(&products)

	if len(products) == 0 {
		t.Errorf("Expected 0 products, got %d", len(products))
	}

	if products[0].Name != "bmw" {
		t.Errorf("Expected bmw, got %s", products[0].Name)
	}
}

func TestServerCreateProduct(t *testing.T) {
	mockHandler := ProductHandlerMock{}
	server := NewServer(mockHandler)

	product := repository.Product{
		Name:     "bmw",
		Category: "araba",
		Price:    1,
	}

	body, _ := json.Marshal(product)

	request, _ := http.NewRequest("POST", "http://localhost:8080/products", bytes.NewBuffer(body))
	request.Header.Add("content-type", "application/json")

	resp, err := server.Test(request)
	if err != nil {
		t.Error(err)
	}

	var createdProduct repository.Product
	json.NewDecoder(resp.Body).Decode(&createdProduct)

	if createdProduct.Name != "bmw" {
		t.Errorf("Expected bmw, got %s", createdProduct.Name)
	}

	if createdProduct.Category != "araba" {
		t.Errorf("Expected araba, got %s", createdProduct.Category)
	}
}
