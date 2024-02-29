package mw

import (
	"back/internal/database"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type UserData struct {
	AuthId      string `json:"auth_id"`
	PasswordWeb string `json:"password_web"`
}

type User struct {
	UserData `json:"user_data"`
	Name     string `json:"name"`
	License  int    `json:"license"`
	Token    string `json:"jwt_token"`
}

type HTTPError struct {
	ErrorCode    int
	ErrorMessage string
}

func CheckUser(u UserData) (User, error) {
	var user User
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		row := connPostgres.QueryRow(context.Background(), `SELECT AUTHID, PASSWORD_WEB, "NAME", LICENSE 
															FROM parus.userlist
															WHERE AUTHID = $1 AND PASSWORD_WEB = $2
															`, u.AuthId, u.PasswordWeb)
		err := row.Scan(&user.AuthId, &user.PasswordWeb, &user.Name, &user.License)
		if err != nil {
			fmt.Println("Ошибка получения данных: ", err)
		}
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			row := connOracle.QueryRow(`SELECT AUTHID, PASSWORD_WEB, NAME, LICENSE 
															FROM parus.userlist
															WHERE AUTHID = :1 AND PASSWORD_WEB = :2
															`, u.AuthId, u.PasswordWeb)
			err := row.Scan(&user.AuthId, &user.PasswordWeb, &user.Name, &user.License)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}
		}
	}
	user.Token = "there is generate jwt token"

	return user, nil
}
