package models

import "time"

type Order struct {
	Id         int       `json:"-"`
	UserId     int       `json:"-"`
	Date       time.Time `json:"date"`
	Status     string    `json:"status"`
	TotalPrice int       `json:"total_price"`
	Adress     string    `json:"adress"`
}

type OrderItems struct {
	Id          int `json:"-"`
	OrderId     int `json:"-"`
	ProductId   int `json:"-"`
	Quantity    int `json:"quantity"`
	PriceForOne int `json:"price_for_one"`
	TotalPrice  int `json:"total_price"`
}
