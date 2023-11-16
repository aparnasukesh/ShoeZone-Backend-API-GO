package controller

import (
	handler "github.com/aparnasukesh/shoezone/pkg/controller/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", handler.RegisterUser)
		userGroup.POST("/register/validate", handler.RegisterValidate)
		userGroup.POST("/login", handler.UserLogin)
	}

	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/user", handler.AdminAuthRequired)
		adminGroup.GET("/user/:id")
	}

}
