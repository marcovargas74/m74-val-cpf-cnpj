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

			var fakeTransfer TransferBank
			recebido, err := fakeTransfer.MakeTransfer(&origAccount, &dstAccount, tt.inValueTrasfer)
			if err != nil {
				fmt.Println(err)
			}
			checkIfTruOrFalse(t, recebido, tt.wantValue)

			checkIfEqualFloat(t, origAccount.Balance, tt.wantValueOri)
			checkIfEqualFloat(t, dstAccount.Balance, tt.wantValueDst)
		})

	}

}
