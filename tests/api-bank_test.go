package m74bankAPI

import (
	"testing"

	bankAPI "github.com/marcovargas74/m74-bank-api"
)

const (
	erroMsg = "Valor esperado %v, resultado encontrado %v"
)

func TestGetIsProduction(t *testing.T) {
	if bankAPI.GetIsProduction() == true {
		t.Error(erroMsg)
	}

}

func TestSetIsProduction(t *testing.T) {
	valorEsperado := true
	bankAPI.SetIsProduction(valorEsperado)

	valorRetornado := bankAPI.GetIsProduction()
	if valorRetornado != valorEsperado {
		t.Errorf(erroMsg, valorEsperado, valorRetornado)
	}

	valorEsperado = false
	bankAPI.SetIsProduction(valorEsperado)
	valorRetornado = bankAPI.GetIsProduction()
	if valorRetornado != valorEsperado {
		t.Errorf(erroMsg, valorEsperado, valorRetornado)
	}

}
