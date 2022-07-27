package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CEP struct {
	ID         uint64 `json:"id,omitempty"`
	Logradouro string `json:"nome,omitempty"`
	Bairro     string `json:"nick,omitempty"`
	Cidade     string `json:"email,omitempty"`
	Estado     string `json:"senha,omitempty"`
}

func BuscarCep(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	codigo, erro := strconv.ParseUint(parametros["codigo"], 10, 64)
	if erro != nil {
		Erro(w, http.StatusBadRequest, erro)
		return
	}

	var cep CEP
	cep.ID = codigo
	cep.Logradouro = "Rua Carlos Von Koseritz, 1346/305"
	cep.Bairro = "São João"
	cep.Cidade = "Porto Alegre"
	cep.Estado = "RS"

	fmt.Printf("CEP pesquisado: %d\n", codigo)

	JSON(w, 200, cep)
}

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
