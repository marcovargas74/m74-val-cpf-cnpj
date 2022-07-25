package main

import (
	"log"

	myquery "github.com/marcovargas74/m74-val-cpf-cnpj/src/cpf-cnpj"

	validator "github.com/marcovargas74/m74-val-cpf-cnpj/src/api-validator"
)

func init() {
	myquery.CreateStatus()
	myquery.SetUsingMongoDocker(myquery.SetDockerRun)
	myquery.CreateDB()
}

func main() {
	log.Printf("======== API VALIDATOR Version %s \n", validator.GetVersion())
	validator.StartAPI("dev")

}
