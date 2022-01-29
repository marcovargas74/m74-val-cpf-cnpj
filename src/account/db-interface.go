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
	DBSourceOpen = "root:my-secret-pw@tcp(localhost:3307)/"
	DBSource     = "root:my-secret-pw@tcp(localhost:3307)/bankAPI" //root:Mysql#my-secret-pw@/bankAPI"
)

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
		log.Fatal(err)
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
	db, err := sql.Open("mysql", DBSource)
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

/*Antes de Usar o Banco deve-se Subir o servidor
service mysqld start
*/

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

//CreateDB Cria banco sql
func CreateDB(isDropTable bool) {
	fmt.Println("Conectado ao Banco...")

	// MYSQL LOCAL
	db, err := sql.Open("mysql", DBSourceOpen)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Conectado ao Banco")
	exec(db, "create database if not exists bankAPI")
	exec(db, "use bankAPI")
	if isDropTable {
		exec(db, "drop table if exists accounts")

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

	//exec(db, `create table IF NOT EXISTS accounts(
	/*ID        string    `json:"id"`
	Name      string    `json:"name" validate:"min=3,max=40"`
	CPF       string    `json:"cpf"`
	Balance   float64   `json:"balance"`
	Secret    string    `json:"secret" `
	CreatedAt time.Time `json:"created_at"`
	*/
	fmt.Println("Conectado ao Banco com sucesso!")
	//fmt.Println(exec(db, ".tables"))

}
