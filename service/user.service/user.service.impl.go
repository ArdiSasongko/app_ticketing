package userservice

import (
	"errors"
	"time"

	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"github.com/ArdiSasongko/app_ticketing/db/model/web"
	"github.com/ArdiSasongko/app_ticketing/helper"
	userrepository "github.com/ArdiSasongko/app_ticketing/repository/user.repository"
	verificationrepository "github.com/ArdiSasongko/app_ticketing/repository/verification.repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo       userrepository.UserRepositoryInterface
	verifyRepo verificationrepository.VerificationEmailInterface
	Token      helper.TokenUseCaseInterface
}

func NewUserService(repo userrepository.UserRepositoryInterface, verifyRepo verificationrepository.VerificationEmailInterface, token helper.TokenUseCaseInterface) *UserService {
	return &UserService{
		repo:       repo,
		verifyRepo: verifyRepo,
		Token:      token,
	}
}

func (service *UserService) Create(user *web.UserRequest) (helper.CustomResponse, error) {
	passHash, errPass := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if errPass != nil {
		return nil, errPass
	}

	newUser := domain.Users{
		Email:    user.Email,
		Name:     user.Name,
		Password: string(passHash),
		Role:     "buyer",
	}

	result, errCreated := service.repo.Create(&newUser)

	if errCreated != nil {
		return nil, errCreated
	}

	verifyToken := domain.EmailVerification{
		UserID:    result.UserID,
		Token:     helper.GeneratedRandomToken(6),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	errVerify := service.verifyRepo.Create(&verifyToken)

	if errVerify != nil {
		return nil, errVerify
	}

	errSend := helper.SendEmail(result.Email, verifyToken.Token)

	if errSend != nil {
		return nil, errSend
	}

	data := helper.CustomResponse{
		"user_id": result.UserID,
		"email":   result.Email,
		"name":    result.Name,
		"role":    result.Role,
	}

	return data, nil
}

func (service *UserService) CreateSeller(user *web.UserRequest) (helper.CustomResponse, error) {
	passHash, errPass := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if errPass != nil {
		return nil, errPass
	}

	newUser := domain.Users{
		Email:    user.Email,
		Name:     user.Name,
		Password: string(passHash),
		Role:     "seller",
	}

	result, errCreated := service.repo.Create(&newUser)

	if errCreated != nil {
		return nil, errCreated
	}

	verifyToken := domain.EmailVerification{
		UserID:    result.UserID,
		Token:     helper.GeneratedRandomToken(6),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	errVerify := service.verifyRepo.Create(&verifyToken)

	if errVerify != nil {
		return nil, errVerify
	}

	errSend := helper.SendEmail(result.Email, verifyToken.Token)

	if errSend != nil {
		return nil, errSend
	}

	data := helper.CustomResponse{
		"user_id": result.UserID,
		"email":   result.Email,
		"name":    result.Name,
		"role":    result.Role,
	}

	return data, nil
}

func (service *UserService) VerifyEmail(token string) (helper.CustomResponse, error) {
	verifyToken, errToken := service.verifyRepo.FindByToken(token)

	if errToken != nil {
		return nil, errors.New("invalid or expired token")
	}

	if verifyToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	user, errUser := service.repo.FindByID(verifyToken.UserID)

	if errUser != nil {
		return nil, errUser
	}

	user.IsVerified = true

	if errUpdate := service.repo.Update(user); errUpdate != nil {
		return nil, errUpdate
	}

	if errDelete := service.verifyRepo.DeleteByToken(verifyToken); errDelete != nil {
		return nil, errDelete
	}

	return helper.CustomResponse{"email": "verified"}, nil
}

func (service *UserService) Login(email, password string) (helper.CustomResponse, error) {
	user, errUser := service.repo.FindByEmail(email)

	if errUser != nil {
		return nil, errUser
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if errPass != nil {
		return nil, errors.New("invalid password")
	}

	expiredTime := time.Now().Local().Add(15 * time.Minute)

	claims := helper.CustomClaims{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			Issuer:    "app-ticketig",
		},
	}

	token, errToken := service.Token.GeneratedToken(claims)

	if errToken != nil {
		return nil, errToken
	}

	data := helper.CustomResponse{
		"token":      token,
		"expired_at": expiredTime,
	}

	return data, nil
}
