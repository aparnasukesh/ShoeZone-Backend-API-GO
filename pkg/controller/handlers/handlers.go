package handler

import (
	"net/http"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	userDate := domain.User{}

	if err := ctx.ShouldBindJSON(&userDate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding error",
			"Error":   err,
		})
		return
	}

	err := usecase.RegisterUser(&userDate)
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

func RegisterValidtae(ctx *gin.Context) {
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
