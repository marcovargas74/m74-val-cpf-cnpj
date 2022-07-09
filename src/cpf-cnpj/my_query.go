package cpfcnpj

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//MyQuery Strutc Main Used in AppValidate
type MyQuery struct {
	ID        string    `json:"id"`
	Number    string    `json:"cpf"`
	IsValid   bool      `json:"is_valid" `
	IsCPF     bool      `json:"is_cpf" `
	IsCNPJ    bool      `json:"is_cnpj" `
	CreatedAt time.Time `json:"created_at"`
}

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
		message := fmt.Sprintf("Can not save cpf/cnpj %v ", q.Number)
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

//GetQuerysByNum return CPF/CNPJ pass number in arg
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
	db.QueryRow("select id, number, is_valid, is_cpf, is_cnpj, createAt from querys where number = ?", findNum).Scan(&aQuery.ID, &aQuery.Number, &aQuery.IsValid, &aQuery.IsCPF, &aQuery.IsCNPJ, &aQuery.CreatedAt)
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

//GetQuerysByType return ALL CPF or CNPJ pass type in arg
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

//DeleteQuerysByNum Delete Number
func (q *MyQuery) DeleteQuerysByNum(w http.ResponseWriter, r *http.Request, findCPFofCNPJ string) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}
	q.deleteQuerysByNum(w, r, findCPFofCNPJ)
}

func (a *MyQuery) deleteQuerysByNum(w http.ResponseWriter, r *http.Request, findNum string) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}

	stmt, err := db.Prepare("delete from querys where number = ?")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail prepare command to access DB")
		return
	}

	_, err = stmt.Exec(findNum)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail To delete register in access DB")
		return
	}

	tx.Commit()

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "SUCCESS TO DELETE CPF/CNPJ")

}

//MarshalJSON format the date
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
