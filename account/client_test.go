package account

import (
	"testing"
)

/*const (
	erroMsg = "Valor esperado %v, resultado encontrado %v"
)*/

func TestSetClient(t *testing.T) {

	checkResult := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf(erroMsg, resultado, esperado)
		}
	}

	t.Run("test function with true", func(t *testing.T) {
		valorEsperado := "123"

		client := Client{}
		client.CPF = "111.111.111-11"
		client.Secret = valorEsperado

		valorRetornado := client.Secret

		checkResult(t, valorRetornado, valorEsperado)
	})

}
