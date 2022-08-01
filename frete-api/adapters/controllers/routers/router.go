package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucianoortizsilva/frete/frete-api/adapters/controllers/frete"
)

func InicializarRotas() *mux.Router {
	r := mux.NewRouter()

	freteController := frete.FreteController{}

	r.HandleFunc("/v1/fretes", freteController.SolicitarFrete).Methods(http.MethodPost)

	return r
}
