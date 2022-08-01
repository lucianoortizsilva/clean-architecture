package entities

import (
	"errors"
	"testing"
)

func TestCalcularValorFreteRegiaoSulPesoMaiorQueVinte(t *testing.T) {
	calculo := NewCalculoValorFrete()
	calculo.PesoTotalPedido = 21
	calculo.Regiao = "SUL"
	valorFreteCalculado, _ := calculo.Calcular()
	valorFreteEsperado := 44.50
	if valorFreteCalculado != valorFreteEsperado {
		t.Errorf("Gerado %f, Esperado %f", valorFreteCalculado, valorFreteEsperado)
	}
}

func TestCalcularValorFreteRegiaoSulPesoMenorQueVinte(t *testing.T) {
	calculo := NewCalculoValorFrete()
	calculo.PesoTotalPedido = 19
	calculo.Regiao = "SUL"
	valorFreteCalculado, _ := calculo.Calcular()
	valorFreteEsperado := 114.00
	if valorFreteCalculado != valorFreteEsperado {
		t.Errorf("Gerado %f, Esperado %f", valorFreteCalculado, valorFreteEsperado)
	}
}

func TestCalcularValorFreteRegiaoSulPesoIgualVinte(t *testing.T) {
	calculo := NewCalculoValorFrete()
	calculo.PesoTotalPedido = 20
	calculo.Regiao = "SUL"
	valorFreteCalculado, _ := calculo.Calcular()
	valorFreteEsperado := 114.00
	if valorFreteCalculado != valorFreteEsperado {
		t.Errorf("Gerado %f, Esperado %f", valorFreteCalculado, valorFreteEsperado)
	}
}

func TestCalcularValorFreteRegiaoSudestePesoMaiorQueCinquenta(t *testing.T) {
	calculo := NewCalculoValorFrete()
	calculo.PesoTotalPedido = 51
	calculo.Regiao = "SUDESTE"
	valorFreteCalculado, erroGerado := calculo.Calcular()
	valorFreteEsperado := 0.00
	erroEsperado := errors.New("Peso excedido para a região SUDESTE").Error()
	if valorFreteCalculado != valorFreteEsperado || erroGerado.Error() != erroEsperado {
		t.Errorf("Gerado %f, Esperado %f", valorFreteCalculado, valorFreteEsperado)
		t.Errorf("Gerado %s, Esperado %s", erroGerado.Error(), erroEsperado)
	}
}
