package models

type Product struct {
	Id          int    `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	CategoryId  int    `json:"categoryId"`
	Img         string `json:"img"`
}

type Category struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
