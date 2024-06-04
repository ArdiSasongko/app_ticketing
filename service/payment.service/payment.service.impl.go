package paymentservice

import (
	"errors"
	"time"

	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
	eventrepository "github.com/ArdiSasongko/app_ticketing/repository/event.repository"
	historyrepository "github.com/ArdiSasongko/app_ticketing/repository/history.repository"
	orderrepository "github.com/ArdiSasongko/app_ticketing/repository/order.repository"
	paymentrepository "github.com/ArdiSasongko/app_ticketing/repository/payment.repository"
	orderservice "github.com/ArdiSasongko/app_ticketing/service/order.service"
)

type PaymentService struct {
	paymentRepo  paymentrepository.PaymentRepoInterface
	historyRepo  historyrepository.HistoryRepoInterface
	orderRepo    orderrepository.OrderRepositoryInterface
	eventRepo    eventrepository.EventRepoInterface
	orderService orderservice.OrderServiceInterface
}

func NewPaymentService(
	paymentRepo paymentrepository.PaymentRepoInterface,
	historyRepo historyrepository.HistoryRepoInterface,
	orderRepo orderrepository.OrderRepositoryInterface,
	eventRepo eventrepository.EventRepoInterface,
	orderService orderservice.OrderServiceInterface,
) *PaymentService {
	return &PaymentService{
		paymentRepo:  paymentRepo,
		historyRepo:  historyRepo,
		orderRepo:    orderRepo,
		eventRepo:    eventRepo,
		orderService: orderService,
	}
}

func (s *PaymentService) CreatePayment(payment web.PaymentRequest) (helper.CustomResponse, error) {
	// Check order ID
	order, err := s.orderRepo.GetOrderById(payment.OrderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	// check order status
	statusOrder := s.orderService.CheckOrderStatus(payment.OrderID)

	if statusOrder != nil {
		return nil, statusOrder
	}

	// Check if amount is sufficient
	if payment.Amount < order.TotalPrice {
		return nil, errors.New("amount not enough")
	}

	// Create payment
	paymentData := domain.Payments{
		OrderID:       payment.OrderID,
		Amount:        payment.Amount,
		PaymentMethod: payment.PaymentMethod,
		Status:        "pending",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	result, err := s.paymentRepo.Create(&paymentData)
	if err != nil {
		return nil, err
	}

	// Simulate payment process
	time.Sleep(5 * time.Second) // Simulate payment process
	result.Status = "completed"

	errStatus := s.paymentRepo.UpdateStatus(result.PaymentID, result.Status)
	if errStatus != nil {
		return nil, errStatus
	}

	// Update order status
	errOrder := s.orderRepo.UpdateOrderStatus(payment.OrderID, "completed")
	if errOrder != nil {
		return nil, errOrder
	}

	// Update event ticket
	ticket, errTicket := s.eventRepo.FetchTicket(order.TicketID)
	if errTicket != nil {
		return nil, errTicket
	}
	newQuantity := ticket.Quantity - order.Quantity
	if newQuantity < 0 {
		return nil, errors.New("not enough tickets available")
	}

	errUpdateTicket := s.eventRepo.DecreaseTicket(order.TicketID, newQuantity)
	if errUpdateTicket != nil {
		return nil, errUpdateTicket
	}

	// Create history
	historyData := domain.History{
		UserID:    order.BuyerID,
		EventID:   order.EventID,
		Action:    "purchased",
		CreatedAt: time.Now(),
	}

	errHistory := s.historyRepo.Create(&historyData)
	if errHistory != nil {
		return nil, errHistory
	}

	data := helper.CustomResponse{
		"message": "payment success",
	}

	return data, nil
}
