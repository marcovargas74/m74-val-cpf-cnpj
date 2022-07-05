package m74validatorapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	account "github.com/marcovargas74/m74-val-cpf-cnpj/src/cpf-cnpj"
)

type handler func(w http.ResponseWriter, r *http.Request)

//BasicAuth Used To authentic the user to access API
func BasicAuth(pass handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {
		user, passw, _ := r.BasicAuth()
		fmt.Printf("login [%s] [%s] \n", user, passw)

		if user == "" || passw == "" {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "authorization failed", http.StatusUnauthorized)
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

		r.Header.Add("token", userLogin.ID)
		pass(w, r)
	}
}
