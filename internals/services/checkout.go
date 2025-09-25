package services

import (
	"context"

	"github.com/mahdifr17/CheckoutService/internals/models"
	"github.com/mahdifr17/CheckoutService/internals/services/appmodel"
)

type CheckoutService interface {
	ProcessCheckout(ctx context.Context, req appmodel.CheckoutCommand) (appmodel.OrderSummary, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
}

type checkoutService struct {
	productRepo ProductRepository
}

func NewCheckoutService(productRepo ProductRepository) CheckoutService {
	return &checkoutService{productRepo: productRepo}
}

func (s *checkoutService) ProcessCheckout(ctx context.Context, req appmodel.CheckoutCommand) (appmodel.OrderSummary, error) {
	return appmodel.OrderSummary{}, nil
}

func (s *checkoutService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	return s.productRepo.GetAll(ctx)
}
