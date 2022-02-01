package account

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
	DBSourceLocal = "root:my-secret-pw@tcp(localhost:3307)/bankAPI" //root:Mysql#my-secret-pw@/bankAPI"

	//DBSourceOpenDocker Const used to Open Docker db
	DBSourceOpenDocker = "root:my-secret-pw@tcp(mysql-api)/" //mysql-api é o nome do serviço no docker-composer

	//DBSourceDocker Const used to acces Docker db
	DBSourceDocker = "root:my-secret-pw@tcp(mysql-api)/bankAPI" //root:Mysql#my-secret-pw@/bankAPI"

)

//AddrOpenDB VAR used to open and to access BD
var AddrOpenDB string

//AddrDB VAR used to to access BD (selects/Update/Insert)
var AddrDB string

//ShowAccountAll mostra todos as contas
func ShowAccountAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "fail to access DB")
		return
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

/*Antes de Usar o Banco deve-se Subir o servidor
service mysqld start
*/

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		log.Print(err)
	}
	return result
}

//CreateDB Cria banco sql
func CreateDB(isDropTable bool) {
	// MYSQL LOCAL
	AddrOpenDB = DBSourceOpenDocker
	AddrDB = DBSourceDocker

	db, err := sql.Open("mysql", AddrOpenDB)
	if err != nil {
		log.Printf("FALHA ao conectar ao Banco Mysql LOCAL...")
		AddrOpenDB = DBSourceOpenLocal
		AddrDB = DBSourceLocal
		db, err = sql.Open("mysql", AddrOpenDB)
		if err != nil {
			log.Printf("FALHA ao conectar ao Banco Mysql Local IP 127.0.0.1")
			log.Print(err)
		}

	}

	defer db.Close()

	fmt.Println("Conectado ao Banco")
	exec(db, "create database if not exists bankAPI")
	exec(db, "use bankAPI")
	if isDropTable {
		exec(db, "drop table if exists accounts")
		exec(db, "drop table if exists transfers")
	}

	exec(db, `create table IF NOT EXISTS accounts(
	idx integer auto_increment,
	id varchar(40) ,
	nome varchar(80),
	cpf varchar(15),
	balance float,
	secret varchar(80),
    createAt datetime,
	PRIMARY KEY (idx)
	)`)

	exec(db, `create table IF NOT EXISTS transfers(
		idx integer auto_increment,
		id varchar(40) ,
		ori varchar(40),
		dest varchar(40),
		amount float,
		createAt datetime,
		PRIMARY KEY (idx)
		)`)

	fmt.Println("Conectado ao Banco com sucesso!")

}
