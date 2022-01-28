package m74bankapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"github.com/marcovargas74/m74-bank-api/src/account"
	//"github.com/marcovargas74/m74-bank-api/src/account"
	account "github.com/marcovargas74/m74-bank-api/src/account"
	"gopkg.in/validator.v2"
)

const (
	serverPort = ":5000"
)

type ServerBank struct {
	//Armazenamento ArmazenamentoAccount
	http.Handler
}

func (s *ServerBank) CallbackAccounts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	var accountJSON account.Account
	if err := json.NewDecoder(r.Body).Decode(&accountJSON); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if errs := validator.Validate(accountJSON); errs != nil {
		fmt.Printf("INVALIDO %v\n", errs) // do something
		w.WriteHeader(http.StatusBadRequest)
	}

	defer r.Body.Close()

	//json.Unmarshal(mariaJSON, &accountFromJSON)
	//fmt.Printf("Name:%s  cpf:%s balance %.2f\n", accountJSON.Name, accountJSON.CPF, accountJSON.Balance)

	switch r.Method {
	case http.MethodPost:
		accountJSON.SaveAccount(w, r)
	case http.MethodGet:
		accountJSON.GetAccounts(w, r)
		//message := fmt.Sprintf("GET %v", r.URL)
		//fmt.Printf("accountID GET %s", accountID)
		//fmt.Fprint(w, message)
		//w.WriteHeader(http.StatusOK)
		//s.mostraPontuacao(w, accountID)
	default:
		message := fmt.Sprintf("CallbackAccounts data in %v", r.URL)
		//fmt.Printf("accountID %s", accountID)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusOK)
	}

}

func (s *ServerBank) CallbackLogin(w http.ResponseWriter, r *http.Request) {

	client := r.URL.Path[len("/login/"):]
	fmt.Printf("Login Body: %v\n", r.Body)

	w.Header().Set("content-type", "application/json")

	switch r.Method {
	case http.MethodPost:
		message := fmt.Sprintf("POST %v", r.URL)
		fmt.Printf("client POST %s\n", client)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		message := fmt.Sprintf("%v", r.URL)
		fmt.Printf("client GET %s", client)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusOK)
	default:
		message := fmt.Sprintf("CallbackLogin Get data in %v", r.URL)
		fmt.Printf("client GET %s", client)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusOK)
	}

}

func (s *ServerBank) CallbackTransfer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	transfer := r.URL.Path[len("/transfers/"):]
	fmt.Printf("tranfer: %v\n", r.Body)

	switch r.Method {
	case http.MethodPost:
		message := fmt.Sprintf("POST %v", r.URL)
		fmt.Printf("transfer Post %s\n", transfer)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusOK)
	case http.MethodGet:
		message := fmt.Sprintf("GET %v", r.URL)
		fmt.Printf("transfer GET %s", transfer)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusOK)
	default:
		message := fmt.Sprintf("CallbackTransfer data in %v", r.URL)
		fmt.Printf("transfer GET %s", transfer)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusOK)
	}

}

//Codigo Antigo

func (s *ServerBank) DefaultEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	fmt.Printf("Default data in %v\n", r.URL)
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	fmt.Fprint(w, "Endpoint not found")
}

/*
 * BANK INICIA AQUI
 */
//NewServerBank Cria Servidor
func NewServerBank() *ServerBank {

	server := new(ServerBank)
	//s.Armazenamento = Armazenamento

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(server.DefaultEndpoint))
	router.Handle("/accounts", http.HandlerFunc(server.CallbackAccounts))
	//api.Handle("/accounts/{account_id}/balance", g.buildFindBalanceAccountAction()).Methods(http.MethodGet)

	router.Handle("/login/", http.HandlerFunc(server.CallbackLogin))
	router.Handle("/transfers/", http.HandlerFunc(server.CallbackTransfer))

	server.Handler = router
	return server

}

//StartAPI inicia o servidor http
func StartAPI(modo string) {
	servidor := NewServerBank()

	if err := http.ListenAndServe(serverPort, servidor); err != nil {
		log.Fatalf("Não foi possivel ouvir na porta 5000 %v", err)
	}
}
