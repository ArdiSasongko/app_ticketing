package verificationrepository

import "github.com/ArdiSasongko/app_ticketing/db/model/domain"

type VerificationEmailInterface interface {
	Create(token *domain.EmailVerification) error
	FindByToken(token string) (*domain.EmailVerification, error)
	DeleteByToken(token *domain.EmailVerification) error
}
