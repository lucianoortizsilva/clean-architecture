package main

import (
	"fmt"
	"net/http"

	"github.com/lucianoortizsilva/cep/cep-api/src/config"
)

func main() {
	config.InicializarVariaveisAmbiente()
	fmt.Printf("CEP-API Iniciado na porta: %d\n", config.PORT_HTTP)
	http.ListenAndServe(fmt.Sprintf(":%d", config.PORT_HTTP), nil)
}
