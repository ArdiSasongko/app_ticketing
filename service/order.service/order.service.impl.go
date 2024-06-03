package orderservice

import (
	"time"

	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
	orderrepository "github.com/ArdiSasongko/app_ticketing/repository/order.repository"
)

type OrderService struct {
	repo orderrepository.OrderRepositoryInterface
}

func NewOrderService(repo orderrepository.OrderRepositoryInterface) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) Create(orderReq web.OrderRequest) (helper.CustomResponse, error) {
	// convert orderReq to domain.Orders
	order := domain.Orders{
		BuyerID:    orderReq.BuyerID,
		EventID:    orderReq.EventID,
		TicketID:   orderReq.TicketID,
		Quantity:   orderReq.Quantity,
		TotalPrice: orderReq.TotalPrice,
		ExpiredAt:  time.Now().Add(15 * time.Minute),
		Status:     "pending",
	}

	// create order
	result, err := s.repo.Create(&order)

	if err != nil {
		return nil, err
	}

	data := helper.CustomResponse{
		"order": result,
	}

	return data, nil
}
