package m74validatorapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	cpfcnpj "github.com/marcovargas74/m74-val-cpf-cnpj/src/cpf-cnpj"
)

const (
	serverPort = ":5000"
)

//ServerValidator is struct to start server
type ServerValidator struct {
	http.Handler
}

//CallbackQuerys function Used to handle endpoint /cpfs and /cnpjs
func (s *ServerValidator) CallbackQuerys(w http.ResponseWriter, r *http.Request) {

	var aQueryJSON cpfcnpj.MyQuery
	fmt.Printf("CallbackQuerys TOKEN: %v\n", r.Header.Get("token"))

	cpfcnpj.CreateDB(false)

	switch r.Method {
	case http.MethodPost:
		aQueryJSON.SaveQuery(w, r)
	case http.MethodGet:
		aQueryJSON.GetQuerys(w, r)
	default:
		message := fmt.Sprintf("MethodNotAllowed in %v", r.URL)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

/*
//CallbackFindAccountID function Used to handle endpoint /accounts/{}
func (s *ServerValidator) CallbackFindAccountID(w http.ResponseWriter, r *http.Request) {

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

//CallbackLogin function Used to handle endpoint /login - just for test
func (s *ServerValidator) CallbackLogin(w http.ResponseWriter, r *http.Request) {

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
*/
//DefaultEndpoint function Used to handle endpoint /- can be a load a page in html to configure
func (s *ServerValidator) DefaultEndpoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	fmt.Printf("Default data in %v\n", r.URL)
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	fmt.Fprint(w, "Endpoint not found")
}

//NewServerValidator Create Server
func NewServerValidator(mode string) *ServerValidator {

	server := new(ServerValidator)

	routerG := mux.NewRouter()
	//routerG.HandleFunc("/login/", server.CallbackLogin)
	routerG.HandleFunc("/", BasicAuth(server.DefaultEndpoint))
	routerG.HandleFunc("/cpfs", server.CallbackQuerys)
	routerG.HandleFunc("/cnpjs", server.CallbackQuerys)
	//routerG.HandleFunc("/accounts/{account_id}/balance", BasicAuth(server.CallbackFindAccountID))
	//routerG.HandleFunc("/status", BasicAuth(server.CallbackTransfer))

	// if mode == "dev" {
	// 	routerG.HandleFunc("/transfers/{account_id}", server.CallbackTransferByID)
	// }
	server.Handler = routerG
	return server

}

//StartAPI http starter server
func StartAPI(mode string) {
	servidor := NewServerValidator(mode)

	if err := http.ListenAndServe(serverPort, servidor); err != nil {
		log.Printf("Fail to conect in a port-> 5000 %v", err)
	}
}
