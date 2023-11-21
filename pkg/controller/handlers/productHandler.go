package handler

import (
	"net/http"
	"strconv"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/usecase"
	"github.com/gin-gonic/gin"
)

func AddCategories(ctx *gin.Context) {
	category := domain.Category{}

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding Error",
			"Error":   err.Error(),
		})
		return
	}
	err := usecase.AddCategories(&category)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Category add failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":  true,
		"Message":  "Category Successfully added",
		"Error":    nil,
		"Category": category,
	})
}

func UpdateCategory(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Category update failed",
			"Error":   err.Error(),
		})
		return
	}
	usecase.UpdateCategory(id)
}
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
