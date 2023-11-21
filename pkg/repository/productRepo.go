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

func AddCategories(categoryData *domain.Category) error {
	err := db.DB.Create(&categoryData)
	if err != nil {
		return err.Error
	}
	return nil
}

func UpdateCategory(id int) (domain.Category, error) {

	updatedData := domain.Category{}

	err := db.DB.Where("id ?", id).First(&updatedData)
}
