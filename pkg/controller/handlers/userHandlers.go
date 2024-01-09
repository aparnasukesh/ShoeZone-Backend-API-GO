package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

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

// User-Products----------------------------------------------------------------------------------------------------
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
			"Message": "Categoriess not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":    true,
		"Message":    "Categories found",
		"Error":      nil,
		"Categories": categories,
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
	err = usecase.AddToCart(cartProduct, id)
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

// User - Wish List------------------------------------------------------------------------------------------------
func AddToWishList(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	userId, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Invalid User id",
			"Error":   err.Error(),
		})
		return
	}
	idstr := ctx.Param("productid")
	productId, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Product add to wishlist failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.AddToWishList(userId, productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Product add to wishlist failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Product add to wishlist successfull",
		"Error":   false,
	})
}

func DeleteWishlistItem(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	userId, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Delete Item from wish list failed",
			"Error":   err.Error(),
		})
		return
	}
	idstr := ctx.Param("productid")
	productId, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Delete Item from wish list failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.DeleteWishlistItem(userId, productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Delete Item from wish list failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Delete item from wish list successfull",
		"Error":   false,
	})

}

func WishListItems(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	userId, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Failed to get wishlist items",
			"Error":   err.Error(),
		})
		return
	}
	products, err := usecase.WishListItems(userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Failed to get wishlist items",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":  true,
		"Message":  "List wishlist items successfull",
		"Error":    false,
		"Products": products,
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
func CartItemsOrderSummary(ctx *gin.Context) {
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
	res, err := usecase.GetCartItemsOrderSummary(id)
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

func OrderCartItems(ctx *gin.Context) {
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
	address := ctx.DefaultQuery("address_id", "0")
	addressId, err := strconv.Atoi(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
	}
	err = usecase.OrderCartItems(id, addressId, coupon)
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

func VerifyPayment(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	addressID, err := strconv.Atoi(ctx.Query("addressid"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	orderid := ctx.Query("order_id")
	paymentid := ctx.Query("payment_id")
	signature := ctx.Query("signature")
	totalamount := ctx.Query("total")
	coupon := ctx.Query("coupon")
	order_TableId := ctx.Query("order_TableId")

	fmt.Println("order_TableId :", order_TableId)
	fmt.Println("coupon:", coupon)
	fmt.Println("user id :", userID)
	fmt.Println("order id : ", orderid)
	fmt.Println("payment id  : ", paymentid)
	fmt.Println("signature: ", signature)
	fmt.Println("totalamount : ", totalamount)

	orderTableId, err := strconv.Atoi(order_TableId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order Failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.RazorpaySuccess(userID, orderTableId, signature, paymentid, orderid, coupon, addressID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Order Failed",
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

func RazorpaySuccess(ctx *gin.Context) {
	pid := ctx.Query("id")

	if pid == "" {
		ctx.JSON(400, gin.H{
			"Error": "Payment ID is missing",
		})
		return
	}

	temp, err := template.ParseFiles("./pkg/templates/success.html")
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
		return
	}
	data := map[string]interface{}{
		"paymentid": pid,
	}
	err = temp.Execute(ctx.Writer, data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
	}
}

func RazorPayFailed(ctx *gin.Context) {
	temp, err := template.ParseFiles("./pkg/templates/failed.html")
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
		return
	}
	err = temp.Execute(ctx.Writer, nil)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
	}
}

func OrderCartItemsRazorpay(ctx *gin.Context) {
	token, err := ctx.Cookie("UserAuthorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}

	userId, err := usecase.GetUserIDFromToken(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	coupon := ctx.DefaultQuery("coupon_name", "")
	address := ctx.DefaultQuery("address_id", "0")
	addressId, err := strconv.Atoi(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
	}
	paymentDetails, err := usecase.OrderCartItemsRazorpay(userId, coupon, addressId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	temp, err := template.ParseFiles("./pkg/templates/app.html")
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
		return
	}
	if strings.Trim(paymentDetails.Coupon, `""`) == "" {
		paymentDetails.Coupon = "No Coupon Applied"
	}

	data := map[string]interface{}{
		"userid":        paymentDetails.UserID,
		"address_id":    paymentDetails.AddressID,
		"totalprice":    paymentDetails.TotalAmount,
		"paymentid":     paymentDetails.PaymentID,
		"coupon":        paymentDetails.Coupon,
		"order_TableId": paymentDetails.Order_TableID,
	}

	err = temp.Execute(ctx.Writer, data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": err,
		})
	}
}

func OrderItemByID(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	userId, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	productidStr := ctx.DefaultQuery("productid", "0")
	productId, err := strconv.Atoi(productidStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	coupon := ctx.DefaultQuery("coupon_name", "")
	quantityStr := ctx.DefaultQuery("quantity", "0")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	address := ctx.DefaultQuery("address_id", "0")
	addressId, err := strconv.Atoi(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
	}
	err = usecase.OrderItemByID(userId, productId, quantity, coupon, addressId)
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
		"Message": "Order successfull",
		"Error":   false,
	})

}

func OrderSummary(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	userId, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Failed to get order summary",
			"Error":   err.Error(),
		})
		return
	}
	orderidStr := ctx.DefaultQuery("order_id", "0")
	orderId, err := strconv.Atoi(orderidStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Failed to get order summary",
			"Error":   err.Error(),
		})
		return
	}
	orderSummary, err := usecase.OrderSummary(userId, orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Failed to get order summary",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":       true,
		"Message":       "Get order summary successfull",
		"Error":         false,
		"Order Summary": orderSummary,
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
	userId, err := usecase.GetUserIDFromToken(authorization)
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
			"Message": "Order Cancellation failed",
			"Error":   err.Error(),
		})
		return
	}
	orders, err := usecase.OrderCancel(userId, orderId)
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

func AdminOrderCancel(ctx *gin.Context) {
	userIdStr := ctx.DefaultQuery("user_id", "0")
	userId, err := strconv.Atoi(userIdStr)
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
			"Message": "Order Cancellation failed",
			"Error":   err.Error(),
		})
		return
	}
	orders, err := usecase.OrderCancel(userId, orderId)
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

func WalletPaymentCartItems(ctx *gin.Context) {
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
	address := ctx.DefaultQuery("address_id", "0")
	addressId, err := strconv.Atoi(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
	}
	err = usecase.WalletPaymentCartItems(id, coupon, addressId)
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

func WalletPaymentOrderItemByID(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	userId, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	productidStr := ctx.DefaultQuery("productid", "0")
	productId, err := strconv.Atoi(productidStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	coupon := ctx.DefaultQuery("coupon_name", "")
	quantityStr := ctx.DefaultQuery("quantity", "0")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
		return
	}
	address := ctx.DefaultQuery("address_id", "0")
	addressId, err := strconv.Atoi(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Order failed",
			"Error":   err.Error(),
		})
	}
	err = usecase.WalletPaymentOrderItemByID(userId, productId, quantity, coupon, addressId)
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
		"Message": "Order successfull",
		"Error":   false,
	})

}

// User - Invoice -------------------------------------------------------------------------------------------------
func InvoiceDetails(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	userId, err := usecase.GetUserIDFromToken(authorization)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Failed to get invoice details",
			"Error":   err.Error(),
		})
		return
	}
	orderIdStr := ctx.DefaultQuery("order_id", "0")
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Failed to get invoice details",
			"Error":   err.Error(),
		})
		return
	}
	details, err := usecase.InvoiceDetails(userId, orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Failed to get invoice details",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":         true,
		"Message":         "Successfully get invoice details",
		"Error":           false,
		"Invoice Details": details,
	})

}

func InvoiceDownload(ctx *gin.Context) {

	ctx.Header("Content-Disposition", "attachment; filename=invoice.pdf")
	ctx.Header("Content-Type", "application/pdf")
	ctx.File("./data/invoice.pdf")

}

// Admin - Sales Report-------------------------------------------------------------------------------------------
func SalesReport(ctx *gin.Context) {
	fromDate := ctx.DefaultQuery("from_date", "")
	toDate := ctx.DefaultQuery("to_date", "")

	parseFromDate, err := time.Parse("2006-01-02", fromDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Getting Sales report failed",
			"Error":   err.Error(),
		})
		return
	}
	parseToDate, err := time.Parse("2006-01-02", toDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Getting Sales report failed",
			"Error":   err.Error(),
		})
		return
	}

	salesDetails, err := usecase.SalesReport(parseFromDate, parseToDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Failed to get sales report details",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":              true,
		"Message":              "Successfully get sales report",
		"Error":                false,
		"Sales Report Details": salesDetails,
	})
}

func SalesReportDownload(ctx *gin.Context) {

	ctx.Header("Content-Disposition", "attachment; filename=salesReport.pdf")
	ctx.Header("Content-Type", "application/pdf")
	ctx.File("./data/salesReport.pdf")

}
