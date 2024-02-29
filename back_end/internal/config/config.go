package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

const ServerConfigPath = "./configs/server.yaml"

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func MustLoad() *Server {
	if _, err := os.Stat(ServerConfigPath); os.IsNotExist(err) {
		log.Fatalf("server config file %s does not exists", ServerConfigPath)
	}

	var srv Server
	if err := cleanenv.ReadConfig(ServerConfigPath, &srv); err != nil {
		log.Fatalf("can not read server config %s", err)
	}

	return &srv
}