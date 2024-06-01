package eventrepository

import "github.com/ArdiSasongko/app_ticketing/db/model/domain"

type EventRepoInterface interface {
	Create(event *domain.Events) (*domain.Events, error)
	FetchAll() ([]domain.Events, error)
}
