package util

import (
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

// func ProductPriceCalculation(quantities []int, prices []float64) []float64 {
// 	total_price := make([]float64, len(quantities))
// 	for i := range quantities {
// 		total_price[i] = float64(quantities[i]) * prices[i]
// 	}

// 	return total_price

// }

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

func BuildOrder(orderItems []domain.OrderItem, user domain.User, orderItemID, orderID uint) domain.Order {
	var orders domain.Order
	var totalAmount float64 = 0
	for _, orderItem := range orderItems {
		totalAmount += orderItem.TotalPrice
	}

	orders.UserID = user.ID
	orders.TotalAmount = totalAmount
	orders.OrderStatus = "Pending status"
	orders.AddressID = user.DefaultAddressID
	orders.OrderDate = time.Now()
	orders.PaymentMethod = "Cash On Delivery"
	orders.OrderItemID = orderItemID
	orders.BookingID = orderID

	return orders

}

// func UpdateProductStockQuantity(cartItem []domain.Cart) error {
// 	product := domain.Product{}

// }
