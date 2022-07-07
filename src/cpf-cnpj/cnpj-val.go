package cpfcnpj

import "regexp"

//IsValidCNPJ Check if cnpj is valid
func IsValidCNPJ(cnpjToCheck string) bool {
	var cnpjRegexp = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$`)
	isValid := true

	if len(cnpjToCheck) < 14 {
		return false
	}

	if !cnpjRegexp.MatchString(cnpjToCheck) {
		return false
	}
	return isValid
}
