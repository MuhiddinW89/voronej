package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Product() ProductRepoI
}

type ProductRepoI interface {
	Create(ctx context.Context, req *models.CreateProduct) (int, error)
	GetById(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error)
}
