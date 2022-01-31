package m74bankapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	account "github.com/marcovargas74/m74-bank-api/src/account"
)

/*
// LoginBank
type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

//StructAndJson teste do struct
func StructAndJson() {
	myLogin := Login{CPF: "111.111.111-11", Secret: "111"}
	loginJson, _ := json.Marshal(myLogin)
	fmt.Println(string(loginJson))
	//Convert Json To struct
	var myNewLogin Login
	json.Unmarshal(loginJson, &myNewLogin)
	fmt.Println(myNewLogin.Secret)

}*/

type handler func(w http.ResponseWriter, r *http.Request)

//BasicAuth Used To authentic the user to access API
func BasicAuth(pass handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {
		user, passw, _ := r.BasicAuth()
		fmt.Printf("login [%s] [%s] \n", user, passw)

		if user == "" || passw == "" {
			w.WriteHeader(http.StatusUnauthorized)
			//w.WriteHeader(http.StatusNetworkAuthenticationRequired)
			http.Error(w, "authorization failed", http.StatusUnauthorized)
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

		r.Header.Add("token", userLogin.ID)
		pass(w, r)
	}
}
