package cpfcnpj

import "testing"

func TestIsValidCNPJ(t *testing.T) {

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
			give:      "Valid CNPJ Test if arg is Invalid",
			wantValue: false,
			inFindID:  "000.000.000-11",
		},
		{
			give:      "Valid CNPJ Test if arg is a CPF",
			wantValue: false,
			inFindID:  "111.111.111-11",
		},
		{
			give:      "Valid CNPJ Test if arg is Zeros Numbers",
			wantValue: false,
			inFindID:  "00.000.000/0000-00",
		},
		{
			give:      "Valid CNPJ Test if arg is a Valid CNPJ",
			wantValue: true,
			inFindID:  "36.562.098/0001-18",
		},

		{
			give:      "Valid CNPJ Test if arg is a Valid CNPJ",
			wantValue: true,
			inFindID:  "12.074.074/0001-51",
		},
		{
			give:      "Valid CNPJ Test if arg is a Valid CNPJ",
			wantValue: true,
			inFindID:  "24.572.400/0001-30",
		},
		{
			give:      "Valid CNPJ Test if arg is a Valid CNPJ",
			wantValue: true,
			inFindID:  "47.425.683/0001-92",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := true

			if !IsValidCNPJ(tt.inFindID) {
				result = false
			}
			CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}
