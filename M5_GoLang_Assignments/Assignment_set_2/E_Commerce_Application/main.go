package main

import (
	"E_Commerce_Application/config"
	"E_Commerce_Application/controller"
	"E_Commerce_Application/middleware"
	"E_Commerce_Application/repository"
	"E_Commerce_Application/service"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	config.InitializeDatabase()

	// Initialize repositories, services, and controllers
	productRepo := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/product", productController.AddProduct)
	mux.HandleFunc("/product/", productController.GetProductByID)
	mux.HandleFunc("/product/update", productController.UpdateProductStock)
	mux.HandleFunc("/product/delete", productController.RemoveProduct)

	// Apply middleware
	handler := middleware.JWTAuthentication(mux)       // JWT Authentication
	handler = middleware.ValidateProductInput(handler) // Input Validation

	// Start the server
	log.Println("Server is running on port :8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
