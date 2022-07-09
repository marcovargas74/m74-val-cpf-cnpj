package cpfcnpj

import "testing"

func TestIsValidFormatCNPJ(t *testing.T) {

	tests := []struct {
		give        string
		wantValue   bool
		cnpjToCheck string
	}{
		{
			give:        "Valid Format CPF Test if arg is Empty",
			wantValue:   false,
			cnpjToCheck: "",
		},
		{
			give:        "Valid Format CPF Test if arg is Invalid",
			wantValue:   false,
			cnpjToCheck: "b1080263",
		},
		{
			give:        "Valid Format CPF Test char - is not correct",
			wantValue:   false,
			cnpjToCheck: "000.000.00000-",
		},
		{
			give:        "Valid Format CPF Test char . is not correct",
			wantValue:   false,
			cnpjToCheck: "111111.111-11",
		},
		{
			give:        "Valid CPF Test if arg is a Valid CPF",
			wantValue:   false,
			cnpjToCheck: "00.000.000/0000-00",
		},
		{
			give:        "Valid CPF Test if arg is a Valid CPF",
			wantValue:   true,
			cnpjToCheck: "36.562.098/0001-18",
		},
		{
			give:        "Valid CPF Test if arg is a Valid CPF",
			wantValue:   true,
			cnpjToCheck: "12.074.074/0001-51",
		},
		{
			give:        "Valid CPF Test if arg is a Valid CPF",
			wantValue:   true,
			cnpjToCheck: "24.572.400/0001-30",
		},
		{
			give:        "Valid CPF Test if arg is a Valid CPF",
			wantValue:   true,
			cnpjToCheck: "47.425.683/0001-92",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := true

			if !isValidFormatCNPJ(tt.cnpjToCheck) {
				result = false
			}
			CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}

func TestMultiplyNumDigCNPJ(t *testing.T) {

	tests := []struct {
		give        string
		wantValue1  uint64
		wantValue2  uint64
		cnpjToCheck string
	}{
		{
			give:        "Get Digits To check if arg is Zeros Numbers",
			wantValue1:  8,
			wantValue2:  1,
			cnpjToCheck: "11222333000181",
		},
		{
			give:        "Get Digits To check if arg is Zeros Numbers",
			wantValue1:  1,
			wantValue2:  8,
			cnpjToCheck: "36562098000118",
		},
		{
			give:        "Get Digits To check if arg is Zeros Numbers",
			wantValue1:  6,
			wantValue2:  1,
			cnpjToCheck: "11444777000161",
		},
		{
			give:        "Get Digits To check if arg is Zeros Numbers",
			wantValue1:  3,
			wantValue2:  0,
			cnpjToCheck: "24572400000130",
		},
		{
			give:        "Get Digits To check if arg is Zeros Numbers",
			wantValue1:  9,
			wantValue2:  2,
			cnpjToCheck: "47425683000192",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			//Dig1, Dig2 := MultiplyNumDigCPF(tt.cnpjToCheck)
			//CheckIfEqualInt(t, Dig1, tt.wantValue1)
			//CheckIfEqualInt(t, Dig2, tt.wantValue2)

			Dig1 := MultiplyNumDigCNPJ(tt.cnpjToCheck, SizeToValidDig1CNPJ)
			CheckIfEqualInt(t, Dig1, tt.wantValue1)

			Dig2 := MultiplyNumDigCNPJ(tt.cnpjToCheck, SizeToValidDig2CNPJ)
			CheckIfEqualInt(t, Dig2, tt.wantValue2)

		})

	}

}

func TestIsValidCNPJ(t *testing.T) {

	tests := []struct {
		give        string
		wantValue   bool
		cnpjToCheck string
	}{
		{
			give:        "Valid CNPJ Test if arg is Empty",
			wantValue:   false,
			cnpjToCheck: "",
		},
		{
			give:        "Valid CNPJ Test if arg is Invalid",
			wantValue:   false,
			cnpjToCheck: "b1080263",
		},
		{
			give:        "Valid CNPJ Test if arg is Invalid",
			wantValue:   false,
			cnpjToCheck: "000.000.000-11",
		},
		{
			give:        "Valid CNPJ Test if arg is a CPF",
			wantValue:   false,
			cnpjToCheck: "111.111.111-11",
		},
		{
			give:        "Valid CNPJ Test if arg is Zeros Numbers",
			wantValue:   false,
			cnpjToCheck: "00.000.000/0000-00",
		},
		{
			give:        "Valid CNPJ Test if arg is a Valid CNPJ",
			wantValue:   true,
			cnpjToCheck: "36.562.098/0001-18",
		},

		{
			give:        "Valid CNPJ Test if arg is a Valid CNPJ",
			wantValue:   true,
			cnpjToCheck: "12.074.074/0001-51",
		},
		{
			give:        "Valid CNPJ Test if arg is a Valid CNPJ",
			wantValue:   true,
			cnpjToCheck: "24.572.400/0001-30",
		},
		{
			give:        "Valid CNPJ Test if arg is a Valid CNPJ",
			wantValue:   true,
			cnpjToCheck: "47.425.683/0001-92",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := true

			if !IsValidCNPJ(tt.cnpjToCheck) {
				result = false
			}
			CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}
