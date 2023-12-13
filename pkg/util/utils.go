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

func BuildOrderSummary(userCartDetails []domain.Cart) domain.OrderSummary {
	var orderSummary domain.OrderSummary
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

func BuildOrderItem(userCartDetails []domain.Cart, userId int) ([]domain.OrderItem, uint, error) {
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

func BuildOrder(orderItems []domain.OrderItem, user domain.User, orderItemID, orderID uint, coupon domain.Coupon) domain.Order {
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
	orders.AddressID = user.DefaultAddressID
	orders.OrderDate = time.Now()
	orders.PaymentMethod = "Cash On Delivery"
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
	user.DefaultAddressID = updatedUser.DefaultAddressID
	user.Dateofbirth = updatedUser.Dateofbirth

	if password == "" {
		user.Password = "password not updated"
	} else {
		user.Password = password
	}

	return *user, nil
}
