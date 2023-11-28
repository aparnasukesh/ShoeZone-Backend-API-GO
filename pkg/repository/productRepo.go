package repository

import (
	"errors"

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

func GetCategories() ([]domain.Category, error) {
	categories := []domain.Category{}
	res := db.DB.Find(&categories)
	if res.Error != nil {
		return nil, res.Error
	}
	return categories, nil
}

func GetCategoriesUser() ([]domain.Category, error) {
	categories := []domain.Category{}
	res := db.DB.Table("categories").Select("category_name").Find(&categories)
	if res.Error != nil {
		return nil, res.Error
	}
	return categories, nil

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

func GetBrands() ([]domain.Brand, error) {
	brands := []domain.Brand{}
	res := db.DB.Find(&brands)
	if res.Error != nil {
		return nil, res.Error
	}
	return brands, nil
}

func GetBrandsUser() ([]domain.Brand, error) {
	brands := []domain.Brand{}
	res := db.DB.Table("brands").Select("brand_name").Find(&brands)
	if res.Error != nil {
		return nil, res.Error
	}
	return brands, nil
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
	product := []domain.Product{}

	if err := db.DB.Preload("ProductCategory").Preload("ProductBrand").Limit(limit).Offset(offset).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductByID(id int) (*domain.Product, error) {
	product := domain.Product{}
	result := db.DB.Preload("ProductCategory").Preload("ProductBrand").First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func GetProductByBrandID(limit, offset, id int) ([]domain.Product, error) {
	product := []domain.Product{}

	if err := db.DB.Preload("ProductCategory").Preload("ProductBrand").Limit(limit).Offset(offset).Where("product_brand_id=?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductByName(name string) (*domain.Product, error) {
	product := domain.Product{}
	if err := db.DB.Preload("ProductCategory").Preload("ProductBrand").Where("product_name ILIKE ?", name).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func GetProductByCategoryID(limit, offset, id int) ([]domain.Product, error) {
	product := []domain.Product{}

	if err := db.DB.Preload("ProductCategory").Preload("ProductBrand").Where("product_category_id=?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetBrandIdByBrandName(name string) (int, error) {
	brand := domain.Brand{}
	res := db.DB.Where("brand_name ILIKE ?", name).First(&brand)
	id := brand.ID
	if res.Error != nil {
		return 0, res.Error
	}
	return int(id), nil
}

func GetProductByBrandName(id int) ([]domain.Product, error) {
	product := []domain.Product{}
	if err := db.DB.Where("product_brand_id=?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func GetCategoryIDByCategoryName(name string) (int, error) {
	category := domain.Category{}
	res := db.DB.Where("category_name ILIKE ?", name).First(&category)
	id := category.ID
	if res.Error != nil {
		return 0, res.Error
	}
	return int(id), nil
}

func GetProductByCategoryName(id int) ([]domain.Product, error) {
	product := []domain.Product{}
	if err := db.DB.Where("product_category_id=?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func CheckProductQuantity(cartProduct *domain.Cart) (bool, error) {
	product := domain.Product{}
	if err := db.DB.Where("id=?", cartProduct.ProductID).First(&product).Error; err != nil {
		return false, err
	}
	if cartProduct.Quantity <= product.StockQuantity {
		return true, nil
	}
	return false, errors.New("Out of stock")
}
func AddToCart(cartProduct *domain.Cart, id int) error {
	res := db.DB.Where("user_id = ? AND product_id = ?", id, cartProduct.ProductID).First(&cartProduct)

	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			cartProduct.UserID = id
			result := db.DB.Create(&cartProduct)
			if result.Error != nil {
				return result.Error
			}
		} else {
			return res.Error
		}
	} else {
		cartProduct.Quantity += cartProduct.Quantity
		result := db.DB.Save(&cartProduct)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func CartList(id int) ([]domain.Cart, error) {
	cartProducts := []domain.Cart{}
	if err := db.DB.Where("user_id=?", id).Find(&cartProducts).Error; err != nil {
		return nil, err
	}
	return cartProducts, nil

}

func DeleteCartItem(id, productID int) error {
	cartProduct := domain.Cart{}
	if err := db.DB.Where("user_id= ? AND product_id=?", id, productID).Delete(&cartProduct).Error; err != nil {
		return err
	}
	return nil
}

func AddAddress(userAdd *domain.Address, id int) error {
	userAdd.UserID = id
	if err := db.DB.Create(&userAdd).Error; err != nil {
		return err
	}
	return nil
}
func GetPricesAndQuantities(productIDs []int, userID int) ([]int, []float64, error) {
	quantities := make([]int, len(productIDs))
	prices := make([]float64, len(productIDs))

	var cartProducts []domain.Cart
	if err := db.DB.Where("user_id=? AND product_id IN (?)", userID, productIDs).Preload("CartProduct").Find(&cartProducts).Error; err != nil {
		return nil, nil, err
	}

	for i, cartProduct := range cartProducts {
		quantities[i] = cartProduct.Quantity
		prices[i] = cartProduct.CartProduct.Price
	}

	return quantities, prices, nil
}

func GetProductIDFromCart(userId int) ([]int, error) {
	cartProduct := []domain.Cart{}
	cartproduct := make([]int, len(cartProduct))
	if err := db.DB.Where("user_id=?", userId).Find(&cartProduct).Error; err != nil {
		return nil, err
	}
	for _, val := range cartProduct {
		cartproduct = append(cartproduct, val.ProductID)
	}
	return cartproduct, nil
}

func GetCartDetails(userID int) ([]domain.Cart, error) {
	var userCartDetails []domain.Cart
	res := db.DB.Where("user_id = ?", userID).Preload("CartProduct").Find(&userCartDetails)
	if res.Error != nil {
		return nil, errors.New("database fetching error")
	}

	return userCartDetails, nil
}
