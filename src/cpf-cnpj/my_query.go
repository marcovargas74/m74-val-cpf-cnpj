package cpfcnpj

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

type MyQuery struct {
	ID        string    `json:"id"`
	Number    string    `json:"cpf"`
	IsValid   bool      `json:"is_valid" `
	IsCPF     bool      `json:"is_cpf" `
	IsCNPJ    bool      `json:"is_cnpj" `
	CreatedAt time.Time `json:"created_at"`
}

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

//NewUUID Cria um novo UUID valido
func NewUUID() string {
	uuidNew, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
	}
	return uuidNew.String()
}

//IsValidUUID Check if IUUID is valid
func IsValidUUID(uuidVal string) bool {
	_, err := uuid.FromString(uuidVal)
	return err == nil
}

/*
//newAccount Create a new account
func newAccount(name, cpf, secret string, balance float64) Account {
	return Account{
		ID:        NewUUID(),
		Name:      name,
		CPF:       cpf,
		Secret:    secret,
		Balance:   balance,
		CreatedAt: time.Now(),
	}
}
*/
//SaveQuery main fuction to save a new query in system
func (q *MyQuery) SaveQuery(w http.ResponseWriter, r *http.Request, newCPFofCNPJ string, isCPF bool) {

	q.Number = newCPFofCNPJ

	q.IsCNPJ = false
	q.IsCPF = false
	if isCPF {
		q.IsCPF = true
		if !IsValidCPF(q.Number) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Something gone wrong: Invalid CPF\n")
			log.Printf("Something gone wrong: Invalid CPF:%s\n", q.Number)
			q.IsValid = false
			return
		}
	} else {
		q.IsCNPJ = true
		if !IsValidCNPJ(q.Number) {
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprint(w, "Something gone wrong: Invalid CNPJ")
			log.Printf("Something gone wrong: Invalid CNPJ:%s\n", q.Number)
			q.IsValid = false
			return
		}
	}

	q.IsValid = true
	q.ID = NewUUID()
	q.CreatedAt = time.Now()
	fmt.Printf("UUIDv4: %s\n", q.ID)

	if !IsValidUUID(q.ID) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Something gone wrong: Invalid ID\n")
		log.Printf("Something gone wrong: Invalid ID:%s\n", q.ID)
		return
	}

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] number[%s] \n", r.UserAgent(), q.Number)
		w.WriteHeader(http.StatusOK)
		return
	}

	if !q.saveQueryInDB() {
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("Can not save account from %v", q.ID)
		fmt.Fprint(w, message)
		return
	}

	//json, err := json.Marshal(q)
	json, err := q.MarshalJSON()

	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}

func (q *MyQuery) saveQueryInDB() bool {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}

	stmt, err := tx.Prepare("insert into querys(id, number, is_valid, is_cpf, is_cnpj, createAt) values(?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec(q.ID, q.Number, q.IsValid, q.IsCPF, q.IsCNPJ, q.CreatedAt)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	tx.Commit()
	return true
}

//GetQuerys show All querys save in system
func (q *MyQuery) GetQuerys(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}
	q.ShowQueryAll(w, r)
}

//ShowQueryAll Show all querys
func (q *MyQuery) ShowQueryAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, number, is_valid, is_cpf, is_cnpj, createAt from querys")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "DB is Empty")
		return
	}
	defer rows.Close()

	var queryList []MyQuery
	for rows.Next() {
		var aQuery MyQuery
		rows.Scan(&aQuery.ID, &aQuery.Number, &aQuery.IsValid, &aQuery.IsCPF, &aQuery.IsCNPJ, &aQuery.CreatedAt)
		queryList = append(queryList, aQuery)
	}

	json, err := json.Marshal(queryList)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

//GetAccountByID return account pass token ID in arg
func (q *MyQuery) GetQuerysByNum(w http.ResponseWriter, r *http.Request, findCPFofCNPJ string) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}
	q.showQuerysByNum(w, r, findCPFofCNPJ)
}

func (a *MyQuery) showQuerysByNum(w http.ResponseWriter, r *http.Request, findNum string) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}
	defer db.Close()

	var aQuery MyQuery
	db.QueryRow("select id, number, is_valid, is_cpf, is_cnpj, createAt from querys where id = ?", findNum).Scan(&aQuery.ID, &aQuery.Number, &aQuery.IsValid, &aQuery.IsCPF, &aQuery.IsCNPJ, &aQuery.CreatedAt)
	json, err := aQuery.MarshalJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	log.Printf("DADOS DO DB id[%s] data[%s]\n", findNum, string(json))

}

//GetAccountByID return account pass token ID in arg
func (q *MyQuery) GetQuerysByType(w http.ResponseWriter, r *http.Request, isCPF bool) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}
	q.showQuerysByType(w, r, isCPF)
}

func (a *MyQuery) showQuerysByType(w http.ResponseWriter, r *http.Request, isCPF bool) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, number, is_valid, is_cpf, is_cnpj, createAt from querys where is_cpf = ?", isCPF)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}

	defer rows.Close()

	var queryList []MyQuery
	for rows.Next() {
		var aQuery MyQuery
		rows.Scan(&aQuery.ID, &aQuery.Number, &aQuery.IsValid, &aQuery.IsCPF, &aQuery.IsCNPJ, &aQuery.CreatedAt)
		queryList = append(queryList, aQuery)
	}

	json, err := json.Marshal(queryList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(json))
	log.Printf("DADOS DO DB - showQuerysByType IsCPF[%v] data[%s]\n", isCPF, string(json))
}

func (q *MyQuery) MarshalJSON() ([]byte, error) {
	type Alias MyQuery
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(q),
		CreatedAt: q.CreatedAt.Format("02-Jan-06 15:04:05"),
	})
}

/*





//ShowAccountAll show All accounts in a Bank
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

	json, err := json.Marshal(usuarios)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
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
	json, err := json.Marshal(aBalance)
	if err != nil {
		log.Println(err)
	}
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

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}

	stmt, err := tx.Prepare("update accounts set balance = ? where id = ?")
	if err != nil {
		log.Println(err)
	}

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
*/

/*
//SaveQuery main fuction to save a new query in system
func (q *MyQuery) SaveQueryBody(w http.ResponseWriter, r *http.Request) {

	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	q.IsValid = true
	log.Printf("\nSaveQuery Number:%s\n", q.Number)
	if errs := validator.Validate(q); errs != nil {
		log.Printf("INVALIDO %v\n", errs)
		w.WriteHeader(http.StatusBadRequest)
		q.IsValid = false
	}

	defer r.Body.Close()
	q.IsValid = true
	q.ID = NewUUID()
	q.CreatedAt = time.Now()
	fmt.Printf("UUIDv4: %s\n", q.ID)

	if !IsValidUUID(q.ID) {
		log.Printf("Something gone wrong: Invalid ID:%s\n", q.ID)
		fmt.Fprint(w, "Something gone wrong: Invalid ID\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !q.saveQueryInDB() {
		message := fmt.Sprintf("Can not save account from %v", q.ID)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(q)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	w.WriteHeader(http.StatusOK)

}*/
