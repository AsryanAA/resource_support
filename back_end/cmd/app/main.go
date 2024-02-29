package main

import (
	"back/internal/app"
	"log"
)

// точка входа в приложение
func main() {
	err := app.Run()
	if err != nil {
		log.Fatalf("can not start the resource support service, %s", err)
	}
}