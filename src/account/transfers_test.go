package account

import (
	"fmt"
	"testing"
)

func checkIfTruOrFalse(t *testing.T, gotValue, waitValue bool) {
	t.Helper()
	if gotValue != waitValue {
		t.Errorf(erroMsg, gotValue, waitValue)
	}
}

func checkIfEqualFloat(t *testing.T, gotValue, waitValue float64) {
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
			checkIfTruOrFalse(t, recebido, tt.wantValue)

			checkIfEqualFloat(t, origAccount.Balance, tt.wantValueOri)
			checkIfEqualFloat(t, dstAccount.Balance, tt.wantValueDst)
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
		{
			give:          "Testa Busca por ID Valido",
			wantValue:     true,
			inFindID:      "3b582666-4600-4640-ab9c-bc01b66def0e",
			inValueAmount: 10.00,
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
			checkIfTruOrFalse(t, result, tt.wantValue)
			checkIfEqualFloat(t, aTransfer.Amount, tt.inValueAmount)
		})

	}

}
