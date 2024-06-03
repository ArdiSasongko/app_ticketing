package orderrepository

import "github.com/ArdiSasongko/app_ticketing/db/model/domain"

type OrderRepositoryInterface interface {
	Create(order *domain.Orders) (*domain.Orders, error)
}
