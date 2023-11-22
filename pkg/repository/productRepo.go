package repository

import (
	"github.com/aparnasukesh/shoezone/pkg/db"
	"github.com/aparnasukesh/shoezone/pkg/domain"
	"gorm.io/gorm"
)

// Category -----------------------------------------------------------------------------------------------------------------------------

func AddCategories(categoryData *domain.Category) error {
	err := db.DB.Create(&categoryData)
	if err != nil {
		return err.Error
	}
	return nil
}

func UpdateCategory(updateCategory domain.Category, id int) error {
	if err := db.DB.Model(&updateCategory).Where("id = ?", id).Update("category_name", updateCategory.CategoryName).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id int) error {
	categoryData := domain.Category{}

	err := db.DB.Where("id=?", id).Delete(&categoryData)
	if err != nil {
		return err.Error
	}
	return nil
}

// Brand----------------------------------------------------------------------------------------------------------------------------------

func AddBrand(newBrand domain.Brand) error {
	err := db.DB.Create(&newBrand)
	if err != nil {
		return err.Error
	}
	return nil
}

func UpdateBrand(updatedBrand domain.Brand, id int) error {

	if err := db.DB.Model(&updatedBrand).Where("id = ?", id).Update("brand_name", updatedBrand.BrandName).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBrand(id int) error {
	brandData := domain.Brand{}
	if err := db.DB.Delete(&brandData, id); err != nil {
		return err.Error
	}
	return nil
}

// Product --------------------------------------------------------------------------------------------------------------------------------

func FindProductByProductName(newProduct *domain.Product) (*domain.Product, error) {
	productData := &domain.Product{}
	result := db.DB.Where("product_name ILIKE ?", newProduct.ProductName).First(productData)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return productData, nil

}
func AddProduct(productData *domain.Product) error {
	err := db.DB.Create(&productData)
	if err != nil {
		return err.Error
	}
	return nil
}

func UpdateProduct(updateProduct domain.Product, id int) error {
	if err := db.DB.Model(&updateProduct).Where("id=?", id).Updates(&updateProduct).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int) error {
	productData := domain.Product{}
	result := db.DB.Delete(&productData, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetProducts(limit, offset int) ([]domain.Product, error) {
	products := []domain.Product{}
	if err := db.DB.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
