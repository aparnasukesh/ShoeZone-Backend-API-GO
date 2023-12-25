package usecase

import (
	"errors"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/razorpay/razorpay-go"
)

func RazorPay(paymentDetails domain.RazorPay) (*domain.RazorPay, error) {

	client := razorpay.NewClient("rzp_test_m7lKivrgzyXezp", "9xTriOeqo6SNVBA0eRojnNIi")
	data := map[string]interface{}{
		"amount":   paymentDetails.TotalAmount,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return nil, err
	}

	value, ok := body["id"].(string)
	if !ok {
		return nil, errors.New("Razorpay error")
	}
	paymentDetails.PaymentID = value
	return &paymentDetails, nil
}
