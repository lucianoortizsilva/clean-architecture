package frete

import (
	"errors"
	"strings"
)

type FreteDTO struct {
	Codigo          string  `json:"codigo,omitempty"`
	PedidoId        string  `json:"pedidoId,omitempty"`
	PesoTotalPedido float64 `json:"pesoTotalPedido,omitempty"`
	Status          string  `json:"status,omitempty"`
	Regiao          string  `json:"regiao,omitempty"`
	Cep             string  `json:"cep,omitempty"`
}

func (dto *FreteDTO) Validar() error {

	if strings.TrimSpace(dto.Codigo) == "" {
		return errors.New("o codigo é obrigatório e não pode estar em branco")
	}

	if strings.TrimSpace(dto.PedidoId) == "" {
		return errors.New("o pedidoId é obrigatório e não pode estar em branco")
	}

	if strings.TrimSpace(dto.Cep) == "" {
		return errors.New("o cep é obrigatório e não pode estar em branco")
	}

	if len(strings.TrimSpace(dto.Cep)) != 8 {
		return errors.New("o cep deve conter 8 caracteres")
	}

	if dto.PesoTotalPedido <= 0 {
		return errors.New("o peso total do pedido deve ser maior que zero")
	}

	return nil
}
