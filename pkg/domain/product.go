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
	ProductCategory   Catogery `gorm:"foreignkey:ProductCategoryID"`
	ProductBrand      Brand    `gorm:"foreignkey:BrandID"`
}

type Catogery struct {
	gorm.Model
	CatogeryName string   `json:"product_name" validate:"required,max=24"` 
	CatogeryId int
}

type Brand struct {
	gorm.Model
	BrandName string   `json:"product_name" validate:"required,max=24"` 
	BrandID int
}
