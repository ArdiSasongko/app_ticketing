package orderrepository

import (
	"time"

	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
)

type OrderRepositoryInterface interface {
	Create(order *domain.Orders) (*domain.Orders, error)
	CanceledOrder(curentTime time.Time) error
	GetOrderById(order_id int) (*domain.Orders, error)
	UpdateOrderStatus(orderId int, status string) error
}
