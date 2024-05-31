package userrepository

import (
	"errors"
	"strings"

	"github.com/ArdiSasongko/app_ticketing/db/model/domain"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) Create(user *domain.Users) (*domain.Users, error) {
	if err := repo.DB.Create(user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, errors.New("email already used")
		}
		return nil, err
	}

	return user, nil
}

func (repo *UserRepo) FindByEmail(email string) (*domain.Users, error) {
	var user domain.Users
	if err := repo.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepo) FindByID(id int) (*domain.Users, error) {
	var user domain.Users
	if err := repo.DB.Where("user_id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepo) Update(user *domain.Users) error {
	if err := repo.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}
