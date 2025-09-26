package domain

type (
	Order struct {
		ID       string      `json:"id"`
		Subtotal float64     `json:"subtotal"`
		Discount float64     `json:"discount"`
		Total    float64     `json:"total"`
		Items    []OrderItem `json:"items"`
	}

	OrderItem struct {
		ID       string   `json:"id"`
		Product  *Product `json:"product"`
		Quantity int      `json:"quantity"`
		Price    float64  `json:"price"`
	}
)
