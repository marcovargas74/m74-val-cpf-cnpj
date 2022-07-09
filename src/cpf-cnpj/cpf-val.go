package cpfcnpj

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
	"unicode"
)

const (
	NumTotalDigCPF  = 14
	SizeToValidDig1 = 9
	SizeToValidDig2 = 10
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

func formatToValidateCPF(cpfToFormat string) string {

	cpfClean := bytes.NewBufferString("")
	for _, digit := range cpfToFormat {
		if unicode.IsDigit(digit) {
			cpfClean.WriteRune(digit)
		}
	}

	return cpfClean.String()
}

func allDigitsIsEqual(cpfToCheck string) bool {

	if len(cpfToCheck) < 10 {
		return false
	}

	for pos, _ := range cpfToCheck {
		if cpfToCheck[0] != cpfToCheck[pos] {
			return false
		}

	}

	return true
}

func getVerifyingDigits(cpfToCheck string) (uint64, uint64) {
	size := len(cpfToCheck)
	str_d2 := cpfToCheck[size-1:]
	str_d1 := cpfToCheck[size-2 : size-1]

	int_dig1, err := strconv.Atoi(str_d1)
	if err != nil {
		log.Print(err)
		return 0, 0
	}

	int_dig2, err := strconv.Atoi(str_d2)
	if err != nil {
		log.Print(err)
		return 0, 0
	}

	return uint64(int_dig1), uint64(int_dig2)
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
	compereWithDig1 := 11 - restDivision
	if restDivision < 2 {
		compereWithDig1 = 0
	}

	//fmt.Printf("comperToDig1 [%d]\n\nFIM\n", compereWithDig1)
	return uint64(compereWithDig1)
}

func ValidateVerifierDigit(sumCpf, digToCheck uint64) bool {
	return sumCpf == digToCheck
}

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

//IsValidCPF Check if cpf is valid
func IsValidCPF(cpfToCheck string) bool {

	if !isValidFormatCPF(cpfToCheck) {
		return false
	}

	cpfFormated := formatToValidateCPF(cpfToCheck)
	if allDigitsIsEqual(cpfFormated) {
		return false
	}

	return !isValidCPFOnlyValid(cpfFormated)

}
