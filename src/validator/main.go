package main

import (
	"log"

	myquery "github.com/marcovargas74/m74-val-cpf-cnpj/src/cpf-cnpj"

	validator "github.com/marcovargas74/m74-val-cpf-cnpj/src/api-validator"
)

func init() {
	//myquery.CreateDB(false)
	myquery.SetUsingMongoDB(myquery.SetMongoDB) //SetSqlDB
	if myquery.GetIsUsingMongoDB() {
		myquery.InitDBMongo(myquery.SetLocalRun) //SetDockerRun
	} else {
		myquery.CreateStatus()
	}

}

func main() {
	log.Printf("======== API VALIDATOR Version %s \n", validator.GetVersion())
	validator.StartAPI("dev")

}
