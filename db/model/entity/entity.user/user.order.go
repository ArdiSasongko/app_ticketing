package entityuser

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"github.com/ArdiSasongko/app_ticketing/db/model/entity"
)

type UserEntityOrder struct {
	UserID int         `json:"user_id"`
	Email  string      `json:"email"`
	Name   string      `json:"name"`
	Orders interface{} `json:"orders"`
}

func ToUserEntityOrder(user domain.Users) UserEntityOrder {
	var orders []entity.OrderEntity

	if len(user.Orders) > 0 {
		for _, v := range user.Orders {
			orders = append(orders, entity.OrderEntity{
				OrderID:    v.OrderID,
				TicketID:   v.TicketID,
				Quantity:   v.Quantity,
				TotalPrice: v.TotalPrice,
				Status:     v.Status,
				ExpiredAt:  v.ExpiredAt,
			})
		}

		return UserEntityOrder{
			UserID: user.UserID,
			Email:  user.Email,
			Name:   user.Name,
			Orders: orders,
		}
	}

	return UserEntityOrder{
		UserID: user.UserID,
		Email:  user.Email,
		Name:   user.Name,
		Orders: "never order before",
	}
}
