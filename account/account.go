package account

import (
	"encoding/json"
	"fmt"
)

const (
	aConst = "ola"
)

/*TODO: refatorar para que a struct
  seja composta pela strucr login + client*/

type Account struct {
	Client
	Balance    float64 `json:"balance"`
	Created_at string  `json:"created_at"` //TODO change to date*/
}

func (a Account) setId(name string) {
	a.Name = "Maria"
}

func StructAndJson() {
	accountMaria := Account{Balance: 0, Created_at: "17-01-2022"}
	accountMaria.ID = "abc"
	accountMaria.CPF = "111.111.111-11"
	accountMaria.Secret = "111"
	accountMaria.Name = "Maria"
	mariaJson, _ := json.Marshal(accountMaria)
	fmt.Println(string(mariaJson))
	//Convert Json To struct
	var accountFromJson Account
	json.Unmarshal(mariaJson, &accountFromJson)
	fmt.Println(accountFromJson.Name)
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
