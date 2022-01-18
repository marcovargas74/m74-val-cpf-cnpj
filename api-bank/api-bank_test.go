package m74bankAPI

import (
	"testing"
)

const (
	erroMsg = "Valor esperado %v, resultado encontrado %v"
)

func TestSetIsProduction(t *testing.T) {

	checkResult := func(t *testing.T, resultado, esperado bool) {
		t.Helper()
		if resultado != esperado {
			t.Errorf(erroMsg, resultado, esperado)
		}
	}

	t.Run("test function with true", func(t *testing.T) {
		valorEsperado := true
		SetIsProduction(valorEsperado)
		valorRetornado := GetIsProduction()

		checkResult(t, valorRetornado, valorEsperado)
	})

	t.Run("test function with false", func(t *testing.T) {
		valorEsperado := false
		SetIsProduction(valorEsperado)
		valorRetornado := GetIsProduction()

		checkResult(t, valorRetornado, valorEsperado)
	})

}
