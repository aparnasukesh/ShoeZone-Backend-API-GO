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

func AddCategories(categoryData *domain.Category) error {
	err := repository.AddCategories(categoryData)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(id int) {

}
