package service

import (
	"E_Commerce_Application/model"
	"E_Commerce_Application/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

// NewProductService creates a new instance of ProductService.
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}

// AddProduct creates a new product.
func (s *ProductService) AddProduct(product *model.Product) error {
	return s.repo.CreateProduct(product)
}

// GetProductByID fetches a product by ID.
func (s *ProductService) GetProductByID(id int) (*model.Product, error) {
	return s.repo.GetProduct(id)
}

// UpdateProductStock updates the stock of a product.
func (s *ProductService) UpdateProductStock(id int, stock int) error {
	return s.repo.UpdateProductStock(id, stock)
}

// RemoveProduct deletes a product by ID.
func (s *ProductService) RemoveProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
