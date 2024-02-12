package service

import (
	"github.com/MrSaveliy/store/models"
	"github.com/MrSaveliy/store/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseTokern(token string) (int, error)
	// CreateProfile(profile models.Profile, userId int) (int, error)
}

type Product interface {
	CreateProduct(product models.Product) (int, error)
}

type Category interface {
}

type Order interface {
}

type Service struct {
	Authorization
	Product
	Category
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Product:       NewProductService(repos.Product),
	}
}
