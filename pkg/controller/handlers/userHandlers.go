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
		//"Data":    userData,
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

// User-Products------------------------------------------------
func GetProducts(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "0")
	limitStr := ctx.DefaultQuery("limit", "0")
	pageNum, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "No product found",
			"Error":   err.Error(),
		})
		return
	}
	product, err := usecase.GetProductByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
	idStr := ctx.DefaultQuery("id", "0")
	limitStr := ctx.DefaultQuery("limit", "0")
	pageStr := ctx.DefaultQuery("page", "0")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Products not found",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":  true,
		"Message":  "Products Details",
		"Error":    nil,
		"Products": products,
	})
}

func GetProductByName(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "")

	product, err := usecase.GetProductByName(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
	idStr := ctx.DefaultQuery("id", "0")
	limitStr := ctx.DefaultQuery("limit", "0")
	pageStr := ctx.DefaultQuery("page", "0")

	id, _ := strconv.Atoi(idStr)
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	product, err := usecase.GetProductCategoryID(limit, offset, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
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
		ctx.JSON(http.StatusBadRequest, gin.H{
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
