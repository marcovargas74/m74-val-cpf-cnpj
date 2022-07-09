package cpfcnpj

import "testing"

func TestFormatToValidateCNPJ(t *testing.T) {

	tests := []struct {
		give        string
		cnpjToCheck string
		wantValue   string
	}{
		{
			give:        "CNPJ format for a string with only digits",
			cnpjToCheck: "000.000.000-00",
			wantValue:   "00000000000",
		},
		{
			give:        "CNPJ format for a string with only digits",
			cnpjToCheck: "00.000.000/0000-00",
			wantValue:   "00000000000000",
		},

		{
			give:        "CNPJ format for a string with only digits",
			cnpjToCheck: "36.562.098/0001-18",
			wantValue:   "36562098000118",
		},
		{
			give:        "CNPJ format for a string with only digits",
			cnpjToCheck: "12.074.074/0001-51",
			wantValue:   "12074074000151",
		},
		{
			give:        "CNPJ format for a string with only digits",
			cnpjToCheck: "24.572.400/0001-30",
			wantValue:   "24572400000130",
		},
		{
			give:        "CNPJ format for a string with only digits",
			cnpjToCheck: "47.425.683/0001-92",
			wantValue:   "47425683000192",
		},
	}
	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := formatToValidate(tt.cnpjToCheck)
			CheckIfEqualString(t, result, tt.wantValue)
		})

	}

}

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

func TestGetVerifyingDigitsCNPJ(t *testing.T) {

	tests := []struct {
		give       string
		wantValue1 uint64
		wantValue2 uint64
		cpfToCheck string
	}{
		{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 0,
			wantValue2: 0,
			cpfToCheck: "00.000.000/0000-00",
		},
		{
			give:       "Get Digits To check ",
			wantValue1: 1,
			wantValue2: 8,
			cpfToCheck: "36.562.098/0001-18",
		},

		{
			give:       "Get Digits To check ",
			wantValue1: 5,
			wantValue2: 1,
			cpfToCheck: "12.074.074/0001-51",
		},
		{
			give:       "Get Digits To check ",
			wantValue1: 3,
			wantValue2: 0,
			cpfToCheck: "24.572.400/0001-30",
		},
		{
			give:       "Get Digits To check ",
			wantValue1: 9,
			wantValue2: 2,
			cpfToCheck: "47.425.683/0001-92",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			Dig1, Dig2 := getVerifyingDigits(tt.cpfToCheck)
			CheckIfEqualInt(t, Dig1, tt.wantValue1)
			CheckIfEqualInt(t, Dig2, tt.wantValue2)
		})

	}

}

func TestMultiplyNumDigCNPJ(t *testing.T) {

	tests := []struct {
		give       string
		wantValue1 uint64
		wantValue2 uint64
		cpfToCheck string
	}{
		{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 3,
			wantValue2: 5,
			cpfToCheck: "11144477735",
		},
		/*{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 2,
			wantValue2: 5,
			cpfToCheck: "52998224725",
		},
		{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 0,
			wantValue2: 0,
			cpfToCheck: "00000000000",
		},
		{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 1,
			wantValue2: 1,
			cpfToCheck: "11111111111",
		},

		{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 8,
			wantValue2: 6,
			cpfToCheck: "83846172286",
		},
		{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 7,
			wantValue2: 7,
			cpfToCheck: "31339602377",
		},
		{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 9,
			wantValue2: 9,
			cpfToCheck: "68251194199",
		},
		{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 4,
			wantValue2: 0,
			cpfToCheck: "28875224340",
		},
		{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 0,
			wantValue2: 1,
			cpfToCheck: "48416241201",
		},*/
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			//Dig1, Dig2 := MultiplyNumDigCPF(tt.cpfToCheck)
			//CheckIfEqualInt(t, Dig1, tt.wantValue1)
			//CheckIfEqualInt(t, Dig2, tt.wantValue2)

			Dig1 := MultiplyNumDigCNPJ(tt.cpfToCheck, SizeToValidDig1)
			CheckIfEqualInt(t, Dig1, tt.wantValue1)

			Dig2 := MultiplyNumDigCNPJ(tt.cpfToCheck, SizeToValidDig2)
			CheckIfEqualInt(t, Dig2, tt.wantValue2)

		})

	}

}

/*
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
*/
