package util

import (
	"errors"
	"strconv"
	"time"

	"github.com/aparnasukesh/shoezone/pkg/domain"
)

func CreateCartResponse(response *[]domain.CartResponse, cartProducts []domain.Cart) {
	for _, val := range cartProducts {
		cartResponse := domain.CartResponse{
			UserID:    val.UserID,
			ProductID: val.ProductID,
			Quantity:  val.Quantity,
		}

		*response = append(*response, cartResponse)
	}
}

func BuildCartItemsOrderSummary(userCartDetails []domain.Cart) domain.CartItemsOrderSummary {
	var orderSummary domain.CartItemsOrderSummary
	orderSummary.TotalPrice = 0

	for _, cartItem := range userCartDetails {
		total_amount := float64(cartItem.Quantity) * cartItem.CartProduct.Price
		orderProduct := domain.OrderProduct{
			ProductName: cartItem.CartProduct.ProductName,
			Price:       cartItem.CartProduct.Price,
			Quantity:    cartItem.Quantity,
			TotalAmount: total_amount,
		}

		orderSummary.TotalPrice += (float64(cartItem.Quantity) * cartItem.CartProduct.Price)

		orderSummary.Products = append(orderSummary.Products, orderProduct)
	}

	return orderSummary
}

func BuildOrderSummary(order []domain.Order, orderitem []domain.OrderItem, userId, orderId int) (domain.OrderSummary, error) {
	var orderSummary domain.OrderSummary
	orderSummary.UserID = userId
	orderSummary.OrderID = uint(orderId)
	orderSummary.PaymentMethod = order[0].PaymentMethod
	orderSummary.TotalPrice = order[0].TotalAmount
	orderSummary.DiscountPrice = order[0].DiscountPrice
	orderSummary.AmountPayable = order[0].AmountPayable

	for _, orderItem := range orderitem {
		total_amount := float64(orderItem.Quantity) * orderItem.UnitPrice
		orderProduct := domain.OrderProduct{
			ProductName: orderItem.Product.ProductName,
			Price:       orderItem.UnitPrice,
			Quantity:    int(orderItem.Quantity),
			TotalAmount: total_amount,
		}
		orderSummary.Products = append(orderSummary.Products, orderProduct)
	}
	return orderSummary, nil

}
func BuildOrderCartItems(userCartDetails []domain.Cart, userId int) ([]domain.OrderItem, uint, error) {
	var orderItems []domain.OrderItem

	OrderIdStr, err := GenCaptchaCode()
	if err != nil {
		return nil, 0, err
	}
	orderId, err := strconv.Atoi(OrderIdStr)
	if err != nil {
		return nil, 0, err
	}

	for _, cartItem := range userCartDetails {
		totalPrice := float64(cartItem.Quantity) * cartItem.CartProduct.Price

		orderItem := domain.OrderItem{
			UserID:     uint(userId),
			OrderID:    uint(orderId),
			ProductID:  uint(cartItem.ProductID),
			Quantity:   uint(cartItem.Quantity),
			Product:    cartItem.CartProduct,
			UnitPrice:  cartItem.CartProduct.Price,
			TotalPrice: totalPrice,
		}

		orderItems = append(orderItems, orderItem)
	}

	return orderItems, uint(orderId), nil
}

func BuildOrderItemByID(userId, quantity int, product domain.Product) (*domain.OrderItem, uint, error) {
	orderItem := &domain.OrderItem{}
	OrderIdStr, err := GenCaptchaCode()
	if err != nil {
		return nil, 0, err
	}
	orderId, err := strconv.Atoi(OrderIdStr)
	if err != nil {
		return nil, 0, err
	}
	orderItem.UserID = uint(userId)
	orderItem.OrderID = uint(orderId)
	orderItem.ProductID = product.ID
	orderItem.Quantity = uint(quantity)
	orderItem.Product = product
	orderItem.UnitPrice = product.Price
	orderItem.TotalPrice = float64(quantity) * product.Price

	return orderItem, uint(orderId), nil

}

func BuildOrder(orderItems []domain.OrderItem, user domain.User, orderItemID, orderID uint, coupon domain.Coupon, addressId int) domain.Order {
	var orders domain.Order
	var totalAmount float64 = 0
	var discountAmount float64 = 0
	var amountPayable float64 = 0
	for _, orderItem := range orderItems {
		totalAmount += orderItem.TotalPrice
	}
	offerAmount := float64(coupon.DiscountPercentage)
	discountAmount = (offerAmount / 100) * totalAmount
	amountPayable = totalAmount - discountAmount
	orders.UserID = user.ID
	orders.TotalAmount = totalAmount
	orders.DiscountPrice = discountAmount
	orders.AmountPayable = amountPayable
	orders.CouponName = coupon.Code
	orders.OrderStatus = "Pending status"
	orders.AddressID = uint(addressId)
	orders.OrderDate = time.Now()
	orders.PaymentMethod = "Cash On Delivery"
	orders.OrderItemID = orderItemID
	orders.BookingID = orderID

	return orders

}

func BuildOrderRazorpay(orderItems []domain.OrderItem, user domain.User, orderItemID, orderID uint, coupon domain.Coupon, addressId int) domain.Order {
	var orders domain.Order
	var totalAmount float64 = 0
	var discountAmount float64 = 0
	var amountPayable float64 = 0
	for _, orderItem := range orderItems {
		totalAmount += orderItem.TotalPrice
	}
	offerAmount := float64(coupon.DiscountPercentage)
	discountAmount = (offerAmount / 100) * totalAmount
	amountPayable = totalAmount - discountAmount
	orders.UserID = user.ID
	orders.TotalAmount = totalAmount
	orders.DiscountPrice = discountAmount
	orders.AmountPayable = amountPayable
	orders.CouponName = coupon.Code
	orders.OrderStatus = "Pending status"
	orders.AddressID = uint(addressId)
	orders.OrderDate = time.Now()
	orders.PaymentMethod = "Razorpay"
	orders.OrderItemID = orderItemID
	orders.BookingID = orderID

	return orders
}

func BuildOrderbyProductID(orderItem *domain.OrderItem, user domain.User, orderID, orderId uint, coupon domain.Coupon, addressId int) domain.Order {
	var orders domain.Order
	var discountAmount float64 = 0
	var amountPayable float64 = 0
	offerAmount := float64(coupon.DiscountPercentage)
	discountAmount = (offerAmount / 100) * orderItem.TotalPrice
	amountPayable = orderItem.TotalPrice - discountAmount
	orders.UserID = user.ID
	orders.TotalAmount = orderItem.TotalPrice
	orders.DiscountPrice = discountAmount
	orders.AmountPayable = amountPayable
	orders.CouponName = coupon.Code
	orders.OrderStatus = "Pending status"
	orders.AddressID = uint(addressId)
	orders.OrderDate = time.Now()
	orders.PaymentMethod = "Cash On Delivery"
	orders.OrderItemID = orderID
	orders.BookingID = orderId

	return orders
}

func BuildOrderByWalletPaymentProductID(orderItem *domain.OrderItem, user domain.User, orderID, orderId uint, coupon domain.Coupon, addressId int) domain.Order {
	var orders domain.Order
	var discountAmount float64 = 0
	var amountPayable float64 = 0
	offerAmount := float64(coupon.DiscountPercentage)
	discountAmount = (offerAmount / 100) * orderItem.TotalPrice
	amountPayable = orderItem.TotalPrice - discountAmount
	orders.UserID = user.ID
	orders.TotalAmount = orderItem.TotalPrice
	orders.DiscountPrice = discountAmount
	orders.AmountPayable = amountPayable
	orders.CouponName = coupon.Code
	orders.OrderStatus = "Pending status"
	orders.AddressID = uint(addressId)
	orders.OrderDate = time.Now()
	orders.PaymentMethod = "Wallet Payment"
	orders.OrderItemID = orderID
	orders.BookingID = orderId

	return orders
}
func BuildOrderByWalletPayment(orderItems []domain.OrderItem, user domain.User, orderItemID, orderID uint, coupon domain.Coupon, addressId int) domain.Order {

	var orders domain.Order
	var totalAmount float64 = 0
	var discountAmount float64 = 0
	var amountPayable float64 = 0
	for _, orderItem := range orderItems {
		totalAmount += orderItem.TotalPrice
	}
	offerAmount := float64(coupon.DiscountPercentage)
	discountAmount = (offerAmount / 100) * totalAmount
	amountPayable = totalAmount - discountAmount
	orders.UserID = user.ID
	orders.TotalAmount = totalAmount
	orders.DiscountPrice = discountAmount
	orders.AmountPayable = amountPayable
	orders.CouponName = coupon.Code
	orders.OrderStatus = "Pending status"
	orders.AddressID = uint(addressId)
	orders.OrderDate = time.Now()
	orders.PaymentMethod = "Wallet Payment"
	orders.OrderItemID = orderItemID
	orders.BookingID = orderID

	return orders

}
func BuildOrderResponse(orderRes []domain.OrderResponse, orders []domain.Order) []domain.OrderResponse {
	for _, val := range orders {
		order := domain.OrderResponse{
			ID:            val.ID,
			CreatedAt:     val.CreatedAt,
			OrderStatus:   val.OrderStatus,
			OrderDate:     val.OrderDate,
			AddressID:     val.AddressID,
			PaymentMethod: val.PaymentMethod,
			BookingID:     val.BookingID,
			OrderItemID:   val.OrderItemID,
			TotalAmount:   val.TotalAmount,
		}
		orderRes = append(orderRes, order)
	}
	return orderRes
}

func BuildOrderItemResponse(orderItemRes []domain.OrderItemResponse, orders []domain.Order, orderItem []domain.OrderItem) []domain.OrderItemResponse {
	addressId := orders[0].AddressID
	for i := range orderItem {
		orderItem := domain.OrderItemResponse{
			ID:          orderItem[i].ID,
			CreatedAt:   orderItem[i].CreatedAt,
			ProductID:   int(orderItem[i].ProductID),
			ProductName: orderItem[i].Product.ProductName,
			Quantity:    int(orderItem[i].Quantity),
			UnitPrice:   orderItem[i].UnitPrice,
			Address:     int(addressId),
			TotalPrice:  orderItem[i].TotalPrice,
		}
		orderItemRes = append(orderItemRes, orderItem)
	}
	return orderItemRes
}

func CouponValidate(coupon *domain.Coupon) (*domain.Coupon, error) {
	currentTime := time.Now()
	if currentTime.Before(coupon.ExpiryDate) == true && coupon.RemainingUses > 0 {
		return coupon, nil
	}
	return nil, errors.New("Coupon not valid")
}

func BuildUserCoupon(userId int, coupon domain.Coupon) domain.UserCoupon {
	userCoupon := &domain.UserCoupon{}
	userCoupon.UserID = uint(userId)
	userCoupon.CouponID = coupon.ID
	userCoupon.Used = true
	userCoupon.UsedDate = time.Now()
	return *userCoupon
}

func GetProductIDsFromCart(cartItem []domain.Cart) ([]int, []int) {
	var productIds []int
	var quantities []int

	for _, id := range cartItem {
		productIds = append(productIds, id.ProductID)
		quantities = append(quantities, id.Quantity)
	}
	return productIds, quantities
}

func BuildUserProfileUpdate(updatedUser domain.User, password string) (domain.UserProfileUpdate, error) {
	user := &domain.UserProfileUpdate{}

	user.Username = updatedUser.Username
	user.Phone = updatedUser.Phone
	user.Email = updatedUser.Email
	user.Gender = updatedUser.Gender
	user.Dateofbirth = updatedUser.Dateofbirth

	if password == "" {
		user.Password = "password not updated"
	} else {
		user.Password = password
	}

	return *user, nil
}

func BuildWishListResponse(products []domain.WishList) ([]domain.WishListResponse, error) {
	response := []domain.WishListResponse{}

	for _, product := range products {
		res := domain.WishListResponse{
			UserID:      product.UserID,
			ProductID:   product.ProductID,
			ProductName: product.WishListProduct.ProductName,
			Price:       product.WishListProduct.Price,
		}
		response = append(response, res)
	}

	return response, nil
}

func BuildProfileDetails(profile *domain.UserProfileUpdate) *domain.ProfileDetails {
	profileDetails := domain.ProfileDetails{}

	profileDetails.Username = profile.Username
	profileDetails.Phone = profile.Phone
	profileDetails.Email = profile.Email
	profileDetails.Dateofbirth = profile.Dateofbirth
	profileDetails.Gender = profile.Gender

	return &profileDetails
}
