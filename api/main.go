package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	fmt.Printf("Escutando na porta %d", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
