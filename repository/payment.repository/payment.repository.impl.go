package paymentrepository

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"gorm.io/gorm"
)

type PaymentRepo struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{
		DB: db,
	}
}

func (r *PaymentRepo) Create(payment *domain.Payments) (*domain.Payments, error) {
	if err := r.DB.Create(&payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}

func (r *PaymentRepo) UpdateStatus(payment_id int, status string) error {
	if err := r.DB.Model(&domain.Payments{}).Where("payment_id = ?", payment_id).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
