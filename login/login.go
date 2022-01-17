package login

import (
	"encoding/json"
	"fmt"
)

type login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

func structAndJson() {
	myLogin := login{"111.111.111-11", "111"}
	loginJson, _ := json.Marshal(myLogin)
	fmt.Println(string(loginJson))
	//Convert Json To struct
	var myNewLogin login
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
