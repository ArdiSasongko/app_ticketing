package domain

import "time"

type Users struct {
	UserID     int        `gorm:"column:user_id;primaryKey;autoIncrement"`
	Email      string     `gorm:"column:email;unique"`
	Password   string     `gorm:"column:password_hash"`
	Name       string     `gorm:"column:name"`
	Role       string     `gorm:"column:role"`
	IsVerified bool       `gorm:"column:is_verified;default:false"`
	History    []*History `gorm:"foreignKey:user_id;references:user_id"`
	Orders     []*Orders  `gorm:"foreignKey:buyer_id;references:user_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
