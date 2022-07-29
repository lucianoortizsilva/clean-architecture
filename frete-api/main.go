package main

import (
	"fmt"
	"net/http"

	"github.com/lucianoortizsilva/frete/frete-api/adapters/http/config"
	"github.com/lucianoortizsilva/frete/frete-api/adapters/http/router"
)

func main() {
	config.InicializarVariaveisAmbiente()
	rotas := router.InicializarRotas()

	fmt.Printf("App frete-api Iniciada na porta: %d\n", config.PORT_HTTP)

	http.ListenAndServe(fmt.Sprintf(":%d", config.PORT_HTTP), rotas)
}
