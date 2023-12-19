package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName       string   `json:"product_name" validate:"required,max=24"`
	Description       string   `json:"description"`
	Price             float64  `json:"price"`
	StockQuantity     int      `json:"stock_quantity"`
	ProductCategoryID int      `json:"category_id"`
	ProductBrandID    int      `json:"brand_id"`
	ProductCategory   Category `gorm:"foreignkey:ProductCategoryID"`
	ProductBrand      Brand    `gorm:"foreignkey:ProductBrandID"`
}

type Category struct {
	gorm.Model
	CategoryName string `json:"category_name" validate:"required,max=24"`
}

type Brand struct {
	gorm.Model
	BrandName string `json:"brand_name" validate:"required,max=24"`
}

type Cart struct {
	gorm.Model
	UserID      int     `json:"user_id"`
	ProductID   int     `json:"product_id"`
	Quantity    int     `json:"quantity"`
	CartUser    User    `gorm:"foreignkey:UserID"`
	CartProduct Product `gorm:"foreignkey:ProductID"`
}

type CartResponse struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type WishList struct {
	gorm.Model
	UserID          int     `json:"user_id"`
	ProductID       int     `json:"product_id"`
	WishListUser    User    `gorm:"foreignkey:UserID"`
	WishListProduct Product `gorm:"foreignkey:ProductID"`
}
