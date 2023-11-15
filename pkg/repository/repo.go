package repository

import (
	"github.com/aparnasukesh/shoezone/pkg/db"
	"github.com/aparnasukesh/shoezone/pkg/domain"
)

func CreateUser(userData *domain.User) error {
	err := db.DB.Create(userData)
	if err != nil {
		return err.Error
	}
	return nil
}
