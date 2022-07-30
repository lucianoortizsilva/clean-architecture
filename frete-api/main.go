package main

import (
	"fmt"
	"net/http"

	"github.com/lucianoortizsilva/frete/frete-api/adapters/controllers/configuration"
)

func main() {
	configuration.LoadEnvHttp()
	rotas := configuration.InicializarRotas()

	fmt.Printf("App frete-api Iniciada na porta: %d\n", configuration.PORT_HTTP)

	http.ListenAndServe(fmt.Sprintf(":%d", configuration.PORT_HTTP), rotas)
}
