package historyrepository

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"gorm.io/gorm"
)

type HistoryRepo struct {
	DB *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepo {
	return &HistoryRepo{
		DB: db,
	}
}

func (r *HistoryRepo) Create(history *domain.History) error {
	return r.DB.Create(&history).Error
}
