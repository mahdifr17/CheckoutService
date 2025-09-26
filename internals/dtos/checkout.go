package dtos

type (
	CheckoutRequest struct {
		ListItem []CheckoutItem `json:"list_item" validation:"required"`
	}

	CheckoutItem struct {
		ProductID string `json:"product_id" validation:"required"`
		Quantity  int    `json:"quantity" validation:"required"`
	}

	CheckoutResponse struct {
		TotalPrice float64 `json:"total_price"`
	}
)
