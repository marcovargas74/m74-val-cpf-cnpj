package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/validator.v2"
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
func StructAndJSON() string {
	//var create time.Time
	accountMaria := Account{Balance: 0, CreatedAt: time.Now()}
	accountMaria.ID = "abc"
	accountMaria.CPF = "111.111.111-11"
	accountMaria.Secret = "111"
	accountMaria.Name = "Maria"
	mariaJSON, _ := json.Marshal(accountMaria)
	fmt.Println(string(mariaJSON))
	return (string(mariaJSON))

	//Convert Json To struct
	/*var accountFromJSON Account
	json.Unmarshal(mariaJSON, &accountFromJSON)
	fmt.Println(accountFromJSON.Name)*/
}

func NewUUID() string {
	uuidNew, _ := uuid.NewV4()
	//fmt.Printf("UUIDv4: %s\n", u1)
	return uuidNew.String()
	//return gouuid.NewV4().String()

}

func IsValidUUID(uuidVal string) bool {
	//_, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	_, err := uuid.FromString(uuidVal)
	return err == nil
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

´
*/

func (a *Account) SaveAccount(w http.ResponseWriter, r *http.Request) {

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("\nSaveAccount Name:%s  cpf:%s balance %.2f\n", a.Name, a.CPF, a.Balance)
	if errs := validator.Validate(a); errs != nil {
		fmt.Printf("INVALIDO %v\n", errs) // do something
		w.WriteHeader(http.StatusBadRequest)
	}

	defer r.Body.Close()

	a.ID = NewUUID()
	a.CreatedAt = time.Now()
	fmt.Printf("UUIDv4: %s\n", a.ID)

	// Parsing UUID from string input
	//IsValidUUID
	///*u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if !IsValidUUID(a.ID) {
		fmt.Printf("Something gone wrong:")
	}

	//fmt.Printf("Successfully parsed: %s", u2)
	/*
		var account = domain.NewAccount(
			domain.AccountID(domain.NewUUID()),
			input.Name,
			input.CPF,
			domain.Money(input.Balance),
			time.Now(),
		)

		account, err := a.repo.Create(ctx, account)
		if err != nil {
			return a.presenter.Output(domain.Account{}), err
		}

		return a.presenter.Output(account), nil*/
	a.SaveAccountInDB()
	message := fmt.Sprintf("POST %v ID:%v", r.URL, a.ID)
	fmt.Fprint(w, message)
	w.WriteHeader(http.StatusOK)

}

func (a *Account) SaveAccountInDB() {
	db, err := sql.Open("mysql", "root:Mysql#2510@/bankAPI")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into accounts(id, nome, cpf, balance, secret, createAt) values(?,?,?,?,?,?)")

	_, err = stmt.Exec(a.ID, a.Name, a.CPF, a.Balance, a.Secret, a.CreatedAt)
	//stmt.Exec(2001, "Carlos")
	//_, err = stmt.Exec(1, "Tiago") //chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}

func (a *Account) GetAccounts(w http.ResponseWriter, r *http.Request) {
	//message := StructAndJSON()
	a.ShowAccountAll(w, r)

	//fmt.Fprint(w, message)
	w.WriteHeader(http.StatusOK)
}

func (a *Account) GetAccountByID(w http.ResponseWriter, r *http.Request, ID string) {

	fmt.Printf("   -->GetAccountByID [%s] \n", ID)
	if !IsValidUUID(ID) {
		fmt.Printf("Something gone wrong:")
		fmt.Fprint(w, "Something gone wrong: Invalid ID\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//message := StructAndJSON()
	a.ShowAccountByID(w, r, ID)

	//fmt.Fprint(w, message)
	w.WriteHeader(http.StatusOK)
}

//ShowAccountAll mostra todos as contas
func (a *Account) ShowAccountAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:Mysql#2510@/bankAPI")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from accounts")
	defer rows.Close()

	var usuarios []Account
	for rows.Next() {
		var usuario Account
		rows.Scan(&usuario.ID, &usuario.Name)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}

func (a *Account) ShowAccountByID(w http.ResponseWriter, r *http.Request, findID string) {
	db, err := sql.Open("mysql", "root:Mysql#2510@/bankAPI")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u Account
	db.QueryRow("select id, nome from accounts where id = ?", findID).Scan(&u.ID, &u.Name)
	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")

	fmt.Printf("DADOS DO BANOC id[%s] data[%s]\n", findID, string(json))

	fmt.Fprint(w, string(json))
}
