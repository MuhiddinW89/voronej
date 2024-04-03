package models

type Product struct {
	ProductId   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

type ProductPrimaryKey struct {
	ProductId int `json:"product_id"`
}

type CreateProduct struct {
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

type UpdateProduct struct {
	ProductId   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}
