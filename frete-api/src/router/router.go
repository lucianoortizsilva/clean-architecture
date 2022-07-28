package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucianoortizsilva/frete/frete-api/src/controller"
)

func InicializarRotas() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/fretes", controller.SolicitarFrete).Methods(http.MethodPost)
	return r
}
