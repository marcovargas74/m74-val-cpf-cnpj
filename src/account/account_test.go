package account

import (
	"testing"
)

const (
	erroMsg = "Valor esperado %v, resultado encontrado %v"
)

func TestSetAccount(t *testing.T) {

	checkResult := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf(erroMsg, resultado, esperado)
		}
	}

	t.Run("test function with true", func(t *testing.T) {
		valorEsperado := "abc"

		accountMaria := Account{Balance: 0, Created_at: "17-01-2022"}
		accountMaria.ID = valorEsperado

		valorRetornado := accountMaria.ID

		checkResult(t, valorRetornado, valorEsperado)
	})

}
