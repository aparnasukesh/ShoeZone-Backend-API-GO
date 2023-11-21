package handler

import (
	"net/http"
	"strconv"

	"github.com/aparnasukesh/shoezone/pkg/domain"
	"github.com/aparnasukesh/shoezone/pkg/usecase"
	"github.com/gin-gonic/gin"
)

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

func GetUserByID(ctx *gin.Context) {
	idstr := ctx.Param("id")
	Id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Get User Failed",
			"Error":   err.Error(),
		})
		return
	}
	var res *domain.User
	res, err = usecase.GetUserByID(Id)
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"Success": false,
			"Message": "User not found ",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success":  true,
		"Message":  "Get User Successful",
		"Error":    nil,
		"UserData": res,
	})

}

func BlockUser(ctx *gin.Context) {
	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Get User Failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.BlockUser(id)
	if err != nil {
		ctx.JSON(http.StatusAlreadyReported, gin.H{
			"Success": false,
			"Message": "Block user failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "User blocked successfully",
		"Error":   nil,
	})
	return

}
func UnblockUser(ctx *gin.Context) {

	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Get User Failed",
			"Error":   err.Error(),
		})
		return
	}
	err = usecase.UnblockUser(id)
	if err != nil {
		ctx.JSON(http.StatusAlreadyReported, gin.H{
			"Success": false,
			"Message": "UnBlock user failed",
			"Error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Success": true,
		"Message": "User Unblocked successfully",
		"Error":   nil,
	})
	return

}
