package usecase

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/repository"
	"github.com/aparnasukesh/shoezone/pkg/util"
)

// Category Management----------------------------------------------------------------------------

func AddCategories(categoryData *domain.Category) error {
	err := repository.AddCategories(categoryData)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(updateCategory domain.Category, id int) error {
	err := repository.UpdateCategory(updateCategory, id)
	if err != nil {
		return err
	}
	return nil

}

func DeleteCategory(id int) error {
	err := repository.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}

func GetCategories() ([]domain.Category, error) {
	categories, err := repository.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoriesUser() ([]domain.Category, error) {
	categories, err := repository.GetCategoriesUser()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// Brand------------------------------------------------------------------------------------------

func AddBrand(newBrand domain.Brand) error {
	err := repository.AddBrand(newBrand)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBrand(updatedBrand domain.Brand, id int) error {
	err := repository.UpdateBrand(updatedBrand, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBrand(id int) error {
	if err := repository.DeleteBrand(id); err != nil {
		return err
	}
	return nil
}

func GetBrands() ([]domain.Brand, error) {
	brands, err := repository.GetBrands()
	if err != nil {
		return nil, err
	}
	return brands, nil

}

func GetBrandsUser() ([]domain.Brand, error) {
	brands, err := repository.GetBrandsUser()
	if err != nil {
		return nil, err
	}
	return brands, nil
}

// Product Management------------------------------------------------------------------------------

func AddProduct(productData *domain.Product) error {
	res, err := repository.FindProductByProductName(productData)
	if err != nil && res == nil {
		err = repository.AddProduct(productData)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Product already existed")
}

func UpdateProduct(updateProduct domain.Product, id int) error {
	err := repository.UpdateProduct(updateProduct, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(id int) error {
	if err := repository.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}

func ProductImageUpload(filename, uploadPath string, id int) error {
	err := repository.ProductImageUpload(filename, uploadPath, id)
	if err != nil {
		return err
	}
	return nil
}

func ProductImageViewByID(id int) (string, error) {
	path, err := repository.ProductImageViewByID(id)
	if err != nil {
		return "", err
	}
	return path, nil
}

// User- Products---------------------------------------------------------------------

func GetProducts(limit, offset int) ([]domain.Product, error) {
	products, err := repository.GetProducts(limit, offset)
	if err != nil {
		return nil, err
	}
	return products, err
}

func GetProductByID(id int) (*domain.Product, error) {
	product, err := repository.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductByName(name string) (*domain.Product, error) {
	product, err := repository.GetProductByName(name)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductByBrandID(limit, offset, id int) ([]domain.Product, error) {
	product, err := repository.GetProductByBrandID(limit, offset, id)
	if err != nil {
		return nil, err
	}
	if len(product) == 0 {
		return nil, errors.New("No Products")
	}
	return product, nil
}

func GetProductCategoryID(limit, offset, id int) ([]domain.Product, error) {
	product, err := repository.GetProductByCategoryID(limit, offset, id)
	if err != nil {
		return nil, err
	}
	if len(product) == 0 {
		return nil, errors.New("No Products")
	}
	return product, nil
}

func GetProductByBrandName(name string) ([]domain.Product, error) {
	id, _ := repository.GetBrandIdByBrandName(name)

	if id == 0 {
		return nil, errors.New("Brand not exist")
	}
	products, err := repository.GetProductByBrandName(id)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("No Products")
	}
	return products, nil
}

func GetProductByCategoryName(name string) ([]domain.Product, error) {
	id, _ := repository.GetCategoryIDByCategoryName(name)
	if id == 0 {
		return nil, errors.New("Category not exist")
	}
	product, err := repository.GetProductByCategoryName(id)
	if err != nil {
		return nil, err
	}
	if len(product) == 0 {
		return nil, errors.New("No Products")
	}
	return product, nil
}

// Get- User Id from token--------------------------------------------------------------------------------------

func GetUserIDFromToken(authorization string) (int, error) {
	tokenParts := strings.Split(authorization, "Bearer ")
	token, err := util.VerifyJWT(tokenParts[1])
	if err != nil {
		return 0, err
	}
	userID, err := util.GetUserID(token)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// User - Cart ---------------------------------------------------------------------------------------------------

func AddToCart(cartProduct domain.Cart, id int) error {
	res, err := repository.CheckProductQuantity(cartProduct)
	if res == false || err != nil {
		return err
	}
	err = repository.AddToCart(cartProduct, id)
	if err != nil {
		return err
	}
	return nil
}

func CartList(id int) ([]domain.CartResponse, error) {
	resPonse := []domain.CartResponse{}
	cartProducts, err := repository.CartList(id)
	if err != nil {
		return nil, err
	}

	util.CreateCartResponse(&resPonse, cartProducts)
	return resPonse, nil
}

func DeleteCartItem(id, productID int) error {
	err := repository.DeleteCartItem(id, productID)
	if err != nil {
		return err
	}
	return nil
}

// User - Wish List------------------------------------------------------------------------------------------------
func AddToWishList(userId, productId int) error {
	err := repository.AddToWishList(userId, productId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteWishlistItem(userId, productId int) error {
	err := repository.DeleteWishlistItem(userId, productId)
	if err != nil {
		return err
	}
	return nil
}

func WishListItems(userId int) ([]domain.WishListResponse, error) {
	products, err := repository.WishListItems(userId)
	if err != nil {
		return nil, err
	}
	response, err := util.BuildWishListResponse(products)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// User - Order - Cash on Delivery ----------------------------------------------------------------------------------------------------

func GetCartItemsOrderSummary(userID int) (domain.CartItemsOrderSummary, error) {
	userCartDetails, err := repository.GetCartDetails(userID)
	if err != nil {
		return domain.CartItemsOrderSummary{}, err
	}

	orderSummary := util.BuildCartItemsOrderSummary(userCartDetails)
	orderSummary.UserID = userID

	return orderSummary, nil
}

func OrderCartItems(userId, addressId int, coupon string) error {
	userCartDetails, err := repository.GetCartDetails(userId)
	if err != nil {
		return err
	}

	orderItem, orderId, err := util.BuildOrderCartItems(userCartDetails, userId)
	if err != nil {
		return err
	}

	err = repository.CreateOrderCartItems(orderItem)
	if err != nil {
		return err
	}

	user, err := repository.GetUserByID(userId)
	if err != nil {
		return err
	}

	orderID, err := repository.GetOrderItemByUserIdAndOrderId(uint(userId), orderId)
	if err != nil {
		return nil
	}

	couponData := &domain.Coupon{}
	usercoupon := domain.UserCoupon{}
	validCoupon := &domain.Coupon{}

	if coupon != "" {
		couponData, err = repository.GetCouponByCouponName(coupon)
	}
	if couponData != nil && err == nil {
		validCoupon, err = util.CouponValidate(couponData)
		if err != nil {
			return err
		}
	}

	err = repository.CheckCouponUsedByUser(userId, validCoupon)
	if err != nil {
		usercoupon = util.BuildUserCoupon(userId, *validCoupon)
		err = repository.CreateUserCoupon(usercoupon)
		if err != nil {
			return err
		}
		err = repository.UpdateCouponRemainingUses(validCoupon)
		if err != nil {
			return err
		}
	} else {
		validCoupon.Code = ""
		validCoupon.DiscountPercentage = 0
	}

	err = repository.CheckValidAddress(userId, addressId)
	if err != nil {
		return errors.New("Enter a valid address or enter a new address")
	}

	order := util.BuildOrder(orderItem, *user, orderID, orderId, *validCoupon, addressId)
	err = repository.Order(order)
	if err != nil {
		return err
	}

	productIDs, quantities := util.GetProductIDsFromCart(userCartDetails)
	err = repository.UpdateProductStockQuantity(productIDs, quantities)
	if err != nil {
		return err
	}

	err = repository.DeleteCartItemByUserID(uint(userId))
	if err != nil {
		return err
	}

	for i := 0; i < len(productIDs); i++ {
		repository.DeleteWishlistItem(userId, productIDs[i])
	}
	return nil

}

func OrderItemByID(userId, productId, quantity int, coupon string, addressId int) error {
	productDetails, err := repository.GetProductByID(productId)
	if err != nil {
		return err
	}
	if productDetails.StockQuantity < quantity {
		return errors.New("Product out of stock")
	}
	orderItem, orderId, err := util.BuildOrderItemByID(userId, quantity, *productDetails)
	if err != nil {
		return err
	}
	err = repository.CreateOrderItems(*orderItem)
	if err != nil {
		return err
	}
	user, err := repository.GetUserByID(userId)
	if err != nil {
		return err
	}

	orderID, err := repository.GetOrderItemByUserIdAndOrderId(uint(userId), orderId)
	if err != nil {
		return nil
	}

	couponData := &domain.Coupon{}
	usercoupon := domain.UserCoupon{}
	validCoupon := &domain.Coupon{}

	if coupon != "" {
		couponData, err = repository.GetCouponByCouponName(coupon)
	}
	if couponData != nil && err == nil {
		validCoupon, err = util.CouponValidate(couponData)
		if err != nil {
			return err
		}
	}

	err = repository.CheckCouponUsedByUser(userId, validCoupon)
	if err != nil {
		usercoupon = util.BuildUserCoupon(userId, *validCoupon)
		err = repository.CreateUserCoupon(usercoupon)
		if err != nil {
			return err
		}
		err = repository.UpdateCouponRemainingUses(validCoupon)
		if err != nil {
			return err
		}
	} else {
		validCoupon.Code = ""
		validCoupon.DiscountPercentage = 0
	}

	err = repository.CheckValidAddress(userId, addressId)
	if err != nil {
		return errors.New("Enter a valid address or enter a new address")
	}

	order := util.BuildOrderbyProductID(orderItem, *user, orderID, orderId, *validCoupon, addressId)
	err = repository.CheckCartItemByUserIdAndProductId(userId, productId)
	if err == nil {
		err := repository.DeleteCartItem(userId, productId)
		if err != nil {
			return err
		}
	}
	err = repository.Order(order)
	if err != nil {
		return err
	}
	err = repository.UpdateProductStock(productId, quantity)
	if err != nil {
		return err
	}
	err = repository.CheckItemPresentInWishList(userId, productId)
	if err == nil {
		err = repository.DeleteWishlistItem(userId, productId)
		if err != nil {
			return err
		}
	}

	return nil

}

// User - Order - Razorpay------------------------------------------------------------------------------------------
func OrderCartItemsRazorpay(userId int, coupon string, addressId int) (*domain.RazorPay, error) {
	userCartDetails, err := repository.GetCartDetails(userId)
	if err != nil {
		return nil, err
	}

	orderItem, orderId, err := util.BuildOrderCartItems(userCartDetails, userId)
	if err != nil {
		return nil, err
	}

	err = repository.CreateOrderCartItems(orderItem)
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserByID(userId)
	if err != nil {
		return nil, err
	}
	orderItemID, err := repository.GetOrderItemByUserIdAndOrderId(uint(userId), orderId)
	if err != nil {
		return nil, err
	}

	couponData := &domain.Coupon{}
	validCoupon := &domain.Coupon{}

	if coupon != "" {
		couponData, err = repository.GetCouponByCouponName(coupon)
	}
	if couponData != nil && err == nil {
		validCoupon, err = util.CouponValidate(couponData)
		if err != nil {
			return nil, err
		}
	}

	err = repository.CheckCouponUsedByUser(userId, validCoupon)
	if err != nil {
		err = repository.UpdateCouponRemainingUses(validCoupon)
		if err != nil {
			return nil, err
		}
	} else {
		validCoupon.Code = ""
		validCoupon.DiscountPercentage = 0
	}

	err = repository.CheckValidAddress(userId, addressId)
	if err != nil {
		return nil, errors.New("Enter a valid address or enter a new address")
	}

	order := util.BuildOrderRazorpay(orderItem, *user, orderItemID, orderId, *validCoupon, addressId)
	paymentDetails := domain.RazorPay{
		UserID:        userId,
		AddressID:     addressId,
		Order_TableID: int(orderId),
		Coupon:        coupon,
		TotalAmount:   order.AmountPayable,
	}
	res, err := RazorPay(paymentDetails)
	if err != nil {
		return nil, err
	}
	if err := repository.CreateRazorpayPayment(*res); err != nil {
		return nil, err
	}
	return res, nil

}

func RazorpaySuccess(userId, order_TableId int, signature, paymentid, orderid, coupon string, addressID int) error {
	err := repository.UpdateRazorpay(userId, signature, paymentid, orderid)
	if err != nil {
		return err
	}
	userCartDetails, err := repository.GetCartDetails(userId)
	if err != nil {
		return err
	}
	orderItem, orderId, err := repository.OrderItemsByUserIDandOrderTableID(userId, order_TableId)
	if err != nil {
		return err
	}
	user, err := repository.GetUserByID(userId)
	if err != nil {
		return err
	}

	orderID, err := repository.GetOrderItemByUserIdAndOrderId(uint(userId), uint(orderId))
	if err != nil {
		return nil
	}

	couponData := &domain.Coupon{}
	usercoupon := domain.UserCoupon{}
	validCoupon := &domain.Coupon{}

	if coupon == "No Coupon Applied" {
		validCoupon.Code = ""
		validCoupon.DiscountPercentage = 0
	} else {
		couponData, err = repository.GetCouponByCouponName(coupon)
		if couponData != nil && err == nil {
			validCoupon, err = util.CouponValidate(couponData)
			if err != nil {
				return err
			}
		}

		err = repository.CheckCouponUsedByUser(userId, validCoupon)
		if err != nil {
			usercoupon = util.BuildUserCoupon(userId, *validCoupon)
			err = repository.CreateUserCoupon(usercoupon)
			if err != nil {
				return err
			}
			err = repository.UpdateCouponRemainingUses(validCoupon)
			if err != nil {
				return err
			}
		} else {
			validCoupon.Code = ""
			validCoupon.DiscountPercentage = 0
		}
	}

	order := util.BuildOrderRazorpay(orderItem, *user, orderID, uint(orderId), *validCoupon, addressID)
	err = repository.Order(order)
	if err != nil {
		return err
	}

	productIDs, quantities := util.GetProductIDsFromCart(userCartDetails)
	err = repository.UpdateProductStockQuantity(productIDs, quantities)
	if err != nil {
		return err
	}

	err = repository.DeleteCartItemByUserID(uint(userId))
	if err != nil {
		return err
	}

	for i := 0; i < len(productIDs); i++ {

		err = repository.CheckItemPresentInWishList(userId, productIDs[i])
		if err == nil {
			repository.DeleteWishlistItem(userId, productIDs[i])
		}

	}
	return nil

}

// User - Order - Wallet Payment----------------------------------------------------------------------------------
func WalletPaymentCartItems(userId int, coupon string, addressId int) error {
	userCartDetails, err := repository.GetCartDetails(userId)
	if err != nil {
		return err
	}

	orderItem, orderId, err := util.BuildOrderCartItems(userCartDetails, userId)
	if err != nil {
		return err
	}

	err = repository.CreateOrderCartItems(orderItem)
	if err != nil {
		return err
	}

	user, err := repository.GetUserByID(userId)
	if err != nil {
		return err
	}

	orderID, err := repository.GetOrderItemByUserIdAndOrderId(uint(userId), orderId)
	if err != nil {
		return nil
	}

	couponData := &domain.Coupon{}
	usercoupon := domain.UserCoupon{}
	validCoupon := &domain.Coupon{}

	if coupon != "" {
		couponData, err = repository.GetCouponByCouponName(coupon)
	}
	if couponData != nil && err == nil {
		validCoupon, err = util.CouponValidate(couponData)
		if err != nil {
			return err
		}
	}

	err = repository.CheckCouponUsedByUser(userId, validCoupon)
	if err != nil {
		usercoupon = util.BuildUserCoupon(userId, *validCoupon)
		err = repository.CreateUserCoupon(usercoupon)
		if err != nil {
			return err
		}
		err = repository.UpdateCouponRemainingUses(validCoupon)
		if err != nil {
			return err
		}
	} else {
		validCoupon.Code = ""
		validCoupon.DiscountPercentage = 0
	}

	err = repository.CheckValidAddress(userId, addressId)
	if err != nil {
		return errors.New("Enter a valid address or enter a new address")
	}

	order := util.BuildOrderByWalletPayment(orderItem, *user, orderID, orderId, *validCoupon, addressId)
	err = repository.Order(order)
	if err != nil {
		return err
	}
	repository.UpdateWalletAmouont(order, userId)
	productIDs, quantities := util.GetProductIDsFromCart(userCartDetails)
	err = repository.UpdateProductStockQuantity(productIDs, quantities)
	if err != nil {
		return err
	}

	err = repository.DeleteCartItemByUserID(uint(userId))
	if err != nil {
		return err
	}

	for i := 0; i < len(productIDs); i++ {
		err = repository.CheckItemPresentInWishList(userId, productIDs[i])
		if err == nil {
			repository.DeleteWishlistItem(userId, productIDs[i])
		}

	}
	return nil
}

func WalletPaymentOrderItemByID(userId, productId, quantity int, coupon string, addressId int) error {
	productDetails, err := repository.GetProductByID(productId)
	if err != nil {
		return err
	}
	if productDetails.StockQuantity < quantity {
		return errors.New("Product out of stock")
	}
	orderItem, orderId, err := util.BuildOrderItemByID(userId, quantity, *productDetails)
	if err != nil {
		return err
	}
	err = repository.CreateOrderItems(*orderItem)
	if err != nil {
		return err
	}
	user, err := repository.GetUserByID(userId)
	if err != nil {
		return err
	}

	orderID, err := repository.GetOrderItemByUserIdAndOrderId(uint(userId), orderId)
	if err != nil {
		return nil
	}
	couponData := &domain.Coupon{}
	usercoupon := domain.UserCoupon{}
	validCoupon := &domain.Coupon{}

	if coupon != "" {
		couponData, err = repository.GetCouponByCouponName(coupon)
	}
	if couponData != nil && err == nil {
		validCoupon, err = util.CouponValidate(couponData)
		if err != nil {
			return err
		}
	}

	err = repository.CheckCouponUsedByUser(userId, validCoupon)
	if err != nil {
		usercoupon = util.BuildUserCoupon(userId, *validCoupon)
		err = repository.CreateUserCoupon(usercoupon)
		if err != nil {
			return err
		}
		err = repository.UpdateCouponRemainingUses(validCoupon)
		if err != nil {
			return err
		}
	} else {
		validCoupon.Code = ""
		validCoupon.DiscountPercentage = 0
	}

	err = repository.CheckValidAddress(userId, addressId)
	if err != nil {
		return errors.New("Enter a valid address or enter a new address")
	}

	order := util.BuildOrderByWalletPaymentProductID(orderItem, *user, orderID, orderId, *validCoupon, addressId)

	repository.UpdateWalletAmouont(order, userId)
	err = repository.CheckCartItemByUserIdAndProductId(userId, productId)
	if err == nil {
		err := repository.DeleteCartItem(userId, productId)
		if err != nil {
			return err
		}
	}
	err = repository.Order(order)
	if err != nil {
		return err
	}
	err = repository.UpdateProductStock(productId, quantity)
	if err != nil {
		return err
	}
	err = repository.CheckItemPresentInWishList(userId, productId)
	if err == nil {
		err = repository.DeleteWishlistItem(userId, productId)
		if err != nil {
			return err
		}
	}
	return nil
}

func OrderSummary(userId, orderId int) (*domain.OrderSummary, error) {
	order, orderItem, err := repository.ViewOrdersByID(userId, orderId)
	if err != nil {
		return nil, err
	}
	ordersummary, err := util.BuildOrderSummary(order, orderItem, userId, orderId)
	if err != nil {
		return nil, err
	}
	return &ordersummary, nil

}

func ViewOrders(id int) ([]domain.OrderResponse, error) {
	orderRes := []domain.OrderResponse{}
	orders, err := repository.ViewOrders(id)
	if err != nil {
		return nil, err
	}

	orderResponse := util.BuildOrderResponse(orderRes, orders)
	return orderResponse, nil
}

func ViewOrdersByID(userId, orderId int) ([]domain.OrderItemResponse, error) {
	orderItemRes := []domain.OrderItemResponse{}

	order, orderItem, err := repository.ViewOrdersByID(userId, orderId)
	if err != nil {
		return nil, err
	}

	orderItemResponse := util.BuildOrderItemResponse(orderItemRes, order, orderItem)
	return orderItemResponse, nil

}

func ViewOrdersByUserID(id int) ([]domain.OrderResponse, error) {
	orderRes := []domain.OrderResponse{}

	orders, err := repository.ViewOrdersByUserID(id)
	if err != nil {
		return nil, err
	}
	orderResponse := util.BuildOrderResponse(orderRes, orders)
	return orderResponse, nil
}

func ViewOrderItemsByUserID(userId, orderId int) ([]domain.OrderItemResponse, error) {
	orderItemRes := []domain.OrderItemResponse{}

	order, orderItem, err := repository.ViewOrdersByID(userId, orderId)
	if err != nil {
		return nil, err
	}

	orderItemResponse := util.BuildOrderItemResponse(orderItemRes, order, orderItem)
	return orderItemResponse, nil
}

func OrderCancel(userId, orderId int) (*domain.Order, error) {
	orders, err := repository.OrderCancel(userId, orderId)
	if err != nil {
		return nil, err
	}

	_, orderItems, _ := repository.ViewOrdersByID(userId, orderId)

	productIds, quantities := repository.GetProductIDsFromOrderItems(orderItems)

	err = repository.ProductStockUpdationAfterCancellation(productIds, quantities)
	if err != nil {
		return nil, err
	}
	if orders.PaymentMethod != "Cash On Delivery" {
		err := repository.RefundWalletAmount(*orders, userId)
		if err != nil {
			return nil, err
		}

	}
	return orders, nil
}

func OrderReturn(orderId, userId int) error {
	_, err := repository.OrderReturn(orderId, userId)
	if err != nil {
		return err
	}
	return nil
}

func ReturnConfirmation(userId, orderId int) error {
	orders, err := repository.ReturnConfirmation(userId, orderId)
	if err != nil {
		return err
	}
	_, orderItems, err := repository.ViewOrdersByID(userId, orderId)
	if err != nil {
		return err
	}
	productIds, quantities := repository.GetProductIDsFromOrderItems(orderItems)

	err = repository.RefundWalletAmount(*orders, userId)
	if err != nil {
		return errors.New("Refund failed")
	}

	err = repository.ProductStockUpdationAfterCancellation(productIds, quantities)
	if err != nil {
		return err
	}
	return nil

}

// Admin - Coupon------------------------------------------------------------------------------------------------

func AddCoupon(coupon *domain.Coupon) error {
	res, err := repository.FindCouponByCode(coupon)
	if err != nil && res == nil {
		err := repository.AddCoupon(coupon)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Coupon Already Exist")
	}

	return nil
}

func DeleteCoupon(id int) error {
	err := repository.DeleteCoupon(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCoupon(updateCoupon domain.Coupon, id int) error {
	err := repository.UpdateCoupon(updateCoupon, id)
	if err != nil {
		return err
	}
	return nil
}

func ViewCoupons() ([]domain.Coupon, error) {
	coupon, err := repository.ViewCoupons()
	if err != nil {
		return nil, err
	}
	return coupon, nil
}

// Wallet----------------------------------------------------------------------------------------------------------
func AddAmountToWallet(data *domain.Wallet, userId int) error {
	err := repository.CreateWalletAmount(data, userId)
	if err != nil {
		return err
	}
	return nil
}

// Admin - Change Order Status-------------------------------------------------------------------------------------
func ChangeOrderStatus() error {
	err := repository.ChangeOrderStatus()
	if err != nil {
		return err
	}
	return nil
}

// User - Invoice -------------------------------------------------------------------------------------------------
func InvoiceDetails(userId, orderId int) (*domain.Invoice, error) {
	order, err := repository.OrderByOrderID(userId, orderId)
	if err != nil {
		return nil, err
	}

	_, oredrItems, err := repository.ViewOrdersByID(userId, orderId)
	if err != nil {
		return nil, err
	}
	user, err := repository.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	address, err := repository.ViewAddressByID(order.AddressID)
	invoiceDetails, err := util.BuildInvoiceDetails(*order, oredrItems, *user, *address)

	tmpl, err := template.New("invoice").Parse(util.InvoiceTemplate)
	if err != nil {

		return nil, err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, invoiceDetails)
	if err != nil {

		return nil, err

	}

	htmlFile := "./pkg/templates/invoice.html"
	err = ioutil.WriteFile(htmlFile, buf.Bytes(), 0644)
	if err != nil {
		return nil, err

	}

	cmd := exec.Command("wkhtmltopdf", htmlFile, "./data/invoice.pdf")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return nil, err

	}
	return &invoiceDetails, nil
}

// Admin - Sales Report---------------------------------------------------------------------------------------------
func SalesReport(parseFromDate, parseToDate time.Time) (*domain.SalesReport, error) {
	orders, orderIds, err := repository.SalesReport(parseFromDate, parseToDate)
	if err != nil {
		return nil, err
	}

	orderItems, err := repository.GetOrderItemsByOrderIDs(orderIds)
	sales := util.BuildSales(orders)

	salesDetails, err := util.BuildSalesReport(orders, orderItems, sales, parseFromDate, parseToDate)

	tmpl, err := template.New("salesReport").Parse(util.SalesReportTemplate)
	if err != nil {

		return nil, err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, salesDetails)
	if err != nil {

		return nil, err

	}

	htmlFile := "./pkg/templates/salesReport.html"
	err = ioutil.WriteFile(htmlFile, buf.Bytes(), 0644)
	if err != nil {
		return nil, err

	}

	cmd := exec.Command("wkhtmltopdf", htmlFile, "./data/salesReport.pdf")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return nil, err

	}
	return &salesDetails, nil
}

// Admin - Dashboard---------------------------------------------------------------------------------------------
func GetDashBoard() (float64, int, error) {
	totalSales, totalOrders, err := repository.GetDashBoard()
	if err != nil {
		return 0, 0, err
	}
	return totalSales, totalOrders, nil
}

func GetYearlySales(year int) (float64, error) {
	totalSales, err := repository.GetYearlySales(year)
	if err != nil {
		return 0, err
	}
	return totalSales, nil
}

func GetMonthlySales(year, month int) (float64, error) {
	totalSales, err := repository.GetMonthlySales(year, month)
	if err != nil {
		return 0, err
	}
	return totalSales, nil
}

func GetWeeklySales(year, week int) (float64, error) {
	totalSales, err := repository.GetWeeklySales(year, week)
	if err != nil {
		return 0, err
	}
	return totalSales, nil
}

func GetTotalSalesAmountToday() (float64, error) {
	totalSales, err := repository.GetTotalSalesAmountToday()
	if err != nil {
		return 0, err
	}
	return totalSales, nil
}

func GetTotalOrdersToday() (int, error) {
	totalOrders, err := repository.GetTotalOrdersToday()
	if err != nil {
		return 0, err
	}
	return totalOrders, nil
}
