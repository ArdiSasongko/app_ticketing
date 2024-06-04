package web

type PaymentRequest struct {
	OrderID       int     `json:"order_id"`
	Amount        float64 `validate:"required" json:"amount"`
	PaymentMethod string  `validate:"required" json:"payment_method"`
}
