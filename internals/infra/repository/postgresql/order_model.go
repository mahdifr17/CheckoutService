package postgresql

import "github.com/mahdifr17/CheckoutService/internals/domain"

type (
	Order struct {
		BaseModel
		Subtotal    float64 `gorm:"column:subtotal"`
		Discount    float64 `gorm:"column:discount"`
		Total       float64 `gorm:"column:total"`
		PromotionID *string `gorm:"column:promotion_id"`

		// Relationships
		Items     []OrderItem `gorm:"foreignKey:OrderID"`
		Promotion *Promotion  `gorm:"foreignKey:PromotionID"`
	}

	OrderItem struct {
		BaseModel
		OrderID   string  `gorm:"column:order_id"`
		ProductID string  `gorm:"column:product_id"`
		Quantity  int     `gorm:"column:quantity"`
		Price     float64 `gorm:"column:price"` // price snapshot

		// Relationships
		Product *Product `gorm:"foreignKey:ProductID"`
		Order   *Order   `gorm:"foreignKey:OrderID"`
	}
)

func (o *Order) ToDomain() *domain.Order {
	items := make([]domain.OrderItem, len(o.Items))
	for _, item := range o.Items {
		items = append(items, *item.ToDomain())
	}
	return &domain.Order{
		ID:       o.ID,
		Subtotal: o.Subtotal,
		Discount: o.Discount,
		Total:    o.Total,
		Items:    items,
	}
}

func OrderFromDomain(do domain.Order) *Order {
	items := make([]OrderItem, 0, len(do.Items))
	for _, item := range do.Items {
		items = append(items, *OrderItemFromDomain(item))
	}
	return &Order{
		BaseModel: BaseModel{
			ID: do.ID,
		},
		Subtotal: do.Subtotal,
		Discount: do.Discount,
		Total:    do.Total,
		Items:    items,
	}
}

func (oi *OrderItem) ToDomain() *domain.OrderItem {
	return &domain.OrderItem{
		ID:       oi.ID,
		Product:  oi.Product.ToDomain(),
		Quantity: oi.Quantity,
		Price:    oi.Price,
	}
}

func OrderItemFromDomain(doi domain.OrderItem) *OrderItem {
	return &OrderItem{
		BaseModel: BaseModel{
			ID: doi.ID,
		},
		ProductID: doi.Product.ID,
		Quantity:  doi.Quantity,
		Price:     doi.Price,
		Product:   ProductFromDomain(*doi.Product),
	}
}
