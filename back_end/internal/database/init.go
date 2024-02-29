package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/sijms/go-ora/v2"
	"log"
)

const FilePath = "./configs/db_config.json"

var DB, _ = InitDataBase()

func InitDataBase() (interface{}, error) {
	configDBConnect := Init()
	err := configDBConnect.Get(FilePath)
	if err != nil {
		log.Fatal("Не удалось считать данные подключения к БД: ", err)
	}
	if configDBConnect.DataBaseManagementSystem == "Postgres" {
		connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", configDBConnect.User, configDBConnect.Password,
			configDBConnect.Host, configDBConnect.Port, configDBConnect.DBName)
		conn, err := pgx.Connect(context.Background(), connectionString)
		if err != nil {
			log.Fatal("Не удалось подключитьсяк базе данных: ", err)
		}
		return conn, nil
	} else {
		connectionString := fmt.Sprintf("oracle://%s:%s@%s:%d/%s", configDBConnect.User, configDBConnect.Password,
			configDBConnect.Host, configDBConnect.Port, configDBConnect.DBName)
		conn, err := sql.Open("oracle", connectionString)
		if err != nil {
			log.Fatal("Не удалось подключитьсяк базе данных: ", err)
		}
		return conn, nil
	}
}