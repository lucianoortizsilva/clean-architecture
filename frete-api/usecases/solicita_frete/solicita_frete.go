package solicita_frete

import "github.com/lucianoortizsilva/frete/frete-api/entities"

type SolicitaFrete struct {
	CalculoValorFreteRepository entities.CalculoValorFreteRepository
}

func NewSolicitaFrete(calculoValorFreteRepository entities.CalculoValorFreteRepository) *SolicitaFrete {
	return &SolicitaFrete{CalculoValorFreteRepository: calculoValorFreteRepository}
}

func (s *SolicitaFrete) Executar(input SolicitaFreteDtoInput) (SolicitaFreteDtoOutput, error) {

	calculoValorFrete := entities.NewCalculoValorFrete()
	calculoValorFrete.Regiao = input.Regiao
	calculoValorFrete.PesoTotalPedido = input.PesoTotalPedido
	valorFreteCalculado := calculoValorFrete.Calcular()

	output := SolicitaFreteDtoOutput{
		StatusSolicitacaoFrete: "APROVADO",
		ValorFrete:             valorFreteCalculado,
	}

	return output, nil

}
