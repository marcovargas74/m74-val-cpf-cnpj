package main

import (
	"log"

	account "github.com/marcovargas74/m74-val-cpf-cnpj/src/cpf-cnpj"

	validator "github.com/marcovargas74/m74-val-cpf-cnpj/src/api-validator"
)

func init() {
	account.CreateDB(false)

}

func main() {
	log.Printf("======== API VALIDATOR Version %s \n", validator.GetVersion())
	validator.StartAPI("dev")

}