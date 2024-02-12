package repository

import (
	"fmt"

	"github.com/MrSaveliy/store/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}

// func (r *AuthPostgres) CreateProfile(profile models.Profile, userId int) (int, error) {
// 	var id int
// 	query := fmt.Sprintf("INSERT INTO %s (name, surname, patronymic, adress, number, date, user_id) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", profileTable)
// 	row := r.db.QueryRow(query, profile.Name, profile.Surname, profile.Patronymic, profile.Adress, profile.Number, profile.Date, profile.UserId)
// 	if err := row.Scan(&id); err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }
