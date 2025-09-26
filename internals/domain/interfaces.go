package domain

import (
	"context"
)

type (
	ProductRepository interface {
		GetAll(ctx context.Context) ([]Product, error)
		GetByID(ctx context.Context, id string) (*Product, error)
		Save(ctx context.Context, product *Product) error
	}

	PromotionRepository interface {
		GetByID(ctx context.Context, id string) (*Promotion, error)
	}

	OrderRepository interface {
		GetAll(ctx context.Context) ([]Order, error)
		Create(ctx context.Context, order *Order) error
	}
)
