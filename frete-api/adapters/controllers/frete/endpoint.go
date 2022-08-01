package frete

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lucianoortizsilva/frete/frete-api/adapters/controllers"
	"github.com/lucianoortizsilva/frete/frete-api/adapters/repository"
	"github.com/lucianoortizsilva/frete/frete-api/drivers"
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

	// Conectar no banco de dados mongodb
	clientMongoDB, err := drivers.ConectarServidorMongoDB()
	if err != nil {
		controllers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repository := repository.NewSolicitaFreteDB(clientMongoDB)
	useCaseSolicitaFrete := usecases.NewSolicitaFrete(repository)

	var input usecases.SolicitaFreteDtoInput
	input.PesoTotalPedido = dto.PesoTotalPedido
	input.PedidoId = dto.PedidoId
	input.Codigo = dto.Codigo
	input.Regiao = dto.Regiao
	input.Cep = dto.Cep

	output, err := useCaseSolicitaFrete.Executar(input)

	if err == nil {
		controllers.JSON(w, http.StatusCreated, output)
	} else {
		controllers.Erro(w, http.StatusConflict, err)
	}

}
