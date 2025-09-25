package repository

import (
	"context"

	"github.com/mahdifr17/CheckoutService/internals/models"
	"github.com/mahdifr17/CheckoutService/internals/services"
	"gorm.io/gorm"
)

type psqlProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) services.ProductRepository {
	return &psqlProductRepository{db: db}
}

func (r *psqlProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *psqlProductRepository) GetByID(ctx context.Context, id string) (*models.Product, error) {
	return nil, nil
}

func (r *psqlProductRepository) Create(ctx context.Context, product *models.Product) error {
	return nil
}

func (r *psqlProductRepository) Update(ctx context.Context, product *models.Product) error {
	return nil
}

func (r *psqlProductRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r *psqlProductRepository) UpdateStock(ctx context.Context, id string, quantity int) error {
	return nil
}
