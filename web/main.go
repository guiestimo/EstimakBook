package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
)

func main() {
	fmt.Println("Rodando WebAPp!")

	r := router.Gerar()
	//fmt.Printf("Escutando na porta %d", ":3000")
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", ":3000"), r))
	log.Fatal(http.ListenAndServe(":3000", r))
}
