package main

import (
	"fmt"
	"net/http"

	"github.com/lucianoortizsilva/frete/frete-api/adapters/controllers/configuration"
	"github.com/lucianoortizsilva/frete/frete-api/drivers"
)

func main() {

	rotas := configuration.InicializarRotas()

	drivers.ConectarServidorMongoDB()

	var PORT_HTTP = 8080

	fmt.Printf("App frete-api Iniciada na porta: %d\n", PORT_HTTP)

	http.ListenAndServe(fmt.Sprintf(":%d", PORT_HTTP), rotas)
}
