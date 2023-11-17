package handler

import (
	"net/http"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func AddProduct(ctx *gin.Context) {

	productData := domain.Product{}

	if err := ctx.ShouldBindJSON(&productData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Succes":   false,
			"Messsage": "Binding Error",
			"Error":    err.Error(),
		})
		return
	}
	err := usecase.AddProduct(&productData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Product add failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Success": true,
		"Message": "Product Successfully Added",
		"Error":   nil,
		"Product": productData,
	})

}
