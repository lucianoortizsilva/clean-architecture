package frete

import (
	"errors"
	"strings"
)

type FreteDTO struct {
	PedidoId        string  `json:"pedidoId,omitempty"`
	PesoTotalPedido float64 `json:"pesoTotalPedido,omitempty"`
	Regiao          string  `json:"regiao,omitempty"`
}

func (dto *FreteDTO) Validar() error {

	if strings.TrimSpace(dto.PedidoId) == "" {
		return errors.New("o pedidoId é obrigatório e não pode estar em branco")
	}

	if strings.TrimSpace(dto.Regiao) == "" {
		return errors.New("a região é obrigatória e não pode estar em branco")
	}

	if dto.PesoTotalPedido <= 0 {
		return errors.New("o peso total do pedido deve ser maior que zero")
	}

	return nil
}
