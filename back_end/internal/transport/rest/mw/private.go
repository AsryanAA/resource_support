package mw

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("jwt_token")
		if err != nil {
			fmt.Println("Пользователь не авторизован", err)
			_ = ctx.AbortWithError(http.StatusUnauthorized, err)
		}
		ctx.Next()
	}
}
