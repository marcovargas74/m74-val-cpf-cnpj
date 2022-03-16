package m74bankapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	account "github.com/marcovargas74/m74-bank-api/src/account"
)

const (
	serverPort = ":5000"
)

//ServerBank is struct to start server
type ServerBank struct {
	http.Handler
}

//CallbackAccounts function Used to handle endpoint /accounts
func (s *ServerBank) CallbackAccounts(w http.ResponseWriter, r *http.Request) {

	var accountJSON account.Account
	fmt.Printf("CallbackAccounts TOKEN: %v\n", r.Header.Get("token"))

	account.CreateDB(false)

	switch r.Method {
	case http.MethodPost:
		accountJSON.SaveAccount(w, r)
	case http.MethodGet:
		accountJSON.GetAccounts(w, r)
	default:
		message := fmt.Sprintf("MethodNotAllowed in %v", r.URL)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

//CallbackFindAccountID function Used to handle endpoint /accounts/{}
func (s *ServerBank) CallbackFindAccountID(w http.ResponseWriter, r *http.Request) {

	accountID := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("accountID [%s] \n", accountID["account_id"])

	if r.Method == http.MethodGet {
		var accountJSON account.Account
		accountJSON.ShowBalanceByID(w, r, accountID["account_id"])
		return
	}

	message := fmt.Sprintf("MethodNotAllowed in %v", r.URL)
	fmt.Fprint(w, message)
	w.WriteHeader(http.StatusBadRequest)

}

//CallbackTransfer function Used to handle endpoint /transfers/{}
func (s *ServerBank) CallbackTransfer(w http.ResponseWriter, r *http.Request) {

	var aTransfer account.TransferBank
	tokenAccount := r.Header.Get("token")
	log.Printf("CallbackTransfer TOKEN: %v\n", r.Header.Get("token"))

	switch r.Method {
	case http.MethodPost:
		aTransfer.SaveTransfer(w, r, tokenAccount)
	case http.MethodGet:
		aTransfer.GetTransfersByID(w, r, tokenAccount)
	default:
		message := fmt.Sprintf("MethodNotAllowed in %v", r.URL)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

//CallbackTransferByID function Used to handle endpoint /transfers/{}
func (s *ServerBank) CallbackTransferByID(w http.ResponseWriter, r *http.Request) {

	var aTransfer account.TransferBank

	accountID := mux.Vars(r)
	log.Printf("accountID [%s] \n", accountID["account_id"])

	switch r.Method {
	case http.MethodPost:
		aTransfer.SaveTransfer(w, r, accountID["account_id"])
	case http.MethodGet:
		aTransfer.GetTransfersByID(w, r, accountID["account_id"])
	default:
		message := fmt.Sprintf("MethodNotAllowed in %v", r.URL)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

//CallbackLogin function Used to handle endpoint /login - just for test
func (s *ServerBank) CallbackLogin(w http.ResponseWriter, r *http.Request) {

	account.CreateDB(false)
	user, passw, _ := r.BasicAuth()

	if user == "" || passw == "" || r.Method != http.MethodPost {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Authentication Required")
		return
	}

	userLogin, err := account.GetAccountByCPF(user)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "PASS NOT FOUND")
		return
	}

	if account.HashToSecret(userLogin.Secret) != passw {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "FORBIDDEN - ACCESS UNAUTHORIZED")
		return
	}
	json, _ := json.Marshal(userLogin.ID)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	w.WriteHeader(http.StatusOK)
}

//DefaultEndpoint function Used to handle endpoint /- can be a load a page in html to configure
func (s *ServerBank) DefaultEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	fmt.Printf("Default data in %v\n", r.URL)
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	fmt.Fprint(w, "Endpoint not found")
}

//NewServerBank Create Server
func NewServerBank(mode string) *ServerBank {

	server := new(ServerBank)

	routerG := mux.NewRouter()
	routerG.HandleFunc("/login/", server.CallbackLogin)
	routerG.HandleFunc("/", BasicAuth(server.DefaultEndpoint))
	routerG.HandleFunc("/accounts", server.CallbackAccounts)
	routerG.HandleFunc("/accounts/{account_id}/balance", BasicAuth(server.CallbackFindAccountID))
	routerG.HandleFunc("/transfers", BasicAuth(server.CallbackTransfer))

	if mode == "dev" {
		routerG.HandleFunc("/transfers/{account_id}", server.CallbackTransferByID)
	}
	server.Handler = routerG
	return server

}

//StartAPI http starter server
func StartAPI(mode string) {
	servidor := NewServerBank(mode)

	if err := http.ListenAndServe(serverPort, servidor); err != nil {
		log.Printf("Não foi possivel ouvir na porta 5000 %v", err)
	}
}
