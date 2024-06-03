package domain

import "time"

type Payments struct {
	PaymentID     int     `gorm:"column:payment_id;primaryKey;autoIncrement"`
	OrderID       int     `gorm:"column:order_id"`
	Amount        float64 `gorm:"column:amount"`
	PaymentMethod string  `gorm:"column:payment_method"`
	Status        string  `gorm:"column:status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
