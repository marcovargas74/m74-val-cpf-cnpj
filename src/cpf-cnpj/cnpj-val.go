package cpfcnpj

import (
	"fmt"
	"regexp"
)

const (
	SizeToValidTotalCNPJDig1 = 12
	SizeToValidTotalCNPJDig2 = SizeToValidTotalCNPJDig1 + 1

	SizeToValidDig1CNPJ       = 4
	SizeToValidDig2CNPJ       = SizeToValidDig1CNPJ + 1
	SizeToValidDigDefaultCNPJ = SizeToValidDig1CNPJ + SizeToValidDig2CNPJ
)

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

	indexLastDigitToCheck := SizeToValidTotalCNPJDig1
	if numIndexFinal == SizeToValidDig2CNPJ {
		indexLastDigitToCheck++
	}

	str_cnpj_without_verifyDigit := cpfToCheckOnlyNumber[:indexLastDigitToCheck]
	digitMultiplier := SizeToValidDigDefaultCNPJ

	str_to_sum1 := str_cnpj_without_verifyDigit[4:indexLastDigitToCheck]
	fmt.Printf("str[%s] FinalIndex[%d]multiplicationResult [%d] \n", str_to_sum1, numIndexFinal, digitMultiplier)

	multiplicationResult := 0
	for _, nextDigit := range str_to_sum1 {
		fmt.Printf("digito[%d] multi[%d]\n", runeToInt(nextDigit), digitMultiplier)

		multiplicationResult += runeToInt(nextDigit) * digitMultiplier
		digitMultiplier--
	}

	fmt.Printf("Parte 1 multiplicationResult [%d] \n", multiplicationResult)

	///==== PARTE 2 ====
	digitMultiplier = numIndexFinal + 1

	str_to_sum1 = str_cnpj_without_verifyDigit[:numIndexFinal]
	fmt.Printf("str[%s] FinalIndex[%d]multiplicationResult [%d] \n", str_to_sum1, numIndexFinal, digitMultiplier)

	//multiplicationResult = 0
	for _, nextDigit := range str_to_sum1 {
		fmt.Printf("digito[%d] multi[%d]\n", runeToInt(nextDigit), digitMultiplier)

		multiplicationResult += runeToInt(nextDigit) * digitMultiplier
		digitMultiplier--
	}

	fmt.Printf("\nParte 2 multiplicationResult [%d] \n", multiplicationResult)

	restDivision := multiplicationResult % 11
	compereWithDig1 := 11 - restDivision

	fmt.Printf("str[%s] multiplicationResult [%d] resto[%d]\n", str_to_sum1, multiplicationResult, restDivision)

	if restDivision < 2 {
		compereWithDig1 = 0
	}

	fmt.Printf("comperToDig1 [%d]\n\nFIM\n", compereWithDig1)
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
