package postgresql

import (
	"app/api/models"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) Create(ctx context.Context, req *models.CreateProduct) (int, error) {
	var (
		query string
		id    int
	)

	query = `
		INSERT INTO product(
			product_name,
			price
		)
		VALUES (
			 $1, $2, $3
		) 
		RETURNING product_id
	`

	err := p.db.QueryRow(ctx, query,
		req.ProductName,
		req.Price,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *productRepo) GetById(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {
	var (
		query   string
		product models.Product
	)

	query = `
		SELECT
			p.product_id,
			p.product_name,
			p.price
		FROM product AS p	
		WHERE product_id = $1
	`
	err := p.db.QueryRow(ctx, query, req.ProductId).Scan(
		&product.ProductId,
		&product.ProductName,
		&product.Price,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
