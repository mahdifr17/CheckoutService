package services

import (
	"context"

	"github.com/mahdifr17/CheckoutService/internals/models"
)

type (
	ProductRepository interface {
		GetAll(ctx context.Context) ([]models.Product, error)
		GetByID(ctx context.Context, id string) (*models.Product, error)
		Create(ctx context.Context, product *models.Product) error
		Update(ctx context.Context, product *models.Product) error
		Delete(ctx context.Context, id string) error
		UpdateStock(ctx context.Context, id string, quantity int) error
	}
)
