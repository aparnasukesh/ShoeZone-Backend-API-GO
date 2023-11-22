package handler

import (
	"net/http"
	"strconv"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/usecase"
	"github.com/gin-gonic/gin"
)

// Category Handlers---------------------------------------------------------------------------
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
	updateCategory := domain.Category{}
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

	if err = ctx.ShouldBindJSON(&updateCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Category update failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.UpdateCategory(updateCategory, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Category update failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":          true,
		"Message":          "Category successfully updated",
		"Error":            nil,
		"Updated Category": updateCategory,
	})
}

func DeleteCategory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Detele category failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.DeleteCategory(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": true,
			"Message": "Category delete failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Category Deleted",
		"Error":   nil,
	})
}

// Brand Handlers---------------------------------------------------------------------------------

func AddBrand(ctx *gin.Context) {
	newBrand := domain.Brand{}

	if err := ctx.ShouldBindJSON(&newBrand); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Succcess": false,
			"Message":  "Binding Error",
			"Error":    err.Error(),
		})
		return
	}
	err := usecase.AddBrand(newBrand)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Brand Create failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Success":   true,
		"Message":   "Brand successfully created",
		"Error":     nil,
		"New Brand": newBrand,
	})
}

func UpdateBrand(ctx *gin.Context) {
	updatedBrand := domain.Brand{}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Brand Update failed",
			"Error":   err.Error(),
		})
		return
	}

	if err = ctx.ShouldBindJSON(&updatedBrand); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding error",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.UpdateBrand(updatedBrand, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Brand update failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":       true,
		"Message":       "Brand successfully updated",
		"Error":         nil,
		"Updated Brand": updatedBrand,
	})
}

func DeleteBrand(ctx *gin.Context) {
	//brandData := domain.Brand{}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Delete brand failed",
			"Error":   err.Error(),
		})
		return
	}

	err = usecase.DeleteBrand(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Delete brand error",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{
		"Success": true,
		"Message": "Brand successfully deleted",
		"Error":   nil,
	})
}

// Product Handlers--------------------------------------------------------------------------------
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

func UpdateProduct(ctx *gin.Context) {
	updatedProduct := domain.Product{}
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Product Update failed",
			"Error":   err.Error(),
		})
		return
	}
	if err = ctx.ShouldBindJSON(&updatedProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding Error",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.UpdateProduct(updatedProduct, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Product update failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":         true,
		"Message":         "Product successfully updated",
		"Error":           nil,
		"Updated Product": updatedProduct,
	})

}

func DeleteProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Product delete failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Product delete failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Product successfully deleted",
		"Error":   nil,
	})
}
