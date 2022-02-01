package account

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

//Balance used to create a Json Return to endpoint
type Balance struct {
	Value float64 `json:"balance"`
}

/*
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
	fmt.Println(accountFromJSON.Name)* /
}
*/

//IsValidCPF Check if cpf is valid
func IsValidCPF(cpf string) bool {
	isValid := true
	if len(cpf) != 14 {
		isValid = false
	}
	return isValid
}

//NewUUID Cria um novo UUID valido
func NewUUID() string {
	uuidNew, _ := uuid.NewV4()
	return uuidNew.String()
	//return gouuid.NewV4().String()

}

//IsValidUUID Check if IUUID is valid
func IsValidUUID(uuidVal string) bool {
	//_, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	_, err := uuid.FromString(uuidVal)
	return err == nil
}

//SecretToHash change a string in a HAshValue
func SecretToHash(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))

}

//HashToSecret change a HAshValue in a visible string
func HashToSecret(hashIn string) string {
	passw, _ := base64.StdEncoding.DecodeString(hashIn)
	return string(passw)
}

func newAccount(name, cpf, secret string, balance float64) Account {
	//fmt.Printf("->Name: %s CreateAt: \n ", name)

	return Account{
		ID:        NewUUID(),
		Name:      name,
		CPF:       cpf,
		Secret:    secret,
		Balance:   balance,
		CreatedAt: time.Now(),
	}
}

//SaveAccount main fuction to save a new account in system
func (a *Account) SaveAccount(w http.ResponseWriter, r *http.Request) {

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("\nSaveAccount Name:%s  cpf:%s balance %.2f\n", a.Name, a.CPF, a.Balance)
	if errs := validator.Validate(a); errs != nil {
		log.Printf("INVALIDO %v\n", errs)
		w.WriteHeader(http.StatusBadRequest)
	}

	defer r.Body.Close()

	a.ID = NewUUID()
	a.CreatedAt = time.Now()
	fmt.Printf("UUIDv4: %s\n", a.ID)

	if !IsValidUUID(a.ID) {
		log.Printf("Something gone wrong: Invalid ID:%s\n", a.ID)
		fmt.Fprint(w, "Something gone wrong: Invalid ID\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !a.saveAccountInDB() {
		message := fmt.Sprintf("Can not save account from %v", a.ID)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(a)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	w.WriteHeader(http.StatusOK)

}

func (a *Account) saveAccountInDB() bool {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into accounts(id, nome, cpf, balance, secret, createAt) values(?,?,?,?,?,?)")
	secretHash := SecretToHash(a.Secret)
	_, err = stmt.Exec(a.ID, a.Name, a.CPF, a.Balance, secretHash, a.CreatedAt)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	tx.Commit()
	return true
}

//GetAccounts show All account save in system
func (a *Account) GetAccounts(w http.ResponseWriter, r *http.Request) {
	a.ShowAccountAll(w, r)
	w.WriteHeader(http.StatusOK)
}

//GetAccountByID return account pass token ID in arg
func (a *Account) GetAccountByID(w http.ResponseWriter, r *http.Request, ID string) {

	//log.Printf("   -->GetAccountByID [%s] \n", ID)
	if !IsValidUUID(ID) {
		log.Printf("Something gone wrong: Invalid ID:%s\n", ID)
		fmt.Fprint(w, "Something gone wrong: Invalid ID\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a.showAccountByID(w, r, ID)
	w.WriteHeader(http.StatusOK)
}

//ShowAccountAll mostra todos as contas
func (a *Account) ShowAccountAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return

	}
	defer db.Close()

	rows, err := db.Query("select id, nome, cpf, balance, secret, createAt from accounts")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return

	}
	defer rows.Close()

	var usuarios []Account
	for rows.Next() {
		var account Account
		rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Balance, &account.Secret, &account.CreatedAt)
		usuarios = append(usuarios, account)
	}

	json, _ := json.Marshal(usuarios)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}

func (a *Account) showAccountByID(w http.ResponseWriter, r *http.Request, findID string) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}
	defer db.Close()

	var account Account
	db.QueryRow("select id, nome, cpf, balance, secret, createAt from accounts where id = ?", findID).Scan(&account.ID, &account.Name, &account.CPF, &account.Balance, &account.Secret, &account.CreatedAt)
	account.Secret = "*****"
	json, _ := json.Marshal(account)

	w.Header().Set("Content-Type", "application/json")

	log.Printf("DADOS DO DB id[%s] data[%s]\n", findID, string(json))

	fmt.Fprint(w, string(json))
}

//ShowBalanceByID return Balance account pass token ID in arg
func (a *Account) ShowBalanceByID(w http.ResponseWriter, r *http.Request, findID string) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return

	}
	defer db.Close()

	var aBalance Balance
	db.QueryRow("select balance from accounts where id = ?", findID).Scan(&aBalance.Value)
	json, _ := json.Marshal(aBalance)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

//UpdateBalanceByID update balance value account in DB
func UpdateBalanceByID(accID string, newTransationValue float64) bool {

	accountInBD, err := GetAccountByID(accID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	accountInBD.Balance = newTransationValue

	log.Printf("<<-id:%s val %.2f\n", accountInBD.ID, accountInBD.Balance)
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update accounts set balance = ? where id = ?")
	_, err = stmt.Exec(accountInBD.Balance, accountInBD.ID)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	tx.Commit()
	return true

}

//GetAccountByID return account pass token ID in arg
func GetAccountByID(findID string) (Account, error) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		return Account{}, err
	}
	defer db.Close()

	var account Account
	db.QueryRow("select id, nome, cpf, balance, secret, createAt from accounts where id = ?", findID).Scan(&account.ID, &account.Name, &account.CPF, &account.Balance, &account.Secret, &account.CreatedAt)
	fmt.Printf("DADOS DO BANOC id[%s] data[%s]\n", findID, account.CPF)
	return account, nil
}

//GetAccountByCPF Retorna a conta passando o CPF como parametro
func GetAccountByCPF(findCPF string) (Account, error) {

	if !IsValidCPF(findCPF) {
		return Account{}, fmt.Errorf("CPF inválido: %s", findCPF)
	}

	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		return Account{}, err
	}
	defer db.Close()

	var account Account
	db.QueryRow("select id, nome, cpf, balance, secret, createAt from accounts where cpf = ?", findCPF).Scan(&account.ID, &account.Name, &account.CPF, &account.Balance, &account.Secret, &account.CreatedAt)
	fmt.Printf("DADOS DO BANOC id[%s] data[%s]\n", findCPF, account.CPF)
	return account, nil
}
