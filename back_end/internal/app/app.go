package app

import (
	"back/internal/config"
	"back/internal/transport/rest"
)

func Run() error {
	srv := config.MustLoad()
	err := rest.InitServer(srv.Host, srv.Port)

	return err
}