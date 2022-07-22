package cpfcnpj

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//Used blank to can used mysql commands
	_ "github.com/go-sql-driver/mysql"
)

const (
	//DBSourceOpenLocal Const used to Open Local db
	DBSourceOpenLocal = "root:my-secret-pw@tcp(localhost:3307)/"

	//DBSourceLocal Const used to acces Local db
	DBSourceLocal = "root:my-secret-pw@tcp(localhost:3307)/validatorAPP" //root:Mysql#my-secret-pw@/validatorAPP"

	//DBSourceOpenDocker Const used to Open Docker db
	DBSourceOpenDocker = "root:my-secret-pw@tcp(mysql-api)/" //mysql-api é o nome do serviço no docker-composer

	//DBSourceDocker Const used to acces Docker db
	DBSourceDocker = "root:my-secret-pw@tcp(mysql-api)/validatorAPP" //root:Mysql#my-secret-pw@/validatorAPP"

	//DBisDropTableSQL Clear TAbles
	DBisDropTableSQL = false
)

//AddrOpenDB VAR used to open and to access BD
var AddrOpenDB string

//AddrDB VAR data source name
var AddrDB string

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		log.Print(err)
	}
	return result
}

//InitDBSQL Connect To SQL Database
func InitDBSQL(isDropTable bool) {

	AddrOpenDB = DBSourceOpenDocker
	AddrDB = DBSourceDocker

	db, err := sql.Open("mysql", AddrOpenDB)
	if err != nil {
		log.Printf("Failed to connect to db Local Mysql...")
		AddrOpenDB = DBSourceOpenLocal
		AddrDB = DBSourceLocal
		db, err = sql.Open("mysql", AddrOpenDB)
		if err != nil {
			log.Printf("Failed to connect to db Local Mysql IP 127.0.0.1")
			log.Print(err)
		}

	}

	defer db.Close()

	fmt.Println("Successfully connected to the DB")
	exec(db, "create database if not exists validatorAPP")
	exec(db, "use validatorAPP")
	if isDropTable {
		exec(db, "drop table if exists querys")
	}

	exec(db, `create table IF NOT EXISTS querys(
	idx integer auto_increment,
	id varchar(40) ,
	number varchar(18),
	is_valid boolean,
	is_cpf boolean,
	is_cnpj boolean,
    createAt datetime,
	PRIMARY KEY (idx)
	)`)

	fmt.Println("Successfully connected to the DB!")

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

func (q *MyQuery) showQueryAll(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusNotFound)
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

func (a *MyQuery) showQuerysByType(w http.ResponseWriter, r *http.Request, isCPF bool) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Fail to access DB")
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, number, is_valid, is_cpf, is_cnpj, createAt from querys where is_cpf = ?", isCPF)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Not Found elements to this Type ")
		return
	}

	defer rows.Close()

	var queryList []MyQuery
	for rows.Next() {
		var aQuery MyQuery
		rows.Scan(&aQuery.ID, &aQuery.Number, &aQuery.IsValid, &aQuery.IsCPF, &aQuery.IsCNPJ, &aQuery.CreatedAt)
		queryList = append(queryList, aQuery)
	}

	if len(queryList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found elements to this Type ")
		return
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

func (a *MyQuery) deleteQuerysByNum(w http.ResponseWriter, r *http.Request, findNum string) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Fail to access DB")
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
		fmt.Fprint(w, "Fail prepare command to access DB")
		return
	}

	result, err := stmt.Exec(findNum)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Fail To delete register in access DB")
		return
	}

	numElement, err := result.RowsAffected()
	if numElement == 0 || err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found elements to this Type - Delete")
		return
	}

	tx.Commit()

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "SUCCESS TO DELETE CPF/CNPJ")
}
