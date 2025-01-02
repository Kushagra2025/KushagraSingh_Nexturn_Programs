package controller

import (
	"Blog_Management_System/model"
	"Blog_Management_System/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type BlogController struct {
	BlogService *service.BlogService
}

func NewBlogController(BlogService *service.BlogService) *BlogController {
	return &BlogController{BlogService: BlogService}
}

func (controller *BlogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdBlog, err := controller.BlogService.CreateBlog(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdBlog)
}

func (controller *BlogController) GetBlog(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	blog, err := controller.BlogService.GetBlog(blogID)
	if err != nil {
		http.Error(w, "ID not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blog)
}

func (controller *BlogController) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := controller.BlogService.GetAllBlogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blogs)
}

func (controller *BlogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var blog model.Blog
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blog.ID = blogID
	updatedBlog, err := controller.BlogService.UpdateBlog(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBlog)
}

func (controller *BlogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	blogID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = controller.BlogService.DeleteBlog(blogID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Blog Deleted Successfully."}`))
}
