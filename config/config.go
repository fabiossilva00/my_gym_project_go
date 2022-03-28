package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	URLConexaoDB string
	NameDatabase string
	Porta        string
}

const uriDB = "mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority"
const portaEcho = ":%s"

func LoadConfig() *Configuration {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Erro ao carregar godotenv:", err)
	}

	porta := fmt.Sprintf(portaEcho,
		os.Getenv("API_PORT"),
	)

	nameDb := os.Getenv("DB_DATABASE")

	uri := fmt.Sprintf(uriDB,
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	return &Configuration{
		URLConexaoDB: uri,
		NameDatabase: nameDb,
		Porta:        porta,
	}
}
