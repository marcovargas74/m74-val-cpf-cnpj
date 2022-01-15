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

	checkResult := func(t *testing.T, resultado, esperado bool) {
		t.Helper()
		if resultado != esperado {
			t.Errorf(erroMsg, resultado, esperado)
		}
	}

	t.Run("test function with true", func(t *testing.T) {
		valorEsperado := true
		bankAPI.SetIsProduction(valorEsperado)
		valorRetornado := bankAPI.GetIsProduction()

		checkResult(t, valorRetornado, valorEsperado)
	})

	t.Run("test function with false", func(t *testing.T) {
		valorEsperado := false
		bankAPI.SetIsProduction(valorEsperado)
		valorRetornado := bankAPI.GetIsProduction()

		checkResult(t, valorRetornado, valorEsperado)
	})

}
