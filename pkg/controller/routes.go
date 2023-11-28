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

		userGroup.GET("/products", handler.GetProducts)
		userGroup.GET("/product/:id", handler.GetProductByID)
		userGroup.GET("/product", handler.GetProductByName)

		userGroup.GET("/brands", handler.GetBrandsUser)
		userGroup.GET("/product/brand/:id", handler.GetProductByBrandID)
		userGroup.GET("/product/brand", handler.GetProductByBrandName)

		userGroup.GET("/categories", handler.GetCategoriesUser)
		userGroup.GET("/product/category/:id", handler.GetProductByCategoryID)
		userGroup.GET("/product/category", handler.GetProductByCategoryName)

	}

	userAuthGroup := r.Group("/user")
	{
		userAuthGroup.Use(handler.UserAuthRequired)

		userAuthGroup.POST("/cart", handler.AddToCart)
		userAuthGroup.GET("/cart", handler.CartList)
		userAuthGroup.DELETE("/cart/:id", handler.DeleteCartItem)

		userAuthGroup.POST("/address", handler.AddAddress)
		userAuthGroup.GET("/address", handler.ViewAddress)

		userAuthGroup.PATCH("/profile", handler.EditUserProfile)
		userAuthGroup.GET("/profile/:id", handler.ProfileDetails)

		userAuthGroup.GET("/ordersummary", handler.OrderSummary)

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
