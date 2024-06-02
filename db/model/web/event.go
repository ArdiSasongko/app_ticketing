package web

import (
	"github.com/ArdiSasongko/app_ticketing/helper"
)

type EventRequest struct {
	Title       string            `validate:"required" json:"title"`
	Description string            `validate:"required" json:"description"`
	Location    string            `validate:"required" json:"location"`
	StartTime   helper.CustomTime `validate:"required" json:"start_time"`
	EndTime     helper.CustomTime `validate:"required" json:"end_time"`
	Tickets     []TicketRequest   `validate:"required" json:"tickets"`
}

type TicketRequest struct {
	Category string  `validate:"required" json:"category"`
	Price    float64 `validate:"required,gt=0" json:"price"`
	Quantity int     `validate:"required,gte=0" json:"quantity"`
}

type EventUpdateRequest struct {
	Title       *string               `json:"title"`
	Description *string               `json:"description"`
	Location    *string               `json:"location"`
	StartTime   *helper.CustomTime    `json:"start_time"`
	EndTime     *helper.CustomTime    `json:"end_time"`
	Tickets     []TicketUpdateRequest `json:"tickets"`
}

type TicketUpdateRequest struct {
	TicketID int      `json:"ticket_id"`
	Category *string  `json:"category"`
	Price    *float64 `json:"price"`
	Quantity *int     `json:"quantity"`
}
