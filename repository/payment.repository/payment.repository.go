package paymentrepository

import "github.com/ArdiSasongko/app_ticketing/db/model/domain"

type PaymentRepoInterface interface {
	Create(payment *domain.Payments) (*domain.Payments, error)
	UpdateStatus(payment_id int, status string) error
}
