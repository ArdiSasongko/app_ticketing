package entityevent

import "github.com/ArdiSasongko/app_ticketing/db/model/domain"

type EventEntity struct {
	EventID  int    `json:"event_id"`
	SellerID int    `json:"seller_id"`
	Title    string `json:"title"`
}

type EventDetailEntity struct {
	EventID     int    `json:"event_id"`
	SellerID    int    `json:"seller_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

// menampilakn semua data event yang ada
func ToEventEntity(event domain.Events) EventEntity {
	return EventEntity{
		EventID:  event.EventID,
		SellerID: event.SellerID,
		Title:    event.Title,
	}
}

func ToEventEntityList(event []domain.Events) []EventEntity {
	var result []EventEntity
	for _, v := range event {
		result = append(result, ToEventEntity(v))
	}
	return result
}
