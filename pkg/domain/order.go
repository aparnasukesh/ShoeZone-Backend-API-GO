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

type OrderSummary struct {
	UserID     int            `json:"user_id"`
	TotalPrice float64        `json:"total_price"`
	Products   []OrderProduct `json:"products"`
}
