package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApiUrl   = ""   // Representa a URL para comunicação com a API
	ApiPort  = 0    // Porta onde a aplicação web está rodando
	HashHey  []byte // HashKey é utilizada para autenticar o cookie
	BlockKey []byte // BlockKey é utilizada para criptografar os dados do cookie
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	ApiPort, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	ApiUrl = os.Getenv("API_URL")
	HashHey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
