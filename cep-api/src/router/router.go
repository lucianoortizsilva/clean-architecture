package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucianoortizsilva/cep/cep-api/src/controller"
)

func InicializarRotas() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cep/{codigo}", controller.BuscarCep).Methods(http.MethodGet)
	return r
}
