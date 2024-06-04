package orderservice

import (
	"time"

	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
	eventrepository "github.com/ArdiSasongko/app_ticketing/repository/event.repository"
	orderrepository "github.com/ArdiSasongko/app_ticketing/repository/order.repository"
)

type OrderService struct {
	repo      orderrepository.OrderRepositoryInterface
	eventRepo eventrepository.EventRepoInterface
}

func NewOrderService(repo orderrepository.OrderRepositoryInterface, eventRepo eventrepository.EventRepoInterface) *OrderService {
	return &OrderService{
		repo:      repo,
		eventRepo: eventRepo,
	}
}

func (s *OrderService) Create(orderReq web.OrderRequest) (helper.CustomResponse, error) {
	// get ticket price
	ticket, err := s.eventRepo.FetchTicket(orderReq.TicketID)

	if err != nil {
		return nil, err
	}
	// counting ticket price
	totalPrice := float64(orderReq.Quantity) * ticket.Price

	// convert orderReq to domain.Orders
	order := domain.Orders{
		BuyerID:    orderReq.BuyerID,
		EventID:    orderReq.EventID,
		TicketID:   orderReq.TicketID,
		Quantity:   orderReq.Quantity,
		TotalPrice: totalPrice,
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

// cancelled order automatically
func (s *OrderService) CanceledOrder() error {
	currentTime := time.Now()
	return s.repo.CanceledOrder(currentTime)
}
