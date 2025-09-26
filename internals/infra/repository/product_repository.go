package repository

import (
	"context"

	"github.com/mahdifr17/CheckoutService/internals/domain"
	"github.com/mahdifr17/CheckoutService/internals/infra/repository/postgresql"
	"gorm.io/gorm"
)

type psqlProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &psqlProductRepository{db: db}
}

func (r *psqlProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var (
		products []postgresql.Product
		out      []domain.Product
	)
	err := r.db.Find(&products).Error
	for _, v := range products {
		out = append(out, *v.ToDomain())
	}
	return out, err
}

func (r *psqlProductRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	var product postgresql.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return product.ToDomain(), nil
}

func (r *psqlProductRepository) Save(ctx context.Context, product *domain.Product) error {
	productModel := postgresql.ProductFromDomain(*product)
	return r.db.Save(productModel).Error
}
