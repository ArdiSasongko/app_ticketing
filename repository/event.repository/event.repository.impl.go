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

func (r *EventRepo) Create(event *domain.Events) (*domain.Events, error) {
	if err := r.DB.Create(&event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventRepo) FetchAll() ([]domain.Events, error) {
	var events []domain.Events
	if err := r.DB.Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}
