package domain

import "time"

type History struct {
	HistoryID int    `gorm:"column:history_id;primaryKey;autoIncrement"`
	UserID    int    `gorm:"column:user_id"`
	EventID   int    `gorm:"column:event_id"`
	Action    string `gorm:"column:action"`
	CreatedAt time.Time
}
