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
	SetMongoDB = true
	SetSqlDB   = false

	SetDockerRun = true
	SetLocalRun  = false
)

//IsUsingMongoDB //If Using MongoDB this var is True
var IsUsingMongoDB bool

//GetIsUsingMongoDB Get If Using MongoDB
func GetIsUsingMongoDB() bool {
	return IsUsingMongoDB
}

//SetUsingMongoDB set If Using MongoDB
func SetUsingMongoDB(isMongoDB bool) {
	IsUsingMongoDB = isMongoDB
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

	var result bool
	if IsUsingMongoDB {
		result = q.saveQueryInMongoDB()
	} else {
		result = q.saveQueryInDB()
	}
	//result = q.saveQueryInDB()

	if !result {
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("Can not save cpf/cnpj %v ", q.Number)
		fmt.Fprint(w, message)
		return
	}

	//json, err := json.Marshal(q)
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

	if IsUsingMongoDB {
		q.showQueryAllMongoDB(w, r)
	} else {
		q.showQueryAll(w, r)
	}

	//q.showQueryAll(w, r)
}

//GetQuerysByNum return CPF/CNPJ pass number in arg
func (q *MyQuery) GetQuerysByNum(w http.ResponseWriter, r *http.Request, findCPFofCNPJ string) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	if IsUsingMongoDB {
		q.showQuerysByNumMongoDB(w, r, findCPFofCNPJ)
	} else {
		q.showQuerysByNum(w, r, findCPFofCNPJ)
	}

}

//GetQuerysByType return ALL CPF or CNPJ pass type in arg
func (q *MyQuery) GetQuerysByType(w http.ResponseWriter, r *http.Request, isCPF bool) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	if IsUsingMongoDB {
		q.showQuerysByTypeMongoDB(w, r, isCPF)
	} else {
		q.showQuerysByType(w, r, isCPF)
	}

}

//DeleteQuerysByNum Delete Number
func (q *MyQuery) DeleteQuerysByNum(w http.ResponseWriter, r *http.Request, findCPFofCNPJ string) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	if IsUsingMongoDB {
		q.deleteQuerysByNumMongoDB(w, r, findCPFofCNPJ)
	} else {
		q.deleteQuerysByNum(w, r, findCPFofCNPJ)
	}
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
