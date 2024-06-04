package orderrepository

import (
	"time"

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

func (r *OrderRepo) CanceledOrder(curentTime time.Time) error {
	return r.DB.Model(&domain.Orders{}).Where("expires_at <= ? AND status = ?", curentTime, "pending").
		Update("status", "canceled").Error
}

func (r *OrderRepo) GetOrderById(order_id int) (*domain.Orders, error) {
	var order domain.Orders
	if err := r.DB.First(&order, order_id).Error; err != nil {
		return &domain.Orders{}, err
	}

	return &order, nil
}

func (r *OrderRepo) UpdateOrderStatus(orderId int, status string) error {
	return r.DB.Model(&domain.Orders{}).Where("order_id = ?", orderId).Update("status", status).Error
}
