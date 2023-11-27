package util

import "github.com/aparnasukesh/shoezone/pkg/domain"

func CreateCartResponse(response *[]domain.CartResponse, cartProducts []domain.Cart) {
	for _, val := range cartProducts {
		cartResponse := domain.CartResponse{
			ID:        val.ID,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			UserID:    val.UserID,
			ProductID: val.ProductID,
			Quantity:  val.Quantity,
		}

		*response = append(*response, cartResponse)
	}
}
