package entities

type CalculoValorFreteRepository interface {
	Insert(PesoTotalPedido float64, Regiao string) error
}
