package usecases

import (
	"github.com/lucianoortizsilva/frete/frete-api/entities"
)

type SolicitaFrete struct {
	SolicitaFreteRepository SolicitaFreteRepository
}

type SolicitaFreteRepository interface {
	Insert(PedidoId string, Regiao string, PesoTotalPedido float64) error
}

type SolicitaFreteDtoInput struct {
	PedidoId        string
	PesoTotalPedido float64
	Regiao          string
}

type SolicitaFreteDtoOutput struct {
	PedidoId   string  `json:"pedidoId,omitempty"`
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
		erro := s.SolicitaFreteRepository.Insert(input.PedidoId, input.Regiao, input.PesoTotalPedido)
		if erro == nil {
			return solicitado(input.PedidoId, valorFreteCalculado)
		} else {
			return SolicitaFreteDtoOutput{}, erro
		}
	} else {
		return SolicitaFreteDtoOutput{}, erro
	}

}

func solicitado(PedidoId string, ValorFrete float64) (SolicitaFreteDtoOutput, error) {
	output := SolicitaFreteDtoOutput{
		PedidoId:   PedidoId,
		ValorFrete: ValorFrete,
		Status:     "SOLICITADO",
	}
	return output, nil
}
