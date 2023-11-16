package controller

import (
	handler "github.com/aparnasukesh/shoezone/pkg/controller/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.POST("/user/register", handler.RegisterUser)
	r.POST("/user/register/validate", handler.RegisterValidtae)
}
