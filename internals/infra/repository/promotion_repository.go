package repository

import (
	"context"

	"github.com/mahdifr17/CheckoutService/internals/domain"
	"github.com/mahdifr17/CheckoutService/internals/infra/repository/postgresql"
	"gorm.io/gorm"
)

type psqlPromotionRepository struct {
	db *gorm.DB
}

func NewPromotionRepository(db *gorm.DB) domain.PromotionRepository {
	return &psqlPromotionRepository{db: db}
}

func (r *psqlPromotionRepository) GetByID(ctx context.Context, id string) (*domain.Promotion, error) {
	var promo postgresql.Promotion
	err := r.db.Find(&promo).Error
	return promo.ToDomain(), err
}
