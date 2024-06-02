package eventrepository

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"gorm.io/gorm"
)

type EventRepo struct {
	DB *gorm.DB
}

func NewEventRepo(db *gorm.DB) *EventRepo {
	return &EventRepo{
		DB: db,
	}
}

// for creating event
func (r *EventRepo) Create(event *domain.Events) (*domain.Events, error) {
	if err := r.DB.Create(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

// for fetching all event
func (r *EventRepo) FetchAll() ([]domain.Events, error) {
	var events []domain.Events
	if err := r.DB.Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

// for fetching event
func (r *EventRepo) FetchEvent(eventID int) (*domain.Events, error) {
	var event domain.Events
	if err := r.DB.Preload("Tickets").First(&event, eventID).Error; err != nil {
		return &domain.Events{}, err
	}

	return &event, nil
}

// for updating event
func (r *EventRepo) UpdateEvent(eventID int, updateEvent domain.Events) (*domain.Events, error) {
	if err := r.DB.Model(&domain.Events{}).Where("event_id = ?", eventID).Updates(&updateEvent).Error; err != nil {
		return nil, err
	}

	return &updateEvent, nil
}

// for fetching ticket
func (r *EventRepo) FetchTicket(ticketID int) (*domain.Tickets, error) {
	var ticket domain.Tickets
	if err := r.DB.First(&ticket, ticketID).Error; err != nil {
		return &domain.Tickets{}, err
	}

	return &ticket, nil
}

// for updating ticket
func (r *EventRepo) UpdateTicket(ticketID int, updateTicket domain.Tickets) (*domain.Tickets, error) {
	if err := r.DB.Model(&domain.Tickets{}).Where("ticket_id = ?", ticketID).Updates(&updateTicket).Error; err != nil {
		return nil, err
	}

	return &updateTicket, nil
}
