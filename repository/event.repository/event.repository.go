package eventrepository

import "github.com/ArdiSasongko/app_ticketing/db/model/domain"

type EventRepoInterface interface {
	Create(event *domain.Events) (*domain.Events, error)
	FetchAll() ([]domain.Events, error)
	FetchEvent(eventID int) (*domain.Events, error)
	UpdateEvent(eventID int, updateEvent domain.Events) (*domain.Events, error)
	FetchTicket(ticketID int) (*domain.Tickets, error)
	UpdateTicket(ticketID int, updateTicket domain.Tickets) (*domain.Tickets, error)
	DeleteEvent(eventID int) error
	DeleteTicket(ticketID int) error
	FetchEventBySellerID(sellerID int) ([]domain.Events, error)
	DecreaseTicket(ticketID int, amount int) error
}
