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
		/*{
			give:      "Valid CPF Test if arg is Zeros Numbers",
			wantValue: true,
			inFindID:  "000.000.000-00",
		},
		{
			give:      "Valid CPF Test if arg is a Valid CPF",
			wantValue: true,
			inFindID:  "111.111.111-11",
		},

		{
			give:      "Valid CPF Test if arg is a Valid CPF",
			wantValue: true,
			inFindID:  "838.461.722-86",
		},
		{
			give:      "Valid CPF Test if arg is a Valid CPF",
			wantValue: true,
			inFindID:  "313.396.023-77",
		},
		{
			give:      "Valid CPF Test if arg is a Valid CPF",
			wantValue: true,
			inFindID:  "682.511.941-99",
		},*/
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
		/*{
			give:       "Get Digits To check if arg is Zeros Numbers",
			wantValue1: 1,
			wantValue2: 1,
			cpfToCheck: "111.111.111-11",
		},
		/*
			{
				give:       "Get Digits To check if arg is Zeros Numbers",
				wantValue1: 8,
				wantValue2: 6,
				cpfToCheck: "838.461.722-86",
			},
			{
				give:       "Get Digits To check if arg is Zeros Numbers",
				wantValue1: 7,
				wantValue2: 7,
				cpfToCheck: "313.396.023-77",
			},
			{
				give:       "Get Digits To check if arg is Zeros Numbers",
				wantValue1: 9,
				wantValue2: 9,
				cpfToCheck: "682.511.941-99",
			},*/
	}

	for _, tt := range tests {

		t.Run(tt.give, func(t *testing.T) {
			//result := true

			Dig1, Dig2 := getVerifyingDigits(tt.cpfToCheck)
			print(Dig1, Dig2)
			CheckIfEqualInt(t, Dig1, tt.wantValue1)
			CheckIfEqualInt(t, Dig2, tt.wantValue2)

			//if !getVerifyingDigits(tt.cpfToCheck) {
			//	result = false
			//}
			//CheckIfEqualBool(t, result, tt.wantValue)
		})

	}

}
