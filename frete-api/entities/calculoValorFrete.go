package entities

type CalculoValorFrete struct {
	PesoTotalPedido float64
	Regiao          string
}

func NewCalculoValorFrete() *CalculoValorFrete {
	return &CalculoValorFrete{}
}

func (consulta *CalculoValorFrete) Calcular() float64 {

	if consulta.Regiao == "NORTE" && consulta.PesoTotalPedido > 10 {
		return 66.00
	}

	if consulta.Regiao == "SUL" && consulta.PesoTotalPedido > 20 {
		return 44.50
	}

	if consulta.Regiao == "NORDESTE" && consulta.PesoTotalPedido > 15 {
		return 77.00
	}

	return 114.00
}
