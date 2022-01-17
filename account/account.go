package account

const (
	aConst = "ola"
)


struct 

func getCliente(name string) string {

	if name == "Maria" {
		return "20"
	}

	if name == "Pedro" {
		return "10"
	}
	return ""

}


/*

A entidade Account possui os seguintes atributos:

id
name
cpf
secret
balance
created_at
Espera-se as seguintes ações:

GET /accounts - obtém a lista de contas
GET /accounts/{account_id}/balance - obtém o saldo da conta
POST /accounts - cria uma Account
Regras para esta rota

balance pode iniciar com 0 ou algum valor para simplificar
secret deve ser armazenado como hash


*/