package cpfcnpj

import "regexp"

//IsValidCPF Check if cpf is valid
func IsValidCPF(cpfToCheck string) bool {
	var CPFRegexp = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)

	isValid := true
	if len(cpfToCheck) != 14 {
		isValid = false
	}

	if !CPFRegexp.MatchString(cpfToCheck) {
		return false
	}

	return isValid
}
