package domain

import (
	"context"
	"time"
)

// Product is representing the Product data struct
type Product struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    int64     `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ProductUsecase represent the article's usecases
type ProductUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Product, string, error)
	GetByID(ctx context.Context, id int64) (Product, error)
	Store(context.Context, *Product) error
}

// ProductRepository represent the article's repository contract
type ProductRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Product, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Product, error)
	Store(ctx context.Context, a *Product) error
}
