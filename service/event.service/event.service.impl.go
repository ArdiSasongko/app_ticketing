package eventservice

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	entityevent "github.com/ArdiSasongko/app_ticketing/db/model/entity/entity.event"
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
	eventrepository "github.com/ArdiSasongko/app_ticketing/repository/event.repository"
)

type EventService struct {
	Repo eventrepository.EventRepoInterface
}

func NewEventService(repo eventrepository.EventRepoInterface) *EventService {
	return &EventService{
		Repo: repo,
	}
}

// Create is a function to create event
func (s *EventService) Create(sellerID int, event web.EventRequest) (helper.CustomResponse, error) {
	// convert web.EventRequest to domain.Event
	eventReq := domain.Events{
		SellerID:    sellerID,
		Title:       event.Title,
		Description: event.Description,
		Location:    event.Location,
		StartTime:   event.StartTime.Time,
		EndTime:     event.EndTime.Time,
	}

	for _, ticket := range event.Tickets {
		eventReq.Tickets = append(eventReq.Tickets, &domain.Tickets{
			Category: ticket.Category,
			Price:    ticket.Price,
			Quantity: ticket.Quantity,
		})
	}
	result, err := s.Repo.Create(&eventReq)

	if err != nil {
		return nil, err
	}

	data := helper.CustomResponse{
		"title":       result.Title,
		"description": result.Description,
		"location":    result.Location,
		"start_time":  result.StartTime,
		"end_time":    result.EndTime,
	}

	return data, nil
}

// FetchAll is a function to fetch all event
func (s *EventService) FetchAll() ([]entityevent.EventEntity, error) {
	result, err := s.Repo.FetchAll()

	if err != nil {
		return []entityevent.EventEntity{}, err
	}

	return entityevent.ToEventEntityList(result), nil
}

// FetchEvent is a function to fetch event
func (s *EventService) FetchEvent(eventId int) (entityevent.EventDetailEntity, error) {
	result, err := s.Repo.FetchEvent(eventId)

	if err != nil {
		return entityevent.EventDetailEntity{}, err
	}

	return entityevent.ToDetailTicket(*result), nil
}

// UpdateEvent is a function to update event
func (s *EventService) UpdateEvent(eventID int, eventUpdate web.EventUpdateRequest) (helper.CustomResponse, error) {
	// Fetch the existing event
	event, err := s.Repo.FetchEvent(eventID)
	if err != nil {
		return nil, err
	}

	// Update the event fields only if they are provided in the request
	if eventUpdate.Title != nil {
		event.Title = *eventUpdate.Title
	}
	if eventUpdate.Description != nil {
		event.Description = *eventUpdate.Description
	}
	if eventUpdate.Location != nil {
		event.Location = *eventUpdate.Location
	}
	if eventUpdate.StartTime != nil {
		event.StartTime = eventUpdate.StartTime.Time
	}
	if eventUpdate.EndTime != nil {
		event.EndTime = eventUpdate.EndTime.Time
	}

	// Save the updated event
	_, errUpdate := s.Repo.UpdateEvent(eventID, *event)
	if errUpdate != nil {
		return nil, errUpdate
	}

	// Optionally update tickets
	for _, ticketReq := range eventUpdate.Tickets {
		ticket, err := s.Repo.FetchTicket(ticketReq.TicketID)
		if err != nil {
			return nil, err
		}

		// Update the ticket fields only if they are provided in the request
		if ticketReq.Category != nil {
			ticket.Category = *ticketReq.Category
		}
		if ticketReq.Price != nil {
			ticket.Price = *ticketReq.Price
		}
		if ticketReq.Quantity != nil {
			ticket.Quantity = *ticketReq.Quantity
		}

		_, err = s.Repo.UpdateTicket(ticketReq.TicketID, *ticket)
		if err != nil {
			return nil, err
		}
	}

	// Prepare the response data
	data := helper.CustomResponse{
		"update": event,
	}

	return data, nil
}

// DeleteEvent is a function to delete event and the related tickets
func (s *EventService) DeleteEvent(eventID int) (helper.CustomResponse, error) {
	err := s.Repo.DeleteEvent(eventID)

	if err != nil {
		return nil, err
	}

	data := helper.CustomResponse{
		"message": "Event deleted successfully",
	}
	return data, nil
}

// DeleteTicket is a function to delete ticket
func (s *EventService) DeleteTicket(ticketID int) (helper.CustomResponse, error) {
	err := s.Repo.DeleteTicket(ticketID)

	if err != nil {
		return nil, err
	}

	data := helper.CustomResponse{
		"message": "Ticket deleted successfully",
	}

	return data, nil
}
