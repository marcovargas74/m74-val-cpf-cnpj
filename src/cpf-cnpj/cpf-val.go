package cpfcnpj

import "regexp"

func FormatoToValidateCPF(cpfToCheck string) string {
	return cpfToCheck
}

func getVerifyingDigits(cpfToCheck string) (uint64, uint64) {
	//retorna os digitos verificadores
	return 0, 0
}

//Multiplica os digitos do cpf por 10 ou 11
func MultiplyNumDigCPF(cpfToCheck string) (uint64, uint64) {
	return 0, 0
}

func ValidateVerifierDigit(sumCpf, digToCheck uint64) bool {
	return sumCpf == digToCheck
}

func isValidCPFOnlyValid(cpfToCheck string) bool {

	validDigit1, validDigit2 := getVerifyingDigits(cpfToCheck)
	print(validDigit1, validDigit2)

	sumDig1, sumDig2 := MultiplyNumDigCPF(cpfToCheck)
	print(sumDig1, sumDig2)

	if !ValidateVerifierDigit(sumDig1, validDigit1) {
		return false
	}

	return !ValidateVerifierDigit(sumDig2, validDigit2)
}

//IsValidCPF Check if cpf is valid
func IsValidCPF(cpfToCheck string) bool {
	var CPFRegexp = regexp.MustCompile(`^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)

	isValid := true
	if len(cpfToCheck) != 14 {
		return false
	}

	//FormatoToValidateCPF()
	if !CPFRegexp.MatchString(cpfToCheck) {
		return false
	}

	if !isValidCPFOnlyValid(cpfToCheck) {
		return false
	}

	return isValid
}

/*

Tomes como exemplo o CPF fictício : 111.444.777-05.

a - Cálculo do primeiro dígito

O primeiro passo é calcular o primeiro dígito verificador, e para isso, separamos os primeiros 9 dígitos do CPF (111.444.777) e multiplicamos cada um dos números, da direita para a esquerda por números crescentes a partir do número 2, como no exemplo abaixo:

1	1	1	4	4	4	7	7	7
10	9	8	7	6	5	4	3	2
10	9	8	28	24	20	28	21	14
Multiplicamos cada digito do CPF pelo respectivo número e somamos cada um dos resultados : 10+9+8+28+24+20+28+21+14 = 162

Pegamos o resultado obtido 162 e dividimos por 11.  Consideramos como quociente apenas o valor inteiro.

162 / 11  =    14  com resto 8

- Se o resto da divisão for menor que 2, então o dígito é igual a 0 (Zero).
- Se o resto da divisão for maior ou igual a 2, então o dígito verificador é igual a 11 menos o resto da divisão (11 - resto).

No nosso exemplo temos que o resto é 8 então faremos 11-8 = 3

Logo o primeiro dígito verificador é 3. Então podemos escrever o CPF com os dois dígitos calculados :  111.444.777-3X

b - Cálculo do segundo dígito

Para  calcular o segundo dígito vamos usar o primeiro digito já calculado. Vamos montar a mesma tabela de multiplicação usada no cálculo do primeiro dígito. Só que desta vez usaremos na segunda linha os valores 11,10,9,8,7,6,5,4,3,2 já que estamos incluindo mais um digito no cálculo(o primeiro dígito calculado):

1	1	1	4	4	4	7	7	7	3
11	10	9	8	7	6	5	4	3	2
11	10	9	32	28	24	35	28	21	6
Novamente, efetuamos somamos o resultado da multiplicação : 11 + 10 + 9 + 32 + 28 + 24 + 35 + 28 + 21 + 6 = 204

Dividimos o total do somatório por 11 e consideramos o resto da divisão.

204 / 11  =  18  e  resto 6

Após obter o resto da divisão, precisamos aplicar a mesma regra que utilizamos para obter o primeiro dígito:

- Se o resto da divisão for menor que 2, então o dígito é igual a 0 (Zero).
- Se o resto da divisão for maior ou igual a 2, então o dígito é igual a 11 menos o resto da divisão (11 - resto).

11-6= 5   logo 5 é o nosso segundo dígito verificador.

Logo o nosso CPF fictício será igual a : 111.444.777-35.

*/
