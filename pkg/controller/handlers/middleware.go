package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aparnasukesh/shoezone/pkg/util"
	"github.com/gin-gonic/gin"
)

func AdminAuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		// Response unthorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization Failed",
			"Error":   "Authorization header is missing",
		})
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		// Response Unathorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization Failed",
			"Error":   "Bearer token is missing or malformed",
		})
		return
	}

	verifiedToken, err := util.VerifyJWT(token[1])
	if err != nil {
		// JWT token verification failed
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization Failed",
			"Error":   err.Error(),
		})
		return
	}

	isAdmin, err := util.Getrole(verifiedToken)
	if err != nil {
		// Getrole error
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   err.Error(),
		})
		return

	}

	admin := fmt.Sprintf("%v", isAdmin)
	fmt.Println(isAdmin)

	if admin != "true" {
		// Unathorized
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization Failed",
			"Error":   "User is not an Admin",
		})
		return
	}
	ctx.Next()
}
