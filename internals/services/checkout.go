package services

import (
	"context"

	"github.com/mahdifr17/CheckoutService/internals/domain"
)

type CheckoutService interface {
	ProcessCheckout(ctx context.Context, input ProcessCheckoutInput) (*domain.Order, error)
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
}

type checkoutService struct {
	productRepo   domain.ProductRepository
	promotionRepo domain.PromotionRepository
	orderRepo     domain.OrderRepository
}

func NewCheckoutService(
	productRepo domain.ProductRepository,
	promotionRepo domain.PromotionRepository,
	orderRepo domain.OrderRepository,
) CheckoutService {
	return &checkoutService{
		productRepo:   productRepo,
		promotionRepo: promotionRepo,
		orderRepo:     orderRepo,
	}
}

type (
	ProcessCheckoutItemm struct {
		ProductID string
		Quantity  int
	}

	ProcessCheckoutInput struct {
		Items []ProcessCheckoutItemm
	}
)

func (s *checkoutService) ProcessCheckout(ctx context.Context, input ProcessCheckoutInput) (*domain.Order, error) {
	var items []domain.OrderItem
	subtotal := 0.0

	for _, reqItem := range input.Items {
		product, err := s.productRepo.GetByID(ctx, reqItem.ProductID)
		if err != nil {
			return nil, err
		}
		itemPrice := product.Price * float64(reqItem.Quantity)

		items = append(items, domain.OrderItem{
			Product:  product,
			Quantity: reqItem.Quantity,
			Price:    itemPrice,
		})
		subtotal += itemPrice
	}

	// [ ] handle promo
	total := subtotal

	order := &domain.Order{
		Subtotal: subtotal,
		Discount: 0,
		Total:    total,
		Items:    items,
	}

	if err := s.orderRepo.Create(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *checkoutService) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	return s.productRepo.GetAll(ctx)
}
