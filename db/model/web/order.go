package web

type OrderRequest struct {
	EventID    int     `json:"event_id"`
	TicketID   int     `json:"ticket_id"`
	BuyerID    int     `json:"buyer_id"`
	Quantity   int     `validate:"required,min=1" json:"quantity"`
	TotalPrice float64 `validate:"required,min=0" json:"total_price"`
}
