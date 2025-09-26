package postgresql

import "github.com/mahdifr17/CheckoutService/internals/domain"

type (
	Product struct {
		BaseModel
		SKU      string  `gorm:"column:sku;uniqueIndex;not null"`
		Name     string  `gorm:"column:name"`
		Price    float64 `gorm:"column:price"`
		Quantity int     `gorm:"column:quantity"`

		// Relationships
		OrderItems []OrderItem `gorm:"foreignKey:ProductID"`
	}
)

func (p *Product) ToDomain() *domain.Product {
	return &domain.Product{
		ID:       p.ID,
		SKU:      p.SKU,
		Name:     p.Name,
		Price:    p.Price,
		Quantity: p.Quantity,
	}
}

func ProductFromDomain(dp domain.Product) *Product {
	return &Product{
		BaseModel: BaseModel{
			ID: dp.ID,
		},
		SKU:      dp.SKU,
		Name:     dp.Name,
		Price:    dp.Price,
		Quantity: dp.Quantity,
	}
}
