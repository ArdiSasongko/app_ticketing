package helper

import (
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUseCaseInterface interface {
	GeneratedToken(claims CustomClaims) (string, error)
	InvalidToken(token string) error
	IsTokenInvalid(token string) bool
}

type TokenUseCaseImpl struct {
	blacklist map[string]time.Time
	mu        sync.Mutex
}

type CustomClaims struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewTokenUseCase() *TokenUseCaseImpl {
	return &TokenUseCaseImpl{
		blacklist: make(map[string]time.Time),
	}
}

func (t *TokenUseCaseImpl) GeneratedToken(claims CustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, errToken := plainToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if errToken != nil {
		return "", errToken
	}

	return token, nil
}

func (t *TokenUseCaseImpl) InvalidToken(token string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	expirateTime := time.Now().Add(time.Hour * 24)
	t.blacklist[token] = expirateTime

	return nil
}

func (t *TokenUseCaseImpl) IsTokenInvalid(token string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	expireTime, exists := t.blacklist[token]

	if !exists {
		return false
	}

	if time.Now().After(expireTime) {
		delete(t.blacklist, token)
		return false
	}

	return true
}
