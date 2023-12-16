package handler

import (
	"net/http"
	"strconv"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/usecase"
	"github.com/aparnasukesh/shoezone/pkg/util"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {

	userData := domain.User{}

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding error",
			"Error":   err,
		})
		return
	}

	err := usecase.RegisterUser(&userData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Registering the User Failed",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Redirect: http://localhost:8000/user/register/validate",
		"Error":   nil,
	})

}

func RegisterValidate(ctx *gin.Context) {
	userData := domain.User{}

	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding Error",
			"Error":   err.Error(),
		})
		return
	}
	err := usecase.RegisterValidate(&userData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Succes":  false,
			"Message": "Register validate Error",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "User Registered Successfull",
		"Error":   nil,
		//"Data":    userData,
	})
}

func Login(ctx *gin.Context) {
	userData := domain.User{}
	if err := ctx.ShouldBindJSON(&userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding error",
			"Error":   err.Error(),
		})
		return
	}
	err, res := usecase.UserLogin(&userData)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Login failed",
			"Error":   err.Error(),
		})
		return
	}

	token, err := util.GenerateJWT(*res)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Login failed",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Succesfully Login",
		"Error":   nil,
		"Token":   token,
	})

}

// User-Products------------------------------------------------
func GetProducts(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "0")
	limitStr := ctx.DefaultQuery("limit", "0")
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "List Products failed",
			"Error":   err.Error(),
		})
		return
	}
	limitNum, _ := strconv.Atoi(limitStr)

	offset := pageNum * limitNum
	products, err := usecase.GetProducts(limitNum, offset)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "List products failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":  true,
		"Message":  "Product Details",
		"Error":    nil,
		"Products": products,
	})
}

func GetProductByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "No product found",
			"Error":   err.Error(),
		})
		return
	}
	product, err := usecase.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "No product found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Product Details",
		"Error":   nil,
		"Product": product,
	})
}

func GetProductByBrandID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	limitStr := ctx.DefaultQuery("limit", "0")
	pageStr := ctx.DefaultQuery("page", "0")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "No product found",
			"Error":   err.Error(),
		})
		return
	}
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	products, err := usecase.GetProductByBrandID(limit, offset, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Products not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":  true,
		"Message":  "Product found Successfull",
		"Error":    nil,
		"Products": products,
	})
}

func GetProductByName(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")

	product, err := usecase.GetProductByName(name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Product not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Product Details",
		"Error":   nil,
		"Product": product,
	})
}

func GetProductByCategoryID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	limitStr := ctx.DefaultQuery("limit", "0")
	pageStr := ctx.DefaultQuery("page", "0")

	id, _ := strconv.Atoi(idStr)
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	product, err := usecase.GetProductCategoryID(limit, offset, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Product not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":  true,
		"Message":  "Product Details",
		"Error":    nil,
		"Products": product,
	})
}

func GetBrandsUser(ctx *gin.Context) {
	brands, err := usecase.GetBrandsUser()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Brands not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Brands found",
		"Error":   nil,
		"Brands":  brands,
	})
}

func GetCategoriesUser(ctx *gin.Context) {
	categories, err := usecase.GetCategoriesUser()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Brands not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Brands found",
		"Error":   nil,
		"Brands":  categories,
	})
}

func GetProductByBrandName(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")
	product, err := usecase.GetProductByBrandName(name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Product not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":  true,
		"Message":  "Product found successfull",
		"Products": product,
	})
}

func GetProductByCategoryName(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")
	product, err := usecase.GetProductByCategoryName(name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Product Not  found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":  true,
		"Message":  "Product found successfull",
		"Error":    nil,
		"Products": product,
	})
}

// User-Carts----------------------------------------------------------------------------------------------
func AddToCart(ctx *gin.Context) {

	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Product add to cart failed",
			"Error":   err.Error(),
		})
		return
	}

	cartProduct := domain.Cart{}

	if err := ctx.ShouldBindJSON(&cartProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Bindin Error",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.AddToCart(&cartProduct, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Product add to cart failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"Message": "Product added to cart successfully",
		"Error":   nil,
		"CartProduct": gin.H{
			"ID":        cartProduct.ID,
			"UserID":    cartProduct.UserID,
			"ProductID": cartProduct.ProductID,
			"Quantity":  cartProduct.Quantity,
		},
	})

}

func CartList(ctx *gin.Context) {

	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "No products found",
			"Error":   err.Error(),
		})
		return
	}
	cartProduct, err := usecase.CartList(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Products not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusFound, gin.H{
		"Success":   true,
		"Message":   "Product Successfully found",
		"Error":     nil,
		"Cart List": cartProduct,
	})
}
func DeleteCartItem(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Failed to remove cart item",
			"Error":   err.Error(),
		})
		return
	}
	productIdstr := ctx.Param("id")
	productID, err := strconv.Atoi(productIdstr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Failed to remove cart item",
			"Error":   err.Error(),
		})
		return
	}

	err = usecase.DeleteCartItem(id, productID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Failed to remove cart item",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Item Successfully deleted from the cart",
		"Error":   nil,
	})
}

// User - Profile--------------------------------------------------------------------------------------------------
func AddAddress(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Add Address failed",
			"Error":   err.Error(),
		})
		return
	}
	userAdd := domain.Address{}
	if err := ctx.ShouldBindJSON(&userAdd); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Bindin Error",
			"Error":   err.Error(),
		})
		return
	}

	err = usecase.AddAddress(&userAdd, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Add Address failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Success": true,
		"Message": "Address successfully added",
		"Error":   nil,
		"Address": gin.H{
			"Id":       userAdd.ID,
			"UserID":   userAdd.UserID,
			"Street":   userAdd.Street,
			"City":     userAdd.City,
			"Pin Code": userAdd.PINCode,
			"State":    userAdd.State,
			"Country":  userAdd.Country,
		},
	})
}

func EditUserProfile(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Update user profile failed",
			"Error":   err.Error(),
		})
		return
	}
	user := domain.UserProfileUpdate{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding error",
			"Error":   err.Error(),
		})
		return
	}
	updatedUser, err := usecase.EditUserProfile(user, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Update user profile failed",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success":         true,
		"Message":         "User profile updated successfully",
		"Error":           nil,
		"Updated profile": updatedUser,
	})
}

func ForgotPassword(ctx *gin.Context) {
	email := ctx.DefaultQuery("email", "")
	err := usecase.ForgotPassword(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Failed to send otp",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Otp send successfully",
		"Error":   false,
	})
}

func ResetPassword(ctx *gin.Context) {
	email := ctx.DefaultQuery("email", "")
	data := domain.ResetPassword{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding error",
			"Error":   err.Error(),
		})
		return
	}
	err := usecase.ResetPassword(data, email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Reset password failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":      true,
		"Message":      "Reset password successfull",
		"Error":        false,
		"New Password": data.NewPassword,
	})
}

func ProfileDetails(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Fetching user profile details failed",
			"Error":   err.Error(),
		})
		return
	}
	userDetails, err := usecase.ProfileDetails(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Details not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":         true,
		"Message":         "User profile details found successfully",
		"Error":           nil,
		"Profile Details": userDetails,
	})
}

func ViewAddress(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View Address failed",
			"Error":   err.Error(),
		})
		return
	}
	userAdd, err := usecase.ViewAddress(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View Address failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":              true,
		"Message":              "Address found successfully",
		"Error":                nil,
		"User Address Details": userAdd,
	})
}

// User - Order---------------------------------------------------------------------------------------------------
func OrderSummary(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View Order Summary failed",
			"Error":   err.Error(),
		})
		return
	}
	res, err := usecase.GetOrderSummary(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Found Order Summary failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":       true,
		"Message":       "Found Order summary succeefull",
		"Error":         nil,
		"Order Summary": res,
	})
}

func OrderItem(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	coupon := ctx.DefaultQuery("coupon_name", "")
	err = usecase.OrderItem(id, coupon)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Order Successfull",
		"Error":   nil,
	})
}

func ViewOrders(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View Orders Failed",
			"Error":   err.Error(),
		})
		return
	}

	orders, err := usecase.ViewOrders(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View orders failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "View orders successfull",
		"Error":   nil,
		"Orders":  orders,
	})
}

func ViewOrdersByID(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View Order falied",
			"Error":   err.Error(),
		})
		return
	}
	orderIdStr := ctx.DefaultQuery("booking_id", "0")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View Order Failed",
			"Error":   err.Error(),
		})
		return
	}
	orders, err := usecase.ViewOrdersByID(id, orderId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View order failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":     true,
		"Message":     "View order successfull",
		"Error":       false,
		"Order Items": orders,
	})

}

func ViewOrdersByUserID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View orders failed",
			"Error":   err.Error(),
		})
		return
	}
	orders, err := usecase.ViewOrdersByUserID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View orders failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "View orders successfull",
		"Error":   false,
		"Orders":  orders,
	})
}

func ViewOrderItemsByUserID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View orders failed",
			"Error":   err.Error(),
		})
		return
	}
	orderIdStr := ctx.DefaultQuery("booking_id", "0")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View Order Failed",
			"Error":   err.Error(),
		})
		return
	}
	orders, err := usecase.ViewOrderItemsByUserID(id, orderId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View orders failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "View orders successfull",
		"Error":   false,
		"Orders":  orders,
	})
}

func OrderCancel(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order Cancellation failed",
			"Error":   err.Error(),
		})
		return
	}
	orderIdStr := ctx.DefaultQuery("booking_id", "0")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View Order Failed",
			"Error":   err.Error(),
		})
		return
	}
	orders, err := usecase.OrderCancel(id, orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Order Cancellation failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":      true,
		"Message":      "Order Cancelled Successfully",
		"Error":        false,
		"Order Status": orders.OrderStatus,
	})
}

func OrderReturn(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Return order failed",
			"Error":   err.Error(),
		})
		return
	}
	orderIdStr := ctx.DefaultQuery("booking_id", "0")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Return order failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.OrderReturn(orderId, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Return order failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Return request accepted",
		"Error":   false,
	})
}

func WalletPayment(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	id, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	coupon := ctx.DefaultQuery("coupon_name", "")
	err = usecase.WalletPayment(id, coupon)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Order Successfull",
		"Error":   nil,
	})
}
