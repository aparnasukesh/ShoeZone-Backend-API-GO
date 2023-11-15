package usecase

import (
	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/repository"
)

func CreateUser(userData *domain.User) error {
	//signup logic

	err := repository.CreateUser(userData)
	if err != nil {
		return err
	}
	return nil
}
