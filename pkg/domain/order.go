package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	TotalAmount   float64   `json:"total_amount"`
	OrderStatus   string    `json:"order_status"`
	OrderDate     time.Time `json:"order_date"`
	AddressID     uint      `json:"address_id"`
	PaymentMethod string    `json:"payment_method"`
	CouponName    string    `json:"coupon_name"`
	DiscountPrice float64   `json:"discount_price"`
	AmountPayable float64   `json:"amount_payable"`
	BookingID     uint      `json:"booking_id"`
	OrderItemID   uint      `json:"order_item_id"`
	OrderItems    OrderItem `gorm:"foreignKey:OrderItemID" json:"order_items"`
	Address       Address   `gorm:"foreignKey:AddressID" json:"address"`
}

type OrderResponse struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	OrderStatus   string    `json:"order_status"`
	OrderDate     time.Time `json:"order_date"`
	AddressID     uint      `json:"address_id"`
	PaymentMethod string    `json:"payment_method"`
	BookingID     uint      `json:"booking_id"`
	OrderItemID   uint      `json:"order_item_id"`
	TotalAmount   float64   `json:"total_amount"`
}

type OrderItem struct {
	gorm.Model
	UserID     uint    `json:"user_id"`
	ProductID  uint    `json:"product_id"`
	Quantity   uint    `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
	OrderID    uint    `json:"order_id"`
	Product    Product `gorm:"foreignKey:ProductID" json:"product"`
}

type OrderItemResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ProductID   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	UnitPrice   float64   `json:"unit_price"`
	Address     int       `json:"address"`
	TotalPrice  float64   `json:"total_price"`
}

type OrderProduct struct {
	ProductName string `json:"product_name"`
	Price       float64
	Quantity    int `json:"quantity"`
	TotalAmount float64
}

type CartItemsOrderSummary struct {
	UserID     int            `json:"user_id"`
	TotalPrice float64        `json:"total_price"`
	Products   []OrderProduct `json:"products"`
}

type OrderSummary struct {
	UserID        int            `json:"user_id"`
	OrderID       uint           `json:"order_id"`
	Products      []OrderProduct `json:"products"`
	PaymentMethod string
	TotalPrice    float64 `json:"total_price"`
	DiscountPrice float64 `json:"discount_price"`
	AmountPayable float64 `json:"amount_payable"`
}

// Coupon
type Coupon struct {
	gorm.Model
	Code               string    `json:"code"`
	DiscountPercentage int       `json:"discount_percentage"`
	ExpiryDate         time.Time `json:"expiry_date"`
	UsageLimit         int       `json:"usage_limit"`
	RemainingUses      int       `json:"remaining_uses"`
}

type UserCoupon struct {
	gorm.Model
	UserID   uint      `json:"user_id"`
	CouponID uint      `json:"coupon_id"`
	Used     bool      `json:"used"`
	UsedDate time.Time `json:"used_date,omitempty"`
}

type UserCouponResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UserID     uint      `json:"user_id"`
	CouponID   uint      `json:"coupon_id"`
	Used       bool      `json:"used"`
	UsedDate   time.Time `json:"used_date,omitempty"`
	CouponInfo Coupon    `gorm:"foreignKey:CouponID" json:"coupon_info"`
}

// Wallet
type Wallet struct {
	gorm.Model
	UserID          uint      `json:"user_id" gorm:"uniqueIndex"`
	Balance         float64   `json:"balance"`
	LastTransaction time.Time `json:"last_transaction"`
}

type RazorPayPayment struct {
	gorm.Model
	UserID            int
	OrderID           string
	PaymentID         string
	TotalAmount       float64
	RazorPaySignature string
	PaymentStatus     string
}

// Razorpay
type RazorPay struct {
	UserID        int
	AddressID     int
	Order_TableID int
	Coupon        string
	PaymentID     string
	TotalAmount   float64
}

// Invoice
type Invoice struct {
	Name          string
	Email         string
	PaymentMethod string
	Date          time.Time
	OrderID       int
	Address       []UserAddress
	Products      []OrderProduct
	TotalAmount   float64
}

// Sales - Report
type Sales struct {
	OrderStatus string
	OrderID     int
	TotalAmount float64
}

type SalesReport struct {
	FromDate      time.Time
	ToDate        time.Time
	Sale          []Sales
	Products      []OrderProduct
	TotalQuantity int
	TotalAmount   float64
}
