package paymentservice

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
)

type PaymentServiceInterface interface {
	CreatePayment(payment web.PaymentRequest) (helper.CustomResponse, error)
}
