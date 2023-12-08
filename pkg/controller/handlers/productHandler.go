package handler

import (
	"net/http"
	"strconv"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/usecase"
	"github.com/gin-gonic/gin"
)

// Admin- Category Handlers---------------------------------------------------------------------------
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
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

func GetCategories(ctx *gin.Context) {
	categories, err := usecase.GetCategories()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "No categories found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":    true,
		"Message":    "Category deatils",
		"Error":      nil,
		"Categories": categories,
	})
}

// Admin-Brand Handlers---------------------------------------------------------------------------------

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
		ctx.JSON(http.StatusInternalServerError, gin.H{
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

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
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

func GetBrands(ctx *gin.Context) {
	brands, err := usecase.GetBrands()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Brands not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":       true,
		"Message":       "Brand Details",
		"Error":         nil,
		"Brand Details": brands,
	})

}

// Admin-Product Handlers--------------------------------------------------------------------------------
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
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

// Admin - Coupon Handlers----------------------------------------------------------------------------------------
func AddCoupon(ctx *gin.Context) {
	newCoupon := domain.Coupon{}

	if err := ctx.ShouldBindJSON(&newCoupon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding Error",
			"Error":   err.Error(),
		})
		return
	}
	err := usecase.AddCoupon(&newCoupon)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Add Coupon failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Success": true,
		"Message": "Add coupon successfull",
		"Error":   false,
	})
}

func DeleteCoupon(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Delete coupon failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.DeleteCoupon(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Delete coupon failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{
		"Success": true,
		"Message": "Coupon deleted succesffully",
		"Error":   false,
	})
}

func UpdateCoupon(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Update coupon failed",
			"Error":   err.Error(),
		})
		return
	}
	updateCoupon := domain.Coupon{}
	if err := ctx.ShouldBindJSON(&updateCoupon); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Binding Error",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.UpdateCoupon(updateCoupon, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Update coupon failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "Coupon updated successfully",
		"Error":   false,
	})
}

func ViewCoupons(ctx *gin.Context) {
	coupons, err := usecase.ViewCoupons()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View coupons failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "View coupons successfull",
		"Error":   false,
		"Coupons": coupons,
	})
}
