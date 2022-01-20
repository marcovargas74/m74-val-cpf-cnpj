package account

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

/*const (
	erroMsg = "Valor esperado %v, resultado encontrado %v"
)*/

func TestSetClient(t *testing.T) {

	tests := []struct {
		give      string
		wantValue string
		inData    string
	}{
		{
			give:      "test function criando client com cpf",
			wantValue: "123",
			inData:    "111.111.111-11",
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			client := Client{}
			client.CPF = tt.inData
			client.Secret = tt.wantValue
			valorRetornado := client.Secret
			assert.Equal(t, valorRetornado, tt.wantValue)
		})

	}

}

func TestGetCPFbyName(t *testing.T) {

	tests := []struct {
		give       string
		wantValue  string
		inDataCPF  string
		inDataName string
	}{
		{
			give:       "test function criando client com cpf",
			wantValue:  "111.111.111-11",
			inDataCPF:  "111.111.111-11",
			inDataName: "Maria",
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			client := Client{ID: "1233"}
			client.CPF = tt.inDataCPF
			client.Name = tt.inDataName

			//
			valorRetornado := client.GetCPFbyName(tt.inDataName)
			assert.Equal(t, valorRetornado, tt.wantValue)
		})

	}

}

/*
//GetCPF Get
func (c Client) GetCPFbyName(name string) string {
	return c.CPF
}

//GetCPF Get
func (c Client) GetCPFbyID(ID string) string {
	return c.CPF
}* /

func useDataLogin() {
	client := Client{}
	client.CPF = "111.111.111-11"
	client.Secret = "111"
	fmt.Println(client.CPF)
}
*/
