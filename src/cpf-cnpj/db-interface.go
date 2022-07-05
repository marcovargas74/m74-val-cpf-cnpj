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

)

//AddrOpenDB VAR used to open and to access BD
var AddrOpenDB string

//AddrDB VAR data source name
var AddrDB string

//ShowQueryAll Show all querys
func ShowQueryAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from querys")
	if err != nil {
		log.Println(err)
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
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		log.Print(err)
	}
	return result
}

//CreateDB Create SQL dataBase
func CreateDB(isDropTable bool) {
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
