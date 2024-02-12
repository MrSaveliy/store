package models

import "time"

type Profile struct {
	Id         int       `json:"-"`
	UserId     int       `json:"-"`
	Name       string    `json:"name" binding:"required"`
	Surname    string    `json:"surname" binding:"required"`
	Patronymic string    `json:"patronymic" binding:"required"`
	Adress     string    `json:"adress" binding:"required"`
	Number     string    `json:"number" binding:"required"`
	Date       time.Time `json:"date" binding:"required"`
}

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
