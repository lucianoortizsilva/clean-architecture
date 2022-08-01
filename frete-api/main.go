package main

import (
	"fmt"
	"net/http"

	"github.com/lucianoortizsilva/frete/frete-api/adapters/controllers/routers"
	"github.com/lucianoortizsilva/frete/frete-api/drivers"
	"github.com/lucianoortizsilva/frete/frete-api/drivers/configuration"
)

func main() {
	configuration.CarregarVariaveisAmbiente()
	drivers.ConectarServidorMongoDB()
	rotas := routers.InicializarRotas()
	fmt.Printf("\nfrete-api Iniciada na porta: %d\n", configuration.PORT_HTTP)
	http.ListenAndServe(fmt.Sprintf(":%d", configuration.PORT_HTTP), rotas)
}
