package postgresql

import "github.com/mahdifr17/CheckoutService/internals/domain"

type (
	Promotion struct {
		BaseModel
		Name string `gorm:"column:name"`
		Type string `json:"type"`
	}
)

func (p *Promotion) ToDomain() *domain.Promotion {
	return &domain.Promotion{
		ID:   p.ID,
		Name: p.Name,
		Type: domain.PromotionType(p.Type),
	}
}

func PromotionFromDomain(dp domain.Promotion) *Promotion {
	return &Promotion{
		BaseModel: BaseModel{
			ID: dp.ID,
		},
		Name: dp.Name,
		Type: string(dp.Type),
	}
}
