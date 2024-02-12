package repository

import (
	"github.com/MrSaveliy/store/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
	// CreateProfile(profile models.Profile, userId int) (int, error)
}

type Product interface {
	CreateProduct(product models.Product) (int, error)
}

type Category interface {
}

type Order interface {
}

type Repository struct {
	Authorization
	Product
	Category
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Product:       NewProductPostgres(db),
	}
}
