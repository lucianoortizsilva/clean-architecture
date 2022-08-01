package configuration

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	URL_CONEXAO_MONGODB = ""
	PORT_HTTP           = 0
)

func CarregarVariaveisAmbiente() {
	var erro error

	if erro = godotenv.Load("_local.env", "_docker.env"); erro != nil {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println("ERRO AO CARREGAR VARIÁVEIS DE AMBIENTE")
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		log.Fatalf("Erro: %s", erro)
	}

	PORT_HTTP, erro = strconv.Atoi(os.Getenv("PORT_HTTP"))
	if erro != nil {
		PORT_HTTP = 9000
	}

	URL_CONEXAO_MONGODB = os.Getenv("URL_CONEXAO_MONGODB")

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(">>>VARIÁVEIS CARREGAGADAS COM SUCESSO<<<")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

	fmt.Printf("PORT_HTTP: %d", PORT_HTTP)
	fmt.Println("")
	fmt.Printf("URL_CONEXAO_MONGODB: %s", URL_CONEXAO_MONGODB)

}
