package entityuser

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"github.com/ArdiSasongko/app_ticketing/db/model/entity"
)

type UserEntitHistory struct {
	UserID  int         `json:"user_id"`
	Email   string      `json:"email"`
	Name    string      `json:"name"`
	History interface{} `json:"history"`
}

func ToUserEntitHistory(user domain.Users) UserEntitHistory {
	var history []entity.HistoryEntity

	if len(user.History) > 0 {
		for _, v := range user.History {
			history = append(history, entity.HistoryEntity{
				HistoryID: v.HistoryID,
				UserID:    v.UserID,
				EventID:   v.EventID,
				Action:    v.Action,
				CreatedAt: v.CreatedAt,
			})
		}

		return UserEntitHistory{
			UserID:  user.UserID,
			Email:   user.Email,
			Name:    user.Name,
			History: history,
		}
	}

	return UserEntitHistory{
		UserID:  user.UserID,
		Email:   user.Email,
		Name:    user.Name,
		History: "never history before",
	}
}
