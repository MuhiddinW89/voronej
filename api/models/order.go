package models

import "time"

type Orders struct {
	OrderId      int       `json:"order_id"`
	OrderDate    time.Time `json:"order_date"` //<-
	CustomerName string    `json:"customer_name"`
	TotalPrice   float64   `json:"total_price"`
}

type OrderPrimaryKey struct {
	OrderId int `json:"order_id"`
}

type CreateOrder struct {
	OrderDate    time.Time `json:"order_date"` //<-
	CustomerName string    `json:"customer_name"`
	TotalPrice   float64   `json:"total_price"`
}

type UpdateOrder struct {
	OrderId      int       `json:"order_id"`
	OrderDate    time.Time `json:"order_date"` //<-
	CustomerName string    `json:"customer_name"`
	TotalPrice   float64   `json:"total_price"`
}
