package domain

type (
	Product struct {
		ID       string  `json:"id"`
		SKU      string  `json:"sku"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int     `json:"quantity"`
	}
)
