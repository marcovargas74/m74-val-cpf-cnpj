package cpfcnpj

import (
	"testing"
)

func TestFormatToValidateCPF(t *testing.T) {

	tests := []struct {
		give       string
		cpfToCheck string
		wantValue  string
	}{
		{
			give:       "CPF format for a string with only digits",
			wantValue:  "00000000000",
			cpfToCheck: "000.000.000-00",
		},
		{
			give:       "CPF format for a string with only digits",
			cpfToCheck: "111.111.111-11",
			wantValue:  "11111111111",
		},

		{
			give:       "CPF format for a string with only digits",
			cpfToCheck: "838.461.722-86",
			wantValue:  "83846172286",
		},
		{
			give:       "CPF format for a string with only digits",
			cpfToCheck: "313.396.023-77",
			wantValue:  "31339602377",
		},
		{
			give:       "CPF format for a string with only digits",
			cpfToCheck: "682.511.941-99",
			wantValue:  "68251194199",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := formatToValidate(tt.cpfToCheck)
			CheckIfEqualString(t, result, tt.wantValue)
		})

	}

}

func TestIsValidFormatCPF(t *testing.T) {

	tests := []struct {
		give       string
		wantValue  bool
		cpfToCheck string
	}{
		{
			give:       "Valid Format CPF Test if arg is Empty",
			wantValue:  false,
			cpfToCheck: "",
		},
		{
			give:       "Valid Format CPF Test if arg is Invalid",
			wantValue:  false,
			cpfToCheck: "b1080263",
		},
		{
			give:       "Valid Format CPF Test char - is not correct",
			wantValue:  false,
			cpfToCheck: "000.000.00000-",
		},
		{
			give:       "Valid Format CPF Test char . is not correct",
			wantValue:  false,
			cpfToCheck: "111111.111-11",
		},
		{
			give:       "Valid CPF Test if arg is a Valid CPF",
			wantValue:  true,
			cpfToCheck: "838.461.722-86",
		},
		{
			give:       "Valid CPF Test if arg is a Valid CPF",
			wantValue:  true,
			cpfToCheck: "313.396.023-77",
		},
		{
			give:       "Valid CPF Test if arg is a Valid CPF",
			wantValue:  true,
			cpfToCheck: "682.511.941-99",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := true

			if !isValidFormatCPF(tt.cpfToCheck) {
				result = false
			}
			CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}

func TestGetVerifyingDigits(t *testing.T) {

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
			cpfToCheck: "000.000.000-00",
		},
		{
			give:       "Get Digits To check ",
			wantValue1: 1,
			wantValue2: 1,
			cpfToCheck: "111.111.111-11",
		},

		{
			give:       "Get Digits To check ",
			wantValue1: 8,
			wantValue2: 6,
			cpfToCheck: "838.461.722-86",
		},
		{
			give:       "Get Digits To check ",
			wantValue1: 7,
			wantValue2: 7,
			cpfToCheck: "313.396.023-77",
		},
		{
			give:       "Get Digits To check ",
			wantValue1: 9,
			wantValue2: 9,
			cpfToCheck: "682.511.941-99",
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

func TestMultiplyNumDigCPF(t *testing.T) {

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
		{
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
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			//Dig1, Dig2 := MultiplyNumDigCPF(tt.cpfToCheck)
			//CheckIfEqualInt(t, Dig1, tt.wantValue1)
			//CheckIfEqualInt(t, Dig2, tt.wantValue2)

			Dig1 := MultiplyNumDigCPF(tt.cpfToCheck, SizeToValidDig1CPF)
			CheckIfEqualInt(t, Dig1, tt.wantValue1)

			Dig2 := MultiplyNumDigCPF(tt.cpfToCheck, SizeToValidDig2CPF)
			CheckIfEqualInt(t, Dig2, tt.wantValue2)

		})

	}

}

func TestAllDigitsIsEqualCPF(t *testing.T) {

	tests := []struct {
		give       string
		wantValue  bool
		cpfToCheck string
	}{
		{
			give:       "Test if all digits is Equal - This is a Invalid CPF",
			wantValue:  false,
			cpfToCheck: "",
		},
		{
			give:       "Test if all digits is Equal - This is a Invalid CPF",
			wantValue:  false,
			cpfToCheck: "b1080263",
		},
		{
			give:       "Test if all digits is Equal - This is a Invalid CPF",
			wantValue:  false,
			cpfToCheck: "83846172286",
		},
		{
			give:       "Test if all digits is Equal - This is a Invalid CPF",
			wantValue:  false,
			cpfToCheck: "31339602377",
		},
		{
			give:       "Test if all digits is Equal - This is a Invalid CPF",
			wantValue:  false,
			cpfToCheck: "68251194199",
		},
		{
			give:       "Test if all digits is Equal - This is a Invalid CPF",
			wantValue:  false,
			cpfToCheck: "28875224340",
		},
		{
			give:       "Test if all digits is Equal - This is a Invalid CPF",
			wantValue:  false,
			cpfToCheck: "48416241201",
		},
		{
			give:       "Test if all digits is Equal - This is a Invalid CPF",
			wantValue:  true,
			cpfToCheck: "00000000000",
		},
		{
			give:       "Test if all digits is Equal - This is a Invalid CPF",
			wantValue:  true,
			cpfToCheck: "11111111111",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := allDigitsIsEqual(tt.cpfToCheck)
			CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}

func TestIsValidCPF(t *testing.T) {

	tests := []struct {
		give       string
		wantValue  bool
		cpfToCheck string
	}{
		{
			give:       "Valid CPF Test if arg is Empty",
			wantValue:  false,
			cpfToCheck: "",
		},
		{
			give:       "Valid CPF Test if arg is Invalid",
			wantValue:  false,
			cpfToCheck: "b1080263",
		},
		{
			give:       "Valid CPF Test if arg is Zeros Numbers",
			wantValue:  false,
			cpfToCheck: "000.000.000-00",
		},
		{
			give:       "Valid CPF Test if arg is a Valid CPF",
			wantValue:  false,
			cpfToCheck: "111.111.111-11",
		},
		{
			give:       "Valid CPF Test if arg is a Valid CPF",
			wantValue:  true,
			cpfToCheck: "838.461.722-86",
		},
		{
			give:       "Valid CPF Test if arg is a Valid CPF",
			wantValue:  true,
			cpfToCheck: "313.396.023-77",
		},
		{
			give:       "Valid CPF Test if arg is a Valid CPF",
			wantValue:  true,
			cpfToCheck: "682.511.941-99",
		},
		{
			give:       "Valid CPF Test if arg is a Valid CPF",
			wantValue:  true,
			cpfToCheck: "288.752.243-40",
		},
		{
			give:       "Valid CPF Test if arg is a Valid CPF",
			wantValue:  true,
			cpfToCheck: "484.162.412-01",
		},
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			result := true

			if !IsValidCPF(tt.cpfToCheck) {
				result = false
			}
			CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}
