package handlers

import (
	"back/internal/models"
	"back/internal/transport/rest/mw"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

const validityJWT = 42 // в секундах

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

// SingIn @Summary Возвращает пользователя
// @Description Возвращает данные пользователя
// @Tags User (auth)
// @Param auth_id body string true "AuthId"
// @Param password_web body string true "PasswordWeb"
// @Success 200 {object} mw.User
// @Failure 404 {object} mw.HTTPError
// @Router /auth/login [post]
func SingIn(ctx *gin.Context) {
	var u models.UserData
	if err := ctx.BindJSON(&u); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}
	user, err := mw.CheckUser(u)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
			"errorMessage": err,
			"error":        "Вы не авторизованы",
		})
		return
	}

	expirationTime := time.Now().Add(validityJWT * time.Second)

	claims := &Claims{
		Username: user.Login,
		Password: user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"errorMessage": err,
			"error":        "Некорректные данные",
		})
		return
	}

	// TODO
	ctx.SetCookie("jwt_token", tokenString, validityJWT, "/", "localhost", false, true)
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"jwt_token": tokenString,
		"error":     "",
	})
	return
}
