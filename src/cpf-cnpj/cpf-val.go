package cpfcnpj

import (
	"regexp"
)

const (
	NumTotalDigCPF     = 14
	SizeToValidDig1CPF = 9
	SizeToValidDig2CPF = 10
)

// converts a rune to an int.
func runeToInt(r rune) int {
	return int(r - '0')
}

func isValidFormatCPF(cpfToCheck string) bool {
	var CPFRegexp = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)

	if len(cpfToCheck) != NumTotalDigCPF {
		return false
	}

	return CPFRegexp.MatchString(cpfToCheck)
}

func allDigitsIsEqual(cpfToCheck string) bool {

	if len(cpfToCheck) < 10 {
		return false
	}

	for pos := range cpfToCheck {
		if cpfToCheck[0] != cpfToCheck[pos] {
			return false
		}

	}

	return true
}

//Multiplica os digitos do cpf por 10 ou 11 *O numero não pode ter caracter especial
func MultiplyNumDigCPF(cpfToCheckOnlyNumber string, numIndexFinal int) uint64 {

	str_to_sum1 := cpfToCheckOnlyNumber[:numIndexFinal]
	digitMultiplier := (numIndexFinal + 1)

	multiplicationResult := 0
	for _, nextDigit := range str_to_sum1 {
		multiplicationResult += runeToInt(nextDigit) * digitMultiplier
		digitMultiplier--
	}

	restDivision := multiplicationResult % 11
	compareWithDig1 := 11 - restDivision
	if restDivision < 2 {
		compareWithDig1 = 0
	}

	//fmt.Printf("comperToDig1 [%d]\n\nFIM\n", compereWithDig1)
	return uint64(compareWithDig1)
}

func isValidCPFOnlyValid(cpfToCheck string) bool {

	validDigit1, validDigit2 := getVerifyingDigits(cpfToCheck)
	print(validDigit1, validDigit2)

	sumDig1 := MultiplyNumDigCPF(cpfToCheck, SizeToValidDig1CPF)
	sumDig2 := MultiplyNumDigCPF(cpfToCheck, SizeToValidDig2CPF)
	print(sumDig1, sumDig2)

	if !ValidateVerifierDigit(sumDig1, validDigit1) {
		return false
	}

	return !ValidateVerifierDigit(sumDig2, validDigit2)
}

//IsValidCPF Check if cpf is valid
func IsValidCPF(cpfToCheck string) bool {

	if !isValidFormatCPF(cpfToCheck) {
		return false
	}

	cpfFormated := formatToValidate(cpfToCheck)
	if allDigitsIsEqual(cpfFormated) {
		return false
	}

	return !isValidCPFOnlyValid(cpfFormated)

}
