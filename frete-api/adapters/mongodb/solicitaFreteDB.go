package mongodb

import "fmt"

type SolicitaFreteDB struct{}

func NewSolicitaFreteDB() *SolicitaFreteDB {
	return &SolicitaFreteDB{}
}

func (solicitaFreteDB *SolicitaFreteDB) Salvar(Regiao string, PesoTotalPedido float64) error {
	fmt.Println("Salvou no mongoDB FAKE com sucesso")
	return nil
}
