package service

import (
	"github.com/MrSaveliy/store/models"
	"github.com/MrSaveliy/store/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category models.Category) (int, error) {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetById(categoryId int) (models.Category, error) {
	return s.repo.GetById(categoryId)
}
