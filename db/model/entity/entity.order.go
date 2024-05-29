package entity

import "time"

type OrderEntity struct {
	OrderID    int       `json:"order_id"`
	TicketID   int       `json:"ticket_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	ExpiredAt  time.Time `json:"expired_at"`
}
