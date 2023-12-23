package handler

import (
	"net/http"
	"strings"

	"github.com/aparnasukesh/shoezone/pkg/util"
	"github.com/gin-gonic/gin"
)

// AdminAuthRequired is a middleware to check if the request has admin authorization.
func AdminAuthRequired(ctx *gin.Context) {
	// Retrieve the Authorization header from the request
	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		// Respond with unauthorized if Authorization header is missing
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization Failed",
			"Error":   "Authorization header is missing",
		})
		ctx.Abort() // Abort further processing
		return
	}

	// Split the Authorization header to get the token
	tokenParts := strings.Split(authorization, "Bearer ")
	if len(tokenParts) < 2 {
		// Respond with unauthorized if Bearer token is missing or malformed
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization Failed",
			"Error":   "Bearer token is missing or malformed",
		})
		ctx.Abort() // Abort further processing
		return
	}

	// Verify the JWT token
	verifiedToken, err := util.VerifyJWT(tokenParts[1])
	if err != nil {
		// Respond with unauthorized if JWT token verification fails
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization Failed",
			"Error":   err.Error(),
		})
		ctx.Abort() // Abort further processing
		return
	}

	// Get the role from the token
	isAdmin, err := util.GetRole(verifiedToken)
	if err != nil {
		// Respond with unauthorized if there is an error getting the role
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   err.Error(),
		})
		ctx.Abort() // Abort further processing
		return
	}

	// Check if the user is an admin
	if isAdmin != true {
		// Respond with unauthorized if the user is not an admin
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization Failed",
			"Error":   "User is not an Admin",
		})
		ctx.Abort() // Abort further processing
		return
	}

	// If all checks pass, proceed to the next middleware or handler
	ctx.Next()
}

// UserAuthRequired is a middleware to check if the request has user authorization,
func UserAuthRequired(ctx *gin.Context) {

	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   "Authorization Header is missing",
		})
		ctx.Abort()
		return
	}

	tokenParts := strings.Split(authorization, "Bearer ")
	if len(tokenParts) < 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   "Bearer token is missing or malformed",
		})
		ctx.Abort()
		return
	}

	verifiedToken, err := util.VerifyJWT(tokenParts[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   err.Error(),
		})
		ctx.Abort()
		return
	}
	isAdmin, err := util.GetRole(verifiedToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   err.Error(),
		})
		ctx.Abort()
		return
	}
	if isAdmin != false {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   err.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}

func UserPaymentAuthorization(ctx *gin.Context) {
	token, err := ctx.Cookie("UserAuthorization")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   err.Error(),
		})
		ctx.Abort()
		return
	}
	tokenParts := strings.Split(token, "Bearer ")
	if len(tokenParts) < 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   "Bearer token is missing or malformed",
		})
		ctx.Abort()
		return
	}

	verifiedToken, err := util.VerifyJWT(tokenParts[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   err.Error(),
		})
		ctx.Abort()
		return
	}
	isAdmin, err := util.GetRole(verifiedToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   err.Error(),
		})
		ctx.Abort()
		return
	}
	if isAdmin != false {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Authorization failed",
			"Error":   err.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}
