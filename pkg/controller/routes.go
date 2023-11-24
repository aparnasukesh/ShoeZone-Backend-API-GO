package controller

import (
	handler "github.com/aparnasukesh/shoezone/pkg/controller/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handler.RegisterUser)
		userGroup.POST("/register/validate", handler.RegisterValidate)
		userGroup.POST("/login", handler.UserLogin)

		userGroup.GET("/product", handler.GetProducts)
		userGroup.GET("/product/:id", handler.GetProductByID)
		userGroup.GET("/product/productname", handler.GetProductByName)

		userGroup.GET("/brands", handler.GetBrandsUser)
		userGroup.GET("/product/brand", handler.GetProductByBrandID)
		userGroup.GET("/product/brandname", handler.GetProductByBrandName)

		userGroup.GET("/categories", handler.GetCategoriesUser)
		userGroup.GET("/product/category", handler.GetProductByCategoryID)
		userGroup.GET("/product/categoryname", handler.GetProductByCategoryName)

	}

	adminGroup := r.Group("/admin")
	{
		adminGroup.Use(handler.AdminAuthRequired)

		// User Management
		adminGroup.GET("/user", handler.GetUsers)
		adminGroup.GET("/user/:id", handler.GetUserByID)
		adminGroup.PATCH("/user/block/:id", handler.BlockUser)
		adminGroup.PATCH("/user/unblock/:id", handler.UnblockUser)

		// Category Management
		adminGroup.GET("/categories", handler.GetCategories)
		adminGroup.POST("/category", handler.AddCategories)
		adminGroup.PUT("/category/:id", handler.UpdateCategory)
		adminGroup.DELETE("/category/:id", handler.DeleteCategory)

		// Brand Management
		adminGroup.GET("/brands", handler.GetBrands)
		adminGroup.POST("/brand", handler.AddBrand)
		adminGroup.PUT("/brand/:id", handler.UpdateBrand)
		adminGroup.DELETE("/brand/:id", handler.DeleteBrand)

		// Product Management
		adminGroup.GET("/product", handler.GetProducts)
		adminGroup.POST("/product", handler.AddProduct)
		adminGroup.PUT("/product/:id", handler.UpdateProduct)
		adminGroup.DELETE("/product/:id", handler.DeleteProduct)

	}

}
