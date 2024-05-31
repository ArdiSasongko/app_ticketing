package domain

import "time"

type EmailVerification struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    int       `gorm:"column:user_id"`
	Token     string    `gorm:"column:token;unique"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	ExpiresAt time.Time `gorm:"column:expires_at"`
}

func (EmailVerification) TableName() string {
	return "email_verifications"
}
