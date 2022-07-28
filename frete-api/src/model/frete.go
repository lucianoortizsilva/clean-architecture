package model

import (
	"errors"
	"strings"
)

type Frete struct {
	Codigo   string `json:"codigo,omitempty"`
	PedidoId string `json:"pedidoId,omitempty"`
	Status   string `json:"status,omitempty"`
	Cep      string `json:"cep,omitempty"`
}

func (frete *Frete) Validar() error {

	if strings.TrimSpace(frete.Codigo) == "" {
		return errors.New("o codigo é obrigatório e não pode estar em branco")
	}

	if strings.TrimSpace(frete.PedidoId) == "" {
		return errors.New("o pedidoId é obrigatório e não pode estar em branco")
	}

	if strings.TrimSpace(frete.Cep) == "" {
		return errors.New("o cep é obrigatório e não pode estar em branco")
	}

	if len(strings.TrimSpace(frete.Cep)) != 8 {
		return errors.New("o cep deve conter 8 caracteres")
	}

	return nil
}
