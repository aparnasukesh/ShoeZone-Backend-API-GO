package handler

import (
	"strings"

	"github.com/aparnasukesh/shoezone/pkg/util"
	"github.com/gin-gonic/gin"
)

func AdminAuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		//response unthorized
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		//Unathorized
	}

	verifiedToken, err := util.VerifyJWT(token[1])
	if err != nil {

	}

	isAdmin, err := util.Getrole(verifiedToken)
	if err != nil {
		//getrole errorr
	}

	if !isAdmin {
		//Unathorized
		return
	}
	ctx.Next()
}
