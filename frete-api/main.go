package main

import (
	"fmt"
	"net/http"

	"github.com/lucianoortizsilva/frete/frete-api/src/config"
	"github.com/lucianoortizsilva/frete/frete-api/src/router"
)

func main() {
	config.InicializarVariaveisAmbiente()
	rotas := router.InicializarRotas()

	fmt.Printf("CEP-API Iniciado na porta: %d\n", config.PORT_HTTP)
	http.ListenAndServe(fmt.Sprintf(":%d", config.PORT_HTTP), rotas)
}
