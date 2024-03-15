package mw

import (
	"back/internal/database"
	"back/internal/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"net/http"
)

func CheckUser(u models.UserData) (models.User, error) {
	var user models.User
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		row := connPostgres.QueryRow(context.Background(), `SELECT AUTHID, PASSWORD_WEB, "NAME", LICENSE 
															FROM parus.userlist
															WHERE AUTHID = $1 AND PASSWORD_WEB = $2
															`, u.Login, u.Password)
		err := row.Scan(&user.Login, &user.Password, &user.Name, &user.License)
		if err != nil {
			fmt.Println("Ошибка получения данных: ", err)
		}
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			row := connOracle.QueryRow(`SELECT AUTHID, PASSWORD_WEB, NAME, LICENSE 
															FROM parus.userlist
															WHERE AUTHID = :1 AND PASSWORD_WEB = :2
															`, u.Login, u.Password)
			err := row.Scan(&user.Login, &user.Password, &user.Name, &user.License)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
				return user, err
			}
		}
	}

	return user, nil
}

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t, err := ctx.Cookie("jwt_token")
		fmt.Println(t)
		if err != nil {
			fmt.Println("Пожалуйста авторизуйтесь", err)
			_ = ctx.AbortWithError(http.StatusUnauthorized, err)
		}
		ctx.Next()
	}
}

func ValidationToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("jwt_token")
		fmt.Println(token)
		if token == "" {
			fmt.Println("Пожалуйста авторизуйтесь")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Next()
	}
}
