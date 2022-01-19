package m74bankapi

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

func TestExampleTesteFunc(t *testing.T) {
	tests := []struct {
		entrada  bool
		esperado bool
	}{
		//testes ok
		{true, true},
		//Testes de entradas invalidas
		//{"teste", []byte("teste")},
		//{ 257,0},
	}

	for _, teste := range tests {
		t.Logf("Teste %v", teste)

		//valorEsperado := teste.entrue
		SetIsProduction(teste.entrada)
		//valorRetornado := GetIsProduction()

		//checkResult(t, valorRetornado, valorEsperado)

		//atual := ExampleTesteFunc(teste.entrada)
		valorRetornado := GetIsProduction()
		if valorRetornado != teste.esperado {
			t.Errorf(erroMsg, teste.esperado, valorRetornado)
		}
	}

}

/*
// func TestSplitHostPort(t *testing.T)

tests := []struct{
  give     string
  wantHost string
  wantPort string
}{
  {
    give:     "192.0.2.0:8000",
    wantHost: "192.0.2.0",
    wantPort: "8000",
  },
  {
    give:     "192.0.2.0:http",
    wantHost: "192.0.2.0",
    wantPort: "http",
  },
  {
    give:     ":8000",
    wantHost: "",
    wantPort: "8000",
  },
  {
    give:     "1:8",
    wantHost: "1",
    wantPort: "8",
  },
}

for _, tt := range tests {
  t.Run(tt.give, func(t *testing.T) {
    host, port, err := net.SplitHostPort(tt.give)
    require.NoError(t, err)
    assert.Equal(t, tt.wantHost, host)
    assert.Equal(t, tt.wantPort, port)
  })
}



*/
