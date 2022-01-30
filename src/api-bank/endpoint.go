package m74bankapi

import (
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

	//TOD deve estar autenticada
	// filterID string

	switch r.Method {
	case http.MethodPost:
		aTransfer.SaveTransfer(w, r)
	case http.MethodGet:
		aTransfer.GetTransfers(w, r)
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

	accountID := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Printf("accountID [%s] \n", accountID["account_id"]) //TOD deve estar autenticada
	// filterID string

	switch r.Method {
	case http.MethodPost:
		aTransfer.SaveTransfer(w, r)
	case http.MethodGet:
		aTransfer.GetTransfersByID(w, r, accountID["account_id"])
	default:
		message := fmt.Sprintf("MethodNotAllowed in %v", r.URL)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (s *ServerBank) CallbackLogin(w http.ResponseWriter, r *http.Request) {

	//client := r.URL.Path[len("/login/"):]
	fmt.Printf("Login Body: %v\n", r.Body)

	w.Header().Set("content-type", "application/json")
	if r.Method == http.MethodPost {
		loginData := mux.Vars(r)
		//w.WriteHeader(http.StatusOK)
		fmt.Printf("login [%s] [%s] \n", loginData["cpf"], loginData["cpf"])

		//var accountJSON account.Account
		//countJSON.CheckLoginD(w, r, loginData["account_id"])
		message := fmt.Sprintf("Access Autorized %v", r.URL)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusOK)

		return
	}
	message := fmt.Sprintf("ACCESS DANIED %v", r.URL)
	//fmt.Printf("client GET %s", client)
	fmt.Fprint(w, message)
	w.WriteHeader(http.StatusForbidden)

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

	/*router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(server.DefaultEndpoint))
	router.Handle("/accounts", http.HandlerFunc(server.CallbackAccounts))
	router.Handle("/accounts/{account_id}/balance", http.HandlerFunc(server.CallbackFindAccountID))
	router.Handle("/login/", http.HandlerFunc(server.CallbackLogin))
	router.Handle("/transfers/", http.HandlerFunc(server.CallbackTransfer))*/

	//GORILAS
	routerG := mux.NewRouter()
	//routerG.
	routerG.HandleFunc("/", server.DefaultEndpoint)
	routerG.HandleFunc("/accounts", server.CallbackAccounts)
	routerG.HandleFunc("/accounts/{account_id}/balance", server.CallbackFindAccountID)
	routerG.HandleFunc("/login/", server.CallbackLogin)
	routerG.HandleFunc("/transfers", server.CallbackTransfer)
	routerG.HandleFunc("/transfers/{account_id}", server.CallbackTransferByID)

	server.Handler = routerG

	//server.Handler = router
	return server

}

//StartAPI inicia o servidor http
func StartAPI(modo string) {
	servidor := NewServerBank()

	if err := http.ListenAndServe(serverPort, servidor); err != nil {
		log.Fatalf("Não foi possivel ouvir na porta 5000 %v", err)
	}
}
