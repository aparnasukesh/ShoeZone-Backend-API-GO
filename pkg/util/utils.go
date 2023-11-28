package util

import "github.com/aparnasukesh/shoezone/pkg/domain"

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

func ProductPriceCalculation(quantities []int, prices []float64) []float64 {
	total_price := make([]float64, len(quantities))
	for i := range quantities {
		total_price[i] = float64(quantities[i]) * prices[i]
	}

	return total_price

}

func BuildOrderSummary(userCartDetails []domain.Cart) domain.OrderSummary {
	var orderSummary domain.OrderSummary
	orderSummary.TotalPrice = 0

	for _, cartItem := range userCartDetails {
		orderProduct := domain.OrderProduct{
			ProductName: cartItem.CartProduct.ProductName,
			Price:       cartItem.CartProduct.Price,
			Quantity:    cartItem.Quantity,
			Brandname:   cartItem.CartProduct.ProductBrand.BrandName,
		}

		orderSummary.TotalPrice += (float64(cartItem.Quantity) * cartItem.CartProduct.Price)

		orderSummary.Products = append(orderSummary.Products, orderProduct)
	}

	return orderSummary
}
