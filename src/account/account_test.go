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

		account := Account{Balance: 0, CreatedAt: time.Now()}
		account.Name = valorEsperado
		valorRetornado := account.Name

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

//TODO Passar para um arquivo DB
func TestCreateAccount(t *testing.T) {

	tests := []struct {
		give        string
		wantValue   string
		inDataName  string
		inDataCPF   string
		inDataPassw string
		inDataVal   float64
	}{
		{
			give:        "Testa um nome OK Alice",
			wantValue:   "Alice",
			inDataName:  "Alice",
			inDataCPF:   "111.111.111-11",
			inDataPassw: "@Alice11",
			inDataVal:   11.11,
		},

		{
			give:        "Testa um nome OK Peter",
			wantValue:   "Peter",
			inDataName:  "Peter",
			inDataCPF:   "222.222.222-22",
			inDataPassw: "@Peter22",
			inDataVal:   22.22,
		},
		{
			give:        "Testa um nome OK NOBODY",
			wantValue:   "Nobody",
			inDataName:  "Nobody",
			inDataCPF:   "000.000.000-00",
			inDataPassw: "@Nobody00",
			inDataVal:   00.00,
		},
	}

	//server := NewServerBank()
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			aAccount := NewAccount(tt.inDataName, tt.inDataCPF, tt.inDataPassw, tt.inDataVal)
			fmt.Printf("->Name: %s ID: %s CreateAt: %s\n ", aAccount.Name, aAccount.ID, aAccount.CreatedAt.Format("01-02-2006 15:04:05"))
			assert.Equal(t, aAccount.Name, tt.wantValue)
			//func (a *Account) SaveAccountInDB() bool {

			/*requisicao := newReqEndpointsPOST("/transfers", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			verificaTipoDoConteudo(t, resposta, tipoDoConteudoJSON)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)*/
		})

	}

}

//TODO Passar para um arquivo DB
func TestSaveAccountInDB(t *testing.T) {

	tests := []struct {
		give        string
		wantValue   string
		inDataName  string
		inDataCPF   string
		inDataPassw string
		inDataVal   float64
	}{
		{
			give:        "Testa um nome OK Alice",
			wantValue:   "Alice",
			inDataName:  "Alice",
			inDataCPF:   "111.111.111-11",
			inDataPassw: "@Alice11",
			inDataVal:   11.11,
		},

		{
			give:        "Testa um nome OK Peter",
			wantValue:   "Peter",
			inDataName:  "Peter",
			inDataCPF:   "222.222.222-22",
			inDataPassw: "@Peter22",
			inDataVal:   22.22,
		},
		{
			give:        "Testa um nome OK NOBODY",
			wantValue:   "Nobody",
			inDataName:  "Nobody",
			inDataCPF:   "000.000.000-00",
			inDataPassw: "@Nobody00",
			inDataVal:   00.00,
		},
	}

	//server := NewServerBank()
	CreateDB(true)
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			aAccount := NewAccount(tt.inDataName, tt.inDataCPF, tt.inDataPassw, tt.inDataVal)
			fmt.Printf("->Name: %s ID: %s CreateAt: %s\n ", aAccount.Name, aAccount.ID, aAccount.CreatedAt.Format("01-02-2006 15:04:05"))
			if aAccount.SaveAccountInDB() {
				//log.Error("Cant save data in Bank")
				fmt.Println("Cant save data in Bank")
			}
			//func (a *Account) SaveAccountInDB() bool {

			assert.Equal(t, aAccount.Name, tt.wantValue)
			//func (a *Account) SaveAccountInDB() bool {

			accountInBD := GetAccountByID(aAccount.ID)
			assert.Equal(t, accountInBD.Name, tt.inDataName)

			assert.Equal(t, accountInBD.Balance, tt.inDataVal)
			UpdateBalanceByID(aAccount.ID, tt.inDataVal)
			accountInBD = GetAccountByID(aAccount.ID)
			assert.Equal(t, accountInBD.Balance, (tt.inDataVal * 2))
			/*requisicao := newReqEndpointsPOST("/transfers", tt.inData)
			resposta := httptest.NewRecorder()

			server.ServeHTTP(resposta, requisicao)
			verificaTipoDoConteudo(t, resposta, tipoDoConteudoJSON)

			recebido := resposta.Body.String()
			assert.Equal(t, recebido, tt.wantValue)*/
		})

	}

}
