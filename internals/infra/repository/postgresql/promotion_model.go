package postgresql

import (
	"encoding/json"

	"github.com/mahdifr17/CheckoutService/internals/domain"
)

type (
	Promotion struct {
		BaseModel
		Name   string         `gorm:"column:name"`
		Type   string         `gorm:"column:type"`
		Rules  map[string]any `gorm:"column:rules;type:jsonb;not null"`
		Reward map[string]any `gorm:"column:reward;type:jsonb;not null"`
	}
)

func (p *Promotion) ToDomain() *domain.Promotion {
	var rules domain.PromotionRules
	var reward domain.PromotionReward

	switch domain.PromotionType(p.Type) {
	case domain.BuyNGetM:
		var r domain.BuyNGetMRule
		var rw domain.BuyNGetMReward
		json.Unmarshal([]byte(toJSON(p.Rules)), &r)
		json.Unmarshal([]byte(toJSON(p.Reward)), &rw)
		rules, reward = r, rw

	case domain.BuyXPayY:
		var r domain.BuyXPayYRule
		var rw domain.BuyXPayYReward
		json.Unmarshal([]byte(toJSON(p.Rules)), &r)
		json.Unmarshal([]byte(toJSON(p.Reward)), &rw)
		rules, reward = r, rw

	case domain.DiscountThreshold:
		var r domain.ThresholdDiscountRule
		var rw domain.ThresholdDiscountReward
		json.Unmarshal([]byte(toJSON(p.Rules)), &r)
		json.Unmarshal([]byte(toJSON(p.Reward)), &rw)
		rules, reward = r, rw
	}

	return &domain.Promotion{
		ID:     p.ID,
		Name:   p.Name,
		Type:   domain.PromotionType(p.Type),
		Rules:  rules,
		Reward: reward,
	}
}

// helper to marshal JSONB into string
func toJSON(j map[string]any) string {
	b, _ := json.Marshal(j)
	return string(b)
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
