package cpfcnpj

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//MyQuery Strutc Main Used in AppValidate
type MyQuery struct {
	ID        string    `json:"id" bson:"id"`
	Number    string    `json:"cpf" bson:"cpf"`
	IsValid   bool      `json:"is_valid" bson:"is_valid"`
	IsCPF     bool      `json:"is_cpf" bson:"is_cpf"`
	IsCNPJ    bool      `json:"is_cnpj" bson:"is_cnpj"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

const (
	SetDockerRun = true
	SetLocalRun  = false
)

//IsUsingMongoDocker If Using MongoDB in  a Docker image this var is True
var IsUsingMongoDocker bool

//GetIsUsingMongoDocker Get If Using MongoDB in  a Docker image
func GetIsUsingMongoDocker() bool {
	return IsUsingMongoDocker
}

//SetUsingMongoDocker set If Using MongoDB in  a Docker image
func SetUsingMongoDocker(isMongoDocker bool) {
	IsUsingMongoDocker = isMongoDocker
}

//SaveQuery main fuction to save a new query in system
func (q *MyQuery) SaveQuery(w http.ResponseWriter, r *http.Request, newCPFofCNPJ string, isCPF bool) {

	q.Number = newCPFofCNPJ

	q.IsCNPJ = false
	q.IsCPF = false
	if isCPF {
		q.IsCPF = true
		if !IsValidCPF(q.Number) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Something gone wrong: Invalid CPF\n")
			log.Printf("Something gone wrong: Invalid CPF:%s\n", q.Number)
			q.IsValid = false
			return
		}
	} else {
		q.IsCNPJ = true
		if !IsValidCNPJ(q.Number) {
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprint(w, "Something gone wrong: Invalid CNPJ")
			log.Printf("Something gone wrong: Invalid CNPJ:%s\n", q.Number)
			q.IsValid = false
			return
		}
	}

	q.IsValid = true
	q.ID = NewUUID()
	q.CreatedAt = time.Now()
	fmt.Printf("UUIDv4: %s\n", q.ID)

	if !IsValidUUID(q.ID) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Something gone wrong: Invalid ID\n")
		log.Printf("Something gone wrong: Invalid ID:%s\n", q.ID)
		return
	}

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] number[%s] \n", r.UserAgent(), q.Number)
		w.WriteHeader(http.StatusOK)
		return
	}

	result := q.saveQueryInMongoDB()

	if !result {
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("Can not save cpf/cnpj %v ", q.Number)
		fmt.Fprint(w, message)
		return
	}

	json, err := q.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}

//GetQuerys show All querys save in system
func (q *MyQuery) GetQuerys(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	q.showQueryAllMongoDB(w, r)

}

//GetQuerysByNum return CPF/CNPJ pass number in arg
func (q *MyQuery) GetQuerysByNum(w http.ResponseWriter, r *http.Request, findCPFofCNPJ string) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	q.showQuerysByNumMongoDB(w, r, findCPFofCNPJ)

}

//GetQuerysByType return ALL CPF or CNPJ pass type in arg
func (q *MyQuery) GetQuerysByType(w http.ResponseWriter, r *http.Request, isCPF bool) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	q.showQuerysByTypeMongoDB(w, r, isCPF)

}

//DeleteQuerysByNum Delete Number
func (q *MyQuery) DeleteQuerysByNum(w http.ResponseWriter, r *http.Request, findCPFofCNPJ string) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	q.deleteQuerysByNumMongoDB(w, r, findCPFofCNPJ)
}

//MarshalJSON format the date
func (q *MyQuery) MarshalJSON() ([]byte, error) {
	type Alias MyQuery
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"created_at"`
	}{
		Alias:     (*Alias)(q),
		CreatedAt: q.CreatedAt.Format("02-Jan-06 15:04:05"),
	})
}
