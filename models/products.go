package models

type Product struct {
	Id         int    `json:"-"`
	Name       string `json:"name"`
	Descripton string `json:"Descripton"`
	Price      int    `json:"price"`
	Quantity   int    `json:"quantity"`
	Category   string `json:"category"`
	Img        string `json:"img"`
}

type Category struct {
	Id         int    `json:"-"`
	Name       string `json:"name"`
	Descripton string `json:"Descripton"`
}
