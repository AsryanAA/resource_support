package database

import (
	"encoding/json"
	"log"
	"os"
)

type ConnectDB struct {
	DataBaseManagementSystem string `json:"dbms"` // Postgres or Oracle
	Host                     string `json:"host"`
	Port                     int    `json:"port"`
	DBName                   string `json:"db_name"`
	User                     string `json:"user"`
	Password                 string `json:"password"`
}

// Init constructor ConnectDB
func Init() *ConnectDB {
	return &ConnectDB{}
}

func (db *ConnectDB) Get(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("can not file connect config file: ", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal("can not close connect config file: ", err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&db)
	if err != nil {
		log.Fatal("can not decode connect config file: ", err)
	}

	return nil
}