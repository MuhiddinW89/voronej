package models

type OrderDetail struct {
	OrderDetailId int      `json:"order_detail_id"`
	OrderId       int      `json:"order_id"`
	OrderData     *Orders  `json:"order_data"`
	ProductId     int      `json:"product_id"`
	ProductData   *Product `json:"product_data"`
	Quantity      int      `json:"quantity"`
}

type OrderDetailPrimaryKey struct {
	OrderDetailId int `json:"order_detail_id"`
}

type CreateOrderDetail struct {
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type UpdateOrderDetail struct {
	OrderDetailId int `json:"order_detail_id"`
	OrderId       int `json:"order_id"`
	ProductId     int `json:"product_id"`
	Quantity      int `json:"quantity"`
}
