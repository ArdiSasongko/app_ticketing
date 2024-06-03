package domain

import "time"

type Orders struct {
	OrderID    int       `gorm:"column:order_id;primaryKey;autoIncrement"`
	BuyerID    int       `gorm:"column:buyer_id"`
	EventID    int       `gorm:"column:event_id"`
	TicketID   int       `gorm:"column:ticket_id"`
	Quantity   int       `gorm:"column:quantity"`
	TotalPrice float64   `gorm:"column:total_price"`
	Status     string    `gorm:"column:status"`
	ExpiredAt  time.Time `gorm:"column:expired_at"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
