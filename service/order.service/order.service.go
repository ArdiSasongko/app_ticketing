package orderservice

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
)

type OrderServiceInterface interface {
	Create(orderReq web.OrderRequest) (helper.CustomResponse, error)
}
