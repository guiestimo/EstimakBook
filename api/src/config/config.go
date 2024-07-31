package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	ApiPort          = 0
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	ApiPort, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		ApiPort = 9000 // Porta default
	}

	ConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}
