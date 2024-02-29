package handlers

import (
	"back/internal/transport/rest/mw"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Возвращает пользователя
// @Description Возвращает данные пользователя
// @Tags User (auth)
// @Param auth_id body string true "AuthId"
// @Param password_web body string true "PasswordWeb"
// @Success 200 {object} mw.User
// @Failure 404 {object} mw.HTTPError
// @Router /auth/login [post]
func CheckUser(ctx *gin.Context) {
	var u mw.UserData
	if err := ctx.BindJSON(&u); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	data, err := mw.CheckUser(u)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Чтение не произошло",
		})
		return
	} else {
		ctx.SetCookie("my_jwt_token", data.Token, 3600, "/", "localhost", false, true)
		ctx.Header("Authorization", data.Token)
		ctx.IndentedJSON(http.StatusOK, data)
		return
	}
}
