package frete

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lucianoortizsilva/frete/frete-api/adapters/controllers"
	"github.com/lucianoortizsilva/frete/frete-api/adapters/mongodb"
	"github.com/lucianoortizsilva/frete/frete-api/usecases"
)

type FreteController struct {
	UseCaseSolicitaFrete usecases.SolicitaFrete
}

func (controller FreteController) SolicitarFrete(w http.ResponseWriter, r *http.Request) {

	// Ler o json recebido
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		controllers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Converter o json recebido em um type struct
	var dto FreteDTO
	if err = json.Unmarshal(body, &dto); err != nil {
		controllers.Erro(w, http.StatusBadRequest, err)
		return
	}

	// Validar campos de entrada
	if err = dto.Validar(); err != nil {
		controllers.Erro(w, http.StatusBadRequest, err)
		return
	}

	var input usecases.SolicitaFreteDtoInput
	input.Regiao = dto.Regiao
	input.PesoTotalPedido = dto.PesoTotalPedido

	var repository = mongodb.NewSolicitaFreteDB()

	solicitaFrete := usecases.NewSolicitaFrete(repository)

	output, err := solicitaFrete.Executar(input)

	controllers.JSON(w, http.StatusCreated, output)
}
