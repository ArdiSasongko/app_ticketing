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
}

type EventUpdateRequest struct {
	Location  string `json:"location"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
