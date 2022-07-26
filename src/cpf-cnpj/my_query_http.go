package cpfcnpj

import (
	"fmt"
	"log"
	"net/http"
)

//SaveQueryHTTP main fuction to save a new query in system INTERFACE HTTP W R
func (q *MyQuery) SaveQueryHTTP(w http.ResponseWriter, r *http.Request, newCPForCNPJ string, isCPF bool) {
	code, msg := q.SaveQueryGeneric(newCPForCNPJ, isCPF)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, msg)
}

//GetQuerysHTTP show All querys save in system INTERFACE HTTP W R
func (q *MyQuery) GetQuerysHTTP(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	code, msg := q.GetQuerysGeneric()
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, msg)
}

//GetQuerysByTypeHTTP return ALL CPF or CNPJ pass type in arg INTERFACE HTTP W R
func (q *MyQuery) GetQuerysByTypeHTTP(w http.ResponseWriter, r *http.Request, isCPF bool) {

	code, msg := q.GetQuerysByTypeGeneric(isCPF)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, msg)

}

//GetQuerysByNumGeneric return CPF/CNPJ pass number in arg INTERFACE HTTP W R
func (q *MyQuery) GetQuerysByNumHTTP(w http.ResponseWriter, r *http.Request, findCPForCNPJ string) {

	code, msg := q.GetQuerysByNumGeneric(findCPForCNPJ)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, msg)

}

//DeleteQuerysByNumHTTP Delete Number INTERFACE HTTP W R
func (q *MyQuery) DeleteQuerysByNumHTTP(w http.ResponseWriter, r *http.Request, findCPForCNPJ string, isCPF bool) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	code, msg := q.DeleteQuerysByNumGeneric(findCPForCNPJ, isCPF)
	w.WriteHeader(code)
	fmt.Fprint(w, msg)
}
