package userrepository

import "github.com/ArdiSasongko/app_ticketing/db/model/domain"

type UserRepositoryInterface interface {
	Create(user *domain.Users) (*domain.Users, error)
	FindByEmail(email string) (*domain.Users, error)
	FindByID(id int) (*domain.Users, error)
	Update(user *domain.Users) error
	GetOrders(userId int) (*domain.Users, error)
	GetHistory(userId int) (*domain.Users, error)
}
