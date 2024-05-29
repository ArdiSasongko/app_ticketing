package entity

import "time"

type HistoryEntity struct {
	HistoryID int    `json:"history_id"`
	UserID    int    `json:"user_id"`
	EventID   int    `json:"event_id"`
	Action    string `json:"action"`
	CreatedAt time.Time
}
