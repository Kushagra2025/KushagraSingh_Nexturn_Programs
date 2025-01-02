package main

import (
	"Blog_Management_System/config"
	"Blog_Management_System/controller"
	"Blog_Management_System/middleware"
	"Blog_Management_System/repository"
	"Blog_Management_System/service"
	"log"
	"net/http"
)

func YourProtectedHandler(w http.ResponseWriter, r *http.Request) {
	// This is the handler for the protected route
	w.Write([]byte("This is a protected route! You have access."))
}

func main() {
	// Initialize database connection
	config.InitializeDatabase()

	// Create repository, service, and controller
	blogRepo := repository.NewBlogRepository(config.GetDB())
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	// Initialize HTTP mux (router)
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/blogs", blogController.GetAllBlogs)
	mux.HandleFunc("/blogs/", blogController.GetBlog)
	mux.HandleFunc("/blogs/create", blogController.CreateBlog)
	mux.HandleFunc("/blogs/update", blogController.UpdateBlog)
	mux.HandleFunc("/blogs/delete", blogController.DeleteBlog)

	mux.Handle("/protected", middleware.BasicAuth(http.HandlerFunc(YourProtectedHandler)))

	// Apply middleware:
	// Wrap the mux with the middleware chain (logging, validating, etc.)
	handler := middleware.LogRequest(mux)      // Logging incoming requests
	handler = middleware.ValidateJson(handler) // Validate JSON payloads
	handler = middleware.BasicAuth(handler)    // Authentication
	handler = middleware.ErrorHandler(handler) // Error handling

	// Start the server
	log.Println("Server Started on port :8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
