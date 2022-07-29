package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lucianoortizsilva/frete/frete-api/adapters/http/controllers/model"
)

func SolicitarFrete(w http.ResponseWriter, r *http.Request) {

	// Ler o json recebido
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Converter o json recebido em um type struct
	var frete model.Frete
	if err = json.Unmarshal(body, &frete); err != nil {
		Erro(w, http.StatusBadRequest, err)
		return
	}

	// Validar campos de entrada
	if err = frete.Validar(); err != nil {
		Erro(w, http.StatusBadRequest, err)
		return
	}

	// TODO: SALVAR DA BASE DE DADOS A SOLICITAÇÃO DE FRETE
	frete.Status = "SOLICITADO"

	JSON(w, http.StatusCreated, frete)
}
