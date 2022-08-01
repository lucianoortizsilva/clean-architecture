package entities

import "errors"

type CalculoValorFrete struct {
	PesoTotalPedido float64
	Regiao          string
}

func NewCalculoValorFrete() *CalculoValorFrete {
	return &CalculoValorFrete{}
}

func (consulta *CalculoValorFrete) Calcular() (float64, error) {

	if consulta.Regiao == "SUL" && consulta.PesoTotalPedido > 20 {
		return 44.50, nil
	}

	if consulta.Regiao == "SUDESTE" && consulta.PesoTotalPedido > 50 {
		return 0, errors.New("Peso excedido para a região SUDESTE")
	}

	return 114.00, nil
}
