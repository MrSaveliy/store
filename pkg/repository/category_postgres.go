package repository

import (
	"fmt"

	"github.com/MrSaveliy/store/models"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) CreateCategory(category models.Category) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description) values ($1, $2) RETURNING id", categoriesTable)
	row := r.db.QueryRow(query, category.Name, category.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CategoryPostgres) GetById(categoryId int) (models.Category, error) {
	var category models.Category
	query := fmt.Sprintf("SELECT name, description FROM %s WHERE id=$1", categoriesTable)
	err := r.db.Get(&category, query, categoryId)

	return category, err
}
