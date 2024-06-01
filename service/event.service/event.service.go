package eventservice

import (
	entityevent "github.com/ArdiSasongko/app_ticketing/db/model/entity/entity.event"
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
)

type EventServiceInterface interface {
	Create(sellerID int, event web.EventRequest) (helper.CustomResponse, error)
	FetchAll() ([]entityevent.EventEntity, error)
}
