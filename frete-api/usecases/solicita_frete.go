package usecases

import "github.com/lucianoortizsilva/frete/frete-api/entities"

type SolicitaFreteRepository interface {
	Salvar(Regiao string, PesoTotalPedido float64) error
}

type SolicitaFrete struct {
	SolicitaFreteRepository SolicitaFreteRepository
}

type SolicitaFreteDtoInput struct {
	Regiao          string
	PesoTotalPedido float64
}

type SolicitaFreteDtoOutput struct {
	ValorFrete             float64
	StatusSolicitacaoFrete string
	Erro                   string
}

func NewSolicitaFrete(solicitaFreteRepository SolicitaFreteRepository) *SolicitaFrete {
	return &SolicitaFrete{SolicitaFreteRepository: solicitaFreteRepository}
}

func (s *SolicitaFrete) Executar(input SolicitaFreteDtoInput) (SolicitaFreteDtoOutput, error) {

	calculoValorFrete := entities.NewCalculoValorFrete()
	calculoValorFrete.Regiao = input.Regiao
	calculoValorFrete.PesoTotalPedido = input.PesoTotalPedido
	valorFreteCalculado, erro := calculoValorFrete.Calcular()

	if erro == nil {
		output := SolicitaFreteDtoOutput{
			StatusSolicitacaoFrete: "APROVADO",
			ValorFrete:             valorFreteCalculado,
		}
		s.SolicitaFreteRepository.Salvar(calculoValorFrete.Regiao, calculoValorFrete.PesoTotalPedido)
		return output, nil
	} else {
		output := SolicitaFreteDtoOutput{
			StatusSolicitacaoFrete: "NEGADO",
			Erro:                   erro.Error(),
		}
		return output, erro
	}

}
