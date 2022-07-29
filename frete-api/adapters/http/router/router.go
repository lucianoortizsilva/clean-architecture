package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucianoortizsilva/frete/frete-api/adapters/http/controllers"
)

func InicializarRotas() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/fretes", controllers.SolicitarFrete).Methods(http.MethodPost)
	return r
}
