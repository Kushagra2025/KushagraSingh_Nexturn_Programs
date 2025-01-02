package service

import (
	"Blog_Management_System/model"
	"Blog_Management_System/repository"
)

type BlogService struct {
	BlogRepo *repository.BlogRepository
}

func NewBlogService(blogrepo *repository.BlogRepository) *BlogService {
	return &BlogService{BlogRepo: blogrepo}
}

func (service *BlogService) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	return service.BlogRepo.CreateBlog(blog)
}

func (service *BlogService) GetBlog(id int) (*model.Blog, error) {
	return service.BlogRepo.GetBlog(id)
}

func (service *BlogService) GetAllBlogs() ([]model.Blog, error) {
	return service.BlogRepo.GetAllBlogs()
}

func (service *BlogService) UpdateBlog(user *model.Blog) (*model.Blog, error) {
	return service.BlogRepo.UpdateBlog(user)
}

func (service *BlogService) DeleteBlog(id int) error {
	return service.BlogRepo.DeleteBlog(id)
}
