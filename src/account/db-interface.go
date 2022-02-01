package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	//	"github.com/marcovargas74/m74-bank-api/src/account"
)

const (
	//DBSource       = "root:Mysql#2510@/bankAPI"
	DBSourceOpenLocal = "root:my-secret-pw@tcp(localhost:3307)/"
	DBSourceLocal     = "root:my-secret-pw@tcp(localhost:3307)/bankAPI" //root:Mysql#my-secret-pw@/bankAPI"
	//mysql-api é o nome do serviço no docker-composer
	DBSourceOpenDocker = "root:my-secret-pw@tcp(mysql-api)/"
	DBSourceDocker     = "root:my-secret-pw@tcp(mysql-api)/bankAPI" //root:Mysql#my-secret-pw@/bankAPI"

)

var AddrOpenDB string
var AddrDB string

//DBMySQL Data base Used in BANK
//var DBMySQL *sql.DB

/*
//Usuario :)
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

//UsuarioHandler analisa o request e delega para função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)

	case r.Method == "GET":
		usuarioTodos(w, r)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}

}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:Mysql#2510@/cursogo")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)
	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
*/
//ShowAccountAll mostra todos as contas
func ShowAccountAll(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", AddrDB)
	if err != nil {
		log.Println(err)
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
