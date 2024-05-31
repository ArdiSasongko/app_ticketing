package verificationrepository

import (
	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"gorm.io/gorm"
)

type EmailVerification struct {
	DB *gorm.DB
}

func NewEmailVerification(db *gorm.DB) *EmailVerification {
	return &EmailVerification{
		DB: db,
	}
}

func (r *EmailVerification) Create(token *domain.EmailVerification) error {
	return r.DB.Create(token).Error
}

func (r *EmailVerification) FindByToken(token string) (*domain.EmailVerification, error) {
	var emailVerification domain.EmailVerification
	err := r.DB.Where("token = ?", token).First(&emailVerification).Error
	if err != nil {
		return nil, err
	}
	return &emailVerification, nil
}

func (r *EmailVerification) DeleteByToken(token *domain.EmailVerification) error {
	return r.DB.Delete(token).Error
}
