package cpfcnpj

import (
	"fmt"
	"regexp"
)

/*
const (
	NumTotalDigCPF  = 14
	SizeToValidDig1 = 9
	SizeToValidDig2 = 10
)
*/

/*
// converts a rune to an int.
func runeToInt(r rune) int {
	return int(r - '0')
}
*/
func isValidFormatCNPJ(cnpjToCheck string) bool {
	var cnpjRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)

	if len(cnpjToCheck) < NumTotalDigCPF {
		return false
	}

	return cnpjRegexp.MatchString(cnpjToCheck)
}

//Multiplica os digitos do cpf por 10 ou 11 *O numero não pode ter caracter especial
func MultiplyNumDigCNPJ(cpfToCheckOnlyNumber string, numIndexFinal int) uint64 {

	str_to_sum1 := cpfToCheckOnlyNumber[:numIndexFinal]
	digitMultiplier := (numIndexFinal + 1)

	multiplicationResult := 0
	for _, nextDigit := range str_to_sum1 {
		multiplicationResult += runeToInt(nextDigit) * digitMultiplier
		digitMultiplier--
	}

	restDivision := multiplicationResult % 11
	compereWithDig1 := 11 - restDivision
	if restDivision < 2 {
		compereWithDig1 = 0
	}

	//fmt.Printf("comperToDig1 [%d]\n\nFIM\n", compereWithDig1)
	return uint64(compereWithDig1)
}

/*

func isValidCPFOnlyValid(cpfToCheck string) bool {

	validDigit1, validDigit2 := getVerifyingDigits(cpfToCheck)
	print(validDigit1, validDigit2)

	sumDig1 := MultiplyNumDigCPF(cpfToCheck, SizeToValidDig1)
	sumDig2 := MultiplyNumDigCPF(cpfToCheck, SizeToValidDig2)
	print(sumDig1, sumDig2)

	if !ValidateVerifierDigit(sumDig1, validDigit1) {
		return false
	}

	return !ValidateVerifierDigit(sumDig2, validDigit2)
}


*/

//IsValidCNPJ Check if cnpj is valid
func IsValidCNPJ(cnpjToCheck string) bool {

	if !isValidFormatCNPJ(cnpjToCheck) {
		return false
	}

	cnpjFormated := formatToValidate(cnpjToCheck)
	fmt.Printf("cnpjFormated [%v]\n\nFIM\n", cnpjFormated)
	//return !isValidCPFOnlyValid(cpfFormated)

	return true
}
