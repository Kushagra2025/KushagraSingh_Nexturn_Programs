package controller

import (
	"E_Commerce_Application/model"
	"E_Commerce_Application/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService *service.ProductService
}

// NewProductController creates a new instance of ProductController.
func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{productService}
}

// AddProduct handles the creation of a new product.
func (controller *ProductController) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := controller.productService.AddProduct(&product); err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// GetProductByID handles fetching product details by ID.
func (controller *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/product/"):])
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	product, err := controller.productService.GetProductByID(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(product)
}

// UpdateProductStock handles updating product stock.
func (controller *ProductController) UpdateProductStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/product/"):])
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var stockUpdate struct {
		Stock int `json:"stock"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&stockUpdate); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := controller.productService.UpdateProductStock(id, stockUpdate.Stock); err != nil {
		http.Error(w, "Failed to update stock", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// RemoveProduct handles deleting a product by ID.
func (controller *ProductController) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/product/"):])
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	if err := controller.productService.RemoveProduct(id); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
