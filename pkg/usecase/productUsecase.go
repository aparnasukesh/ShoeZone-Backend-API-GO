package usecase

import (
	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/repository"
)

func AddProduct(productData *domain.Product) error {
	err := repository.AddProduct(productData)
	if err != nil {
		return err
	}
	return nil
}
