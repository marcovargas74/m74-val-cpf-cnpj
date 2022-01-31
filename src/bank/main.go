package main

import (
	"log"

	"github.com/marcovargas74/m74-bank-api/src/account"
	bank "github.com/marcovargas74/m74-bank-api/src/api-bank"
)

var isProduction = false

func init() {
	bank.SetIsProduction(isProduction)
	account.CreateDB(false)

}

func main() {
	log.Printf("======== API BANK Version %s isPruduction=%v\n", bank.GetVersion(), bank.GetIsProduction())
	bank.StartAPI("dev")

}
