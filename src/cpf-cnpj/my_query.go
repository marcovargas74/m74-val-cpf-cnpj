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

//SaveQueryHTTP main fuction to save a new query in system
func (q *MyQuery) SaveQueryHTTP(w http.ResponseWriter, r *http.Request, newCPForCNPJ string, isCPF bool) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	code, msg := q.SaveQueryGeneric(newCPForCNPJ, isCPF)
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, msg)

	/*q.Number = newCPForCNPJ

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
	if result != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := fmt.Sprintf("Can not save cpf/cnpj %v ", q.Number)
		fmt.Fprint(w, message)
		return
	}

	/*json, err := q.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))*/

}

//GetQuerysHTTP show All querys save in system
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

	//q.showQueryAllMongoDB(w, r)

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

//DeleteQuerysByNumHTTP Delete Number INTERFACE HTTP W R
func (q *MyQuery) DeleteQuerysByNumHTTP(w http.ResponseWriter, r *http.Request, findCPFofCNPJ string) {

	if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] \n", r.UserAgent())
		w.WriteHeader(http.StatusOK)
		return
	}

	code, msg := q.DeleteQuerysByNumGeneric(findCPFofCNPJ)
	w.WriteHeader(code)
	fmt.Fprint(w, msg)
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

/*NWEEEEW*/

//ValidCPFQueryGeneric Valid CPF
func (q *MyQuery) ValidCPFQueryGeneric(newCPF string) (int, string) {

	q.Number = newCPF
	q.IsValid = true
	q.IsCNPJ = false

	q.IsCPF = true
	if !IsValidCPF(q.Number) {
		q.IsValid = false
		message := fmt.Sprintf("Something gone wrong: Invalid CPF:%s\n", q.Number)
		log.Println(message)
		return http.StatusBadRequest, message
	}

	return http.StatusOK, "SUCCESS CPF is VALID"
}

//ValidCNPJQueryGeneric Valid CNPJ
func (q *MyQuery) ValidCNPJQueryGeneric(newCNPJ string) (int, string) {

	q.Number = newCNPJ
	q.IsValid = true
	q.IsCPF = false

	q.IsCNPJ = true
	if !IsValidCNPJ(q.Number) {
		q.IsValid = false
		message := fmt.Sprintf("Something gone wrong: Invalid CNPJ:%s\n", q.Number)
		log.Println(message)
		return http.StatusBadRequest, message
	}

	return http.StatusOK, "SUCCESS CNPJ is VALID"
}

//SaveQueryGeneric main fuction to save a new query in system
func (q *MyQuery) SaveQueryGeneric(newCPFofCNPJ string, isCPF bool) (int, string) {

	//q.Number = newCPFofCNPJ
	var code int
	var msg string
	if isCPF {
		code, msg = q.ValidCPFQueryGeneric(newCPFofCNPJ)
	} else {
		code, msg = q.ValidCNPJQueryGeneric(newCPFofCNPJ)
	}

	if code != http.StatusOK {
		return code, msg
	}

	/*q.IsCNPJ = false
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

	q.IsValid = true*/
	q.ID = NewUUID()
	q.CreatedAt = time.Now()
	//fmt.Printf("UUIDv4: %s\n", q.ID)

	if !IsValidUUID(q.ID) {
		//w.WriteHeader(http.StatusBadRequest)
		//fmt.Fprint(w, "Something gone wrong: Invalid ID\n")
		message := fmt.Sprintf("Something gone wrong: Invalid ID:%s\n", q.ID)
		log.Println(message)
		return http.StatusBadRequest, message
	}

	/*if r.UserAgent() == "self_test" {
		log.Printf("Its Only a TEST [%s] number[%s] \n", r.UserAgent(), q.Number)
		message := fmt.Sprintf("Its Only a TEST [%s] number[%s] ", q.Number)
		//w.WriteHeader(http.StatusOK)
		return http.StatusOK
	}*/

	result := q.saveQueryInMongoDB()
	if result != nil {
		message := fmt.Sprintf("Can not save cpf/cnpj %v ", q.Number)
		log.Println(message)
		return http.StatusInternalServerError, message
	}

	json, err := q.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	/*w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))*/
	return http.StatusOK, string(json)
}

//GetQuerysHTTP show All querys save in system
func (q *MyQuery) GetQuerysGeneric() (int, string) {

	/*code, msg := q.GetQuerysGeneric()
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, msg)

	q.showQueryAllMongoDB(w, r)

	err := q.showQueryAllMongoDB(findCPFofCNPJ)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound, err.Error()
	}*/

	return http.StatusOK, "SUCCESS TO DELETE CPF/CNPJ"

}

//DeleteQuerysByNum Delete Number
func (q *MyQuery) DeleteQuerysByNumGeneric(findCPFofCNPJ string) (int, string) {

	err := q.deleteQuerysByNumMongoDB(findCPFofCNPJ)
	if err != nil {
		log.Println(err)
		return http.StatusNotFound, err.Error()
	}

	return http.StatusOK, "SUCCESS TO DELETE CPF/CNPJ"
}
