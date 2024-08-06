package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		// if rota.RequerAutenticacao {
		// 	r.HandleFunc(
		// 		rota.URI,
		// 		middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
		// 	).Methods(rota.Metodo)
		// } else {
		// 	r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		// }

	}

	return router
}
