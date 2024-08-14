package main

import (
	"api/cmd/config"
	"api/cmd/router"
	"fmt"
	"log"
	"net/http"

	_ "api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	config.Carregar()
	r := router.Gerar()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Printf("Escutando na porta %d", config.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
