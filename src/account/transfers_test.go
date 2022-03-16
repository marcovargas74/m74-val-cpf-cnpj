package account

import (
	"fmt"
	"testing"
)

//CheckIfEqualFloat check if resulto is OK type Float
func CheckIfEqualFloat(t *testing.T, gotValue, waitValue float64) {
	t.Helper()
	if gotValue != waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

//CheckIfEqualString check if resulto is OK type string
func CheckIfEqualString(t *testing.T, gotValue, waitValue string) {
	t.Helper()
	if gotValue != waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

//CheckIfEqualBool check if resulto is OK type BOOL
func CheckIfEqualBool(t *testing.T, gotValue, waitValue bool) {
	t.Helper()
	if gotValue != waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

func TestMakeTransfer(t *testing.T) {

	tests := []struct {
		give      string
		wantValue bool

		inValueTrasfer float64
		inNameOri      string
		inValueOri     float64
		wantValueOri   float64

		inNameDst    string
		inValueDst   float64
		wantValueDst float64
	}{
		{
			give:           "Testa Se Origin é negativo",
			wantValue:      false,
			inValueTrasfer: 50.00,
			inNameOri:      "Alice",
			inValueOri:     -10.00,
			wantValueOri:   -10.00,

			inNameDst:    "Peter",
			inValueDst:   10.00,
			wantValueDst: 10.00,
		},
		{
			give:           "Testa Se Origin não tem saldo",
			wantValue:      false,
			inValueTrasfer: 50.00,
			inNameOri:      "Alice2",
			inValueOri:     10.00,
			wantValueOri:   10.00,

			inNameDst:    "Peter2",
			inValueDst:   10.00,
			wantValueDst: 10.00,
		},

		{
			give:           "Testa Se Origin tem saldo",
			wantValue:      true,
			inValueTrasfer: 50.00,
			inNameOri:      "Alice3",
			inValueOri:     100.00,
			wantValueOri:   50.00,

			inNameDst:    "Peter3",
			inValueDst:   10.00,
			wantValueDst: 60.00,
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			origAccount := Account{Name: tt.inNameOri, Balance: tt.inValueOri}
			dstAccount := Account{Name: tt.inNameDst, Balance: tt.inValueDst}

			//var fakeTransfer TransferBank
			fakeTransfer := TransferBank{Amount: tt.inValueTrasfer}
			recebido, err := fakeTransfer.MakeTransfer(&origAccount, &dstAccount)
			if err != nil {
				fmt.Println(err)
			}
			CheckIfEqualBool(t, recebido, tt.wantValue)

			CheckIfEqualFloat(t, origAccount.Balance, tt.wantValueOri)
			CheckIfEqualFloat(t, dstAccount.Balance, tt.wantValueDst)
		})

	}

}

func TestGetTransfer(t *testing.T) {

	tests := []struct {
		give          string
		wantValue     bool
		inFindID      string
		inValueAmount float64
	}{
		{
			give:          "Testa Busca por ID Vazio",
			wantValue:     false,
			inFindID:      "",
			inValueAmount: 00.00,
		},
		{
			give:          "Testa Busca por ID Invalido",
			wantValue:     false,
			inFindID:      "b1080263",
			inValueAmount: 00.00,
		},
		{
			give:          "Testa Busca por ID Inexistente",
			wantValue:     false,
			inFindID:      "b1080263-f5e0-495a-8e70-60a303a7a8d3",
			inValueAmount: 00.00,
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := true
			aTransfer, err := GetTranferByID(tt.inFindID)
			if err != nil {
				fmt.Println(err)
				result = false
			}

			fmt.Println(aTransfer.AccountOriginID)
			CheckIfEqualBool(t, result, tt.wantValue)
			CheckIfEqualFloat(t, aTransfer.Amount, tt.inValueAmount)

		})

	}

}

func TestGetTransferByCPF(t *testing.T) {

	tests := []struct {
		give          string
		wantValue     bool
		inFindID      string
		inValueAmount float64
	}{
		{
			give:          "Testa Busca por un CPF Vazio",
			wantValue:     false,
			inFindID:      "",
			inValueAmount: 00.00,
		},
		{
			give:          "Testa Busca por un CPF Invalido",
			wantValue:     false,
			inFindID:      "b1080263",
			inValueAmount: 00.00,
		},
		{
			give:          "Testa Busca por un CPF Inexistente",
			wantValue:     false,
			inFindID:      "000.000.000-11",
			inValueAmount: 00.00,
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := true

			user, err := GetAccountByCPF(tt.inFindID)
			if err != nil || user.CPF == "" {
				result = false
			}
			CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}
