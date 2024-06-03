package eventservice

import (
	entityevent "github.com/ArdiSasongko/app_ticketing/db/model/entity/entity.event"
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
)

type EventServiceInterface interface {
	Create(sellerID int, event web.EventRequest) (helper.CustomResponse, error)
	FetchAll() ([]entityevent.EventEntity, error)
	FetchEvent(eventId int) (entityevent.EventDetailEntity, error)
	UpdateEvent(eventID int, eventUpdate web.EventUpdateRequest) (helper.CustomResponse, error)
	DeleteEvent(eventID int) (helper.CustomResponse, error)
	DeleteTicket(ticketID int) (helper.CustomResponse, error)
}
