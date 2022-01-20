package m74bankapi

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

const (
	erroMsg = "Valor esperado %v, resultado encontrado %v"
)

func TestGetVersion(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
	}{
		{
			give:      "Test if get version OK",
			wantValue: "2022",
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			version := GetVersion()
			//fmt.Printf("valor: %v\n", valorRetornado[0:4])
			assert.Equal(t, version[0:4], tt.wantValue)
		})

	}

}

func TestSetAnGetIsProduction(t *testing.T) {

	tests := []struct {
		give      string
		wantValue bool
		inData    bool
	}{
		{
			give:      "test function with true",
			wantValue: true,
			inData:    true,
		},
		{
			give:      "test function with false",
			wantValue: false,
			inData:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			SetIsProduction(tt.inData)
			assert.Equal(t, GetIsProduction(), tt.wantValue)
			//checkResultBool(t, tt.wantValue, valorRetornado)
			//checkResult(t, valorRetornado, wantValue)
		})

	}

}

/*
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

}*/
