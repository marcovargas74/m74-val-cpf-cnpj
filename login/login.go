package login

import "fmt"

func login() {
	fmt.Println("login")
}

/*
/login
A entidade Login possui os seguintes atributos:

cpf
secret
Espera-se as seguintes ações:

POST /login - autentica a usuaria
Regras para esta rota

Deve retornar token para ser usado nas rotas autenticadas
*/
