package repository

import (
	"errors"
	"time"

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

func UpdateProductStockQuantity(productIDs, quantities []int) error {
	products := []domain.Product{}

	if err := db.DB.Where("id IN (?)", productIDs).Find(&products).Error; err != nil {
		return err
	}

	if len(products) != len(quantities) {
		return errors.New("lengths of productIDs and quantities are different")
	}

	for i := range products {

		products[i].StockQuantity = products[i].StockQuantity - quantities[i]

		if err := db.DB.Save(&products[i]).Error; err != nil {
			return err
		}
	}

	return nil
}

// Cart-----------------------------------------------------------------------------------------------------------

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

func GetCartDetails(userID int) ([]domain.Cart, error) {
	var userCartDetails []domain.Cart
	res := db.DB.Where("user_id = ?", userID).Preload("CartProduct").Find(&userCartDetails)
	if res.Error != nil {
		return nil, res.Error
	}
	if len(userCartDetails) < 1 {
		return nil, errors.New("No Products In The Cart")
	}
	return userCartDetails, nil
}

func DeleteCartItemByUSerID(userId uint) error {
	cartItem := domain.Cart{}
	if err := db.DB.Where("user_id= ? ", userId).Delete(&cartItem).Error; err != nil {
		return err
	}
	return nil
}

// Order-----------------------------------------------------------------------------------------------------------

func CreateOrderItem(orderItems []domain.OrderItem) error {

	for _, orderItem := range orderItems {
		if err := db.DB.Preload("Product").Create(&orderItem).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetOrderItemByUserIdAndOrderId(userId, orderId uint) (uint, error) {

	orderItems := domain.OrderItem{}
	if err := db.DB.Where("user_id=? AND order_id=?", userId, orderId).First(&orderItems).Error; err != nil {
		return 0, err
	}
	return orderItems.ID, nil
}

func Order(order domain.Order) error {

	if err := db.DB.Preload("OrderItems").Preload("Address").Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func ViewOrders(userId int) ([]domain.Order, error) {
	orders := []domain.Order{}

	if err := db.DB.Where("user_id=?", userId).Find(&orders).Error; err != nil {
		return nil, err
	}
	if len(orders) < 1 {
		return nil, errors.New("No Orders Found")
	}
	return orders, nil
}

func ViewOrdersByID(userId, orderId int) ([]domain.Order, []domain.OrderItem, error) {
	orders := []domain.Order{}
	orderItems := []domain.OrderItem{}
	if err := db.DB.Where("user_id=? AND booking_id=?", userId, orderId).Preload("OrderItems").Preload("Address").Find(&orders).Error; err != nil {
		return nil, nil, err
	}
	if err := db.DB.Where("user_id=? AND order_id=?", userId, orderId).Preload("Product").Find(&orderItems).Error; err != nil {
		return nil, nil, err
	}
	if len(orders) < 1 {
		return nil, nil, errors.New("No Items Found")
	}
	return orders, orderItems, nil
}

func ViewOrdersByUserID(id int) ([]domain.Order, error) {
	orders := []domain.Order{}
	if err := db.DB.Where("user_id=?", id).Preload("OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}

	if len(orders) < 1 {
		return nil, errors.New("No Orders found")
	}
	return orders, nil
}

func OrderCancel(userId, orderId int) (*domain.Order, error) {
	orders := domain.Order{}
	if err := db.DB.Where("user_id =? AND booking_id=?", userId, orderId).First(&orders).Error; err != nil {
		return nil, err
	}
	if orders.OrderStatus != "order cancelled" {
		orders.OrderStatus = "order cancelled"
	} else {
		return nil, errors.New("Order already cancelled")
	}
	if err := db.DB.Save(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil

}

func GetProductIDsFromOrderItems(orderItems []domain.OrderItem) ([]uint, []uint) {
	var productIds []uint
	var quantities []uint

	for _, id := range orderItems {
		productIds = append(productIds, id.ProductID)
		quantities = append(quantities, id.Quantity)
	}
	return productIds, quantities
}

func ProductStockUpdationAfterCancellation(productIDs, quantities []uint) error {
	products := []domain.Product{}

	if err := db.DB.Where("id IN (?)", productIDs).Find(&products).Error; err != nil {
		return err
	}

	if len(products) != len(quantities) {
		return errors.New("lengths of productIDs and quantities are different")
	}

	for i := range products {

		products[i].StockQuantity = products[i].StockQuantity + int(quantities[i])

		if err := db.DB.Save(&products[i]).Error; err != nil {
			return err
		}
	}

	return nil
}

// Admin - Coupon----------------------------------------------------------------------------------------------------
func AddCoupon(coupon *domain.Coupon) error {
	if err := db.DB.Create(coupon).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCoupon(id int) error {
	coupon := domain.Coupon{}
	if err := db.DB.Where("id=?", id).Delete(&coupon).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCoupon(updatecoupon domain.Coupon, id int) error {
	if err := db.DB.Model(&updatecoupon).Where("id=?", id).Updates(&updatecoupon).Error; err != nil {
		return err
	}
	return nil
}

func ViewCoupons() ([]domain.Coupon, error) {
	coupon := []domain.Coupon{}
	if err := db.DB.Find(&coupon).Error; err != nil {
		return nil, err
	}
	return coupon, nil
}

func FindCouponByCode(coupon *domain.Coupon) (*domain.Coupon, error) {
	dbcoupon := &domain.Coupon{}

	res := db.DB.Where("code Like ?", coupon.Code).First(&dbcoupon)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return dbcoupon, nil

}

// User - Coupon---------------------------------------------------------------------------------------------------
func GetCouponByCouponName(code string) (*domain.Coupon, error) {
	coupon := &domain.Coupon{}

	if err := db.DB.Where("code Like ?", code).First(&coupon).Error; err != nil {
		return nil, err
	}
	return coupon, nil
}

// User- User_Coupon ---------------------------------------------------------------------------------------------
func CheckCouponUsedByUser(userId int, coupon *domain.Coupon) error {

	userCoupon := &domain.UserCoupon{}
	if err := db.DB.Where("user_id=?  AND coupon_id=? AND used=?", userId, coupon.ID, true).First(&userCoupon).Error; err != nil {
		return err
	}
	return nil
}

func CreateUserCoupon(usercoupon domain.UserCoupon) error {
	if err := db.DB.Create(&usercoupon).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCouponRemainingUses(coupon *domain.Coupon) error {
	coupon.RemainingUses = coupon.RemainingUses - 1
	if err := db.DB.Model(&coupon).Where("id=?", coupon.ID).Updates(&coupon).Error; err != nil {
		return err
	}
	return nil
}

func CreateWalletAmount(data *domain.Wallet, userId int) error {
	res := db.DB.Where("user_id=?", userId).First(&data)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			data.UserID = uint(userId)
			data.LastTransaction = time.Now()
			if err := db.DB.Create(&data).Error; err != nil {
				return err
			}
		} else {
			return res.Error
		}
	} else {
		data.Balance += data.Balance
		result := db.DB.Save(&data)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
