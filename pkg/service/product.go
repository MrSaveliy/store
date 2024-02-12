package service

import (
	"github.com/MrSaveliy/store/models"
	"github.com/MrSaveliy/store/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product models.Product) (int, error) {
	return s.repo.CreateProduct(product)
}
