package domain

type (
	PromotionType string

	Promotion struct {
		ID   string        `json:"id"`
		Name string        `json:"name"`
		Type PromotionType `json:"type"`
	}
)

const (
	BuyNGetM          PromotionType = "BUY_N_GET_M"
	BuyXPayY          PromotionType = "BUY_X_PAY_Y"
	DiscountThreshold PromotionType = "DISCOUNT_THRESHOLD"
)
