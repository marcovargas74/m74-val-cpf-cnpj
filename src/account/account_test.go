package account

import (
	"fmt"
	"testing"
	"time"

	"gopkg.in/validator.v2"
)

const (
	erroMsg = "Test fail got Value[%v], wait Value [%v]"
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

			CheckIfEqualBool(t, recebido, tt.wantValue)
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

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			aAccount := newAccount(tt.inDataName, tt.inDataCPF, tt.inDataPassw, tt.inDataVal)
			fmt.Printf("->Name: %s ID: %s CreateAt: %s\n ", aAccount.Name, aAccount.ID, aAccount.CreatedAt.Format("01-02-2006 15:04:05"))
			CheckIfEqualString(t, aAccount.Name, tt.wantValue)
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

	CreateDB(false)
	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			aAccount := newAccount(tt.inDataName, tt.inDataCPF, tt.inDataPassw, tt.inDataVal)
			fmt.Printf("->Name: %s ID: %s CreateAt: %s\n ", aAccount.Name, aAccount.ID, aAccount.CreatedAt.Format("01-02-2006 15:04:05"))
			if aAccount.saveAccountInDB() {
				fmt.Println("Cant save data in Bank")
			}
			CheckIfEqualString(t, aAccount.Name, tt.wantValue)

			accountInBD, _ := GetAccountByID(aAccount.ID)
			CheckIfEqualString(t, accountInBD.Name, tt.inDataName)
			CheckIfEqualFloat(t, accountInBD.Balance, tt.inDataVal)

			UpdateBalanceByID(aAccount.ID, (tt.inDataVal * 2))
			accountInBD, _ = GetAccountByID(aAccount.ID)
			CheckIfEqualFloat(t, accountInBD.Balance, (tt.inDataVal * 2))
		})

	}

}

func TestIsValidCPF(t *testing.T) {

	tests := []struct {
		give      string
		wantValue bool
		inFindID  string
	}{
		{
			give:      "Testa Se CPF Vazio",
			wantValue: false,
			inFindID:  "",
		},
		{
			give:      "Testa se CPF Invalido",
			wantValue: false,
			inFindID:  "b1080263",
		},
		{
			give:      "Testa se o CPF invalido",
			wantValue: true,
			inFindID:  "000.000.000-11",
		},
		{
			give:      "Testa Busca por un CPF Valido",
			wantValue: true,
			inFindID:  "111.111.111-11",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := true

			if !IsValidCPF(tt.inFindID) {
				result = false
			}
			CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}
