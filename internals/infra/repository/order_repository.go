package repository

import (
	"context"

	"github.com/mahdifr17/CheckoutService/internals/domain"
	"github.com/mahdifr17/CheckoutService/internals/infra/repository/postgresql"
	"gorm.io/gorm"
)

type psqlOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domain.OrderRepository {
	return &psqlOrderRepository{db: db}
}

func (r *psqlOrderRepository) GetAll(ctx context.Context) ([]domain.Order, error) {
	var (
		orders []postgresql.Order
		out    []domain.Order
	)
	err := r.db.Find(&orders).Error
	for _, v := range orders {
		out = append(out, *v.ToDomain())
	}
	return out, err
}

func (r *psqlOrderRepository) Create(ctx context.Context, order *domain.Order) error {
	orderModel := postgresql.OrderFromDomain(*order)
	return r.db.Save(orderModel).Error
}
