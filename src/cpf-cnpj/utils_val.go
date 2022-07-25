package cpfcnpj

import (
	"bytes"
	"log"
	"strconv"
	"unicode"

	"github.com/gofrs/uuid"
)

//NewUUID Cria um novo UUID valido
func NewUUID() string {
	uuidNew, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
	}
	return uuidNew.String()
}

//IsValidUUID Check if IUUID is valid
func IsValidUUID(uuidVal string) bool {
	_, err := uuid.FromString(uuidVal)
	return err == nil
}

//FormatToValidate Strip special characters return a string with only digits
func FormatToValidate(cpfToFormat string) string {

	cpfClean := bytes.NewBufferString("")
	for _, digit := range cpfToFormat {
		if unicode.IsDigit(digit) {
			cpfClean.WriteRune(digit)
		}
	}

	return cpfClean.String()
}

//GetVerifyingDigits returns check digits
func GetVerifyingDigits(cpfToCheck string) (uint64, uint64) {
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

//ValidateVerifierDigit compare the sum of digits with the check digit
func ValidateVerifierDigit(sumCpf, digToCheck uint64) bool {
	return sumCpf == digToCheck
}

// RuneToInt converts a rune to an int.
func RuneToInt(r rune) int {
	return int(r - '0')
}

//AllDigitsIsEqual Check if all digits is Equal - This is invalid CPF
func AllDigitsIsEqual(cpfToCheck string) bool {

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

//CreateDB Create dataBase
func CreateDB() {
	InitDBMongo(IsUsingMongoDB)
}
