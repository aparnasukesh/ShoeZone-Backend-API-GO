package controller

import (
	handler "github.com/aparnasukesh/shoezone/pkg/controller/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	userGroup := r.Group("/user")
	{
		// User - Sign up, Log in
		userGroup.POST("/register", handler.RegisterUser)
		userGroup.POST("/register/validate", handler.RegisterValidate)
		userGroup.POST("/login", handler.Login)

		// User - Products
		userGroup.GET("/products", handler.GetProducts)
		userGroup.GET("/product/:id", handler.GetProductByID)
		userGroup.GET("/product", handler.GetProductByName)
		userGroup.GET("/product/image/view/:id", handler.ProductImageViewByID)

		// User - Brands - Products
		userGroup.GET("/brands", handler.GetBrandsUser)
		userGroup.GET("/product/brand/:id", handler.GetProductByBrandID)
		userGroup.GET("/product/brand", handler.GetProductByBrandName)

		// User - Categories - Products
		userGroup.GET("/categories", handler.GetCategoriesUser)
		userGroup.GET("/product/category/:id", handler.GetProductByCategoryID)
		userGroup.GET("/product/category", handler.GetProductByCategoryName)

		// User - Password
		userGroup.POST("/forgot/password", handler.ForgotPassword)
		userGroup.POST("/reset/password", handler.ResetPassword)

	}

	userAuthGroup := r.Group("/user")
	{
		userAuthGroup.Use(handler.UserAuthRequired)

		// User - Cart
		userAuthGroup.POST("/cart", handler.AddToCart)
		userAuthGroup.GET("/cart", handler.CartList)
		userAuthGroup.DELETE("/cart/:id", handler.DeleteCartItem)

		// User - Wish Lists
		userAuthGroup.POST("/wishlist/:productid", handler.AddToWishList)
		userAuthGroup.DELETE("/wishlist/:productid", handler.DeleteWishlistItem)
		userAuthGroup.GET("/wishlist", handler.WishListItems)

		// User - Profile
		userAuthGroup.POST("/address", handler.AddAddress)
		userAuthGroup.GET("/address", handler.ViewAddress)
		userAuthGroup.PATCH("/profile", handler.EditUserProfile)
		userAuthGroup.GET("/profile", handler.ProfileDetails)

		// User - Order
		userAuthGroup.GET("/cart/order/summary", handler.CartItemsOrderSummary)
		userAuthGroup.GET("/orders", handler.ViewOrders)
		userAuthGroup.GET("/order/orderid", handler.ViewOrdersByID)
		userAuthGroup.PATCH("/order/cancel", handler.OrderCancel)
		userAuthGroup.GET("/order/summary", handler.OrderSummary)
		userAuthGroup.PATCH("/order/return", handler.OrderReturn)

		// User - Wallet
		userAuthGroup.POST("/wallet", handler.AddAmountToWallet)

		// User - Payment
		payment := userAuthGroup.Group("/payment")
		{
			payment.POST("/cod", handler.OrderCartItems)
			payment.POST("/order/cod", handler.OrderItemByID)
			payment.POST("/wallet", handler.WalletPaymentCartItems)
			payment.POST("/order/wallet", handler.WalletPaymentOrderItemByID)
		}
	}

	userPaymentAuth := r.Group("/user")
	{
		userPaymentAuth.Use(handler.UserPaymentAuthorization)

		userPaymentAuth.GET("/razorpay", handler.OrderCartItemsRazorpay)
		userPaymentAuth.GET("/razorpay/success", handler.VerifyPayment)
		userPaymentAuth.GET("/success", handler.RazorpaySuccess)
		userPaymentAuth.GET("/failed", handler.RazorPayFailed)

	}

	adminGroup := r.Group("/admin")
	{
		adminGroup.POST("/login", handler.Login)

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

		// Product Image Upload
		adminGroup.PATCH("/product/upload/:id", handler.ProductImageUpload)
		adminGroup.GET("/product/image/view/:id", handler.ProductImageViewByID)

		// Order Management
		adminGroup.GET("/orders/:id", handler.ViewOrdersByUserID)
		adminGroup.GET("/orders/orderid/:id", handler.ViewOrderItemsByUserID)
		adminGroup.PATCH("/order/cancel", handler.OrderCancel)
		adminGroup.PATCH("/orderstatus", handler.ChangeOrderStatus)
		adminGroup.PATCH("/order/return", handler.ReturnConfirmation)

		//Admin - Coupon Management
		adminGroup.POST("/coupon", handler.AddCoupon)
		adminGroup.DELETE("/coupon/:id", handler.DeleteCoupon)
		adminGroup.PATCH("/coupon/:id", handler.UpdateCoupon)
		adminGroup.GET("/coupon", handler.ViewCoupons)

	}

}
