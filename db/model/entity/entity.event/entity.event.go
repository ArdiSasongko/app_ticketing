package entityevent

import "github.com/ArdiSasongko/app_ticketing/db/model/domain"

type TicketEntity struct {
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type EventEntity struct {
	EventID  int    `json:"event_id"`
	SellerID int    `json:"seller_id"`
	Title    string `json:"title"`
}

type EventDetailEntity struct {
	EventID     int         `json:"event_id"`
	SellerID    int         `json:"seller_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Location    string      `json:"location"`
	StartTime   string      `json:"start_time"`
	EndTime     string      `json:"end_time"`
	Tickets     interface{} `json:"tickets"`
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

func ToDetailTicket(event domain.Events) EventDetailEntity {
	var tickets []TicketEntity

	if len(event.Tickets) > 0 {
		for _, v := range event.Tickets {
			tickets = append(tickets, TicketEntity{
				Category: v.Category,
				Price:    v.Price,
				Quantity: v.Quantity,
			})
		}

		return EventDetailEntity{
			EventID:     event.EventID,
			SellerID:    event.SellerID,
			Title:       event.Title,
			Description: event.Description,
			Location:    event.Location,
			StartTime:   event.StartTime.String(),
			EndTime:     event.EndTime.String(),
			Tickets:     tickets,
		}
	}

	return EventDetailEntity{
		EventID:     event.EventID,
		SellerID:    event.SellerID,
		Title:       event.Title,
		Description: event.Description,
		Location:    event.Location,
		StartTime:   event.StartTime.String(),
		EndTime:     event.EndTime.String(),
		Tickets:     "ticket not available",
	}
}
