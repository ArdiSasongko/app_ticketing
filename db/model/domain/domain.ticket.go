package domain

import "time"

type Tickets struct {
	TicketID  int     `gorm:"column:ticket_id;primaryKey;autoIncrement"`
	EventID   int     `gorm:"column:event_id"`
	Category  string  `gorm:"column:category"`
	Price     float64 `gorm:"column:price"`
	Quantity  int     `gorm:"column:quantity"`
	Event     *Events `gorm:"foreignKey:EventID;references:EventID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
