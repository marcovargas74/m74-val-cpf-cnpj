package cpfcnpj

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//NumTotalQuery VAR used to store number of queries performed since start
//var NumTotalQuery string
//var UPTime time.Time

//NumTotalQuery VAR used to store number of queries performed since start
var StatusQuery QueryStatus

// //Status used to create a Json Return to endpoint
type QueryStatus struct {
	NumTotalQuery uint64    `json:"num_total_query"`
	StartTime     time.Time `json:"start_time"`
	UpTime        float64   `json:"up_time"`
}

//CreateDB Create SQL dataBase
func CreateStatus() {
	StatusQuery.NumTotalQuery = 0
	StatusQuery.StartTime = time.Now()
}

func UpdateStatus() {
	StatusQuery.NumTotalQuery++
	//StatusQuery.UpTime = time.Now()
}

//CreateDB Create SQL dataBase
func GetNumQuery() uint64 {
	return StatusQuery.NumTotalQuery
}

//CreateDB Create SQL dataBase
func GetUptimeQuery() float64 {
	//newTime := StatusQuery.UpTime
	// Prints time elapse
	//	fmt.Println("time elapse:", time.Since(StatusQuery.UpTime))
	timeElapse := time.Since(StatusQuery.StartTime)

	fmt.Println("time elapse2:", timeElapse)
	return (timeElapse.Seconds())
}

/*
//GetQuerys show All querys save in system
func (q *MyQuery) GetQuerys(w http.ResponseWriter, r *http.Request) {
	q.ShowQueryAll(w, r)
	w.WriteHeader(http.StatusOK)
}
*/
//ShowQueryAll Show all querys
//func (q *QueryStatus) ShowStatus(w http.ResponseWriter, r *http.Request) {
func ShowStatus(w http.ResponseWriter, r *http.Request) {

	StatusQuery.UpTime = GetUptimeQuery()
	json, err := json.Marshal(StatusQuery)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
	w.WriteHeader(http.StatusOK)

	fmt.Println("ShowStatus:", string(json))

}
