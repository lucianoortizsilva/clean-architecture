package solicita_frete

type SolicitaFreteDtoInput struct {
	Regiao          string
	PesoTotalPedido float64
}

type SolicitaFreteDtoOutput struct {
	ValorFrete             float64
	StatusSolicitacaoFrete string
	Erro                   string
}
