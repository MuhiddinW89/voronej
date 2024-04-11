package models

type ProductShelf struct {
	ProductShelfId int      `json:"productshelf_id"`
	ProductId      int      `json:"product_id"`
	ProductData    *Product `json:"product_data"`
	ShelfId        int      `json:"shelf_id"`
	ShelfData      *Shelf   `json:"shelf_data"`
	IsMain         bool     `json:"is_main"`
	Quantity       int      `json:"quantity"`
}

type ProductShelfPrimaryKey struct {
	ProductShelfId int `json:"productshelf_id"`
}

type CreateProductShelf struct {
	ProductId int  `json:"product_id"`
	ShelfId   int  `json:"shelf_id"`
	IsMain    bool `json:"is_main"`
}

type UpdateProductShelf struct {
	ProductShelfId int  `json:"productshelf_id"`
	ProductId      int  `json:"product_id"`
	ShelfId        int  `json:"shelf_id"`
	IsMain         bool `json:"is_main"`
	Quantity       int  `json:"quantity"`
}
