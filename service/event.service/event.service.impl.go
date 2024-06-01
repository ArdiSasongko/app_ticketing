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

func (s *EventService) Create(sellerID int, event web.EventRequest) (helper.CustomResponse, error) {
	eventReq := domain.Events{
		SellerID:    sellerID,
		Title:       event.Title,
		Description: event.Description,
		Location:    event.Location,
		StartTime:   event.StartTime.Time,
		EndTime:     event.EndTime.Time,
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

func (s *EventService) FetchAll() ([]entityevent.EventEntity, error) {
	result, err := s.Repo.FetchAll()

	if err != nil {
		return []entityevent.EventEntity{}, err
	}

	return entityevent.ToEventEntityList(result), nil
}
