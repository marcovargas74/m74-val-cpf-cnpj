package account

import (
	"fmt"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"gopkg.in/validator.v2"
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

		accountMaria := Account{Balance: 0, CreatedAt: time.Now()}
		accountMaria.ID = valorEsperado

		valorRetornado := accountMaria.ID

		checkResult(t, valorRetornado, valorEsperado)
	})

}

func TestValidName(t *testing.T) {

	tests := []struct {
		give      string
		wantValue bool
		inName    string
	}{
		{
			give:      "Testa Se Nome é invalido",
			wantValue: false,
			inName:    "",
		},
		{
			give:      "Testa Se Nome é Valido No",
			wantValue: false,
			inName:    "No",
		},
		{
			give:      "Testa Se Nome é Valido",
			wantValue: true,
			inName:    "Nome",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			recebido := true
			acc := Account{Name: tt.inName, Balance: 0.00}
			if errs := validator.Validate(acc); errs != nil {
				fmt.Printf("INVALIDO %v\n", errs) // do something
				recebido = false
			}

			assert.Equal(t, recebido, tt.wantValue)
		})

	}

}
