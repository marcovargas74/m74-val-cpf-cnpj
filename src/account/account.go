package account

import (
	"encoding/json"
	"fmt"
)

/*
const (
	aConst = "ola"
)*/

/*TODO: refatorar para que a struct
  seja composta pela strucr login + client*/

//Account Struct Used to creat a account
type Account struct {
	Client
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"` //TODO change to date*/
}

func (a Account) setID(name string) {
	a.Name = "Maria"
}

//StructAndJSON Just Test
func StructAndJSON() {
	accountMaria := Account{Balance: 0, CreatedAt: "17-01-2022"}
	accountMaria.ID = "abc"
	accountMaria.CPF = "111.111.111-11"
	accountMaria.Secret = "111"
	accountMaria.Name = "Maria"
	mariaJSON, _ := json.Marshal(accountMaria)
	fmt.Println(string(mariaJSON))
	//Convert Json To struct
	var accountFromJSON Account
	json.Unmarshal(mariaJSON, &accountFromJSON)
	fmt.Println(accountFromJSON.Name)
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
