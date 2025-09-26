package domain

type (
	PromotionType   string
	PromotionRules  interface{}
	PromotionReward interface{}

	Promotion struct {
		ID     string        `json:"id"`
		Name   string        `json:"name"`
		Type   PromotionType `json:"type"`
		Rules  PromotionRules
		Reward PromotionReward
	}

	BuyNGetMRule struct {
		BuyQty int `json:"buy_qty"`
		GetQty int `json:"get_qty"`
	}

	BuyNGetMReward struct {
		FreeProductID uint `json:"free_product_id"`
		FreeQty       int  `json:"free_qty"`
	}

	// Buy X Pay Y
	BuyXPayYRule struct {
		BuyQty int `json:"buy_qty"`
		PayQty int `json:"pay_qty"`
	}

	BuyXPayYReward struct {
		Discounted bool `json:"discounted"` // just a flag, actual saving is BuyQty - PayQty
	}

	// Threshold Discount
	ThresholdDiscountRule struct {
		MinAmount float64 `json:"min_amount"`
	}

	ThresholdDiscountReward struct {
		DiscountPct float64 `json:"discount_pct"`
	}
)

const (
	BuyNGetM          PromotionType = "BUY_N_GET_M"
	BuyXPayY          PromotionType = "BUY_X_PAY_Y"
	DiscountThreshold PromotionType = "DISCOUNT_THRESHOLD"
)
