package cpfcnpj

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

func CreateDBMongo(isDockerRun bool) bool {

	urlMongo := "mongodb://localhost:27017"
	if isDockerRun {
		urlMongo = os.Getenv("MONGO_URL")
	}

	fmt.Printf("[%v]\n", urlMongo)
	session, err := mgo.Dial(urlMongo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ping", session.Ping())
	return true
}

/*
func (q *MyQuery) saveQueryInMongoDB() bool {
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

func (q *MyQuery) showQueryAllMongoDB(w http.ResponseWriter, r *http.Request) {
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

func (a *MyQuery) showQuerysByNumMongoDB(w http.ResponseWriter, r *http.Request, findNum string) {
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

func (a *MyQuery) showQuerysByTypeMongoDB(w http.ResponseWriter, r *http.Request, isCPF bool) {
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

func (a *MyQuery) deleteQuerysByNumMongoDB(w http.ResponseWriter, r *http.Request, findNum string) {
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

*/
