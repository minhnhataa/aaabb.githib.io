package models

type Product struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Descrip string `json:"desc"`
	Price   string `json:"price"`
}
