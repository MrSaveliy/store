package models

type Favorites struct {
	ID        int `json:"-"`
	UserId    int `json:"-"`
	ProductId int `json:"-"`
}
