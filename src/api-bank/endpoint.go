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

type ServerBank struct {
	//Armazenamento ArmazenamentoAccount
	http.Handler
}

func (s *ServerBank) CallbackAccounts(w http.ResponseWriter, r *http.Request) {

	var accountJSON account.Account
	fmt.Printf("CallbackAccounts TOKEN: %v\n", r.Header.Get("token"))

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

func (s *ServerBank) CallbackFindAccountID(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("!!!CallbackFindAccountID GET %v\n", r.URL)
	//w.Header().Set("content-type", "application/json")

	//const logKey = "find_balance_account"
	accountID := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("accountID [%s] \n", accountID["account_id"])
	/*if !domain.IsValidUUID(accountID) {
		var err = response.ErrParameterInvalid
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("invalid parameter")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}*/

	if r.Method == http.MethodGet {
		var accountJSON account.Account
		accountJSON.ShowBalanceByID(w, r, accountID["account_id"])
		return
	}

	message := fmt.Sprintf("MethodNotAllowed in %v", r.URL)
	fmt.Fprint(w, message)
	w.WriteHeader(http.StatusBadRequest)

}

func (s *ServerBank) CallbackTransfer(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("content-type", "application/json")
	//transfer := r.URL.Path[len("/transfers/"):]
	//fmt.Printf("tranfer: %v\n", r.Body)
	var aTransfer account.TransferBank
	tokenAccount := r.Header.Get("token")
	fmt.Printf("CallbackTransfer TOKEN: %v\n", r.Header.Get("token"))

	//TOD deve estar autenticada
	// filterID string

	switch r.Method {
	case http.MethodPost:
		aTransfer.SaveTransfer(w, r, tokenAccount)
	case http.MethodGet:
		//aTransfer.GetTransfers(w, r)
		aTransfer.GetTransfersByID(w, r, tokenAccount)
	default:
		message := fmt.Sprintf("MethodNotAllowed in %v", r.URL)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (s *ServerBank) CallbackTransferByID(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("content-type", "application/json")
	//transfer := r.URL.Path[len("/transfers/"):]
	//fmt.Printf("tranfer: %v\n", r.Body)
	var aTransfer account.TransferBank
	//fmt.Printf("CallbackTransferByID TOKEN: %v\n", r.Header.Get("token"))

	accountID := mux.Vars(r)
	//tokenAccount := r.Header.Get("token")
	//w.WriteHeader(http.StatusOK)
	fmt.Printf("accountID [%s] \n", accountID["account_id"]) //TOD deve estar autenticada
	// filterID string

	switch r.Method {
	case http.MethodPost:
		aTransfer.SaveTransfer(w, r, accountID["account_id"])
	case http.MethodGet:
		aTransfer.GetTransfersByID(w, r, accountID["account_id"])
		//aTransfer.GetTransfersByID(w, r, tokenAccount)
	default:
		message := fmt.Sprintf("MethodNotAllowed in %v", r.URL)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (s *ServerBank) CallbackLogin(w http.ResponseWriter, r *http.Request) {

	//fmt.Printf("Login Body: %v\n", r.Body)
	//fmt.Printf("Login Body: %v\n", r.Header.Get())

	user, passw, _ := r.BasicAuth()
	fmt.Printf("login [%s] [%s] \n", user, passw)

	if user == "" || passw == "" || r.Method != http.MethodPost {
		w.WriteHeader(http.StatusUnauthorized)
		//w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		fmt.Fprint(w, "Authentication Required")
		return
	}

	userLogin, err := account.GetAccountByCPF(user)
	if err != nil {
		log.Fatal(err)
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
func NewServerBank(mode string) *ServerBank {

	server := new(ServerBank)

	//GORILAS
	routerG := mux.NewRouter()
	routerG.HandleFunc("/login/", server.CallbackLogin)
	routerG.HandleFunc("/", BasicAuth(server.DefaultEndpoint))
	//routerG.HandleFunc("/accounts", BasicAuth(server.CallbackAccounts))
	routerG.HandleFunc("/accounts", server.CallbackAccounts)
	routerG.HandleFunc("/accounts/{account_id}/balance", BasicAuth(server.CallbackFindAccountID))
	routerG.HandleFunc("/transfers", BasicAuth(server.CallbackTransfer))

	if mode == "dev" {
		//Incluir endpoint para testar não precisa authenticação(SOMENTE PARA TESTE)
		routerG.HandleFunc("/transfers/{account_id}", server.CallbackTransferByID)
	}
	server.Handler = routerG
	return server

}

//StartAPI inicia o servidor http
func StartAPI(mode string) {
	servidor := NewServerBank(mode)

	if err := http.ListenAndServe(serverPort, servidor); err != nil {
		log.Fatalf("Não foi possivel ouvir na porta 5000 %v", err)
	}
}

//Trash CODE
