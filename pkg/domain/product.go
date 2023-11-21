package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName       string   `json:"product_name" validate:"required,max=24"`
	Description       string   `json:"description"`
	Price             float64  `json:"price"`
	StockQuantity     int      `json:"stock_quantity"`
	ProductCategoryID int      `json:"product_category_id"`
	ProductBrandID    int      `json:"brand_id"`
	ProductCategory   Category `gorm:"foreignkey:ProductCategoryID"`
	ProductBrand      Brand    `gorm:"foreignkey:ProductBrandID"`
}

type Category struct {
	gorm.Model   `json:"category_id"`
	CategoryName string `json:"category_name" validate:"required,max=24"`
	//CategoryID   int    `json:"category_id"`
}

type Brand struct {
	gorm.Model `json:"brand_id"`
	BrandName  string `json:"brand_name" validate:"required,max=24"`
	//BrandID   int    `json:"brand_id"`
}
