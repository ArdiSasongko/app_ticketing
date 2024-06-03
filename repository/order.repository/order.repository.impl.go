package orderrepository

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		DB: db,
	}
}

// create order
func (r *OrderRepo) Create(order *domain.Orders) (*domain.Orders, error) {
	if err := r.DB.Create(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}
