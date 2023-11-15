package handler

import (
	"net/http"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func SignUpUser(ctx *gin.Context) {
	userDate := domain.User{}

	if err := ctx.ShouldBindJSON(&userDate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding error",
			"Error":   err,
		})
		return
	}

	err := usecase.CreateUser(&userDate)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Registering the User Failed",
			"Error":   err,
		})
	}
}
