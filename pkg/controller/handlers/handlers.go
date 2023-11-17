package handler

import (
	"fmt"
	"net/http"

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
		ctx.JSON(http.StatusNotFound, gin.H{
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
		ctx.JSON(http.StatusNotFound, gin.H{
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
		"Data":    userData,
	})
}

func UserLogin(ctx *gin.Context) {
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

	fmt.Println("======================================", res.Isadmin)

	token, err := util.GenerateJWT(*res)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
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
func GetUsers(ctx *gin.Context) {
	res, err := usecase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Get users failed",
			"Error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Success":   true,
		"Message":   "User Details",
		"Error":     nil,
		"User Data": res,
	})

}

// func GetUserByID(ctx *gin.Context) {

// }
