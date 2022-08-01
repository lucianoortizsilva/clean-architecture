package usecases

import (
	"github.com/lucianoortizsilva/frete/frete-api/entities"
)

type SolicitaFrete struct {
	SolicitaFreteRepository SolicitaFreteRepository
}

type SolicitaFreteRepository interface {
	Insert(Codigo string, PedidoId string, Regiao string, Cep string, PesoTotalPedido float64) error
}

type SolicitaFreteDtoInput struct {
	Codigo          string
	PedidoId        string
	PesoTotalPedido float64
	Regiao          string
	Cep             string
}

type SolicitaFreteDtoOutput struct {
	Codigo     string  `json:"codigo,omitempty"`
	ValorFrete float64 `json:"valorFrete,omitempty"`
	Status     string  `json:"status,omitempty"`
	Erro       string  `json:"erro,omitempty"`
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
		erro := s.SolicitaFreteRepository.Insert(input.Codigo, input.PedidoId, input.Regiao, input.Cep, input.PesoTotalPedido)
		if erro == nil {
			return aprovado(input.Codigo, valorFreteCalculado)
		} else {
			return SolicitaFreteDtoOutput{}, erro
		}
	} else {
		return SolicitaFreteDtoOutput{}, erro
	}

}

func aprovado(Codigo string, ValorFrete float64) (SolicitaFreteDtoOutput, error) {
	output := SolicitaFreteDtoOutput{
		Codigo:     Codigo,
		Status:     "APROVADO",
		ValorFrete: ValorFrete,
	}
	return output, nil
}
