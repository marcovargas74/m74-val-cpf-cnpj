package m74bankAPI

import (
	"encoding/json"
	"fmt"
)

// LoginBank
type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

//StructAndJson teste do struct
func StructAndJson() {
	myLogin := Login{CPF: "111.111.111-11", Secret: "111"}
	loginJson, _ := json.Marshal(myLogin)
	fmt.Println(string(loginJson))
	//Convert Json To struct
	var myNewLogin Login
	json.Unmarshal(loginJson, &myNewLogin)
	fmt.Println(myNewLogin.Secret)

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
