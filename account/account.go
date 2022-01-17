package account

import (
	"encoding/json"
	"fmt"
)

const (
	aConst = "ola"
)

/*TODO:criar uma classe clinete e outra accout
  account vai ter cleint + saldo*/

type account struct {
	ID         string  `json:"id"`
	CPF        string  `json:"cpf"`
	Secret     string  `json:"secret"`
	Name       string  `json:"name"`
	Balance    float64 `json:"balance"`
	Created_at string  `json:"created_at"` //TODO change to date
}

func (a account) setId(name string) {
	a.Name = "Maria"
}

func structAndJson() {
	accountMaria := account{"abc", "111.111.111-11", "111", "Maria", 0, "17-01-2022"}
	mariaJson, _ := json.Marshal(accountMaria)
	fmt.Println(string(mariaJson))
	//Convert Json To struct
	var accountFromJson account
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
