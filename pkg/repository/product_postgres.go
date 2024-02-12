package repository

import (
	"fmt"

	"github.com/MrSaveliy/store/models"
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) CreateProduct(product models.Product) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description, price, quantity, category, img) values ($1, $2, $3, $4, $5, $6)", productsTable)
	row := r.db.QueryRow(query, product.Name, product.Descripton, product.Price, product.Quantity, product.Category, product.Img)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
