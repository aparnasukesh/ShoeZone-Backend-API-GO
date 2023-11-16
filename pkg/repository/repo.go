package repository

import (
	"github.com/aparnasukesh/shoezone/pkg/db"
	"github.com/aparnasukesh/shoezone/pkg/domain"
	"gorm.io/gorm"
)

func CreateUser(userData *domain.User) error {
	err := db.DB.Create(userData)
	if err != nil {
		return err.Error
	}
	return nil
}

func FindUserByEmail(userData *domain.User) (*domain.User, error) {
	result := db.DB.Where("email = ?", userData.Email).First(userData)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return userData, nil
}

func DeleteUserByEmail(userData *domain.User) error {
	result := db.DB.Where("email = ? ", userData.Email).Delete(userData)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
