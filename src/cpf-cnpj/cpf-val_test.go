package cpfcnpj

import "testing"

func TestIsValidCPF(t *testing.T) {

	tests := []struct {
		give      string
		wantValue bool
		inFindID  string
	}{
		{
			give:      "Valid CNPJ Test if arg is Empty",
			wantValue: false,
			inFindID:  "",
		},
		{
			give:      "Valid CNPJ Test if arg is Invalid",
			wantValue: false,
			inFindID:  "b1080263",
		},
		{
			give:      "Valid CPF Test if arg is Zeros Numbers",
			wantValue: true,
			inFindID:  "000.000.000-00",
		},
		{
			give:      "Valid CPF Test if arg is a Valid CPF",
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
