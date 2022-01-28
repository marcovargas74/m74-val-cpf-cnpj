package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
const (
	aConst = "ola"
)*/

/*TODO: refatorar para validar todos os dados /
//Account Struct Used to creat a account
type Account struct {
	ID        string  `json:"id" validate:"required,uuid4"`
	Name      string  `json:"name "validate:"min=3,max=40,regexp=^[a-zA-Z]*$""`
	CPF       string  `json:"cpf" required`
	Balance   float64 `json:"balance" validate:"gt=0,required"`
	Secret    string  `json:"secret" validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	CreatedAt string  `json:"created_at"`

}
*/

//Account Struct Used to creat a account
type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"min=3,max=40"`
	CPF       string    `json:"cpf"`
	Balance   float64   `json:"balance"`
	Secret    string    `json:"secret" `
	CreatedAt time.Time `json:"created_at"`
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a Account) ValidName(name string) bool {
	isValid := false
	if name != "" {
		isValid = true
	}

	return isValid
}

//StructAndJSON Just Test
func StructAndJSON() {
	//var create time.Time
	accountMaria := Account{Balance: 0, CreatedAt: time.Now()}
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

func (a *Account) SaveAccount(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("content-type", "application/json")
	//accountID := r.URL.Path[len("/accounts/"):]
	//fmt.Printf("Account Body: %v\n", r.Body)
	//fmt.Printf("Account R Body: %v\n", r)
	//fmt.Printf("Account W Body: %v\n", w)

	//Convert Json To struct
	/*var accountJSON Account
	if err := json.NewDecoder(r.Body).Decode(&accountJSON); err != nil {
		//r.Err.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	//acc := Account{Name: tt.inName, Balance: 0.00}
	if errs := validator.Validate(accountJSON); errs != nil {
		fmt.Printf("INVALIDO %v\n", errs) // do something
		w.WriteHeader(http.StatusBadRequest)
	}*/

	//defer r.Body.Close()

	//json.Unmarshal(mariaJSON, &accountFromJSON)
	fmt.Printf("Name:%s  cpf:%s balance %.2f\n", a.Name, a.CPF, a.Balance)

	message := fmt.Sprintf("POST %v", r.URL)
	//fmt.Printf("account Post %s\n", accountID)
	fmt.Fprint(w, message)
	w.WriteHeader(http.StatusOK)
	//accountJSON.SaveAccount(w, r)

}

func (a *Account) GetAccounts(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("content-type", "application/json")
	//accountID := r.URL.Path[len("/accounts/"):]
	//fmt.Printf("Account Body: %v\n", r.Body)
	//fmt.Printf("Account R Body: %v\n", r)
	//fmt.Printf("Account W Body: %v\n", w)

	//Convert Json To struct
	/*	var accountJSON Account
		if err := json.NewDecoder(r.Body).Decode(&accountJSON); err != nil {
			//r.Err.NewError(err, http.StatusBadRequest).Send(w)
			return
		}//

		//acc := Account{Name: tt.inName, Balance: 0.00}
		if errs := validator.Validate(accountJSON); errs != nil {
			fmt.Printf("INVALIDO %v\n", errs) // do something
			w.WriteHeader(http.StatusBadRequest)
		}

		//defer r.Body.Close()
	*/
	//json.Unmarshal(mariaJSON, &accountFromJSON)
	//fmt.Printf("Name:%s  cpf:%s balance %.2f\n", accountJSON.Name, accountJSON.CPF, accountJSON.Balance)

	message := fmt.Sprintf("GET %v", r.URL)
	//fmt.Printf("account Post %s\n", accountID)
	fmt.Fprint(w, message)
	w.WriteHeader(http.StatusOK)
	//accountJSON.SaveAccount(w, r)

}
