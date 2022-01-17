package m74bankAPI

import (
	"fmt"
	"log"
	"net/http"
)

const (
	serverPort = ":5000"
)

func getPlayerPoints(name string) string {

	if name == "Maria" {
		return "20"
	}

	if name == "Pedro" {
		return "10"
	}
	return ""

}

//ServidorJogador teste
func ServidorJogador(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := r.URL.Path[len("/jogadores/"):]
	fmt.Println("entrada " + player)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, getPlayerPoints(player))
}

func DefaultEndpoint(w http.ResponseWriter, r *http.Request) {

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
func callbackAccount(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("callbackAccount data in %v\n", r.URL)
	fmt.Fprint(w, message)
}

func callbackLogin(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("callbackLogin data in %v\n", r.URL)
	fmt.Fprint(w, message)
}

func callbackTransfer(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("callbackTransfer data in %v\n", r.URL)
	fmt.Fprint(w, message)
}

//StartAPI inicia o servidor http
func StartAPI(modo string) {
	//tratador := http.HandlerFunc(ServidorJogador)
	//log.Fatal(http.ListenAndServe(serverPort, tratador))
	HandleFuncions()
	log.Fatal(http.ListenAndServe(serverPort, nil))
}

//HandleFuncions Inclui os endpoint
func HandleFuncions() {
	http.HandleFunc("/", DefaultEndpoint)
	http.HandleFunc("/jogadores/Maria", ServidorJogador)
	http.HandleFunc("/jogadores/Pedro", ServidorJogador)

	//*TODO endpoint usado no banc
	http.HandleFunc("/accounts", callbackAccount)
	http.HandleFunc("/login", callbackLogin)
	http.HandleFunc("/transfers", callbackTransfer)
}
