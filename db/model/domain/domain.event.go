package domain

import "time"

type Events struct {
	EventID     int        `gorm:"column:event_id;primaryKey;autoIncrement"`
	SellerID    int        `gorm:"column:seller_id"`
	Title       string     `gorm:"column:title"`
	Description string     `gorm:"column:description"`
	Location    string     `gorm:"column:location"`
	StartTime   time.Time  `gorm:"column:start_time"`
	EndTime     time.Time  `gorm:"column:end_time"`
	Tickets     []*Tickets `gorm:"foreignKey:EventID;references:EventID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
