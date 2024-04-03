package models

type Shelf struct {
	ShelfId   int    `json:"shelf_id"`
	ShelfName string `json:"shelf_name"`
}

type ShelfPrimaryKey struct {
	ShelfId int `json:"shelf_id"`
}

type CreateShelf struct {
	ShelfName string `json:"shelf_name"`
}

type UpdateShelf struct {
	ShelfId   int    `json:"shelf_id"`
	ShelfName string `json:"shelf_name"`
}
