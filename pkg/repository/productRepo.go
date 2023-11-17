package repository

import (
	"github.com/aparnasukesh/shoezone/pkg/db"
	"github.com/aparnasukesh/shoezone/pkg/domain"
)

func AddProduct(productData *domain.Product) error {
	err := db.DB.Create(&productData)
	if err != nil {
		return err.Error
	}
	return nil
}
